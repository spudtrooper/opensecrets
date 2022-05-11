package main

import (
	"flag"
	"log"

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/goutil/request"
)

func main() {
	flag.Parse()
	// Options
	printData := true
	printCookies := true
	printPayload := true

	// Data
	uri := request.MakeURL("https://www.opensecrets.org/api",
		request.Param{"id", `LA`},
		request.Param{"method", `getLegislators`},
		request.Param{"apikey", `90a5efe92ce0f29f5a64855d72f1bfb5`},
		request.Param{"output", `json`},
	)

	headers := map[string]string{
		"user-agent": `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36`,
	}

	// Make the request
	var payload interface{}
	var res *request.Response
	var err error

	res, err = request.Get(uri, &payload, request.RequestExtraHeaders(headers))

	if printData {
		log.Printf("data: %s", string(res.Data))
	}
	if printCookies {
		log.Printf("cookies: %v", res.Cookies)
	}
	if printPayload {
		log.Printf("payload: %s", request.MustFormatString(payload))
	}
	check.Err(err)
}
