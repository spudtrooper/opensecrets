// Don't depend on the json attributes, these are only for grabbing data from the wire.
package api

import (
	"strings"

	"github.com/spudtrooper/goutil/request"
)

type GetLegislatorsInfo struct {
	LegislatorInfos []LegislatorInfo
}

type LegislatorInfo struct {
	CID            string `json:"cid"`
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
	FeccandID      string `json:"feccandid"`
	TwitterID      string `json:"twitter_id"`
	YoutubeURL     string `json:"youtube_url"`
	FacebookID     string `json:"facebook_id"`
	Birthdate      string `json:"birthdate"`
}

func (c *core) GetLegislators(id string) (*GetLegislatorsInfo, error) {
	if len(id) == 2 {
		// this is a state
		return c.getLegislatorsForState(strings.ToUpper(id))
	}
	leg, err := c.getLegislatorForCID(id)
	if err != nil {
		return nil, err
	}
	res := &GetLegislatorsInfo{
		LegislatorInfos: []LegislatorInfo{*leg},
	}
	return res, nil
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

func (c *core) getLegislatorForCID(cid string) (*LegislatorInfo, error) {
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

	res := &result.Response.Legislator.Attributes
	return res, nil
}

type GetLegislatorInfo struct {
	LegislatorInfo
}

func (c *core) GetLegislator(cid string) (*GetLegislatorInfo, error) {
	leg, err := c.getLegislatorForCID(cid)
	if err != nil {
		return nil, err
	}
	res := &GetLegislatorInfo{*leg}
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
	CID          string `json:"cid"`
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

//go:generate genopts --function=GetCandSummary "cycle:int"
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

type SimpleCandidate struct {
	CandName string `json:"cand_name"`
	CID      string `json:"cid"`
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
	Recipient    SimpleCandidate
	Contributors []Contributor
}

//go:generate genopts --function=GetCandContrib "cycle:int"
func (c *core) GetCandContrib(cid string, optss ...GetCandContribOption) (*GetCandContribInfo, error) {
	opts := MakeGetCandContribOptions(optss...)

	type resultT struct {
		Response struct {
			Contributors struct {
				Attributes  SimpleCandidate `json:"@attributes"`
				Contributor []struct {
					Attributes Contributor `json:"@attributes"`
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
		Recipient: result.Response.Contributors.Attributes,
	}
	for _, c := range result.Response.Contributors.Contributor {
		res.Contributors = append(res.Contributors, c.Attributes)
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
	Recipient  SimpleCandidate
	Industries []Industry
}

//go:generate genopts --function=GetCandIndustry "cycle:int"
func (c *core) GetCandIndustry(cid string, optss ...GetCandIndustryOption) (*GetCandIndustryInfo, error) {
	opts := MakeGetCandIndustryOptions(optss...)

	type resultT struct {
		Response struct {
			Industries struct {
				Attributes SimpleCandidate `json:"@attributes"`
				Industry   []struct {
					Attributes Industry `json:"@attributes"`
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
		Recipient: result.Response.Industries.Attributes,
	}
	for _, c := range result.Response.Industries.Industry {
		res.Industries = append(res.Industries, c.Attributes)
	}
	return res, nil
}

type GetCandByIndInfo struct {
	CandByIndInfo
}

type CandByIndInfo struct {
	CandName    string `json:"cand_name"`
	CID         string `json:"cid"`
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

//go:generate genopts --function=GetCandByInd "cycle:int"
func (c *core) GetCandByInd(cid, ind string, optss ...GetCandByIndOption) (*GetCandByIndInfo, error) {
	opts := MakeGetCandByIndOptions(optss...)

	type resultT struct {
		Response struct {
			CandIndus struct {
				Attributes CandByIndInfo `json:"@attributes"`
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

	res := &GetCandByIndInfo{
		CandByIndInfo: result.Response.CandIndus.Attributes,
	}
	return res, nil
}

type Sector struct {
	SectorName string `json:"sector_name"`
	Sectorid   string `json:"sectorid"`
	Indivs     string `json:"indivs"`
	Pacs       string `json:"pacs"`
	Total      string `json:"total"`
}

type GetCandSectorInfo struct {
	Candidate SimpleCandidate
	Sectors   []Sector
}

//go:generate genopts --function=GetCandSector "cycle:int"
func (c *core) GetCandSector(cid string, optss ...GetCandSectorOption) (*GetCandSectorInfo, error) {
	opts := MakeGetCandSectorOptions(optss...)

	type resultT struct {
		Response struct {
			Sectors struct {
				Attributes SimpleCandidate `json:"@attributes"`
				Sectors    []struct {
					Attributes Sector `json:"@attributes"`
				} `json:"sector"`
			} `json:"sectors"`
		} `json:"response"`
	}
	var result resultT

	params := request.MakeParamsBuilder().
		Add("cid", cid).
		AddIfNotDefault("cycle", opts.Cycle()).
		Build()
	err := c.get("candSector", &result, params...)
	if err != nil {
		return nil, err
	}

	res := &GetCandSectorInfo{
		Candidate: result.Response.Sectors.Attributes,
	}
	for _, c := range result.Response.Sectors.Sectors {
		res.Sectors = append(res.Sectors, c.Attributes)
	}
	return res, nil
}

type Committee struct {
	CommitteeName string `json:"committee_name"`
	Industry      string `json:"industry"`
	Congno        string `json:"congno"`
	Origin        string `json:"origin"`
	Source        string `json:"source"`
	LastUpdated   string `json:"last_updated"`
}

type CommitteeMember struct {
	MemberName string `json:"member_name"`
	CID        string `json:"cid"`
	Party      string `json:"party"`
	State      string `json:"state"`
	Total      string `json:"total"`
	Indivs     string `json:"indivs"`
	Pacs       string `json:"pacs"`
}

type GetCongCmteIndusInfo struct {
	Committee
	CommitteeMembers []CommitteeMember
}

//go:generate genopts --function=GetCongCmteIndus "congno:int"
func (c *core) GetCongCmteIndus(cmte, indus string, optss ...GetCongCmteIndusOption) (*GetCongCmteIndusInfo, error) {
	opts := MakeGetCongCmteIndusOptions(optss...)

	type resultT struct {
		Response struct {
			Committee struct {
				Attributes Committee `json:"@attributes"`
				Member     []struct {
					Attributes CommitteeMember `json:"@attributes"`
				} `json:"member"`
			} `json:"committee"`
		} `json:"response"`
	}
	var result resultT

	params := request.MakeParamsBuilder().
		Add("cmte", cmte).
		Add("indus", indus).
		AddIfNotDefault("congno", opts.Congno()).
		Build()
	err := c.get("congCmteIndus", &result, params...)
	if err != nil {
		return nil, err
	}

	res := &GetCongCmteIndusInfo{
		Committee: result.Response.Committee.Attributes,
	}
	for _, c := range result.Response.Committee.Member {
		res.CommitteeMembers = append(res.CommitteeMembers, c.Attributes)
	}
	return res, nil
}

type GetOrgInfo struct {
	Orgs []OrgID
}

type OrgID struct {
	OrgID   string `json:"orgid"`
	Orgname string `json:"orgname"`
}

func (c *core) GetOrgs(org string) (*GetOrgInfo, error) {
	type resultT struct {
		Response struct {
			Organization []struct {
				Attributes OrgID `json:"@attributes"`
			} `json:"organization"`
		} `json:"response"`
	}
	var result resultT

	params := request.MakeParamsBuilder().
		Add("org", org).
		Build()
	err := c.get("getOrgs", &result, params...)
	if err != nil {
		return nil, err
	}

	var res GetOrgInfo
	for _, org := range result.Response.Organization {
		res.Orgs = append(res.Orgs, org.Attributes)
	}
	return &res, nil
}

type GetOrgSummaryInfo struct {
	Org Organization
}

type Organization struct {
	Cycle        string `json:"cycle"`
	Orgid        string `json:"orgid"`
	Orgname      string `json:"orgname"`
	Total        string `json:"total"`
	Indivs       string `json:"indivs"`
	Pacs         string `json:"pacs"`
	Soft         string `json:"soft"`
	Tot527       string `json:"tot527"`
	Dems         string `json:"dems"`
	Repubs       string `json:"repubs"`
	Lobbying     string `json:"lobbying"`
	Outside      string `json:"outside"`
	MemsInvested string `json:"mems_invested"`
	GaveToPac    string `json:"gave_to_pac"`
	GaveToParty  string `json:"gave_to_party"`
	GaveTo527    string `json:"gave_to_527"`
	GaveToCand   string `json:"gave_to_cand"`
	Source       string `json:"source"`
}

func (c *core) GetOrgSummary(orgID string) (*GetOrgSummaryInfo, error) {
	type resultT struct {
		Response struct {
			Organization struct {
				Attributes Organization `json:"@attributes"`
			} `json:"organization"`
		} `json:"response"`
	}
	var result resultT

	params := request.MakeParamsBuilder().
		Add("id", orgID).
		Build()
	err := c.get("orgSummary", &result, params...)
	if err != nil {
		return nil, err
	}

	res := &GetOrgSummaryInfo{
		Org: result.Response.Organization.Attributes,
	}
	return res, nil
}

type GetIndependentExpendInfo struct {
	IndependentExpendInfos []IndependentExpendInfo
}

type IndependentExpendInfo struct {
	CmteID   string `json:"cmteid"`
	Pacshort string `json:"pacshort"`
	Suppopp  string `json:"suppopp"`
	Candname string `json:"candname"`
	District string `json:"district"`
	Amount   string `json:"amount"`
	Note     string `json:"note"`
	Party    string `json:"party"`
	Payee    string `json:"payee"`
	Date     string `json:"date"`
	Origin   string `json:"origin"`
	Source   string `json:"source"`
}

func (c *core) GetIndependentExpend() (*GetIndependentExpendInfo, error) {
	type resultT struct {
		Response struct {
			Indexp []struct {
				Attributes IndependentExpendInfo `json:"@attributes"`
			} `json:"indexp"`
		} `json:"response"`
	}
	var result resultT

	err := c.get("independentExpend", &result)
	if err != nil {
		return nil, err
	}

	var res GetIndependentExpendInfo
	for _, ie := range result.Response.Indexp {
		res.IndependentExpendInfos = append(res.IndependentExpendInfos, ie.Attributes)
	}
	return &res, nil
}
