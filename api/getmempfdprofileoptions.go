// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type GetMemPFDprofileOption func(*getMemPFDprofileOptionImpl)

type GetMemPFDprofileOptions interface {
	Year() int
}

func GetMemPFDprofileYear(year int) GetMemPFDprofileOption {
	return func(opts *getMemPFDprofileOptionImpl) {
		opts.year = year
	}
}
func GetMemPFDprofileYearFlag(year *int) GetMemPFDprofileOption {
	return func(opts *getMemPFDprofileOptionImpl) {
		opts.year = *year
	}
}

type getMemPFDprofileOptionImpl struct {
	year int
}

func (g *getMemPFDprofileOptionImpl) Year() int { return g.year }

func makeGetMemPFDprofileOptionImpl(opts ...GetMemPFDprofileOption) *getMemPFDprofileOptionImpl {
	res := &getMemPFDprofileOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeGetMemPFDprofileOptions(opts ...GetMemPFDprofileOption) GetMemPFDprofileOptions {
	return makeGetMemPFDprofileOptionImpl(opts...)
}
