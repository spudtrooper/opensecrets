package model

import (
	"bytes"
	"fmt"
)

// Cycle returns the string Cycle of this Organization
func (o *Organization) Cycle() (string, error) {
	info, err := o.getInfo()
	if err != nil {
		return "", err
	}
	return info.Cycle, nil
}

// Cycle returns Cycle() or panics if there is an error
func (o *Organization) CycleOrPanic() string {
	res, err := o.Cycle()
	if err != nil {
		panic(fmt.Sprintf("error calling Cycle(): %v", err))
	}
	return res
}

// Orgid returns the string Orgid of this Organization
func (o *Organization) Orgid() (string, error) {
	info, err := o.getInfo()
	if err != nil {
		return "", err
	}
	return info.Orgid, nil
}

// Orgid returns Orgid() or panics if there is an error
func (o *Organization) OrgidOrPanic() string {
	res, err := o.Orgid()
	if err != nil {
		panic(fmt.Sprintf("error calling Orgid(): %v", err))
	}
	return res
}

// Orgname returns the string Orgname of this Organization
func (o *Organization) Orgname() (string, error) {
	info, err := o.getInfo()
	if err != nil {
		return "", err
	}
	return info.Orgname, nil
}

// Orgname returns Orgname() or panics if there is an error
func (o *Organization) OrgnameOrPanic() string {
	res, err := o.Orgname()
	if err != nil {
		panic(fmt.Sprintf("error calling Orgname(): %v", err))
	}
	return res
}

// Total returns the string Total of this Organization
func (o *Organization) Total() (string, error) {
	info, err := o.getInfo()
	if err != nil {
		return "", err
	}
	return info.Total, nil
}

// Total returns Total() or panics if there is an error
func (o *Organization) TotalOrPanic() string {
	res, err := o.Total()
	if err != nil {
		panic(fmt.Sprintf("error calling Total(): %v", err))
	}
	return res
}

// Indivs returns the string Indivs of this Organization
func (o *Organization) Indivs() (string, error) {
	info, err := o.getInfo()
	if err != nil {
		return "", err
	}
	return info.Indivs, nil
}

// Indivs returns Indivs() or panics if there is an error
func (o *Organization) IndivsOrPanic() string {
	res, err := o.Indivs()
	if err != nil {
		panic(fmt.Sprintf("error calling Indivs(): %v", err))
	}
	return res
}

// Pacs returns the string Pacs of this Organization
func (o *Organization) Pacs() (string, error) {
	info, err := o.getInfo()
	if err != nil {
		return "", err
	}
	return info.Pacs, nil
}

// Pacs returns Pacs() or panics if there is an error
func (o *Organization) PacsOrPanic() string {
	res, err := o.Pacs()
	if err != nil {
		panic(fmt.Sprintf("error calling Pacs(): %v", err))
	}
	return res
}

// Soft returns the string Soft of this Organization
func (o *Organization) Soft() (string, error) {
	info, err := o.getInfo()
	if err != nil {
		return "", err
	}
	return info.Soft, nil
}

// Soft returns Soft() or panics if there is an error
func (o *Organization) SoftOrPanic() string {
	res, err := o.Soft()
	if err != nil {
		panic(fmt.Sprintf("error calling Soft(): %v", err))
	}
	return res
}

// Tot527 returns the string Tot527 of this Organization
func (o *Organization) Tot527() (string, error) {
	info, err := o.getInfo()
	if err != nil {
		return "", err
	}
	return info.Tot527, nil
}

// Tot527 returns Tot527() or panics if there is an error
func (o *Organization) Tot527OrPanic() string {
	res, err := o.Tot527()
	if err != nil {
		panic(fmt.Sprintf("error calling Tot527(): %v", err))
	}
	return res
}

// Dems returns the string Dems of this Organization
func (o *Organization) Dems() (string, error) {
	info, err := o.getInfo()
	if err != nil {
		return "", err
	}
	return info.Dems, nil
}

// Dems returns Dems() or panics if there is an error
func (o *Organization) DemsOrPanic() string {
	res, err := o.Dems()
	if err != nil {
		panic(fmt.Sprintf("error calling Dems(): %v", err))
	}
	return res
}

// Repubs returns the string Repubs of this Organization
func (o *Organization) Repubs() (string, error) {
	info, err := o.getInfo()
	if err != nil {
		return "", err
	}
	return info.Repubs, nil
}

// Repubs returns Repubs() or panics if there is an error
func (o *Organization) RepubsOrPanic() string {
	res, err := o.Repubs()
	if err != nil {
		panic(fmt.Sprintf("error calling Repubs(): %v", err))
	}
	return res
}

// Lobbying returns the string Lobbying of this Organization
func (o *Organization) Lobbying() (string, error) {
	info, err := o.getInfo()
	if err != nil {
		return "", err
	}
	return info.Lobbying, nil
}

// Lobbying returns Lobbying() or panics if there is an error
func (o *Organization) LobbyingOrPanic() string {
	res, err := o.Lobbying()
	if err != nil {
		panic(fmt.Sprintf("error calling Lobbying(): %v", err))
	}
	return res
}

// Outside returns the string Outside of this Organization
func (o *Organization) Outside() (string, error) {
	info, err := o.getInfo()
	if err != nil {
		return "", err
	}
	return info.Outside, nil
}

