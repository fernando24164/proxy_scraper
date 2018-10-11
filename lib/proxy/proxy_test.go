package proxy

import (
	"testing"
)

func TestNewMethod(t *testing.T) {
	proxy := New()
	if proxy == nil {
		t.Log("Incorrect new method")
		t.Fail()
	}
}

func TestSetIP(t *testing.T) {
	proxy := New()
	proxy.SetIP("192.168.1.1")
	if proxy.IP != "192.168.1.1" {
		t.Log("Incorrect set ip method")
		t.Fail()
	}
}

func TestSetPort(t *testing.T) {
	proxy := New()
	proxy.SetPort("7852")
	if proxy.Port != "7852" {
		t.Log("Incorrect set port method")
		t.Fail()
	}
}

func TestSetProtocol(t *testing.T) {
	proxy := New()
	proxy.SetProtocol("http")
	if proxy.Protocol != "http" {
		t.Log("Incorrect set protocol method")
		t.Fail()
	}
}

func TestProxiesList(t *testing.T) {
	pl := NewList()
	if pl == nil {
		t.Log("Incorrect creation method")
		t.Fail()
	}
}

func TestProxiesAddProxies(t *testing.T) {
	pl := NewList()
	proxy := New()
	proxy.SetIP("192.168.1.1")
	proxy.SetPort("8080")
	pl.AddProxy(proxy)
	if len(*pl) != 1 {
		t.Log("Incorrect Add proxies method")
		t.Fail()
	}
}
