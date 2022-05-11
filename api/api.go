package api

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/flags"
	"github.com/spudtrooper/goutil/request"
)

var (
	apiKEY    = flags.String("api_key", "api key")
	userCreds = flag.String("user_creds", ".user_creds.json", "file with user credentials")
)

// core is an opaque API client for opensecrets.org
type core struct {
	apiKEY string
}

// NewClientFromFlags creates a new client from command line flags
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

// NewClient creates a new client directly from the API Key
func NewClient(apiKEY string) *core {
	return &core{
		apiKEY: apiKEY,
	}
}

// NewClientFromFile creates a new client from a JSON file like `user_creds-example.json`
func NewClientFromFile(credsFile string) (*core, error) {
	creds, err := readCreds(credsFile)
	if err != nil {
		return nil, err
	}
	return &core{
		apiKEY: creds.ApiKey,
	}, nil
}

type creds struct {
	ApiKey string `json:"api_key"`
}

func readCreds(credsFile string) (*creds, error) {
	credsBytes, err := ioutil.ReadFile(credsFile)
	if err != nil {
		return nil, err
	}
	var creds creds
	if err := json.Unmarshal(credsBytes, &creds); err != nil {
		return nil, err
	}
	return &creds, nil
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
	if _, err := request.Get(uri, payload, request.RequestExtraHeaders(headers)); err != nil {
		return err
	}

	return nil

}
