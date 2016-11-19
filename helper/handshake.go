package helper

import (
	"encoding/json"
	"github.com/theatrus/mediate"
	"github.com/theatrus/ooauth2"
	"io/ioutil"
	"log"
	"net/http"
)

type OAuthSettings struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Callback     string `json:"callback"`
}

func LoadSettings(filename string) (*OAuthSettings, error) {
	var settings OAuthSettings
	settingsData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Can't load secret key file - aborting", err)
		return nil, err
	}
	json.Unmarshal(settingsData, &settings)
	return &settings, nil
}

func NewOAuthOptions(settings *OAuthSettings) (*ooauth2.Options, error) {
	var endpoint ooauth2.Option
	endpoint = ooauth2.Endpoint(
		"https://login.eveonline.com/oauth/authorize",
		"https://login.eveonline.com/oauth/token",
	)

	httpClient := &http.Client{}
	httpClient.Transport = mediate.FixedRetries(3,
		mediate.ReliableBody(http.DefaultTransport),
	)

	return ooauth2.New(
		ooauth2.Client(settings.ClientId, settings.ClientSecret),
		ooauth2.RedirectURL(settings.Callback),
		ooauth2.Scope("esi-universe.read_structures.v1"),
		ooauth2.HTTPClient(httpClient),
		endpoint,
	)
}
