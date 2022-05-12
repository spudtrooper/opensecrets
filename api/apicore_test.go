package api

import (
	"fmt"
	"testing"

	"github.com/spudtrooper/goutil/io"
)

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

func TestGetLegislators(t *testing.T) {
	runTest := func(id string) {
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
	}

	runTest("CA")
	runTest("N00007360")
}

func TestGetMemPFDprofile(t *testing.T) {
	runTest := func(cid string, year int) {
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
	}

	runTest("N00007360", 0)
	runTest("N00007360", 2016)
}

func TestGetCandSummary(t *testing.T) {
	runTest := func(cid string, cycle int) {
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
	}

	runTest("N00007360", 0)
	runTest("N00007360", 2016)
}

func TestGetCandContrib(t *testing.T) {
	runTest := func(cid string, cycle int) {
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
	}

	runTest("N00007360", 0)
	runTest("N00007360", 2016)
}
