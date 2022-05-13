package api

import (
	"fmt"
	"testing"

	"github.com/spudtrooper/goutil/io"
)

func TestGetLegislators(t *testing.T) {
	tests := []struct {
		name string
		id   string
	}{
		{
			name: "state",
			id:   "CA",
		},
		{
			name: "person",
			id:   "N00007360",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := tt.id
			client := makeClient(t)
			info, err := client.GetLegislators(id)
			if err != nil {
				t.Fatalf("GetLegislators(%q): %v", id, err)
			}
			found := false
			const firstlast = "Nancy Pelosi"
			for _, li := range info.LegislatorInfos {
				if li.Firstlast == firstlast {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("GetLegislators(%q): did not find expecting firstlast %s", id, firstlast)
			}
		})
	}
}

func TestGetLegislator(t *testing.T) {
	tests := []struct {
		name string
		cid  string
	}{
		{
			name: "person",
			cid:  "N00007360",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cid := tt.cid
			client := makeClient(t)
			li, err := client.GetLegislator(cid)
			if err != nil {
				t.Fatalf("GetLegislator(%q): %v", cid, err)
			}
			if want, got := "Nancy Pelosi", li.Firstlast; want != got {
				t.Errorf("GetLegislators(%q,).Firstlast: got %q, wanted %q", cid, got, want)
			}
		})
	}
}

func TestGetMemPFDprofile(t *testing.T) {
	tests := []struct {
		name string
		cid  string
		year int
	}{
		{
			name: "no year",
			cid:  "N00007360",
		},
		{
			name: "with year",
			cid:  "N00007360",
			year: 2016,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cid, year := tt.cid, tt.year
			client := makeClient(t)
			info, err := client.GetMemPFDprofile(cid, GetMemPFDprofileYear(year))
			if err != nil {
				t.Fatalf("GetMemPFDprofile(%q,%d): %v", cid, year, err)
			}
			if want, got := "Pelosi, Nancy", info.Name; want != got {
				t.Errorf("GetMemPFDprofile(%q,%d).Name: got %q, wanted %q", cid, year, got, want)
			}
			if year != 0 {
				if want, got := fmt.Sprintf("%d", year), info.DataYear; want != got {
					t.Errorf("GetMemPFDprofile(%q,%d).DataYear: got %q, wanted %q", cid, year, got, want)
				}
			}

		})
	}
}

func TestGetCandSummary(t *testing.T) {
	tests := []struct {
		name  string
		cid   string
		cycle int
	}{
		{
			name: "no cycle",
			cid:  "N00007360",
		},
		{
			name:  "with cycle",
			cid:   "N00007360",
			cycle: 2016,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cid, cycle := tt.cid, tt.cycle
			client := makeClient(t)
			info, err := client.GetCandSummary(cid, GetCandSummaryCycle(cycle))
			if err != nil {
				t.Fatalf("GetCandSummary(%q,%d): %v", cid, cycle, err)
			}
			if want, got := "Pelosi, Nancy", info.CandSummaryInfo.CandName; want != got {
				t.Errorf("GetCandSummary(%q,%d).CandName: got %q, wanted %q", cid, cycle, got, want)
			}
			if cycle != 0 {
				if want, got := fmt.Sprintf("%d", cycle), info.CandSummaryInfo.Cycle; want != got {
					t.Errorf("GetCandSummary(%q,%d).Cycle: got %q, wanted %q", cid, cycle, got, want)
				}
			}
		})
	}
}

func TestGetCandContrib(t *testing.T) {
	tests := []struct {
		name  string
		cid   string
		cycle int
	}{
		{
			name: "no cycle",
			cid:  "N00007360",
		},
		{
			name:  "with cycle",
			cid:   "N00007360",
			cycle: 2016,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cid, cycle := tt.cid, tt.cycle
			client := makeClient(t)
			info, err := client.GetCandContrib(cid, GetCandContribCycle(cycle))
			if err != nil {
				t.Fatalf("GetCandContrib(%q,%d): %v", cid, cycle, err)
			}
			if want, got := "Nancy Pelosi (D)", info.Recipient.CandName; want != got {
				t.Errorf("GetCandContrib(%q,%d).CandName: got %q, wanted %q", cid, cycle, got, want)
			}
			if cycle != 0 {
				if want, got := fmt.Sprintf("%d", cycle), info.Recipient.Cycle; want != got {
					t.Errorf("GetCandContrib(%q,%d).Cycle: got %q, wanted %q", cid, cycle, got, want)
				}
			}
		})
	}
}

