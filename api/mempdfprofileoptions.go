// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

//go:generate genopts --prefix=MemPFDprofile --outfile=mempdfprofileoptions.go "year:int"

type MemPFDprofileOption func(*memPFDprofileOptionImpl)

type MemPFDprofileOptions interface {
	Year() int
}

func MemPFDprofileYear(year int) MemPFDprofileOption {
	return func(opts *memPFDprofileOptionImpl) {
		opts.year = year
	}
}
func MemPFDprofileYearFlag(year *int) MemPFDprofileOption {
	return func(opts *memPFDprofileOptionImpl) {
		opts.year = *year
	}
}

type memPFDprofileOptionImpl struct {
	year int
}

func (m *memPFDprofileOptionImpl) Year() int { return m.year }

func makeMemPFDprofileOptionImpl(opts ...MemPFDprofileOption) *memPFDprofileOptionImpl {
	res := &memPFDprofileOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeMemPFDprofileOptions(opts ...MemPFDprofileOption) MemPFDprofileOptions {
	return makeMemPFDprofileOptionImpl(opts...)
}
