package main

import (
	"fmt"
	"proxy_scraper/lib/providers"
)

func main(){
	freeListProviders := providers.New()
	freeListProviders.SetDataResponse()
	proxiesList := freeListProviders.GetProxiesList()
	proxiesList.ToJSON("./proxies.json")
	fmt.Println(proxiesList)
}
