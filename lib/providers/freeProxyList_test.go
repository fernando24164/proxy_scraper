package providers

import (
	"bytes"
	"testing"

	"golang.org/x/net/html"
)

func TestCreation(t *testing.T) {
	proxyProvider := New()
	if proxyProvider == nil {
		t.Log("New method is incorrect")
		t.Fail()
	}
}

func TestGetHeaders(t *testing.T) {
	tableHTML := "<table><tr><th>IP</th><th>Port</th><th>Protocol</th></tr></table>"
	bufferingData := bytes.NewBufferString(tableHTML)
	proxyProvider := New()
	proxyProvider.dataResponse, _ = html.Parse(bufferingData)
	proxyProvider.SetTableHeaders()
	if proxyProvider.headers[0] != "ip" {
		t.Fail()
	}
}

func TestGetHeaderIndex(t *testing.T) {
	tableHTML := "<table><tr><th>IP</th><th>Port</th><th>Protocol</th></tr></table>"
	bufferingData := bytes.NewBufferString(tableHTML)
	proxyProvider := New()
	proxyProvider.dataResponse, _ = html.Parse(bufferingData)
	proxyProvider.SetTableHeaders()
	indexCalculated := proxyProvider.GetHeaderIndex("Ip")
	if indexCalculated != 0 {
		t.Fail()
	}
}

func TestSetMapHeaders(t *testing.T) {
	tableHTML := "<table><tr><th>IP</th><th>Port</th><th>Protocol</th></tr></table>"
	bufferingData := bytes.NewBufferString(tableHTML)
	proxyProvider := New()
	proxyProvider.dataResponse, _ = html.Parse(bufferingData)
	proxyProvider.SetTableHeaders()
	proxyProvider.SetMapHeaders("ip", "port", "protocol")
	if proxyProvider.indexedHeaders["ip"] != 0 {
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
	tableHTML := "<html><table><thead><tr><th>IP Address</th><th>Port</th><th>Code</th><th>Country</th><th>Anonymity</th><th>Google</th><th>Https</th><th>Last Checked</th></tr></thead><tbody><tr><td>118.174.232.181</td><td>38659</td><td>TH</td><td>Thailand</td><td>anonymous</td><td>no</td><td>yes</td><td>6 seconds ago</td></tr></tbody></table><html>"
	bufferingData := bytes.NewBufferString(tableHTML)
	proxyProvider := New()
	proxyProvider.dataResponse, _ = html.Parse(bufferingData)
	proxyProvider.SetTableHeaders()
	proxyProvider.SetMapHeaders("ip address", "port", "https")
	ip := proxyProvider.GetIP()
	if ip != "118.174.232.181" {
		t.Fail()
	}
}

func TestGetPort(t *testing.T) {
	tableHTML := "<html><table><thead><tr><th>IP Address</th><th>Port</th><th>Code</th><th>Country</th><th>Anonymity</th><th>Google</th><th>Https</th><th>Last Checked</th></tr></thead><tbody><tr><td>118.174.232.181</td><td>38659</td><td>TH</td><td>Thailand</td><td>anonymous</td><td>no</td><td>yes</td><td>6 seconds ago</td></tr></tbody></table><html>"
	bufferingData := bytes.NewBufferString(tableHTML)
	proxyProvider := New()
	proxyProvider.dataResponse, _ = html.Parse(bufferingData)
	proxyProvider.SetTableHeaders()
	proxyProvider.SetMapHeaders("ip address", "port", "https")
	port := proxyProvider.GetPort()
	if port != "38659" {
		t.Fail()
	}
}

func TestGetProtocol(t *testing.T) {
	tableHTML := "<html><table><thead><tr><th>IP Address</th><th>Port</th><th>Code</th><th>Country</th><th>Anonymity</th><th>Google</th><th>Https</th><th>Last Checked</th></tr></thead><tbody><tr><td>118.174.232.181</td><td>38659</td><td>TH</td><td>Thailand</td><td>anonymous</td><td>no</td><td>yes</td><td>6 seconds ago</td></tr></tbody></table><html>"
	bufferingData := bytes.NewBufferString(tableHTML)
	proxyProvider := New()
	proxyProvider.dataResponse, _ = html.Parse(bufferingData)
	proxyProvider.SetTableHeaders()
	proxyProvider.SetMapHeaders("ip address", "port", "https")
	protocol := proxyProvider.GetProtocol()
	if protocol != "https" {
		t.Fail()
	}
}
