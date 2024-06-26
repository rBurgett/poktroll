package channel

import (
	"context"
	"sync"

	"github.com/pokt-network/poktroll/pkg/observable"
	"github.com/pokt-network/poktroll/pkg/polylog"
)

var (
	_ observable.ReplayObservable[any] = (*replayObservable[any])(nil)
	_ observable.Observable[any]       = (*replayObservable[any])(nil)
)

type replayObservable[V any] struct {
	// replayBufferSize is the number of notifications to buffer so that they
	// can be replayed to new observers.
	replayBufferSize int
	// replayBufferMu protects replayBuffer from concurrent access/updates.
	replayBufferMu sync.RWMutex
	// replayBuffer holds the last relayBufferSize number of notifications received
	// by this observable. This buffer is replayed to new observers, on subscribing,
	// prior to any new notifications being propagated.
	replayBuffer []V
	// bufferingObsvbl is an observable that emits all buffered values in one
	// notification.
	bufferingObsvbl observable.Observable[[]V]
}

// NewReplayObservable returns a new ReplayObservable with the given replay buffer
// replayBufferSize and the corresponding publish channel to notify it of new values.
func NewReplayObservable[V any](
	ctx context.Context,
	replayBufferSize int,
	opts ...option[V],
) (observable.ReplayObservable[V], chan<- V) {
	obsvbl, publishCh := NewObservable[V](opts...)

	return ToReplayObservable(ctx, replayBufferSize, obsvbl), publishCh
}

// ToReplayObservable returns an observable which replays the last replayBufferSize
// number of values published to the source observable to new observers, before
// publishing new values.
// It should only be used with a srcObservable which contains channelObservers
// (i.e. channelObservable or similar).
func ToReplayObservable[V any](
	ctx context.Context,
	replayBufferSize int,
	srcObsvbl observable.Observable[V],
) observable.ReplayObservable[V] {
	replayObsvbl := &replayObservable[V]{
		replayBufferSize: replayBufferSize,
		replayBuffer:     []V{},
	}

	replayObsvbl.bufferingObsvbl = replayObsvbl.initBufferingObservable(ctx, srcObsvbl)
	return replayObsvbl
}

// Last synchronously returns the last n values from the replay buffer.
// It blocks until n values have been accumulated or its context is canceled,
// If it is canceled before n values are accumulated, it returns all the available
// items at the time of cancellation.
// The values returned are ordered from newest to oldest (i.e. LIFO)
func (ro *replayObservable[V]) Last(ctx context.Context, n int) []V {
	logger := polylog.Ctx(ctx)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// If n is greater than the replay buffer size, return the entire replay buffer.
	if n > ro.replayBufferSize {
		n = ro.replayBufferSize
		logger.Warn().
			Int("requested_replay_buffer_size", n).
			Int("replay_buffer_capacity", ro.replayBufferSize).
			Msg("requested replay buffer size is greater than replay buffer capacity; returning entire replay buffer")
	}

	// Lock any concurrent updates to the replay buffer.
	ro.replayBufferMu.RLock()

	// If the replay buffer has enough values, return the most recent n values.
	if len(ro.replayBuffer) >= n {
		values := ro.replayBuffer[:n]
		ro.replayBufferMu.RUnlock()
		return values
	}

	// If the replay buffer does not have enough values, wait for the source observable
	// to emit enough values to satisfy the request.
	bufferedValuesCh := ro.bufferingObsvbl.Subscribe(ctx).Ch()
	// Initialize latestValues with the values in the replay buffer in case the
	// source observable is closed or the context is canceled before it has a chance
	// to emit an updated buffer of values.
	latestValues := ro.replayBuffer[:]
	// Unlock the replay buffer to allow new values to be added.
	// These new values will be collected in the loop below instead of the replay buffer.
	ro.replayBufferMu.RUnlock()
	// bufferValuesCh emits all buffered values in one notification.
	for values := range bufferedValuesCh {
		// If n is greater than the number of values emitted, update latestValues with
		// the most recent values emitted so far so it could be returned if the context
		// is canceled or the source observable is closed.
		if len(values) >= n {
			latestValues = values[:n]
			break
		}

		// Update latestValues with the most recent values emitted so far.
		latestValues = values[:]
	}

	// Return the most recent n values or all available values if the context is canceled
	// before n values are accumulated.
	return latestValues
}

// Subscribe returns an observer which is notified when the publishCh channel
// receives a value.
// It replays the values stored in the replay buffer in the order of their arrival
// before emitting new values.
func (ro *replayObservable[V]) Subscribe(ctx context.Context) observable.Observer[V] {
	obs, ch := NewObservable[V]()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		// Replay the values stored in the buffer form the oldest to the newest.
		ro.replayBufferMu.RLock()
		for i := len(ro.replayBuffer) - 1; i >= 0; i-- {
			ch <- ro.replayBuffer[i]
		}

		bufferedValuesCh := ro.bufferingObsvbl.Subscribe(ctx).Ch()
		ro.replayBufferMu.RUnlock()
		// Since bufferingObsvbl emits all buffered values in one notification
		// and the replay buffer has already been replayed, only the most recent
		// value needs to be published
		for values := range bufferedValuesCh {
			ch <- values[0]
		}
		// Cancel the context to stop the observer when the source observable is closed.
		cancel()
	}()

	return obs.Subscribe(ctx)
}

// UnsubscribeAll unsubscribes all observers from the replay observable.
func (ro *replayObservable[V]) UnsubscribeAll() {
	ro.bufferingObsvbl.UnsubscribeAll()
}

// initBufferingObservable receives and buffers the last n notifications from
// the a source observable and emits all buffered values at once.
func (ro *replayObservable[V]) initBufferingObservable(
	ctx context.Context,
	srcObsvbl observable.Observable[V],
) observable.Observable[[]V] {
	bufferedObsvbl, bufferedObsvblCh := NewObservable[[]V]()
	ch := srcObsvbl.Subscribe(ctx).Ch()
	subscriptionReady := make(chan struct{})

	go func() {
		subscriptionReady <- struct{}{}
		for value := range ch {
			ro.replayBufferMu.Lock()
			// The newest value is always at the beginning of the replay buffer.
			if len(ro.replayBuffer) < ro.replayBufferSize {
				ro.replayBuffer = append([]V{value}, ro.replayBuffer...)
			} else {
				ro.replayBuffer = append([]V{value}, ro.replayBuffer[:ro.replayBufferSize-1]...)
			}
			// Emit all buffered values at once.
			bufferedObsvblCh <- ro.replayBuffer
			ro.replayBufferMu.Unlock()
		}
	}()

	// Ensure that the source observable has been subscribed to before allowing
	// the replay observable to be subscribed to.
	// It is needed to ensure that no values are missed by the replay observable
	<-subscriptionReady
	return bufferedObsvbl
}
