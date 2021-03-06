package model

import (
	"strings"
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
	var names []string
	found := false
	for _, c := range contribs {
		if c.OrgName() == orgName {
			found = true
			break
		}
		names = append(names, c.OrgName())
	}
	if !found {
		t.Errorf("did not find expecting OrgName: %s, found: %s", orgName, strings.Join(names, ","))
	}
}

func TestIndustries(t *testing.T) {
	f := newFactoryForTesting(t)
	leg := f.NewLegislator("N00007360")
	inds, err := leg.Industries()
	if err != nil {
		t.Fatalf("Industries(): unexpected error: %v", err)
	}
	if len(inds) == 0 {
		t.Errorf("expecting non-empty industries")
	}
	const industryName = "Health Professionals"
	var names []string
	found := false
	for _, ind := range inds {
		if ind.IndustryName() == industryName {
			found = true
			break
		}
		names = append(names, ind.IndustryName())
	}
	if !found {
		t.Errorf("did not find expecting IndustryName: %s, found: %s", industryName, strings.Join(names, ","))
	}
}

func TestSector(t *testing.T) {
	f := newFactoryForTesting(t)
	leg := f.NewLegislator("N00007360")
	secs, err := leg.Sectors()
	if err != nil {
		t.Fatalf("Sectors(): unexpected error: %v", err)
	}
	if len(secs) == 0 {
		t.Errorf("expecting non-empty sectors")
	}
	const sectorName = "Agribusiness"
	var names []string
	found := false
	for _, s := range secs {
		if s.SectorName() == sectorName {
			found = true
			break
		}
		names = append(names, s.SectorName())
	}
	if !found {
		t.Errorf("did not find expecting SectorName: %s, found: %s", sectorName, strings.Join(names, ","))
	}
}
