// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type GetCandContribOption func(*getCandContribOptionImpl)

type GetCandContribOptions interface {
	Cycle() int
}

func GetCandContribCycle(cycle int) GetCandContribOption {
	return func(opts *getCandContribOptionImpl) {
		opts.cycle = cycle
	}
}
func GetCandContribCycleFlag(cycle *int) GetCandContribOption {
	return func(opts *getCandContribOptionImpl) {
		opts.cycle = *cycle
	}
}

type getCandContribOptionImpl struct {
	cycle int
}

func (g *getCandContribOptionImpl) Cycle() int { return g.cycle }

func makeGetCandContribOptionImpl(opts ...GetCandContribOption) *getCandContribOptionImpl {
	res := &getCandContribOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeGetCandContribOptions(opts ...GetCandContribOption) GetCandContribOptions {
	return makeGetCandContribOptionImpl(opts...)
}
