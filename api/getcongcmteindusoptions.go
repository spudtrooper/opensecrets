// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type GetCongCmteIndusOption func(*getCongCmteIndusOptionImpl)

type GetCongCmteIndusOptions interface {
	Congno() int
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

type getCongCmteIndusOptionImpl struct {
	congno int
}

func (g *getCongCmteIndusOptionImpl) Congno() int { return g.congno }

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
