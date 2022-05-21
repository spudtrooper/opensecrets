package model

import (
	"strings"
	"testing"
)

func TestOrgs(t *testing.T) {
	f := newFactoryForTesting(t)
	orgs, err := f.GetOrgs("Goldman")
	if err != nil {
		t.Fatalf("GetOrgs: %v", err)
	}
	if len(orgs) == 0 {
		t.Errorf("expecting > orgs, got none")
	}
	const orgName = "Goldman Sachs"
	var names []string
	found := false
	for _, o := range orgs {
		name, err := o.Orgname()
		if err != nil {
			t.Fatalf("Orgname: %v", err)
		}
		names = append(names, name)
		if name == orgName {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("did not find expecting OrgName from [%s]: %s", strings.Join(names, ","), orgName)
	}
}
