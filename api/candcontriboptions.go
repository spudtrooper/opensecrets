// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

//go:generate genopts --prefix=CandContrib --outfile=candcontriboptions.go "cycle:int"

type CandContribOption func(*candContribOptionImpl)

type CandContribOptions interface {
	Cycle() int
}

func CandContribCycle(cycle int) CandContribOption {
	return func(opts *candContribOptionImpl) {
		opts.cycle = cycle
	}
}
func CandContribCycleFlag(cycle *int) CandContribOption {
	return func(opts *candContribOptionImpl) {
		opts.cycle = *cycle
	}
}

type candContribOptionImpl struct {
	cycle int
}

func (c *candContribOptionImpl) Cycle() int { return c.cycle }

func makeCandContribOptionImpl(opts ...CandContribOption) *candContribOptionImpl {
	res := &candContribOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeCandContribOptions(opts ...CandContribOption) CandContribOptions {
	return makeCandContribOptionImpl(opts...)
}
