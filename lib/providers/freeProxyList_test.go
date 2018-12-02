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
	if proxyProvider.headers[0] != "IP" {
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
	proxyProvider.SetMapHeaders()
	if proxyProvider.indexedHeaders["IP"] != 0 {
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
