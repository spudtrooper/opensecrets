package model

import (
	"testing"
)

func TestLegislator(t *testing.T) {
	f := newFactoryForTesting(t)
	leg := f.NewLegislator("N00007360")
	firstlast, err := leg.Firstlast()
	if err != nil {
		t.Fatalf("unexpected error calling Firstlast: %v", err)
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
		t.Fatalf("unexpected error calling MemPFDprofile: %v", err)
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
		t.Fatalf("unexpected error calling TestCandSummary: %v", err)
	}
	if want, got := "Pelosi, Nancy", p.CandName(); want != got {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
