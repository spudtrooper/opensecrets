// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type GetMemPFDprofileOption func(*getMemPFDprofileOptionImpl)

type GetMemPFDprofileOptions interface {
	Year() int
	HasYear() bool
}

func GetMemPFDprofileYear(year int) GetMemPFDprofileOption {
	return func(opts *getMemPFDprofileOptionImpl) {
		opts.has_year = true
		opts.year = year
	}
}
func GetMemPFDprofileYearFlag(year *int) GetMemPFDprofileOption {
	return func(opts *getMemPFDprofileOptionImpl) {
		if year == nil {
			return
		}
		opts.has_year = true
		opts.year = *year
	}
}

type getMemPFDprofileOptionImpl struct {
	year     int
	has_year bool
}

func (g *getMemPFDprofileOptionImpl) Year() int     { return g.year }
func (g *getMemPFDprofileOptionImpl) HasYear() bool { return g.has_year }

type GetMemPFDprofileParams struct {
	Cid  string `json:"cid" required:"true"`
	Year int    `json:"year"`
}

func (o GetMemPFDprofileParams) Options() []GetMemPFDprofileOption {
	return []GetMemPFDprofileOption{
		GetMemPFDprofileYear(o.Year),
	}
}

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
