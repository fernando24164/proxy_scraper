package proxy

type Proxy struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
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
