// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type GetCandSummaryOption func(*getCandSummaryOptionImpl)

type GetCandSummaryOptions interface {
	Cycle() int
}

func GetCandSummaryCycle(cycle int) GetCandSummaryOption {
	return func(opts *getCandSummaryOptionImpl) {
		opts.cycle = cycle
	}
}
func GetCandSummaryCycleFlag(cycle *int) GetCandSummaryOption {
	return func(opts *getCandSummaryOptionImpl) {
		opts.cycle = *cycle
	}
}

type getCandSummaryOptionImpl struct {
	cycle int
}

func (g *getCandSummaryOptionImpl) Cycle() int { return g.cycle }

func makeGetCandSummaryOptionImpl(opts ...GetCandSummaryOption) *getCandSummaryOptionImpl {
	res := &getCandSummaryOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeGetCandSummaryOptions(opts ...GetCandSummaryOption) GetCandSummaryOptions {
	return makeGetCandSummaryOptionImpl(opts...)
}
