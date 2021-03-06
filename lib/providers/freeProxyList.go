package providers

import (
	"strings"

	"golang.org/x/net/html"

	"proxy_scraper/lib/httpCommand"

	"proxy_scraper/lib/proxy"
)

type proxyProviderFreeProxyList struct {
	url            string
	httpCommand    *httpCommand.HTTPRequest
	dataResponse   *html.Node
	headers        []string
	indexedHeaders map[string]int
}

func New() *proxyProviderFreeProxyList {
	return &proxyProviderFreeProxyList{
		url:            "https://free-proxy-list.net/",
		httpCommand:    httpCommand.NewHTTPRequest(),
		indexedHeaders: make(map[string]int)}
}

func (p proxyProviderFreeProxyList) GetSource() string {
	return p.url
}

func (p *proxyProviderFreeProxyList) SetDataResponse() {
	p.dataResponse = p.httpCommand.Execute(p.url)
	p.SetTableHeaders()
	p.SetMapHeaders("ip address", "port", "https")
}

func (p *proxyProviderFreeProxyList) SetTableHeaders() {
	var answer []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Data == "th" {
			if n.FirstChild != nil {
				thData := strings.ToLower(n.FirstChild.Data)
				answer = append(answer, thData)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(p.dataResponse)
	p.headers = answer
}

func (p *proxyProviderFreeProxyList) GetHeaderIndex(headerName string) int {
	formattedHeader := strings.ToLower(headerName)
	for i := 0; i < len(p.headers); i++ {
		if formattedHeader == strings.ToLower(p.headers[i]) {
			return i
		}
	}
	return -1
}

func (p *proxyProviderFreeProxyList) SetMapHeaders(arguments ...string) {
	for _, arg := range arguments {
		indexCalculated := p.GetHeaderIndex(arg)
		if indexCalculated >= 0 {
			p.indexedHeaders[arg] = indexCalculated
		}
	}
}

func (p *proxyProviderFreeProxyList) GetIP() string {
	var answer string
	var f func(*html.Node)
	lengthHeaders := len(p.headers)
	index := 0
	f = func(n *html.Node) {
		if n.Data == "td" {
			tdData := n.FirstChild.Data
			if index%lengthHeaders == p.indexedHeaders["ip address"] {
				answer = tdData
			}
			index++
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(p.dataResponse)
	return answer
}

func (p *proxyProviderFreeProxyList) GetPort() string {
	var answer string
	var f func(*html.Node)
	lengthHeaders := len(p.headers)
	index := 0
	f = func(n *html.Node) {
		if n.Data == "td" {
			tdData := n.FirstChild.Data
			if index%lengthHeaders == p.indexedHeaders["port"] {
				answer = tdData
			}
			index++
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(p.dataResponse)
	return answer
}

func (p *proxyProviderFreeProxyList) GetProtocol() string {
	var answer string
	var f func(*html.Node)
	lengthHeaders := len(p.headers)
	index := 0
	f = func(n *html.Node) {
		if n.Data == "td" {
			tdData := n.FirstChild.Data
			if index%lengthHeaders == p.indexedHeaders["https"] {
				if strings.Compare(tdData, "yes") == 0 {
					answer = "https"
				} else {
					answer = "http"
				}
			}
			index++
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(p.dataResponse)
	return answer
}

func (p *proxyProviderFreeProxyList) GetProxiesList() *proxy.ProxiesList {
	var f func(*html.Node)
	proxiesList := proxy.NewList()
	var proxyAux *proxy.Proxy
	lengthHeaders := len(p.headers)-1
	index := 0
	f = func(n *html.Node) {
		if n.Data == "tr" {
			proxyAux = proxy.New()
		}
		if n.Data == "td" {
			if n.FirstChild != nil {
				tdData := n.FirstChild.Data
				if index%lengthHeaders == p.indexedHeaders["ip address"] {
					proxyAux.SetIP(tdData)
				}
				if index%lengthHeaders == p.indexedHeaders["port"] {
					proxyAux.SetPort(tdData)
				}
				if index%lengthHeaders == p.indexedHeaders["https"] {
					if strings.Compare(tdData, "yes") == 0 {
						proxyAux.SetProtocol("https")
					} else {
						proxyAux.SetProtocol("http")
					}
				}
				if index%lengthHeaders == 0 {
					proxiesList.AddProxy(proxyAux)
				}
				index++
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(p.dataResponse)
	return proxiesList
}
