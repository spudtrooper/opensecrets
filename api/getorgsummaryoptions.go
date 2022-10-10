// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type GetOrgSummaryOption func(*getOrgSummaryOptionImpl)

type GetOrgSummaryOptions interface {
	Congno() int
	HasCongno() bool
}

func GetOrgSummaryCongno(congno int) GetOrgSummaryOption {
	return func(opts *getOrgSummaryOptionImpl) {
		opts.has_congno = true
		opts.congno = congno
	}
}
func GetOrgSummaryCongnoFlag(congno *int) GetOrgSummaryOption {
	return func(opts *getOrgSummaryOptionImpl) {
		if congno == nil {
			return
		}
		opts.has_congno = true
		opts.congno = *congno
	}
}

type getOrgSummaryOptionImpl struct {
	congno     int
	has_congno bool
}

func (g *getOrgSummaryOptionImpl) Congno() int     { return g.congno }
func (g *getOrgSummaryOptionImpl) HasCongno() bool { return g.has_congno }

type GetOrgSummaryParams struct {
	OrgID  string `json:"org_id" required:"true"`
	Congno int    `json:"congno"`
}

func (o GetOrgSummaryParams) Options() []GetOrgSummaryOption {
	return []GetOrgSummaryOption{
		GetOrgSummaryCongno(o.Congno),
	}
}

func makeGetOrgSummaryOptionImpl(opts ...GetOrgSummaryOption) *getOrgSummaryOptionImpl {
	res := &getOrgSummaryOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeGetOrgSummaryOptions(opts ...GetOrgSummaryOption) GetOrgSummaryOptions {
	return makeGetOrgSummaryOptionImpl(opts...)
}
