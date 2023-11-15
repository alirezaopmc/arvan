package api

import (
	"github.com/alirezaopmc/arvan/api/cdn"
	"github.com/alirezaopmc/arvan/internal/proxy"
)

type API struct {
	BaseURL string
	CDN     *cdn.API
}

func NewAPI(baseURL, apiKey string) *API {
	httpProxy := proxy.NewHttpProxy()
	httpProxy.AddHeader("authorization", apiKey)
	httpProxy.AddHeader("Accept", "application/json")

	return &API{
		BaseURL: baseURL,
		CDN:     cdn.NewAPI(baseURL+"/cdn/4.0", httpProxy),
	}
}
