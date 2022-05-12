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
	id    = flags.String("id", "global ID")
	cid   = flags.String("cid", "candidate ID")
	year  = flags.Int("year", "global year")
	cycle = flags.Int("cycle", "global cycle")
)

func Main(ctx context.Context) error {
	app := minimalcli.Make()
	app.Init()

	coreClient, err := api.NewClientFromFlags()
	if err != nil {
		return err
	}
	client := coreClient //api.MakeExtended(coreClient)

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
		info, err := client.GetMemPFDprofile(*cid, api.MemPFDprofileYear(*year))
		if err != nil {
			return err
		}
		log.Printf("GetMemPFDprofile: %s", mustFormatString(info))
		return nil
	})

	app.Register("GetCandSummary", func(context.Context) error {
		requireStringFlag(cid, "cid")
		info, err := client.GetCandSummary(*cid, api.CandSummaryCycle(*cycle))
		if err != nil {
			return err
		}
		log.Printf("GetCandSummary: %s", mustFormatString(info))
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
