package sdk

import (
	"context"
	"crypto/tls"
	"net/url"

	"cosmossdk.io/depinject"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	grpctypes "github.com/cosmos/gogoproto/grpc"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/pokt-network/poktroll/pkg/client/block"
	"github.com/pokt-network/poktroll/pkg/client/delegation"
	eventsquery "github.com/pokt-network/poktroll/pkg/client/events"
	"github.com/pokt-network/poktroll/pkg/client/query"
	"github.com/pokt-network/poktroll/pkg/crypto/rings"
	"github.com/pokt-network/poktroll/pkg/polylog"
)

// buildDeps builds the dependencies for the POKTRollSDK if they are not provided
// in the config. This is useful for the SDK consumers that do not want or
// cannot provide the dependencies through depinject.
func (sdk *poktrollSDK) buildDeps(
	ctx context.Context,
	config *POKTRollSDKConfig,
) (depinject.Config, error) {
	pocketNodeWebsocketURL := RPCToWebsocketURL(config.QueryNodeUrl)

	cometClient, err := sdkclient.NewClientFromNode(config.QueryNodeUrl.String())
	if err != nil {
		return nil, err
	}

	// Supply the logger & cosmos RPC client.
	deps := depinject.Configs(
		depinject.Supply(
			polylog.Ctx(ctx),
			cometClient,
		),
	)

	// Create and supply the events query client
	eventsQueryClient := eventsquery.NewEventsQueryClient(pocketNodeWebsocketURL)
	deps = depinject.Configs(deps, depinject.Supply(eventsQueryClient))

	blockClientConnRetryLimitOpt := block.WithConnRetryLimit(sdk.config.ConnRetryLimit)

	// Create and supply the block client that depends on the events query client
	blockClient, err := block.NewBlockClient(ctx, deps, blockClientConnRetryLimitOpt)
	if err != nil {
		return nil, err
	}
	deps = depinject.Configs(deps, depinject.Supply(blockClient))

	creds, err := getTransportCreds(config.QueryNodeGRPCUrl)
	if err != nil {
		return nil, err
	}

	// Create and supply the grpc client used by the queriers
	// TODO_TECHDEBT: Configure the grpc client options from the config.
	var grpcClient grpctypes.ClientConn
	grpcClient, err = grpc.Dial(
		config.QueryNodeGRPCUrl.Host,
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		return nil, err
	}
	deps = depinject.Configs(deps, depinject.Supply(grpcClient))

	// Create the account querier and add it to the required dependencies.
	accountQuerier, err := query.NewAccountQuerier(deps)
	if err != nil {
		return nil, err
	}
	deps = depinject.Configs(deps, depinject.Supply(accountQuerier))

	// Create and supply the application querier
	applicationQuerier, err := query.NewApplicationQuerier(deps)
	if err != nil {
		return nil, err
	}
	deps = depinject.Configs(deps, depinject.Supply(applicationQuerier))

	// Create and supply the session querier
	sessionQuerier, err := query.NewSessionQuerier(deps)
	if err != nil {
		return nil, err
	}
	deps = depinject.Configs(deps, depinject.Supply(sessionQuerier))

	delegationClientConnRetryLimitOpt := delegation.WithConnRetryLimit(sdk.config.ConnRetryLimit)

	// Create and supply the delegation client
	delegationClient, err := delegation.NewDelegationClient(ctx, deps, delegationClientConnRetryLimitOpt)
	if err != nil {
		return nil, err
	}
	deps = depinject.Configs(deps, depinject.Supply(delegationClient))

	// Create and supply the ring cache that depends on:
	// the logger, application and account querier and the delegation client
	ringCache, err := rings.NewRingCache(deps)
	if err != nil {
		return nil, err
	}
	deps = depinject.Configs(deps, depinject.Supply(ringCache))

	return deps, nil
}

// getTransportCreds creates TLS or non-TLS credentials based on the url scheme provided
func getTransportCreds(url *url.URL) (credentials.TransportCredentials, error) {
	// Config has forced non-TLS
	if url.Scheme == "http" || url.Scheme == "tcp" {
		return insecure.NewCredentials(), nil
	}

	// Default to TLS
	return credentials.NewTLS(&tls.Config{}), nil
}
