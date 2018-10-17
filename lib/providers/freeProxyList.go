package providers

import (
	"net/http"

	"golang.org/x/net/html"
)

type ProxyProviderFreeProxyList struct {
	URL          string
	Client       *http.Client
	Request      *http.Request
	DataResponse *html.Node
}

func New() *ProxyProviderFreeProxyList {
	return &ProxyProviderFreeProxyList{URL: "https://free-proxy-list.net/",
		Client: &http.Client{}}
}

func (p ProxyProviderFreeProxyList) GetSource() string {
	return p.URL
}

// func (p ProxyProviderFreeProxyList) GetResponse() *html.Node {
// 	client := &http.Client{}
// 	// req, _ := http.NewRequest("GET", "https://free-proxy-list.net/", nil)
// 	req, _ := http.NewRequest("GET", p.URL, nil)

// }
