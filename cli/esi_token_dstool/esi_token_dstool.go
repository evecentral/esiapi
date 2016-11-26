package main

import (
	"flag"
	"fmt"
	"time"

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

func main() {
	flag.Parse()

	settings, err := helper.LoadSettings(settingsFile)
	if err != nil {
		fmt.Printf("Can't load settings file %v\n", settings)
		return
	}

	store, err := helper.NewDatastoreTokenStore(project, tokenName, 60*time.Second)
	if err != nil {
		fmt.Printf("Can't load datastore %v\n", err)
		return
	}

	_, err = helper.InteractiveStartupWithTokenStore(store, settings)
	if err != nil {
		fmt.Printf("Can't do startup %v\n", err)
		return
	}

	token, err := store.ReadToken()
	if err != nil {
		fmt.Printf("Can't load token after startup\n", err)
		return
	}
	fmt.Printf("Token refresh success - %v\n", token)

}
