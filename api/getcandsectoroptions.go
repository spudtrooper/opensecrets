// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type GetCandSectorOption func(*getCandSectorOptionImpl)

type GetCandSectorOptions interface {
	Cycle() int
}

func GetCandSectorCycle(cycle int) GetCandSectorOption {
	return func(opts *getCandSectorOptionImpl) {
		opts.cycle = cycle
	}
}
func GetCandSectorCycleFlag(cycle *int) GetCandSectorOption {
	return func(opts *getCandSectorOptionImpl) {
		opts.cycle = *cycle
	}
}

type getCandSectorOptionImpl struct {
	cycle int
}

func (g *getCandSectorOptionImpl) Cycle() int { return g.cycle }

func makeGetCandSectorOptionImpl(opts ...GetCandSectorOption) *getCandSectorOptionImpl {
	res := &getCandSectorOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeGetCandSectorOptions(opts ...GetCandSectorOption) GetCandSectorOptions {
	return makeGetCandSectorOptionImpl(opts...)
}
