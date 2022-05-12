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

//go:generate genopts --function=GetMemPFDprofile "year:int"
func (c *core) GetMemPFDprofile(cid string, optss ...GetMemPFDprofileOption) (*GetMemPFDprofileInfo, error) {
	opts := MakeGetMemPFDprofileOptions(optss...)

	type resultT struct {
		Response struct {
			MemberProfile struct {
				Attributes MemPFDprofileInfo `json:"@attributes"`
			} `json:"member_profile"`
		} `json:"response"`
	}
	var result resultT

	params := request.MakeParamsBuilder().
		Add("cid", cid).
		AddIfNotDefault("year", opts.Year()).
		Build()
	err := c.get("memPFDprofile", &result, params...)
	if err != nil {
		return nil, err
	}

	res := &GetMemPFDprofileInfo{
		MemPFDprofileInfo: result.Response.MemberProfile.Attributes,
	}
	return res, nil
}

type GetCandSummaryInfo struct {
	CandSummaryInfo CandSummaryInfo
}

type CandSummaryInfo struct {
	CandName     string `json:"cand_name"`
	Cid          string `json:"cid"`
	Cycle        string `json:"cycle"`
	State        string `json:"state"`
	Party        string `json:"party"`
	Chamber      string `json:"chamber"`
	FirstElected string `json:"first_elected"`
	NextElection string `json:"next_election"`
	Total        string `json:"total"`
	Spent        string `json:"spent"`
	CashOnHand   string `json:"cash_on_hand"`
	Debt         string `json:"debt"`
	Origin       string `json:"origin"`
	Source       string `json:"source"`
	LastUpdated  string `json:"last_updated"`
}

//go:generate genopts --prefix=GetCandSummary "cycle:int"
func (c *core) GetCandSummary(cid string, optss ...GetCandSummaryOption) (*GetCandSummaryInfo, error) {
	opts := MakeGetCandSummaryOptions(optss...)

	type resultT struct {
		Response struct {
			Summary struct {
				Attributes CandSummaryInfo `json:"@attributes"`
			} `json:"summary"`
		} `json:"response"`
	}
	var result resultT

	params := request.MakeParamsBuilder().
		Add("cid", cid).
		AddIfNotDefault("cycle", opts.Cycle()).
		Build()
	err := c.get("candSummary", &result, params...)
	if err != nil {
		return nil, err
	}

	res := &GetCandSummaryInfo{
		CandSummaryInfo: result.Response.Summary.Attributes,
	}
	return res, nil
}

type Recipient struct {
	CandName string `json:"cand_name"`
	Cid      string `json:"cid"`
	Cycle    string `json:"cycle"`
	Origin   string `json:"origin"`
	Source   string `json:"source"`
	Notice   string `json:"notice"`
}

type Contributor struct {
	OrgName string `json:"org_name"`
	Total   string `json:"total"`
	Pacs    string `json:"pacs"`
	Indivs  string `json:"indivs"`
}

type GetCandContribInfo struct {
	Recipient    Recipient
	Contributors []Contributor
}

//go:generate genopts --function=GetCandContrib "cycle:int"
func (c *core) GetCandContrib(cid string, optss ...GetCandContribOption) (*GetCandContribInfo, error) {
	opts := MakeGetCandContribOptions(optss...)

	type resultT struct {
		Response struct {
			Contributors struct {
				Recipient   Recipient `json:"@attributes"`
				Contributor []struct {
					Contributor Contributor `json:"@attributes"`
				} `json:"contributor"`
			} `json:"contributors"`
		} `json:"response"`
	}
	var result resultT

	params := request.MakeParamsBuilder().
		Add("cid", cid).
		AddIfNotDefault("cycle", opts.Cycle()).
		Build()
	err := c.get("candContrib", &result, params...)
	if err != nil {
		return nil, err
	}

	res := &GetCandContribInfo{
		Recipient: result.Response.Contributors.Recipient,
	}
	for _, c := range result.Response.Contributors.Contributor {
		res.Contributors = append(res.Contributors, c.Contributor)
	}
	return res, nil
}

type Industry struct {
	IndustryCode string `json:"industry_code"`
	IndustryName string `json:"industry_name"`
	Indivs       string `json:"indivs"`
	Pacs         string `json:"pacs"`
	Total        string `json:"total"`
}

type GetCandIndustryInfo struct {
	Recipient  Recipient
	Industries []Industry
}

//go:generate genopts --function=GetCandIndustry "cycle:int"
func (c *core) GetCandIndustry(cid string, optss ...GetCandIndustryOption) (*GetCandIndustryInfo, error) {
	opts := MakeGetCandIndustryOptions(optss...)

	type resultT struct {
		Response struct {
			Industries struct {
				Recipient  Recipient `json:"@attributes"`
				Industries []struct {
					Industry Industry `json:"@attributes"`
				} `json:"industry"`
			} `json:"industries"`
		} `json:"response"`
	}
	var result resultT

	params := request.MakeParamsBuilder().
		Add("cid", cid).
		AddIfNotDefault("cycle", opts.Cycle()).
		Build()
	err := c.get("candIndustry", &result, params...)
	if err != nil {
		return nil, err
	}

	res := &GetCandIndustryInfo{
		Recipient: result.Response.Industries.Recipient,
	}
	for _, c := range result.Response.Industries.Industries {
		res.Industries = append(res.Industries, c.Industry)
	}
	return res, nil
}

type GetCandByIndustryInfo struct {
	CandByIndustryInfo
}

type CandByIndustryInfo struct {
	CandName    string `json:"cand_name"`
	Cid         string `json:"cid"`
	Cycle       string `json:"cycle"`
	Industry    string `json:"industry"`
	Chamber     string `json:"chamber"`
	Party       string `json:"party"`
	State       string `json:"state"`
	Total       string `json:"total"`
	Indivs      string `json:"indivs"`
	Pacs        string `json:"pacs"`
	Rank        string `json:"rank"`
	Origin      string `json:"origin"`
	Source      string `json:"source"`
	LastUpdated string `json:"last_updated"`
}

//go:generate genopts --function=GetCandByIndustry "cycle:int"
func (c *core) GetCandByIndustry(cid, ind string, optss ...GetCandByIndustryOption) (*GetCandByIndustryInfo, error) {
	opts := MakeGetCandByIndustryOptions(optss...)

	type resultT struct {
		Response struct {
			CandIndus struct {
				Attributes CandByIndustryInfo `json:"@attributes"`
			} `json:"candIndus"`
		} `json:"response"`
	}
	var result resultT

	params := request.MakeParamsBuilder().
		Add("cid", cid).
		Add("ind", ind).
		AddIfNotDefault("cycle", opts.Cycle()).
		Build()
	err := c.get("candIndByInd", &result, params...)
	if err != nil {
		return nil, err
	}

	res := &GetCandByIndustryInfo{
		CandByIndustryInfo: result.Response.CandIndus.Attributes,
	}
	return res, nil
}
