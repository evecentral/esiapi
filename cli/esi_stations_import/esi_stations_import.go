package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"cloud.google.com/go/datastore"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/theatrus/mediate"

	"github.com/evecentral/esiapi"
	"github.com/evecentral/esiapi/client"
	"github.com/evecentral/esiapi/client/universe"
	"github.com/evecentral/esiapi/helper"
)


var settingsFile string
var project string
var tokenName string

func init() {
	flag.StringVar(&settingsFile, "esi.settings", "settings.json", "Default settings file")
	flag.StringVar(&project, "esi.ds.project", "", "Google Cloud Datastore Project")
	flag.StringVar(&tokenName, "esi.token.name", "default-token", "Name of the token")
}

func structureImport(dsClient *datastore.Client, client *client.EVESwaggerInterface) {

	// Do a test and just enumerate all structures
	structureIds, err := client.Universe.GetUniverseStructures(nil)
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(structureIds.Payload))

	stationsLock := sync.Mutex{}
	stations := make([]*esiapi.Station, 0)

	for _, structureId := range structureIds.Payload {
		go func(structureId int64) {
			defer wg.Done()
			ctx := context.Background()
			params := universe.NewGetUniverseStructuresStructureIDParamsWithContext(ctx)
			params.StructureID = structureId

			res, err := client.Universe.GetUniverseStructuresStructureID(params, httptransport.PassThroughAuth)
			if err != nil {
				log.Println(err)
				return
			}

			key := datastore.NameKey("stations", fmt.Sprintf("%d", int(structureId)), nil)

			stationObject := &esiapi.Station{Name: *res.Payload.Name,
				Id:          int(structureId),
				SolarSystem: int(*res.Payload.SolarSystemID),
			}

			_, err = dsClient.Put(ctx, key, stationObject)
			if err != nil {
				log.Printf("station saving failure %v", err)
			}

			stationsLock.Lock()
			defer stationsLock.Unlock()
			stations = append(stations, stationObject)
			log.Printf("structure %s - %d", *res.Payload.Name, *res.Payload.SolarSystemID)
		}(*structureId)
	}
	wg.Wait()

	final, _ := json.Marshal(&stations)
	fmt.Println(string(final))

}

func main() {
	flag.Parse()

	settings, err := helper.LoadSettings(settingsFile)
	if err != nil {
		log.Fatal(err)
	}


	store, err := helper.NewDatastoreTokenStore(project, tokenName, 120*time.Second)
	if err != nil {
		fmt.Printf("Can't load datastore %v\n", err)
		return
	}

	transport, err := helper.InteractiveStartupWithTokenStore(store, settings)
	if err != nil {
		fmt.Printf("Can't do startup %v\n", err)
		return
	}

	cliTransport := httptransport.New("esi.tech.ccp.is", "/latest", []string{"https"})
	cliTransport.Transport = mediate.RateLimit(10, 1*time.Second, transport)

	// Connect to cloud datastore
	ctx := context.Background()

	dsClient, err := datastore.NewClient(ctx, project)
	if err != nil {
		log.Fatal(err)
	}


	client := client.New(cliTransport, strfmt.Default)
	structureImport(dsClient, client)
}
