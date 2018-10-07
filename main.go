package main

import (
	"fmt"
	"log"
	"net/http"
	proxy "proxy_scraper/lib/proxy"
	"regexp"

	"golang.org/x/net/html"
)

var list *proxy.ProxiesList = proxy.NewList()

func findIP(n *html.Node) {
	var proxyObj *proxy.Proxy = proxy.New()
	if n.FirstChild != nil {
		data := fmt.Sprintf("%v", n.FirstChild.Data)
		re := regexp.MustCompile(`\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}`)
		reEval := re.FindAllStringSubmatch(data, -1)
		if len(reEval) > 0 {
			proxyObj.SetIP(reEval[0][0])
			findPort(n, proxyObj)
		}
	}
}

func findPort(n *html.Node, proxy *proxy.Proxy) {
	if n.NextSibling != nil {
		data := fmt.Sprintf("%v", n.NextSibling.FirstChild.Data)
		proxy.SetPort(data)
		list.AddProxy(proxy)
	}
}

func main() {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://free-proxy-list.net/", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64 AppleWebKit/537.36 (KHTML, like Gecko Chrome/60.0.3112.113 Safari/537.36")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Data == "td" {
			findIP(n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
		list.ToJSON("proxies.json")
	}
	f(doc)
}
