// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type GetCongCmteIndusOption func(*getCongCmteIndusOptionImpl)

type GetCongCmteIndusOptions interface {
	Congno() int
	HasCongno() bool
}

func GetCongCmteIndusCongno(congno int) GetCongCmteIndusOption {
	return func(opts *getCongCmteIndusOptionImpl) {
		opts.has_congno = true
		opts.congno = congno
	}
}
func GetCongCmteIndusCongnoFlag(congno *int) GetCongCmteIndusOption {
	return func(opts *getCongCmteIndusOptionImpl) {
		if congno == nil {
			return
		}
		opts.has_congno = true
		opts.congno = *congno
	}
}

type getCongCmteIndusOptionImpl struct {
	congno     int
	has_congno bool
}

func (g *getCongCmteIndusOptionImpl) Congno() int     { return g.congno }
func (g *getCongCmteIndusOptionImpl) HasCongno() bool { return g.has_congno }

type GetCongCmteIndusParams struct {
	Cmte   string `json:"cmte" required:"true"`
	Ind    string `json:"ind" required:"true"`
	Congno int    `json:"congno"`
}

func (o GetCongCmteIndusParams) Options() []GetCongCmteIndusOption {
	return []GetCongCmteIndusOption{
		GetCongCmteIndusCongno(o.Congno),
	}
}

func makeGetCongCmteIndusOptionImpl(opts ...GetCongCmteIndusOption) *getCongCmteIndusOptionImpl {
	res := &getCongCmteIndusOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeGetCongCmteIndusOptions(opts ...GetCongCmteIndusOption) GetCongCmteIndusOptions {
	return makeGetCongCmteIndusOptionImpl(opts...)
}
