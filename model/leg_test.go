package model

import (
	"testing"
)

func TestLegislator(t *testing.T) {
	f := newFactoryForTesting(t)
	leg := f.NewLegislator("N00007360")
	firstlast, err := leg.Firstlast()
	if err != nil {
		t.Fatalf("Firstlast(): unexpected error: %v", err)
	}
	if want, got := "Nancy Pelosi", firstlast; want != got {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestMemPFDprofile(t *testing.T) {
	f := newFactoryForTesting(t)
	leg := f.NewLegislator("N00007360")
	p, err := leg.MemPFDprofile()
	if err != nil {
		t.Fatalf("MemPFDprofile(): unexpected error: %v", err)
	}
	if want, got := "Pelosi, Nancy", p.Name(); want != got {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestCandSummary(t *testing.T) {
	f := newFactoryForTesting(t)
	leg := f.NewLegislator("N00007360")
	p, err := leg.CandSummary()
	if err != nil {
		t.Fatalf("CandSummary(): unexpected error: %v", err)
	}
	if want, got := "Pelosi, Nancy", p.CandName(); want != got {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestContributors(t *testing.T) {
	f := newFactoryForTesting(t)
	leg := f.NewLegislator("N00007360")
	contribs, err := leg.Contributors()
	if err != nil {
		t.Fatalf("Contributors(): unexpected error: %v", err)
	}
	if len(contribs) == 0 {
		t.Errorf("expecting non-empty contributors")
	}
	const orgName = "Facebook Inc"
	found := false
	for _, c := range contribs {
		if c.OrgName() == orgName {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("did not find expecting OrgName: %s", orgName)
	}
}
