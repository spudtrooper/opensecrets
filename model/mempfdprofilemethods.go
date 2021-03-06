package model

import "bytes"

// Name returns the string Name of this MemPFDprofile
func (m *MemPFDprofile) Name() string { return m.info.Name }

// DataYear returns the string DataYear of this MemPFDprofile
func (m *MemPFDprofile) DataYear() string { return m.info.DataYear }

// MemberID returns the string MemberID of this MemPFDprofile
func (m *MemPFDprofile) MemberID() string { return m.info.MemberID }

// NetLow returns the string NetLow of this MemPFDprofile
func (m *MemPFDprofile) NetLow() string { return m.info.NetLow }

// NetHigh returns the string NetHigh of this MemPFDprofile
func (m *MemPFDprofile) NetHigh() string { return m.info.NetHigh }

// PositionsHeldCount returns the string PositionsHeldCount of this MemPFDprofile
func (m *MemPFDprofile) PositionsHeldCount() string { return m.info.PositionsHeldCount }

// AssetCount returns the string AssetCount of this MemPFDprofile
func (m *MemPFDprofile) AssetCount() string { return m.info.AssetCount }

// AssetLow returns the string AssetLow of this MemPFDprofile
func (m *MemPFDprofile) AssetLow() string { return m.info.AssetLow }

// AssetHigh returns the string AssetHigh of this MemPFDprofile
func (m *MemPFDprofile) AssetHigh() string { return m.info.AssetHigh }

// TransactionCount returns the string TransactionCount of this MemPFDprofile
func (m *MemPFDprofile) TransactionCount() string { return m.info.TransactionCount }

// TxLow returns the string TxLow of this MemPFDprofile
func (m *MemPFDprofile) TxLow() string { return m.info.TxLow }

// TxHigh returns the string TxHigh of this MemPFDprofile
func (m *MemPFDprofile) TxHigh() string { return m.info.TxHigh }

// Source returns the string Source of this MemPFDprofile
func (m *MemPFDprofile) Source() string { return m.info.Source }

// Origin returns the string Origin of this MemPFDprofile
func (m *MemPFDprofile) Origin() string { return m.info.Origin }

// UpdateTimestamp returns the string UpdateTimestamp of this MemPFDprofile
func (m *MemPFDprofile) UpdateTimestamp() string { return m.info.UpdateTimestamp }

func (m *MemPFDprofile) String() string {
	var buf bytes.Buffer
	buf.WriteString("MemPFDprofile(")
	buf.WriteString("Name=")
	buf.WriteString(m.Name())
	buf.WriteString(", DataYear=")
	buf.WriteString(m.DataYear())
	buf.WriteString(", MemberID=")
	buf.WriteString(m.MemberID())
	buf.WriteString(", NetLow=")
	buf.WriteString(m.NetLow())
	buf.WriteString(", NetHigh=")
	buf.WriteString(m.NetHigh())
	buf.WriteString(", PositionsHeldCount=")
	buf.WriteString(m.PositionsHeldCount())
	buf.WriteString(", AssetCount=")
	buf.WriteString(m.AssetCount())
	buf.WriteString(", AssetLow=")
	buf.WriteString(m.AssetLow())
	buf.WriteString(", AssetHigh=")
	buf.WriteString(m.AssetHigh())
	buf.WriteString(", TransactionCount=")
	buf.WriteString(m.TransactionCount())
	buf.WriteString(", TxLow=")
	buf.WriteString(m.TxLow())
	buf.WriteString(", TxHigh=")
	buf.WriteString(m.TxHigh())
	buf.WriteString(", Source=")
	buf.WriteString(m.Source())
	buf.WriteString(", Origin=")
	buf.WriteString(m.Origin())
	buf.WriteString(", UpdateTimestamp=")
	buf.WriteString(m.UpdateTimestamp())
	buf.WriteString(")")
	return buf.String()
}
