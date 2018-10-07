package proxy

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ProxiesList []*Proxy

func NewList() *ProxiesList {
	return &ProxiesList{}
}

func (pl *ProxiesList) AddProxy(proxy *Proxy) {
	*pl = append(*pl, proxy)
}

func (pl *ProxiesList) ToJSON(fileName string) {
	file, errFile := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if errFile != nil {
		log.Println("Error happened on file creation!")
	}
	defer file.Close()

	proxiesJSON, _ := json.Marshal(pl)
	ioutil.WriteFile(fileName, proxiesJSON, 0755)
}
