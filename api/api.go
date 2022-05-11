package api

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/flags"
	"github.com/spudtrooper/goutil/request"
)

var (
	apiKEY    = flags.String("api_key", "api key")
	userCreds = flag.String("user_creds", ".user_creds.json", "file with user credentials")
)

// type core represents the core gettr core
type core struct {
	apiKEY string
}

func NewClientFromFlags() (*core, error) {
	if *apiKEY != "" {
		client := NewClient(*apiKEY)
		return client, nil
	}
	if *userCreds != "" {
		client, err := NewClientFromFile(*userCreds)
		if err != nil {
			return nil, err
		}
		return client, nil
	}
	return nil, errors.Errorf("Must set --user & --token or --creds_file")
}

func NewClient(apiKEY string) *core {
	return &core{
		apiKEY: apiKEY,
	}
}

type Creds struct {
	ApiKey string `json:"api_key"`
}

func ReadCredsFromFlags() (Creds, error) {
	return readCreds(*userCreds)
}

func WriteCredsFromFlags(creds Creds) error {
	b, err := json.Marshal(&creds)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(*userCreds, b, 0755); err != nil {
		return err
	}
	log.Printf("wrote to %s", *userCreds)
	return nil
}

func readCreds(credsFile string) (creds Creds, ret error) {
	credsBytes, err := ioutil.ReadFile(credsFile)
	if err != nil {
		ret = err
		return
	}
	if err := json.Unmarshal(credsBytes, &creds); err != nil {
		ret = err
		return
	}
	return
}

func NewClientFromFile(credsFile string) (*core, error) {
	creds, err := readCreds(credsFile)
	if err != nil {
		return nil, err
	}
	return &core{
		apiKEY: creds.ApiKey,
	}, nil
}

func (c *core) createRoute(method string, params ...request.Param) string {
	const host = "www.opensecrets.org"
	base := fmt.Sprintf("http://%s/api", host)
	params = append(params, []request.Param{
		request.MakeParam("method", method),
		request.MakeParam("apikey", c.apiKEY),
		request.MakeParam("output", "json"),
	}...)
	url := request.CreateRoute(base, params...)
	return url
}

func (c *core) get(method string, payload interface{}, params ...request.Param) error {
	uri := c.createRoute(method, params...)

	headers := map[string]string{
		"user-agent": `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36`,
	}
	_, err := request.Get(uri, payload, request.RequestExtraHeaders(headers))
	if err != nil {
		return err
	}

	return nil

}
