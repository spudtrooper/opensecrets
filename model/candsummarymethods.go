package model

// CandName returns the string CandName of this CandSummary
func (c *CandSummary) CandName() string {
	return c.info.CandName
}

// CID returns the string CID of this CandSummary
func (c *CandSummary) CID() string {
	return c.info.CID
}

// Cycle returns the string Cycle of this CandSummary
func (c *CandSummary) Cycle() string {
	return c.info.Cycle
}

// State returns the string State of this CandSummary
func (c *CandSummary) State() string {
	return c.info.State
}

// Party returns the string Party of this CandSummary
func (c *CandSummary) Party() string {
	return c.info.Party
}

// Chamber returns the string Chamber of this CandSummary
func (c *CandSummary) Chamber() string {
	return c.info.Chamber
}

// FirstElected returns the string FirstElected of this CandSummary
func (c *CandSummary) FirstElected() string {
	return c.info.FirstElected
}

// NextElection returns the string NextElection of this CandSummary
func (c *CandSummary) NextElection() string {
	return c.info.NextElection
}

// Total returns the string Total of this CandSummary
func (c *CandSummary) Total() string {
	return c.info.Total
}

// Spent returns the string Spent of this CandSummary
func (c *CandSummary) Spent() string {
	return c.info.Spent
}

// CashOnHand returns the string CashOnHand of this CandSummary
func (c *CandSummary) CashOnHand() string {
	return c.info.CashOnHand
}

// Debt returns the string Debt of this CandSummary
func (c *CandSummary) Debt() string {
	return c.info.Debt
}

// Origin returns the string Origin of this CandSummary
func (c *CandSummary) Origin() string {
	return c.info.Origin
}

// Source returns the string Source of this CandSummary
func (c *CandSummary) Source() string {
	return c.info.Source
}

// LastUpdated returns the string LastUpdated of this CandSummary
func (c *CandSummary) LastUpdated() string {
	return c.info.LastUpdated
}
