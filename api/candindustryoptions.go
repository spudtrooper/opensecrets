// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

//go:generate genopts --prefix=CandIndustry --outfile=candindustryoptions.go "cycle:int"

type CandIndustryOption func(*candIndustryOptionImpl)

type CandIndustryOptions interface {
	Cycle() int
}

func CandIndustryCycle(cycle int) CandIndustryOption {
	return func(opts *candIndustryOptionImpl) {
		opts.cycle = cycle
	}
}
func CandIndustryCycleFlag(cycle *int) CandIndustryOption {
	return func(opts *candIndustryOptionImpl) {
		opts.cycle = *cycle
	}
}

type candIndustryOptionImpl struct {
	cycle int
}

func (c *candIndustryOptionImpl) Cycle() int { return c.cycle }

func makeCandIndustryOptionImpl(opts ...CandIndustryOption) *candIndustryOptionImpl {
	res := &candIndustryOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeCandIndustryOptions(opts ...CandIndustryOption) CandIndustryOptions {
	return makeCandIndustryOptionImpl(opts...)
}
