// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type GetCandSummaryOption struct {
	f func(*getCandSummaryOptionImpl)
	s string
}

func (o GetCandSummaryOption) String() string { return o.s }

type GetCandSummaryOptions interface {
	Cycle() int
	HasCycle() bool
}

func GetCandSummaryCycle(cycle int) GetCandSummaryOption {
	return GetCandSummaryOption{func(opts *getCandSummaryOptionImpl) {
		opts.has_cycle = true
		opts.cycle = cycle
	}, fmt.Sprintf("api.GetCandSummaryCycle(int %+v)}", cycle)}
}
func GetCandSummaryCycleFlag(cycle *int) GetCandSummaryOption {
	return GetCandSummaryOption{func(opts *getCandSummaryOptionImpl) {
		if cycle == nil {
			return
		}
		opts.has_cycle = true
		opts.cycle = *cycle
	}, fmt.Sprintf("api.GetCandSummaryCycle(int %+v)}", cycle)}
}

type getCandSummaryOptionImpl struct {
	cycle     int
	has_cycle bool
}

func (g *getCandSummaryOptionImpl) Cycle() int     { return g.cycle }
func (g *getCandSummaryOptionImpl) HasCycle() bool { return g.has_cycle }

type GetCandSummaryParams struct {
	Cid   string `json:"cid" required:"true"`
	Cycle int    `json:"cycle"`
}

func (o GetCandSummaryParams) Options() []GetCandSummaryOption {
	return []GetCandSummaryOption{
		GetCandSummaryCycle(o.Cycle),
	}
}

func makeGetCandSummaryOptionImpl(opts ...GetCandSummaryOption) *getCandSummaryOptionImpl {
	res := &getCandSummaryOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeGetCandSummaryOptions(opts ...GetCandSummaryOption) GetCandSummaryOptions {
	return makeGetCandSummaryOptionImpl(opts...)
}