func TestGetCandIndustry(t *testing.T) {
	tests := []struct {
		name  string
		cid   string
		cycle int
	}{
		{
			name: "no cycle",
			cid:  "N00007360",
		},
		{
			name:  "with cycle",
			cid:   "N00007360",
			cycle: 2016,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cid, cycle := tt.cid, tt.cycle
			client := makeClient(t)
			info, err := client.GetCandIndustry(cid, GetCandIndustryCycle(cycle))
			if err != nil {
				t.Fatalf("GetCandIndustry(%q,%d): %v", cid, cycle, err)
			}
			if want, got := "Nancy Pelosi (D)", info.Recipient.CandName; want != got {
				t.Errorf("GetCandIndustry(%q,%d).CandName: got %q, wanted %q", cid, cycle, got, want)
			}
			if cycle != 0 {
				if want, got := fmt.Sprintf("%d", cycle), info.Recipient.Cycle; want != got {
					t.Errorf("GetCandIndustry(%q,%d).Cycle: got %q, wanted %q", cid, cycle, got, want)
				}
			}
		})
	}
}

func TestGetCandByInd(t *testing.T) {
	tests := []struct {
		name  string
		cid   string
		ind   string
		cycle int
	}{
		{
			name: "no cycle",
			cid:  "N00007360",
			ind:  "M02",
		},
		{
			name:  "with cycle",
			cid:   "N00007360",
			ind:   "M02",
			cycle: 2016,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cid, ind, cycle := tt.cid, tt.ind, tt.cycle
			client := makeClient(t)
			info, err := client.GetCandByInd(cid, ind, GetCandByIndCycle(cycle))
			if err != nil {
				t.Fatalf("GetCandByInd(%q,%d): %v", cid, cycle, err)
			}
			if want, got := "Pelosi, Nancy", info.CandName; want != got {
				t.Errorf("GetCandByInd(%q,%d).CandName: got %q, wanted %q", cid, cycle, got, want)
			}
			if cycle != 0 {
				if want, got := fmt.Sprintf("%d", cycle), info.Cycle; want != got {
					t.Errorf("GetCandByInd(%q,%d).Cycle: got %q, wanted %q", cid, cycle, got, want)
				}
			}
		})
	}
}

func TestGetCandSector(t *testing.T) {
	tests := []struct {
		name  string
		cid   string
		cycle int
	}{
		{
			name: "no cycle",
			cid:  "N00007360",
		},
		{
			name:  "with cycle",
			cid:   "N00007360",
			cycle: 2016,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cid, cycle := tt.cid, tt.cycle
			client := makeClient(t)
			info, err := client.GetCandSector(cid, GetCandSectorCycle(cycle))
			if err != nil {
				t.Fatalf("GetCandSector(%q,%d): %v", cid, cycle, err)
			}
			if want, got := "Nancy Pelosi (D)", info.Candidate.CandName; want != got {
				t.Errorf("GetCandSector(%q,%d).CandName: got %q, wanted %q", cid, cycle, got, want)
			}
			if cycle != 0 {
				if want, got := fmt.Sprintf("%d", cycle), info.Candidate.Cycle; want != got {
					t.Errorf("GetCandSector(%q,%d).Cycle: got %q, wanted %q", cid, cycle, got, want)
				}
			}
		})
	}
}

func makeClient(t *testing.T) *core {
	f := "../.user_creds.json"
	if !io.FileExists(f) {
		t.Fatalf("add your credentials to %s", f)
	}
	client, err := NewClientFromFile(f)
	if err != nil {
		t.Fatalf("NewClientFromFlags: %v", err)
	}
	return client
}
