// Package handlers is a bridge between the API and handlers to create a CLI/API server.
package handlers

import (
	"context"
	_ "embed"

	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/opensecrets/api"
)

//go:generate minimalcli gsl --input handlers.go --uri_root "github.com/spudtrooper/opensecrets/blob/main/handlers" --output handlers.go.json
//go:embed handlers.go.json
var SourceLocations []byte

func CreateHandlers(client *api.Core) []handler.Handler {
	b := handler.NewHandlerBuilder()

	b.NewHandler("GetLegislator",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.GetLegislatorParams)
			return client.GetLegislator(p.Cid)
		},
		api.GetLegislatorParams{},
	)

	b.NewHandler("GetLegislators",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.GetLegislatorsParams)
			return client.GetLegislators(p.Id)
		},
		api.GetLegislatorsParams{},
	)

	b.NewHandler("GetMemPFDprofile",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.GetMemPFDprofileParams)
			return client.GetMemPFDprofile(p.Cid, p.Options()...)
		},
		api.GetMemPFDprofileParams{},
	)

	b.NewHandler("GetCandSummary",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.GetCandSummaryParams)
			return client.GetCandSummary(p.Cid, p.Options()...)
		},
		api.GetCandSummaryParams{},
	)

	b.NewHandler("GetCandContrib",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.GetCandContribParams)
			return client.GetCandContrib(p.Cid, p.Options()...)
		},
		api.GetCandContribParams{},
	)

	b.NewHandler("GetCandByInd",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.GetCandByIndParams)
			return client.GetCandByInd(p.Cid, p.Ind, p.Options()...)
		},
		api.GetCandByIndParams{},
	)

	b.NewHandler("GetCandSector",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.GetCandSectorParams)
			return client.GetCandSector(p.Cid, p.Options()...)
		},
		api.GetCandSectorParams{},
	)

	b.NewHandler("GetCongCmteIndus",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.GetCongCmteIndusParams)
			return client.GetCongCmteIndus(p.Cmte, p.Ind, p.Options()...)
		},
		api.GetCongCmteIndusParams{},
	)

	b.NewHandler("GetOrgs",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.GetOrgsParams)
			return client.GetOrgs(p.Org)
		},
		api.GetOrgsParams{},
	)

	b.NewHandler("GetOrgSummary",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.GetOrgSummaryParams)
			return client.GetOrgSummary(p.OrgID)
		},
		api.GetOrgSummaryParams{},
	)

	b.NewHandler("GetIndependentExpend",
		func(ctx context.Context, ip any) (any, error) {
			return client.GetIndependentExpend()
		},
		api.GetIndependentExpendParams{},
	)

	return b.Build()
}
