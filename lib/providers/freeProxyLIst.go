package providers

import (
	"time"
	"net/http"

	"golang.org/x/net/html"
)

type ProxyProviderFreeProxyList struct {
	URL          string
	Client       *http.Client
	Request 	 *http.Request
	DataResponse *html.Node
}

func (p ProxyProviderFreeProxyList) GetSource() string {
	return p.URL
}

func (p *ProxyProviderFreeProxyList) SetHeader() {
	p.Client := &http.Client{Timeout=time.Second*30}
	req, _ := http.NewRequest("GET", p.URL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64 AppleWebKit/537.36 (KHTML, like Gecko Chrome/60.0.3112.113 Safari/537.36")
	resp, err := p.client.Do(req)
}

func (p ProxyProviderFreeProxyList) GetResponse() *html.Node {
	client := &http.Client{}
	// req, _ := http.NewRequest("GET", "https://free-proxy-list.net/", nil)
	req, _ := http.NewRequest("GET", p.URL, nil)

}