// Outside returns Outside() or panics if there is an error
func (o *Organization) OutsideOrPanic() string {
	res, err := o.Outside()
	if err != nil {
		panic(fmt.Sprintf("error calling Outside(): %v", err))
	}
	return res
}

// MemsInvested returns the string MemsInvested of this Organization
func (o *Organization) MemsInvested() (string, error) {
	info, err := o.getInfo()
	if err != nil {
		return "", err
	}
	return info.MemsInvested, nil
}

// MemsInvested returns MemsInvested() or panics if there is an error
func (o *Organization) MemsInvestedOrPanic() string {
	res, err := o.MemsInvested()
	if err != nil {
		panic(fmt.Sprintf("error calling MemsInvested(): %v", err))
	}
	return res
}

// GaveToPac returns the string GaveToPac of this Organization
func (o *Organization) GaveToPac() (string, error) {
	info, err := o.getInfo()
	if err != nil {
		return "", err
	}
	return info.GaveToPac, nil
}

// GaveToPac returns GaveToPac() or panics if there is an error
func (o *Organization) GaveToPacOrPanic() string {
	res, err := o.GaveToPac()
	if err != nil {
		panic(fmt.Sprintf("error calling GaveToPac(): %v", err))
	}
	return res
}

// GaveToParty returns the string GaveToParty of this Organization
func (o *Organization) GaveToParty() (string, error) {
	info, err := o.getInfo()
	if err != nil {
		return "", err
	}
	return info.GaveToParty, nil
}

// GaveToParty returns GaveToParty() or panics if there is an error
func (o *Organization) GaveToPartyOrPanic() string {
	res, err := o.GaveToParty()
	if err != nil {
		panic(fmt.Sprintf("error calling GaveToParty(): %v", err))
	}
	return res
}

// GaveTo527 returns the string GaveTo527 of this Organization
func (o *Organization) GaveTo527() (string, error) {
	info, err := o.getInfo()
	if err != nil {
		return "", err
	}
	return info.GaveTo527, nil
}

// GaveTo527 returns GaveTo527() or panics if there is an error
func (o *Organization) GaveTo527OrPanic() string {
	res, err := o.GaveTo527()
	if err != nil {
		panic(fmt.Sprintf("error calling GaveTo527(): %v", err))
	}
	return res
}

// GaveToCand returns the string GaveToCand of this Organization
func (o *Organization) GaveToCand() (string, error) {
	info, err := o.getInfo()
	if err != nil {
		return "", err
	}
	return info.GaveToCand, nil
}

// GaveToCand returns GaveToCand() or panics if there is an error
func (o *Organization) GaveToCandOrPanic() string {
	res, err := o.GaveToCand()
	if err != nil {
		panic(fmt.Sprintf("error calling GaveToCand(): %v", err))
	}
	return res
}

// Source returns the string Source of this Organization
func (o *Organization) Source() (string, error) {
	info, err := o.getInfo()
	if err != nil {
		return "", err
	}
	return info.Source, nil
}

// Source returns Source() or panics if there is an error
func (o *Organization) SourceOrPanic() string {
	res, err := o.Source()
	if err != nil {
		panic(fmt.Sprintf("error calling Source(): %v", err))
	}
	return res
}

func (o *Organization) String() string {
	var buf bytes.Buffer
	buf.WriteString("Organization(")
	buf.WriteString("Cycle=")
	buf.WriteString(o.CycleOrPanic())
	buf.WriteString(", Orgid=")
	buf.WriteString(o.OrgidOrPanic())
	buf.WriteString(", Orgname=")
	buf.WriteString(o.OrgnameOrPanic())
	buf.WriteString(", Total=")
	buf.WriteString(o.TotalOrPanic())
	buf.WriteString(", Indivs=")
	buf.WriteString(o.IndivsOrPanic())
	buf.WriteString(", Pacs=")
	buf.WriteString(o.PacsOrPanic())
	buf.WriteString(", Soft=")
	buf.WriteString(o.SoftOrPanic())
	buf.WriteString(", Tot527=")
	buf.WriteString(o.Tot527OrPanic())
	buf.WriteString(", Dems=")
	buf.WriteString(o.DemsOrPanic())
	buf.WriteString(", Repubs=")
	buf.WriteString(o.RepubsOrPanic())
	buf.WriteString(", Lobbying=")
	buf.WriteString(o.LobbyingOrPanic())
	buf.WriteString(", Outside=")
	buf.WriteString(o.OutsideOrPanic())
	buf.WriteString(", MemsInvested=")
	buf.WriteString(o.MemsInvestedOrPanic())
	buf.WriteString(", GaveToPac=")
	buf.WriteString(o.GaveToPacOrPanic())
	buf.WriteString(", GaveToParty=")
	buf.WriteString(o.GaveToPartyOrPanic())
	buf.WriteString(", GaveTo527=")
	buf.WriteString(o.GaveTo527OrPanic())
	buf.WriteString(", GaveToCand=")
	buf.WriteString(o.GaveToCandOrPanic())
	buf.WriteString(", Source=")
	buf.WriteString(o.SourceOrPanic())
	buf.WriteString(")")
	return buf.String()
}
