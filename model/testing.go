package model

import (
	"testing"

	"github.com/spudtrooper/goutil/io"
	"github.com/spudtrooper/opensecrets/api"
)

func newFactoryForTesting(t *testing.T) *Factory {
	f := "../.user_creds.json"
	if !io.FileExists(f) {
		t.Fatalf("add your credentials to %s", f)
	}
	client, err := api.NewClientFromFile(f)
	if err != nil {
		t.Fatalf("NewClientFromFlags: %v", err)
	}
	return NewFactory(client)
}
