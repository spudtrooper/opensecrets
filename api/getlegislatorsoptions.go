// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type GetLegislatorsOption func(*getLegislatorsOptionImpl)

type GetLegislatorsOptions interface {
}

type getLegislatorsOptionImpl struct {
}

type GetLegislatorsParams struct {
	Id string `json:"id" required:"true"`
}

func (o GetLegislatorsParams) Options() []GetLegislatorsOption {
	return []GetLegislatorsOption{}
}

func makeGetLegislatorsOptionImpl(opts ...GetLegislatorsOption) *getLegislatorsOptionImpl {
	res := &getLegislatorsOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeGetLegislatorsOptions(opts ...GetLegislatorsOption) GetLegislatorsOptions {
	return makeGetLegislatorsOptionImpl(opts...)
}
