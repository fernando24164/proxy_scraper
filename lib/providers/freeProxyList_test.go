package providers

import (
	"bufio"
	"log"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func loadFixtureFile(proxy *proxyProviderFreeProxyList) {
	file, err := os.Open("./fixtures/tableProxies.html")
	if err != nil {
		log.Fatal("It was not possible to read fixture file")
	}

	defer file.Close()

	proxy.dataResponse,_ = html.Parse(bufio.NewReader(file))
}

func TestCreation(t *testing.T) {
	proxyProvider := New()
	if proxyProvider == nil {
		t.Log("New method is incorrect")
		t.Fail()
	}
}

func TestGetHeaders(t *testing.T) {
	proxyProvider := New()
	loadFixtureFile(proxyProvider)
	proxyProvider.SetTableHeaders()
	if proxyProvider.headers[0] != "ip address" {
		t.Fail()
	}
}

func TestGetHeaderIndex(t *testing.T) {
	proxyProvider := New()
	loadFixtureFile(proxyProvider)
	proxyProvider.SetTableHeaders()
	indexCalculated := proxyProvider.GetHeaderIndex("Ip Address")
	if indexCalculated != 0 {
		t.Fail()
	}
}

func TestSetMapHeaders(t *testing.T) {
	proxyProvider := New()
	loadFixtureFile(proxyProvider)
	proxyProvider.SetTableHeaders()
	proxyProvider.SetMapHeaders("ip address", "port", "protocol")
	if proxyProvider.indexedHeaders["ip address"] != 0 {
		t.Fail()
	}
}

func TestCreationDataResponse(t *testing.T) {
	proxyProvider := New()
	proxyProvider.SetDataResponse()
	if proxyProvider.dataResponse == nil {
		t.Log("Method is incorrect")
		t.Fail()
	}
}

func TestGetIP(t *testing.T) {
	proxyProvider := New()
	loadFixtureFile(proxyProvider)
	proxyProvider.SetTableHeaders()
	proxyProvider.SetMapHeaders("ip address", "port", "https")
	ip := proxyProvider.GetIP()
	if ip != "118.174.232.181" {
		t.Fail()
	}
}

func TestGetPort(t *testing.T) {
	proxyProvider := New()
	loadFixtureFile(proxyProvider)
	proxyProvider.SetTableHeaders()
	proxyProvider.SetMapHeaders("ip address", "port", "https")
	port := proxyProvider.GetPort()
	if port != "38659" {
		t.Fail()
	}
}

func TestGetProtocol(t *testing.T) {
	proxyProvider := New()
	loadFixtureFile(proxyProvider)
	proxyProvider.SetTableHeaders()
	proxyProvider.SetMapHeaders("ip address", "port", "https")
	protocol := proxyProvider.GetProtocol()
	if protocol != "https" {
		t.Fail()
	}
}

func TestGetProxyList(t *testing.T) {
	proxyProvider := New()
	loadFixtureFile(proxyProvider)
	proxyProvider.SetTableHeaders()
	proxyProvider.SetMapHeaders("ip address", "port", "https")
	proxyList := proxyProvider.GetProxiesList()
	if len(*proxyList) == 0 {
		t.Fail()
	}
}
