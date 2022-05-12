// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

//go:generate genopts --prefix=CandSummary --outfile=candsummaryoptions.go "cycle:int"

type CandSummaryOption func(*candSummaryOptionImpl)

type CandSummaryOptions interface {
	Cycle() int
}

func CandSummaryCycle(cycle int) CandSummaryOption {
	return func(opts *candSummaryOptionImpl) {
		opts.cycle = cycle
	}
}
func CandSummaryCycleFlag(cycle *int) CandSummaryOption {
	return func(opts *candSummaryOptionImpl) {
		opts.cycle = *cycle
	}
}

type candSummaryOptionImpl struct {
	cycle int
}

func (c *candSummaryOptionImpl) Cycle() int { return c.cycle }

func makeCandSummaryOptionImpl(opts ...CandSummaryOption) *candSummaryOptionImpl {
	res := &candSummaryOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeCandSummaryOptions(opts ...CandSummaryOption) CandSummaryOptions {
	return makeCandSummaryOptionImpl(opts...)
}
