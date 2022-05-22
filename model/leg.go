package model

import (
	"github.com/spudtrooper/opensecrets/api"
)

type Legislator struct {
	base
	factory       *Factory
	cid           CID
	info          *api.LegislatorInfo
	memPFDprofile *MemPFDprofile
	candSummary   *CandSummary
	contributors  []*Contributor
	industries    []*Industry
	sectors       []*Sector
}

func (l *Legislator) getInfo() (*api.LegislatorInfo, error) {
	if l.info == nil {
		info, err := l.client.GetLegislator(string(l.cid))
		if err != nil {
			return nil, err
		}
		l.info = &info.LegislatorInfo
	}
	return l.info, nil
}

func (l *Legislator) MemPFDprofile() (*MemPFDprofile, error) {
	if l.memPFDprofile == nil {
		info, err := l.client.GetMemPFDprofile(string(l.cid))
		if err != nil {
			return nil, err
		}
		l.memPFDprofile = &MemPFDprofile{
			info: &info.MemPFDprofileInfo,
		}
	}
	return l.memPFDprofile, nil
}

func (l *Legislator) CandSummary() (*CandSummary, error) {
	if l.candSummary == nil {
		info, err := l.client.GetCandSummary(string(l.cid))
		if err != nil {
			return nil, err
		}
		l.candSummary = &CandSummary{
			info: &info.CandSummaryInfo,
		}
	}
	return l.candSummary, nil
}

func (l *Legislator) Contributors() ([]*Contributor, error) {
	if l.contributors == nil {
		info, err := l.client.GetCandContrib(string(l.cid))
		if err != nil {
			return nil, err
		}
		contributors := []*Contributor{}
		for _, info := range info.Contributors {
			contributors = append(contributors, &Contributor{info})
		}
		l.contributors = contributors
	}
	return l.contributors, nil
}

func (l *Legislator) Industries() ([]*Industry, error) {
	if l.industries == nil {
		info, err := l.client.GetCandIndustry(string(l.cid))
		if err != nil {
			return nil, err
		}
		industries := []*Industry{}
		for _, info := range info.Industries {
			industries = append(industries, &Industry{info})
		}
		l.industries = industries
	}
	return l.industries, nil
}

func (l *Legislator) Sectors() ([]*Sector, error) {
	if l.industries == nil {
		info, err := l.client.GetCandSector(string(l.cid))
		if err != nil {
			return nil, err
		}
		sectors := []*Sector{}
		for _, info := range info.Sectors {
			sectors = append(sectors, &Sector{info})
		}
		l.sectors = sectors
	}
	return l.sectors, nil
}
