// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type GetCandByIndOption func(*getCandByIndOptionImpl)

type GetCandByIndOptions interface {
	Cycle() int
	HasCycle() bool
}

func GetCandByIndCycle(cycle int) GetCandByIndOption {
	return func(opts *getCandByIndOptionImpl) {
		opts.has_cycle = true
		opts.cycle = cycle
	}
}
func GetCandByIndCycleFlag(cycle *int) GetCandByIndOption {
	return func(opts *getCandByIndOptionImpl) {
		if cycle == nil {
			return
		}
		opts.has_cycle = true
		opts.cycle = *cycle
	}
}

type getCandByIndOptionImpl struct {
	cycle     int
	has_cycle bool
}

func (g *getCandByIndOptionImpl) Cycle() int     { return g.cycle }
func (g *getCandByIndOptionImpl) HasCycle() bool { return g.has_cycle }

type GetCandByIndParams struct {
	Cid   string `json:"cid" required:"true"`
	Ind   string `json:"ind" required:"true"`
	Cycle int    `json:"cycle"`
}

func (o GetCandByIndParams) Options() []GetCandByIndOption {
	return []GetCandByIndOption{
		GetCandByIndCycle(o.Cycle),
	}
}

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
