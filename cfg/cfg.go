package cfg

import (
	"time"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

var (
	// Config - var contains configuration from consul
	Config map[string]interface{}
	Logger *zap.Logger
)

// InitCfgLoop - get configuration from consul
func InitCfgLoop(address string, debug bool, service string) {
	Logger, _ = zap.NewProduction()
	for {
		c, _ := api.NewClient(&api.Config{Address: address})
		kv := c.KV()
		consulConfig, _, err := kv.Get("config/"+service, nil)
		if err != nil {
			Logger.Fatalf(err.Error())
		}

		if consulConfig != nil {
			err = yaml.Unmarshal([]byte(consulConfig.Value), &Config)
			if err != nil {
				Logger.Fatalf(err.Error())
			}
			if debug {
				Logger.Info("config reloaded")
			}
		} else {
			Logger.Info("config reloaded")
		}

		time.Sleep(30 * time.Second)
	}
}

// InitCfg - get configuration from consul
func InitCfg(address string, debug bool, service string) {
	Logger, _ = zap.NewProduction()
	c, _ := api.NewClient(&api.Config{Address: address})
	kv := c.KV()
	consulConfig, _, err := kv.Get("config/"+service, nil)
	if err != nil {
		Logger.Fatalf(err.Error())
	}
	if consulConfig != nil {
		err = yaml.Unmarshal([]byte(consulConfig.Value), &Config)
		if err != nil {
			Logger.Fatalf(err.Error())
		}
		if debug {
			Logger.Info("config reloaded")
		}
	} else {
		Logger.Info("config reloaded")
	}
}
