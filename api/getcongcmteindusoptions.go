// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type GetCongCmteIndusOption struct {
	f func(*getCongCmteIndusOptionImpl)
	s string
}

func (o GetCongCmteIndusOption) String() string { return o.s }

type GetCongCmteIndusOptions interface {
	Congno() int
	HasCongno() bool
}

func GetCongCmteIndusCongno(congno int) GetCongCmteIndusOption {
	return GetCongCmteIndusOption{func(opts *getCongCmteIndusOptionImpl) {
		opts.has_congno = true
		opts.congno = congno
	}, fmt.Sprintf("api.GetCongCmteIndusCongno(int %+v)}", congno)}
}
func GetCongCmteIndusCongnoFlag(congno *int) GetCongCmteIndusOption {
	return GetCongCmteIndusOption{func(opts *getCongCmteIndusOptionImpl) {
		if congno == nil {
			return
		}
		opts.has_congno = true
		opts.congno = *congno
	}, fmt.Sprintf("api.GetCongCmteIndusCongno(int %+v)}", congno)}
}

type getCongCmteIndusOptionImpl struct {
	congno     int
	has_congno bool
}

func (g *getCongCmteIndusOptionImpl) Congno() int     { return g.congno }
func (g *getCongCmteIndusOptionImpl) HasCongno() bool { return g.has_congno }

type GetCongCmteIndusParams struct {
	Cmte   string `json:"cmte" required:"true"`
	Congno int    `json:"congno"`
	Ind    string `json:"ind" required:"true"`
}

func (o GetCongCmteIndusParams) Options() []GetCongCmteIndusOption {
	return []GetCongCmteIndusOption{
		GetCongCmteIndusCongno(o.Congno),
	}
}

func makeGetCongCmteIndusOptionImpl(opts ...GetCongCmteIndusOption) *getCongCmteIndusOptionImpl {
	res := &getCongCmteIndusOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeGetCongCmteIndusOptions(opts ...GetCongCmteIndusOption) GetCongCmteIndusOptions {
	return makeGetCongCmteIndusOptionImpl(opts...)
}
