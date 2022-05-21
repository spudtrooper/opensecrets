package main

import (
	"flag"
	"log"

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/opensecrets/api"
	"github.com/spudtrooper/opensecrets/model"
)

func realMain() {
	log.Printf("creating objects...")
	check.Err(model.CreateObject(model.ObjectProto{
		ObjectName: "Legislator",
		InfoProxy:  api.LegislatorInfo{},
		OutputFile: "model/legmethods.go",
	}))
	check.Err(model.CreateObject(model.ObjectProto{
		ObjectName: "MemPFDprofile",
		InfoProxy:  api.MemPFDprofileInfo{},
		OutputFile: "model/mempfdprofilemethods.go",
		Simple:     true,
	}))
	check.Err(model.CreateObject(model.ObjectProto{
		ObjectName: "CandSummary",
		InfoProxy:  api.CandSummaryInfo{},
		OutputFile: "model/candsummarymethods.go",
		Simple:     true,
	}))
	check.Err(model.CreateObject(model.ObjectProto{
		ObjectName: "Contributor",
		InfoProxy:  api.Contributor{},
		OutputFile: "model/contributormethods.go",
		Simple:     true,
	}))
	check.Err(model.CreateObject(model.ObjectProto{
		ObjectName: "Organization",
		InfoProxy:  api.Organization{},
		OutputFile: "model/orgmethods.go",
	}))
	check.Err(model.CreateObject(model.ObjectProto{
		ObjectName: "Industry",
		InfoProxy:  api.Industry{},
		OutputFile: "model/industrymethods.go",
		Simple:     true,
	}))
}

func main() {
	flag.Parse()
	realMain()
}
