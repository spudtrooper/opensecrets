package model

import (
	"testing"
)

func TestGetLegislators(t *testing.T) {
	f := newFactoryForTesting(t)
	legs, err := f.GetLegislators("CA")
	if err != nil {
		t.Fatalf("unexpected error calling GetLegislators: %v", err)
	}
	found := false
	const firstlast = "Nancy Pelosi"
	for _, leg := range legs {
		fl, err := leg.Firstlast()
		if err != nil {
			t.Fatalf("unexpected error calling Firstlast: %v", err)
		}
		if fl == firstlast {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("did not find expecting firstlast %s", firstlast)
	}
}
