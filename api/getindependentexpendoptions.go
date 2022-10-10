// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type GetIndependentExpendOption func(*getIndependentExpendOptionImpl)

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
		opt(res)
	}
	return res
}

func MakeGetIndependentExpendOptions(opts ...GetIndependentExpendOption) GetIndependentExpendOptions {
	return makeGetIndependentExpendOptionImpl(opts...)
}
