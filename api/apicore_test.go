package api

import (
	"testing"

	"github.com/spudtrooper/goutil/io"
)

func makeClient(t *testing.T) *core {
	f := "../.user_creds.json"
	if !io.FileExists(f) {
		t.Fatalf("add your credentials to %s", f)
	}
	client, err := NewClientFromFile("../.user_creds.json")
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
		const firstlast = "John Kennedy"
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

	runTest("LA")
	runTest("N00026823")
}
