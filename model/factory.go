package model

import "github.com/spudtrooper/opensecrets/api"

type base struct {
	client *api.Core
}

type CID string
type State string

type Factory struct {
	client *api.Core
}

func NewFactory(client *api.Core) *Factory {
	return &Factory{client}
}

func (f *Factory) NewLegislator(cid CID) *Legislator {
	return &Legislator{
		base: base{f.client},
		cid:  cid,
	}
}

func (f *Factory) GetLegislators(state State) ([]*Legislator, error) {
	info, err := f.client.GetLegislators(string(state))
	if err != nil {
		return nil, err
	}
	var res []*Legislator
	for _, leg := range info.LegislatorInfos {
		leg := leg
		res = append(res, &Legislator{
			base: base{f.client},
			cid:  CID(leg.CID),
			info: &leg,
		})
	}
	return res, nil
}
