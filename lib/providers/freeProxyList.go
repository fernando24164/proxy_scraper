package providers

import (
	"golang.org/x/net/html"

	httpCommand "proxy_scraper/lib/httpCommand"
)

type proxyProviderFreeProxyList struct {
	url          string
	httpCommand  *httpCommand.HTTPRequest
	dataResponse *html.Node
}

func New() *proxyProviderFreeProxyList {
	return &proxyProviderFreeProxyList{url: "https://free-proxy-list.net/",
		httpCommand: httpCommand.NewHTTPRequest()}
}

func (p proxyProviderFreeProxyList) GetSource() string {
	return p.url
}

func (p *proxyProviderFreeProxyList) SetDataResponse() {
	p.dataResponse = p.httpCommand.Execute(p.url)
}

// TODO Add test to this function
func (p *proxyProviderFreeProxyList) GetIP() string {
	var answer string
	for c := p.dataResponse.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "td" {
			answer = c.Data
		}
	}
	return answer
}
