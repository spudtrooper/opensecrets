// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type GetLegislatorOption struct {
	f func(*getLegislatorOptionImpl)
	s string
}

func (o GetLegislatorOption) String() string { return o.s }

type GetLegislatorOptions interface {
}

type getLegislatorOptionImpl struct {
}

type GetLegislatorParams struct {
	Cid string `json:"cid" required:"true"`
}

func (o GetLegislatorParams) Options() []GetLegislatorOption {
	return []GetLegislatorOption{}
}

func makeGetLegislatorOptionImpl(opts ...GetLegislatorOption) *getLegislatorOptionImpl {
	res := &getLegislatorOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeGetLegislatorOptions(opts ...GetLegislatorOption) GetLegislatorOptions {
	return makeGetLegislatorOptionImpl(opts...)
}
