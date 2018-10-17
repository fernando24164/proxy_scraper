package providers

type Provider interface {
	GetIp() string
	GetPort() string
	GetProtocol() string
	GetSource() string
}
