package main

import (
	"context"

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/opensecrets/cli"
)

func main() {
	check.Err(cli.Main(context.Background()))
}
