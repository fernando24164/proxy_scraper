package providers

import "testing"

func TestCreation(t *testing.T) {
	proxyProvider := New()
	if proxyProvider == nil {
		t.Log("New method is incorrect")
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
