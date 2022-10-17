// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type GetIndependentExpendOption struct {
	f func(*getIndependentExpendOptionImpl)
	s string
}

func (o GetIndependentExpendOption) String() string { return o.s }

type GetIndependentExpendOptions interface {
}

type getIndependentExpendOptionImpl struct {
}

type GetIndependentExpendParams struct {
}

func (o GetIndependentExpendParams) Options() []GetIndependentExpendOption {
	return []GetIndependentExpendOption{}
}

func makeGetIndependentExpendOptionImpl(opts ...GetIndependentExpendOption) *getIndependentExpendOptionImpl {
	res := &getIndependentExpendOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeGetIndependentExpendOptions(opts ...GetIndependentExpendOption) GetIndependentExpendOptions {
	return makeGetIndependentExpendOptionImpl(opts...)
}
