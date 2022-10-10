package cli

import (
	"context"

	"github.com/spudtrooper/goutil/flags"
	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/opensecrets/api"
	"github.com/spudtrooper/opensecrets/handlers"
)

var (
	id     = flags.String("id", "global ID")
	cid    = flags.String("cid", "candidate ID")
	cmte   = flags.String("cmte", "candidate cmte")
	year   = flags.Int("year", "global year")
	cycle  = flags.Int("cycle", "global cycle")
	congno = flags.Int("congno", "global congno")
	ind    = flags.String("ind", "global ind")
	org    = flags.String("org", "candidate org")
	orgID  = flags.String("org_id", "candidate org ID")
)

func Main(ctx context.Context) error {
	adp := handler.NewCLIAdapter()
	adp.BindStringFlag("id", id)
	adp.BindStringFlag("cid", cid)
	adp.BindStringFlag("cmte", cmte)
	adp.BindIntFlag("year", year)
	adp.BindIntFlag("incycled", cycle)
	adp.BindIntFlag("congno", congno)
	adp.BindStringFlag("ind", ind)
	adp.BindStringFlag("org", org)
	adp.BindStringFlag("org_id", orgID)

	client, err := api.NewClientFromFlags()
	if err != nil {
		return err
	}

	handlers := handlers.CreateHandlers(client)
	app := handler.CreateCLI(adp, handlers...)

	if err := app.Run(ctx); err != nil {
		return err
	}

	return nil

	return nil
}
