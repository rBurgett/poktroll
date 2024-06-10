package types

const (
	DefaultComputeUnitsPerRelay uint64 = 1
)

// NewService creates a new Service instance
func NewService(id string, name string, computeUnitsPerRelay uint64) *Service {
	return &Service{
		Id:                   id,
		Name:                 name,
		ComputeUnitsPerRelay: computeUnitsPerRelay,
	}
}
