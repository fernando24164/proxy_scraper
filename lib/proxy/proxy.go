package proxy

import "regexp"

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
