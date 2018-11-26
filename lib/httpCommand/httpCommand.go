package httpCommand

import (
	"log"
	"net/http"

	"golang.org/x/net/html"
)

type Command interface {
	Execute() *html.Node
}

type HTTPRequest struct {
	client         *http.Client
	request        *http.Request
	cachedResponse *html.Node
	isCached       bool
}

func NewHTTPRequest() *HTTPRequest {
	return &HTTPRequest{client: &http.Client{},
		isCached: false}
}

// TODO Add curry function to set more headers
func (hr *HTTPRequest) SetRequest(url string, arguments ...string) {
	hr.request, _ = http.NewRequest("GET", url, nil)
	hr.request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64 AppleWebKit/537.36 (KHTML, like Gecko Chrome/60.0.3112.113 Safari/537.36")
}

func (hr *HTTPRequest) Execute(url string) *html.Node {
	if !hr.isCached {
		hr.SetRequest(url)
		resp, err := hr.client.Do(hr.request)

		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		dataResponse, err := html.Parse(resp.Body)

		if err != nil {
			log.Fatal(err)
		}
		hr.isCached = true
		hr.cachedResponse = dataResponse
		return dataResponse
	}
	return hr.cachedResponse
}
