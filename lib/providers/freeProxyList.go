package providers

import (
	"log"
	"net/http"

	"golang.org/x/net/html"
)

type proxyProviderFreeProxyList struct {
	url          string
	client       *http.Client
	request      *http.Request
	dataResponse *html.Node
}

func New() *proxyProviderFreeProxyList {
	return &proxyProviderFreeProxyList{url: "https://free-proxy-list.net/",
		client: &http.Client{}}
}

func (p proxyProviderFreeProxyList) GetSource() string {
	return p.url
}

func (p *proxyProviderFreeProxyList) SetRequest() {
	p.request, _ = http.NewRequest("GET", p.url, nil)
	p.request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64 AppleWebKit/537.36 (KHTML, like Gecko Chrome/60.0.3112.113 Safari/537.36")
}

func (p *proxyProviderFreeProxyList) SetDataResponse() {
	resp, err := p.client.Do(p.request)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	p.dataResponse, err = html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
