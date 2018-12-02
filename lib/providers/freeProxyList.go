package providers

import (
	"strings"

	"golang.org/x/net/html"

	httpCommand "proxy_scraper/lib/httpCommand"
)

type proxyProviderFreeProxyList struct {
	url            string
	httpCommand    *httpCommand.HTTPRequest
	dataResponse   *html.Node
	headers        []string
	indexedHeaders map[string]int
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

func (p *proxyProviderFreeProxyList) SetTableHeaders() {
	var answer []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Data == "th" {
			thData := n.FirstChild.Data
			answer = append(answer, thData)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(p.dataResponse)
	p.headers = answer
}

func (p *proxyProviderFreeProxyList) GetHeaderIndex(headerName string) int {
	formatedHeader := strings.ToLower(headerName)
	for i := 0; i < len(p.headers); i++ {
		if formatedHeader == strings.ToLower(p.headers[i]) {
			return i
		}
	}
	return -1
}

func (p *proxyProviderFreeProxyList) SetMapHeaders(arguments ...string) {
	for _, arg := range arguments {
		indexCalculated := p.GetHeaderIndex(arg)
		if indexCalculated > 0 {
			p.indexedHeaders[arg] = indexCalculated
		}
	}
}

func (p *proxyProviderFreeProxyList) GetIP() string {
	var answer string
	p.SetTableHeaders()
	p.SetMapHeaders()
	var f func(*html.Node)
	index := 0
	f = func(n *html.Node) {
		if n.Data == "td" {
			index++
			tdData := n.FirstChild.Data
			if index == p.indexedHeaders["IP"] {
				answer = tdData
			}
		}
		for c := p.dataResponse.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(p.dataResponse)
	return answer
}
