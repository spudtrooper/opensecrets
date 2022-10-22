// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type GetCandByIndOption struct {
	f func(*getCandByIndOptionImpl)
	s string
}

func (o GetCandByIndOption) String() string { return o.s }

type GetCandByIndOptions interface {
	Cycle() int
	HasCycle() bool
}

func GetCandByIndCycle(cycle int) GetCandByIndOption {
	return GetCandByIndOption{func(opts *getCandByIndOptionImpl) {
		opts.has_cycle = true
		opts.cycle = cycle
	}, fmt.Sprintf("api.GetCandByIndCycle(int %+v)", cycle)}
}
func GetCandByIndCycleFlag(cycle *int) GetCandByIndOption {
	return GetCandByIndOption{func(opts *getCandByIndOptionImpl) {
		if cycle == nil {
			return
		}
		opts.has_cycle = true
		opts.cycle = *cycle
	}, fmt.Sprintf("api.GetCandByIndCycle(int %+v)", cycle)}
}

type getCandByIndOptionImpl struct {
	cycle     int
	has_cycle bool
}

func (g *getCandByIndOptionImpl) Cycle() int     { return g.cycle }
func (g *getCandByIndOptionImpl) HasCycle() bool { return g.has_cycle }

type GetCandByIndParams struct {
	Cid   string `json:"cid" required:"true"`
	Cycle int    `json:"cycle"`
	Ind   string `json:"ind" required:"true"`
}

func (o GetCandByIndParams) Options() []GetCandByIndOption {
	return []GetCandByIndOption{
		GetCandByIndCycle(o.Cycle),
	}
}

func makeGetCandByIndOptionImpl(opts ...GetCandByIndOption) *getCandByIndOptionImpl {
	res := &getCandByIndOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeGetCandByIndOptions(opts ...GetCandByIndOption) GetCandByIndOptions {
	return makeGetCandByIndOptionImpl(opts...)
}
