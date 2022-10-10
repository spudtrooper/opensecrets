// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type GetCandSectorOption func(*getCandSectorOptionImpl)

type GetCandSectorOptions interface {
	Cycle() int
	HasCycle() bool
}

func GetCandSectorCycle(cycle int) GetCandSectorOption {
	return func(opts *getCandSectorOptionImpl) {
		opts.has_cycle = true
		opts.cycle = cycle
	}
}
func GetCandSectorCycleFlag(cycle *int) GetCandSectorOption {
	return func(opts *getCandSectorOptionImpl) {
		if cycle == nil {
			return
		}
		opts.has_cycle = true
		opts.cycle = *cycle
	}
}

type getCandSectorOptionImpl struct {
	cycle     int
	has_cycle bool
}

func (g *getCandSectorOptionImpl) Cycle() int     { return g.cycle }
func (g *getCandSectorOptionImpl) HasCycle() bool { return g.has_cycle }

type GetCandSectorParams struct {
	Cid   string `json:"cid" required:"true"`
	Cycle int    `json:"cycle"`
}

func (o GetCandSectorParams) Options() []GetCandSectorOption {
	return []GetCandSectorOption{
		GetCandSectorCycle(o.Cycle),
	}
}

func makeGetCandSectorOptionImpl(opts ...GetCandSectorOption) *getCandSectorOptionImpl {
	res := &getCandSectorOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeGetCandSectorOptions(opts ...GetCandSectorOption) GetCandSectorOptions {
	return makeGetCandSectorOptionImpl(opts...)
}
