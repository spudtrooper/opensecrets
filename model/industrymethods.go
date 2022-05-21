package model

import "bytes"

// IndustryCode returns the string IndustryCode of this Industry
func (i *Industry) IndustryCode() string { return i.info.IndustryCode }

// IndustryName returns the string IndustryName of this Industry
func (i *Industry) IndustryName() string { return i.info.IndustryName }

// Indivs returns the string Indivs of this Industry
func (i *Industry) Indivs() string { return i.info.Indivs }

// Pacs returns the string Pacs of this Industry
func (i *Industry) Pacs() string { return i.info.Pacs }

// Total returns the string Total of this Industry
func (i *Industry) Total() string { return i.info.Total }

func (i *Industry) String() string {
	var buf bytes.Buffer
	buf.WriteString("Industry(")
	buf.WriteString("IndustryCode=")
	buf.WriteString(i.IndustryCode())
	buf.WriteString(", IndustryName=")
	buf.WriteString(i.IndustryName())
	buf.WriteString(", Indivs=")
	buf.WriteString(i.Indivs())
	buf.WriteString(", Pacs=")
	buf.WriteString(i.Pacs())
	buf.WriteString(", Total=")
	buf.WriteString(i.Total())
	buf.WriteString(" )")
	return buf.String()
}
