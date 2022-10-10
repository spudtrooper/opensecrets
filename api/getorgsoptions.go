// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type GetOrgsOption func(*getOrgsOptionImpl)

type GetOrgsOptions interface {
	Congno() int
	HasCongno() bool
}

func GetOrgsCongno(congno int) GetOrgsOption {
	return func(opts *getOrgsOptionImpl) {
		opts.has_congno = true
		opts.congno = congno
	}
}
func GetOrgsCongnoFlag(congno *int) GetOrgsOption {
	return func(opts *getOrgsOptionImpl) {
		if congno == nil {
			return
		}
		opts.has_congno = true
		opts.congno = *congno
	}
}

type getOrgsOptionImpl struct {
	congno     int
	has_congno bool
}

func (g *getOrgsOptionImpl) Congno() int     { return g.congno }
func (g *getOrgsOptionImpl) HasCongno() bool { return g.has_congno }

type GetOrgsParams struct {
	Org    string `json:"org" required:"true"`
	Congno int    `json:"congno"`
}

func (o GetOrgsParams) Options() []GetOrgsOption {
	return []GetOrgsOption{
		GetOrgsCongno(o.Congno),
	}
}

func makeGetOrgsOptionImpl(opts ...GetOrgsOption) *getOrgsOptionImpl {
	res := &getOrgsOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeGetOrgsOptions(opts ...GetOrgsOption) GetOrgsOptions {
	return makeGetOrgsOptionImpl(opts...)
}
