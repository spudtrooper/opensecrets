package model

import "github.com/spudtrooper/opensecrets/api"

type Legislator struct {
	base
	cid           CID
	info          *api.LegislatorInfo
	memPFDprofile *MemPFDprofile
	candSummary   *CandSummary
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
