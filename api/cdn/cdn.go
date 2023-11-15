package cdn

import (
	"github.com/alirezaopmc/arvan/internal/proxy"
)

type API struct {
	BaseURL   string
	HttpProxy *proxy.HttpProxy
}

func NewAPI(baseURL string, httpProxy *proxy.HttpProxy) *API {
	return &API{
		BaseURL:   baseURL,
		HttpProxy: httpProxy,
	}
}
