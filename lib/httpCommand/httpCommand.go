package httpCommand

import (
	"log"
	"net/http"

	"golang.org/x/net/html"
)

type Command interface {
	Execute() string
}

type HTTPRequest struct {
	client  *http.Client
	request *http.Request
}

func (hr *HTTPRequest) SetRequest(url string, arguments ...string) {
	if len(arguments) < 1 {
		hr.request, _ = http.NewRequest("GET", url, nil)
		hr.request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64 AppleWebKit/537.36 (KHTML, like Gecko Chrome/60.0.3112.113 Safari/537.36")
	} else {
		hr.request.Header.Set("User-Agent", arguments[0])
	}
}

func (hr *HTTPRequest) Execute() *html.Node {
	resp, err := hr.client.Do(hr.request)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	dataResponse, err := html.Parse(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return dataResponse
}
