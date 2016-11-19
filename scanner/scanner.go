package main

import (
	"log"
	"flag"
	"sync"

	"github.com/go-openapi/strfmt"
	httptransport "github.com/go-openapi/runtime/client"

	"github.com/evecentral/esiapi/helper"
	"github.com/evecentral/esiapi/client"
	"github.com/evecentral/esiapi/client/universe"
)

var uploadEndpoint string
var settingsFile string
var onlyRegion int

func init() {
	flag.StringVar(&uploadEndpoint, "scanner.upload", "http://localhost", "Default upload endpoint")
	flag.StringVar(&settingsFile, "scanner.settings", "settings.json", "Default settings file")
	flag.IntVar(&onlyRegion, "scanner.region", 0, "Limit to a specific region")
}

func main() {
	flag.Parse()

	settings, err := helper.LoadSettings(settingsFile)
	if err != nil {
		log.Fatal(err)
	}

	// Grab a go-oauth2 http RoundTripper transport
	transport, err := helper.InteractiveStartup("token.json", settings)
	if err != nil {
		log.Fatal(err)
	}

	cliTransport := httptransport.New("esi.tech.ccp.is", "/latest", []string{"https"})
	cliTransport.Transport = transport

	client := client.New(cliTransport, strfmt.Default)

	// Do a test and just enumerate all structures
	structureIds, err := client.Universe.GetUniverseStructures(nil)
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}
	log.Printf("fetching %v entries", len(structureIds.Payload))

	wg.Add(len(structureIds.Payload))
	for _, structureId := range structureIds.Payload {
		go func(structureId int64) {
			defer wg.Done()
			params := universe.NewGetUniverseStructuresStructureIDParams()
			params.StructureID = structureId

			res, err := client.Universe.GetUniverseStructuresStructureID(params, httptransport.PassThroughAuth)
			if err != nil {
				log.Println(err)
				return
			}
			log.Printf("structure %s - %d", *res.Payload.Name, *res.Payload.SolarSystemID)
		}(structureId)
	}
	wg.Wait()
}
