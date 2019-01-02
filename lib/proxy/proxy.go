package proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

var regexIP = `\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`
var regexPort = `[1-65535]`
var regexProtocol = `http|https`

type Proxy struct {
	IP       string `json:"ip"`
	Port     string `json:"port"`
	Protocol string `json:"protocol"`
}

func New() *Proxy {
	return &Proxy{}
}

func (p *Proxy) SetIP(ip string) {
	p.IP = ip
}

func (p *Proxy) SetPort(port string) {
	p.Port = port
}

func (p *Proxy) SetProtocol(protocol string) {
	p.Protocol = protocol
}

//IsValid check if IP address, port and protocol are valid
func (p Proxy) IsValid() bool {
	reIP := regexp.MustCompile(regexIP)
	rePort := regexp.MustCompile(regexPort)
	reProtocol := regexp.MustCompile(regexProtocol)
	return reIP.MatchString(p.IP) &&
		rePort.MatchString(p.Port) &&
		reProtocol.MatchString(p.Protocol)
}

//CheckConnectivity check if proxy just send a response
func (p Proxy) CheckConnectivity() bool {
	urlParse := fmt.Sprintf("%s://%s:%s", p.Protocol, p.IP, p.Port)
	proxyURL, err := url.Parse(urlParse)
	if err != nil {
		log.Println("Cannot parse Proxy URL ")
		return false
	}
	proxyClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
								Timeout: time.Duration(10*time.Second)}
	request, _ := http.NewRequest("GET", "http://example.com", nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64 AppleWebKit/537.36 (KHTML, like Gecko Chrome/60.0.3112.113 Safari/537.36")
	resp, err := proxyClient.Do(request)
	if err != nil {
		log.Println("Cannot get Proxy response")
		return false
	}
	defer resp.Body.Close()
	return true
}
