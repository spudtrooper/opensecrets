// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type GetOrgsOption struct {
	f func(*getOrgsOptionImpl)
	s string
}

func (o GetOrgsOption) String() string { return o.s }

type GetOrgsOptions interface {
	Congno() int
	HasCongno() bool
}

func GetOrgsCongno(congno int) GetOrgsOption {
	return GetOrgsOption{func(opts *getOrgsOptionImpl) {
		opts.has_congno = true
		opts.congno = congno
	}, fmt.Sprintf("api.GetOrgsCongno(int %+v)", congno)}
}
func GetOrgsCongnoFlag(congno *int) GetOrgsOption {
	return GetOrgsOption{func(opts *getOrgsOptionImpl) {
		if congno == nil {
			return
		}
		opts.has_congno = true
		opts.congno = *congno
	}, fmt.Sprintf("api.GetOrgsCongno(int %+v)", congno)}
}

type getOrgsOptionImpl struct {
	congno     int
	has_congno bool
}

func (g *getOrgsOptionImpl) Congno() int     { return g.congno }
func (g *getOrgsOptionImpl) HasCongno() bool { return g.has_congno }

type GetOrgsParams struct {
	Congno int    `json:"congno"`
	Org    string `json:"org" required:"true"`
}

func (o GetOrgsParams) Options() []GetOrgsOption {
	return []GetOrgsOption{
		GetOrgsCongno(o.Congno),
	}
}

func makeGetOrgsOptionImpl(opts ...GetOrgsOption) *getOrgsOptionImpl {
	res := &getOrgsOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeGetOrgsOptions(opts ...GetOrgsOption) GetOrgsOptions {
	return makeGetOrgsOptionImpl(opts...)
}
