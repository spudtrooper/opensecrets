// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type GetCandContribOption struct {
	f func(*getCandContribOptionImpl)
	s string
}

func (o GetCandContribOption) String() string { return o.s }

type GetCandContribOptions interface {
	Cycle() int
	HasCycle() bool
}

func GetCandContribCycle(cycle int) GetCandContribOption {
	return GetCandContribOption{func(opts *getCandContribOptionImpl) {
		opts.has_cycle = true
		opts.cycle = cycle
	}, fmt.Sprintf("api.GetCandContribCycle(int %+v)}", cycle)}
}
func GetCandContribCycleFlag(cycle *int) GetCandContribOption {
	return GetCandContribOption{func(opts *getCandContribOptionImpl) {
		if cycle == nil {
			return
		}
		opts.has_cycle = true
		opts.cycle = *cycle
	}, fmt.Sprintf("api.GetCandContribCycle(int %+v)}", cycle)}
}

type getCandContribOptionImpl struct {
	cycle     int
	has_cycle bool
}

func (g *getCandContribOptionImpl) Cycle() int     { return g.cycle }
func (g *getCandContribOptionImpl) HasCycle() bool { return g.has_cycle }

type GetCandContribParams struct {
	Cid   string `json:"cid" required:"true"`
	Cycle int    `json:"cycle"`
}

func (o GetCandContribParams) Options() []GetCandContribOption {
	return []GetCandContribOption{
		GetCandContribCycle(o.Cycle),
	}
}

func makeGetCandContribOptionImpl(opts ...GetCandContribOption) *getCandContribOptionImpl {
	res := &getCandContribOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeGetCandContribOptions(opts ...GetCandContribOption) GetCandContribOptions {
	return makeGetCandContribOptionImpl(opts...)
}
