package cfg

import (
	"time"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

// InitConsulCfg - get configuration from consul
func InitConsulCfg(address string, debug bool, service string) {
	Logger, _ = zap.NewProduction()
	c, _ := api.NewClient(&api.Config{Address: address})
	kv := c.KV()
	consulConfig, _, err := kv.Get("config/"+service, nil)
	if err != nil {
		Logger.Fatal(err.Error())
	}
	if consulConfig != nil {
		err = yaml.Unmarshal([]byte(consulConfig.Value), &Config)
		if err != nil {
			Logger.Fatal(err.Error())
		}
		if debug {
			Logger.Info("config reloaded")
		}
	} else {
		Logger.Info("config reloaded")
	}
}

// InitConsulCfgLoop - get configuration from consul
func InitConsulCfgLoop(address string, debug bool, service string, refresh int64) {
	Logger, _ = zap.NewProduction()
	for {
		c, _ := api.NewClient(&api.Config{Address: address})
		kv := c.KV()
		consulConfig, _, err := kv.Get("config/"+service, nil)
		if err != nil {
			Logger.Fatal(err.Error())
		}

		if consulConfig != nil {
			err = yaml.Unmarshal([]byte(consulConfig.Value), &Config)
			if err != nil {
				Logger.Fatal(err.Error())
			}
			if debug {
				Logger.Info("config reloaded")
			}
		} else {
			Logger.Info("config reloaded")
		}

		time.Sleep(time.Duration(refresh) * time.Second)
	}
}
