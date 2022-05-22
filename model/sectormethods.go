package model

import "bytes"

// SectorName returns the string SectorName of this Sector
func (s *Sector) SectorName() string { return s.info.SectorName }

// Sectorid returns the string Sectorid of this Sector
func (s *Sector) Sectorid() string { return s.info.Sectorid }

// Indivs returns the string Indivs of this Sector
func (s *Sector) Indivs() string { return s.info.Indivs }

// Pacs returns the string Pacs of this Sector
func (s *Sector) Pacs() string { return s.info.Pacs }

// Total returns the string Total of this Sector
func (s *Sector) Total() string { return s.info.Total }

func (s *Sector) String() string {
	var buf bytes.Buffer
	buf.WriteString("Sector(")
	buf.WriteString("SectorName=")
	buf.WriteString(s.SectorName())
	buf.WriteString(", Sectorid=")
	buf.WriteString(s.Sectorid())
	buf.WriteString(", Indivs=")
	buf.WriteString(s.Indivs())
	buf.WriteString(", Pacs=")
	buf.WriteString(s.Pacs())
	buf.WriteString(", Total=")
	buf.WriteString(s.Total())
	buf.WriteString(")")
	return buf.String()
}
