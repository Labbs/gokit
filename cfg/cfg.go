package cfg

import (
	"log"

	"github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v2"
)

var (
	Config map[string][]string
)

func InitCfg(address string, debug bool, service string) {
	for {
		t := make(map[string][]string)
		c, _ := api.NewClient(&api.Config{Address: address})
		kv := c.KV()
		consulConfig, _, err := kv.Get("config/"+service, nil)
		if err != nil {
			log.Fatalf(err.Error())
		}

		if consulConfig != nil {
			err = yaml.Unmarshal([]byte(consulConfig.Value), &t)
			if err != nil {
				log.Fatalf(err.Error())
			}
			Config = t
			if debug {
				log.Printf("config reload")
			}
		} else {
			log.Printf("config empty")
		}
	}
}
