package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/goutil/flags"
	minimalcli "github.com/spudtrooper/minimalcli/app"
	"github.com/spudtrooper/opensecrets/api"
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
	app := minimalcli.Make()
	app.Init()

	client, err := api.NewClientFromFlags()
	if err != nil {
		return err
	}

	app.Register("GetLegislators", func(context.Context) error {
		requireStringFlag(id, "id")
		info, err := client.GetLegislators(*id)
		if err != nil {
			return err
		}
		log.Printf("GetLegislators: %s", mustFormatString(info))
		return nil
	})

	app.Register("GetMemPFDprofile", func(context.Context) error {
		requireStringFlag(cid, "cid")
		info, err := client.GetMemPFDprofile(*cid, api.GetMemPFDprofileYear(*year))
		if err != nil {
			return err
		}
		log.Printf("GetMemPFDprofile: %s", mustFormatString(info))
		return nil
	})

	app.Register("GetCandSummary", func(context.Context) error {
		requireStringFlag(cid, "cid")
		info, err := client.GetCandSummary(*cid, api.GetCandSummaryCycle(*cycle))
		if err != nil {
			return err
		}
		log.Printf("GetCandSummary: %s", mustFormatString(info))
		return nil
	})

	app.Register("GetCandContrib", func(context.Context) error {
		requireStringFlag(cid, "cid")
		info, err := client.GetCandContrib(*cid, api.GetCandContribCycle(*cycle))
		if err != nil {
			return err
		}
		log.Printf("GetCandContrib: %s", mustFormatString(info))
		return nil
	})

	app.Register("GetCandIndustry", func(context.Context) error {
		requireStringFlag(cid, "cid")
		info, err := client.GetCandIndustry(*cid, api.GetCandIndustryCycle(*cycle))
		if err != nil {
			return err
		}
		log.Printf("GetCandIndustry: %s", mustFormatString(info))
		return nil
	})

	app.Register("GetCandByInd", func(context.Context) error {
		requireStringFlag(cid, "cid")
		requireStringFlag(ind, "ind")
		info, err := client.GetCandByInd(*cid, *ind, api.GetCandByIndCycle(*cycle))
		if err != nil {
			return err
		}
		log.Printf("GetCandByInd: %s", mustFormatString(info))
		return nil
	})

	app.Register("GetCandSector", func(context.Context) error {
		requireStringFlag(cid, "cid")
		info, err := client.GetCandSector(*cid, api.GetCandSectorCycle(*cycle))
		if err != nil {
			return err
		}
		log.Printf("GetCandSector: %s", mustFormatString(info))
		return nil
	})

	app.Register("GetCongCmteIndus", func(context.Context) error {
		requireStringFlag(cmte, "cmte")
		requireStringFlag(ind, "ind")
		info, err := client.GetCongCmteIndus(*cmte, *ind, api.GetCongCmteIndusCongno(*congno))
		if err != nil {
			return err
		}
		log.Printf("GetCongCmteIndus: %s", mustFormatString(info))
		return nil
	})

	app.Register("GetOrgs", func(context.Context) error {
		requireStringFlag(org, "org")
		info, err := client.GetOrgs(*org)
		if err != nil {
			return err
		}
		log.Printf("GetOrgs: %s", mustFormatString(info))
		return nil
	})

	app.Register("GetOrgSummary", func(context.Context) error {
		requireStringFlag(orgID, "org_id")
		info, err := client.GetOrgSummary(*orgID)
		if err != nil {
			return err
		}
		log.Printf("GetOrgSummary: %s", mustFormatString(info))
		return nil
	})

	app.Register("GetIndependentExpend", func(context.Context) error {
		info, err := client.GetIndependentExpend()
		if err != nil {
			return err
		}
		log.Printf("GetIndependentExpend: %s", mustFormatString(info))
		return nil
	})

	if err := app.Run(ctx); err != nil {
		return err
	}

	return nil
}

func requireStringFlag(flag *string, name string) {
	if *flag == "" {
		log.Fatalf("--%s required", name)
	}
}

func mustFormatString(x interface{}) string {
	b, err := json.Marshal(x)
	check.Err(err)
	res, err := prettyPrintJSON(b)
	check.Err(err)
	return res
}

func prettyPrintJSON(b []byte) (string, error) {
	b = []byte(strings.TrimSpace(string(b)))
	if len(b) == 0 {
		return "", nil
	}
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, b, "", "\t"); err != nil {
		return "", errors.Errorf("json.Indent: payload=%q: %v", string(b), err)
	}
	return prettyJSON.String(), nil
}
