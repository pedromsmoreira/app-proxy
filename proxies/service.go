package proxies

import "net/http"

type DoProxy interface {
	Serve(endpoint string, httpMethod string) *http.Response
}

type ProxyService struct {
	proxies *Proxies
}
