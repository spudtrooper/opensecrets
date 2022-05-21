package model

import "bytes"

// OrgName returns the string OrgName of this Contributor
func (c *Contributor) OrgName() string { return c.info.OrgName }

// Total returns the string Total of this Contributor
func (c *Contributor) Total() string { return c.info.Total }

// Pacs returns the string Pacs of this Contributor
func (c *Contributor) Pacs() string { return c.info.Pacs }

// Indivs returns the string Indivs of this Contributor
func (c *Contributor) Indivs() string { return c.info.Indivs }

func (c *Contributor) String() string {
	var buf bytes.Buffer
	buf.WriteString("Contributor(")
	buf.WriteString("OrgName=")
	buf.WriteString(c.OrgName())
	buf.WriteString(", Total=")
	buf.WriteString(c.Total())
	buf.WriteString(", Pacs=")
	buf.WriteString(c.Pacs())
	buf.WriteString(", Indivs=")
	buf.WriteString(c.Indivs())
	buf.WriteString(" )")
	return buf.String()
}
