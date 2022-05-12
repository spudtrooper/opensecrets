// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type GetCandIndustryOption func(*getCandIndustryOptionImpl)

type GetCandIndustryOptions interface {
	Cycle() int
}

func GetCandIndustryCycle(cycle int) GetCandIndustryOption {
	return func(opts *getCandIndustryOptionImpl) {
		opts.cycle = cycle
	}
}
func GetCandIndustryCycleFlag(cycle *int) GetCandIndustryOption {
	return func(opts *getCandIndustryOptionImpl) {
		opts.cycle = *cycle
	}
}

type getCandIndustryOptionImpl struct {
	cycle int
}

func (g *getCandIndustryOptionImpl) Cycle() int { return g.cycle }

func makeGetCandIndustryOptionImpl(opts ...GetCandIndustryOption) *getCandIndustryOptionImpl {
	res := &getCandIndustryOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeGetCandIndustryOptions(opts ...GetCandIndustryOption) GetCandIndustryOptions {
	return makeGetCandIndustryOptionImpl(opts...)
}
