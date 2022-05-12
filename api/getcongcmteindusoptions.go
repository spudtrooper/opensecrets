// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type GetCongCmteIndusOption func(*getCongCmteIndusOptionImpl)

type GetCongCmteIndusOptions interface {
	Cmte() string
	Congno() int
	Indus() string
}

func GetCongCmteIndusCmte(cmte string) GetCongCmteIndusOption {
	return func(opts *getCongCmteIndusOptionImpl) {
		opts.cmte = cmte
	}
}
func GetCongCmteIndusCmteFlag(cmte *string) GetCongCmteIndusOption {
	return func(opts *getCongCmteIndusOptionImpl) {
		opts.cmte = *cmte
	}
}

func GetCongCmteIndusCongno(congno int) GetCongCmteIndusOption {
	return func(opts *getCongCmteIndusOptionImpl) {
		opts.congno = congno
	}
}
func GetCongCmteIndusCongnoFlag(congno *int) GetCongCmteIndusOption {
	return func(opts *getCongCmteIndusOptionImpl) {
		opts.congno = *congno
	}
}

func GetCongCmteIndusIndus(indus string) GetCongCmteIndusOption {
	return func(opts *getCongCmteIndusOptionImpl) {
		opts.indus = indus
	}
}
func GetCongCmteIndusIndusFlag(indus *string) GetCongCmteIndusOption {
	return func(opts *getCongCmteIndusOptionImpl) {
		opts.indus = *indus
	}
}

type getCongCmteIndusOptionImpl struct {
	cmte   string
	congno int
	indus  string
}

func (g *getCongCmteIndusOptionImpl) Cmte() string  { return g.cmte }
func (g *getCongCmteIndusOptionImpl) Congno() int   { return g.congno }
func (g *getCongCmteIndusOptionImpl) Indus() string { return g.indus }

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
