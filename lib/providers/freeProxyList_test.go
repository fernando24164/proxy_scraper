package providers

import "testing"

func TestCreation(t *testing.T) {
	proxyProvider := New()
	if proxyProvider == nil {
		t.Log("New method is incorrect")
		t.Fail()
	}
}
func TestSetRequest(t *testing.T) {
	proxyProvider := New()
	proxyProvider.SetRequest()
	if proxyProvider.request == nil {
		t.Log("Set request method is incorrect")
		t.Fail()
	}
}
func TestSetDataResponse(t *testing.T) {
	proxyProvider := New()
	proxyProvider.SetRequest()
	proxyProvider.SetDataResponse()
	if proxyProvider.dataResponse == nil {
		t.Log("Set data response method is incorrect")
		t.Fail()
	}
}
