package model

import "bytes"

// CandName returns the string CandName of this CandSummary
func (c *CandSummary) CandName() string { return c.info.CandName }

// CID returns the string CID of this CandSummary
func (c *CandSummary) CID() string { return c.info.CID }

// Cycle returns the string Cycle of this CandSummary
func (c *CandSummary) Cycle() string { return c.info.Cycle }

// State returns the string State of this CandSummary
func (c *CandSummary) State() string { return c.info.State }

// Party returns the string Party of this CandSummary
func (c *CandSummary) Party() string { return c.info.Party }

// Chamber returns the string Chamber of this CandSummary
func (c *CandSummary) Chamber() string { return c.info.Chamber }

// FirstElected returns the string FirstElected of this CandSummary
func (c *CandSummary) FirstElected() string { return c.info.FirstElected }

// NextElection returns the string NextElection of this CandSummary
func (c *CandSummary) NextElection() string { return c.info.NextElection }

// Total returns the string Total of this CandSummary
func (c *CandSummary) Total() string { return c.info.Total }

// Spent returns the string Spent of this CandSummary
func (c *CandSummary) Spent() string { return c.info.Spent }

// CashOnHand returns the string CashOnHand of this CandSummary
func (c *CandSummary) CashOnHand() string { return c.info.CashOnHand }

// Debt returns the string Debt of this CandSummary
func (c *CandSummary) Debt() string { return c.info.Debt }

// Origin returns the string Origin of this CandSummary
func (c *CandSummary) Origin() string { return c.info.Origin }

// Source returns the string Source of this CandSummary
func (c *CandSummary) Source() string { return c.info.Source }

// LastUpdated returns the string LastUpdated of this CandSummary
func (c *CandSummary) LastUpdated() string { return c.info.LastUpdated }

func (c *CandSummary) String() string {
	var buf bytes.Buffer
	buf.WriteString("CandSummary(")
	buf.WriteString("CandName=")
	buf.WriteString(c.CandName())
	buf.WriteString(", CID=")
	buf.WriteString(c.CID())
	buf.WriteString(", Cycle=")
	buf.WriteString(c.Cycle())
	buf.WriteString(", State=")
	buf.WriteString(c.State())
	buf.WriteString(", Party=")
	buf.WriteString(c.Party())
	buf.WriteString(", Chamber=")
	buf.WriteString(c.Chamber())
	buf.WriteString(", FirstElected=")
	buf.WriteString(c.FirstElected())
	buf.WriteString(", NextElection=")
	buf.WriteString(c.NextElection())
	buf.WriteString(", Total=")
	buf.WriteString(c.Total())
	buf.WriteString(", Spent=")
	buf.WriteString(c.Spent())
	buf.WriteString(", CashOnHand=")
	buf.WriteString(c.CashOnHand())
	buf.WriteString(", Debt=")
	buf.WriteString(c.Debt())
	buf.WriteString(", Origin=")
	buf.WriteString(c.Origin())
	buf.WriteString(", Source=")
	buf.WriteString(c.Source())
	buf.WriteString(", LastUpdated=")
	buf.WriteString(c.LastUpdated())
	buf.WriteString(" )")
	return buf.String()
}
