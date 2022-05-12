// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type GetCandByIndOption func(*getCandByIndOptionImpl)

type GetCandByIndOptions interface {
	Cycle() int
}

func GetCandByIndCycle(cycle int) GetCandByIndOption {
	return func(opts *getCandByIndOptionImpl) {
		opts.cycle = cycle
	}
}
func GetCandByIndCycleFlag(cycle *int) GetCandByIndOption {
	return func(opts *getCandByIndOptionImpl) {
		opts.cycle = *cycle
	}
}

type getCandByIndOptionImpl struct {
	cycle int
}

func (g *getCandByIndOptionImpl) Cycle() int { return g.cycle }

func makeGetCandByIndOptionImpl(opts ...GetCandByIndOption) *getCandByIndOptionImpl {
	res := &getCandByIndOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeGetCandByIndOptions(opts ...GetCandByIndOption) GetCandByIndOptions {
	return makeGetCandByIndOptionImpl(opts...)
}
