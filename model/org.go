package model

import "github.com/spudtrooper/opensecrets/api"

type Organization struct {
	base
	id   api.OrgID
	info *api.Organization
}

func (o *Organization) getInfo() (*api.Organization, error) {
	if o.info == nil {
		info, err := o.client.GetOrgSummary(o.id.OrgID)
		if err != nil {
			return nil, err
		}
		o.info = &info.Org
	}
	return o.info, nil
}
