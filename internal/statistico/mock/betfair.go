package mock

import (
	bfClient "github.com/statistico/statistico-betfair-go-client"
	"net/http"
)

func BetfairClient(client *http.Client) *bfClient.Client {
	creds := bfClient.InteractiveCredentials{Token:"token"}

	base := bfClient.BaseURLs{
		Betting:  "https://api.betfair.com/test",
	}

	return &bfClient.Client{
		HTTPClient:  client,
		Credentials: creds,
		BaseURLs:    base,
	}
}
