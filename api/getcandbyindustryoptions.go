// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type GetCandByIndustryOption func(*getCandByIndustryOptionImpl)

type GetCandByIndustryOptions interface {
	Cycle() int
}

func GetCandByIndustryCycle(cycle int) GetCandByIndustryOption {
	return func(opts *getCandByIndustryOptionImpl) {
		opts.cycle = cycle
	}
}
func GetCandByIndustryCycleFlag(cycle *int) GetCandByIndustryOption {
	return func(opts *getCandByIndustryOptionImpl) {
		opts.cycle = *cycle
	}
}

type getCandByIndustryOptionImpl struct {
	cycle int
}

func (g *getCandByIndustryOptionImpl) Cycle() int { return g.cycle }

func makeGetCandByIndustryOptionImpl(opts ...GetCandByIndustryOption) *getCandByIndustryOptionImpl {
	res := &getCandByIndustryOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeGetCandByIndustryOptions(opts ...GetCandByIndustryOption) GetCandByIndustryOptions {
	return makeGetCandByIndustryOptionImpl(opts...)
}
