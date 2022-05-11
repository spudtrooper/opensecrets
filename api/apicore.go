package api

import (
	"github.com/spudtrooper/goutil/request"
)

type GetLegislatorsInfo struct {
	LegislatorInfos []LegislatorInfo
}

type LegislatorInfo struct {
	Cid            string `json:"cid"`
	Firstlast      string `json:"firstlast"`
	Lastname       string `json:"lastname"`
	Party          string `json:"party"`
	Office         string `json:"office"`
	Gender         string `json:"gender"`
	FirstElected   string `json:"first_elected"`
	ExitCode       string `json:"exit_code"`
	Comments       string `json:"comments"`
	Phone          string `json:"phone"`
	Fax            string `json:"fax"`
	Website        string `json:"website"`
	Webform        string `json:"webform"`
	CongressOffice string `json:"congress_office"`
	BioguideID     string `json:"bioguide_id"`
	VotesmartID    string `json:"votesmart_id"`
	Feccandid      string `json:"feccandid"`
	TwitterID      string `json:"twitter_id"`
	YoutubeURL     string `json:"youtube_url"`
	FacebookID     string `json:"facebook_id"`
	Birthdate      string `json:"birthdate"`
}

func (c *core) GetLegislators(id string) (*GetLegislatorsInfo, error) {
	if len(id) == 2 {
		// this is a state
		return c.getLegislatorsForState(id)
	}
	return c.getLegislatorsForCID(id)
}

func (c *core) getLegislatorsForState(stateID string) (*GetLegislatorsInfo, error) {
	type resultT struct {
		Response struct {
			Legislator []struct {
				Attributes LegislatorInfo `json:"@attributes"`
			} `json:"legislator"`
		} `json:"response"`
	}
	var result resultT

	err := c.get("getLegislators", &result, request.MakeParam("id", stateID))
	if err != nil {
		return nil, err
	}

	var res GetLegislatorsInfo
	for _, o := range result.Response.Legislator {
		res.LegislatorInfos = append(res.LegislatorInfos, o.Attributes)
	}
	return &res, nil
}

func (c *core) getLegislatorsForCID(cid string) (*GetLegislatorsInfo, error) {
	type resultT struct {
		Response struct {
			Legislator struct {
				Attributes LegislatorInfo `json:"@attributes"`
			} `json:"legislator"`
		} `json:"response"`
	}
	var result resultT

	err := c.get("getLegislators", &result, request.MakeParam("id", cid))
	if err != nil {
		return nil, err
	}

	res := &GetLegislatorsInfo{
		LegislatorInfos: []LegislatorInfo{
			result.Response.Legislator.Attributes,
		},
	}
	return res, nil
}

type GetMemPFDprofileInfo struct {
	MemPFDprofileInfo
}

type MemPFDprofileInfo struct {
	Name               string `json:"name"`
	DataYear           string `json:"data_year"`
	MemberID           string `json:"member_id"`
	NetLow             string `json:"net_low"`
	NetHigh            string `json:"net_high"`
	PositionsHeldCount string `json:"positions_held_count"`
	AssetCount         string `json:"asset_count"`
	AssetLow           string `json:"asset_low"`
	AssetHigh          string `json:"asset_high"`
	TransactionCount   string `json:"transaction_count"`
	TxLow              string `json:"tx_low"`
	TxHigh             string `json:"tx_high"`
	Source             string `json:"source"`
	Origin             string `json:"origin"`
	UpdateTimestamp    string `json:"update_timestamp"`
}

func (c *core) GetMemPFDprofile(cid string, optss ...MemPFDprofileOption) (*GetMemPFDprofileInfo, error) {
	opts := MakeMemPFDprofileOptions(optss...)

	type resultT struct {
		Response struct {
			MemberProfile struct {
				Attributes MemPFDprofileInfo `json:"@attributes"`
			} `json:"member_profile"`
		} `json:"response"`
	}
	var result resultT

	params := []request.Param{
		request.MakeParam("cid", cid),
	}
	if opts.Year() > 0 {
		params = append(params, request.MakeParam("year", opts.Year()))
	}
	err := c.get("memPFDprofile", &result, params...)
	if err != nil {
		return nil, err
	}

	res := &GetMemPFDprofileInfo{
		MemPFDprofileInfo: result.Response.MemberProfile.Attributes,
	}
	return res, nil
}
