package cfg

import (
	"log"
	"time"

	"github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v2"
)

var (
	// Config - var contains configuration from consul
	Config map[string]interface{}
)

// InitCfg - get configuration from consul
func InitCfg(address string, debug bool, service string) {
	for {
		c, _ := api.NewClient(&api.Config{Address: address})
		kv := c.KV()
		consulConfig, _, err := kv.Get("config/"+service, nil)
		if err != nil {
			log.Fatalf(err.Error())
		}

		if consulConfig != nil {
			err = yaml.Unmarshal([]byte(consulConfig.Value), &Config)
			if err != nil {
				log.Fatalf(err.Error())
			}
			if debug {
				log.Printf("config reloaded")
			}
		} else {
			log.Printf("config empty")
		}

		time.Sleep(30 * time.Second)
	}
}
