// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type GetCandIndustryOption struct {
	f func(*getCandIndustryOptionImpl)
	s string
}

func (o GetCandIndustryOption) String() string { return o.s }

type GetCandIndustryOptions interface {
	Cycle() int
	HasCycle() bool
}

func GetCandIndustryCycle(cycle int) GetCandIndustryOption {
	return GetCandIndustryOption{func(opts *getCandIndustryOptionImpl) {
		opts.has_cycle = true
		opts.cycle = cycle
	}, fmt.Sprintf("api.GetCandIndustryCycle(int %+v)", cycle)}
}
func GetCandIndustryCycleFlag(cycle *int) GetCandIndustryOption {
	return GetCandIndustryOption{func(opts *getCandIndustryOptionImpl) {
		if cycle == nil {
			return
		}
		opts.has_cycle = true
		opts.cycle = *cycle
	}, fmt.Sprintf("api.GetCandIndustryCycle(int %+v)", cycle)}
}

type getCandIndustryOptionImpl struct {
	cycle     int
	has_cycle bool
}

func (g *getCandIndustryOptionImpl) Cycle() int     { return g.cycle }
func (g *getCandIndustryOptionImpl) HasCycle() bool { return g.has_cycle }

type GetCandIndustryParams struct {
	Cid   string `json:"cid" required:"true"`
	Cycle int    `json:"cycle"`
}

func (o GetCandIndustryParams) Options() []GetCandIndustryOption {
	return []GetCandIndustryOption{
		GetCandIndustryCycle(o.Cycle),
	}
}

func makeGetCandIndustryOptionImpl(opts ...GetCandIndustryOption) *getCandIndustryOptionImpl {
	res := &getCandIndustryOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeGetCandIndustryOptions(opts ...GetCandIndustryOption) GetCandIndustryOptions {
	return makeGetCandIndustryOptionImpl(opts...)
}
