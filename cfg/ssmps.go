package cfg

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

// InitSSMPSCfg - get configuration from AWS SSM Parameter Store
func InitSSMPSCfg(region string, key string, debug bool) {
	// init logger
	Logger, _ = zap.NewProduction()

	// load aws configuration
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		Logger.Fatal("unable to load SDK config, " + err.Error())
	}
	cfg.Region = region

	// init aws session for AWS SSM
	svc := ssm.New(cfg)

	// create request
	req := svc.GetParameterRequest(&ssm.GetParameterInput{
		Name:           aws.String(key),
		WithDecryption: aws.Bool(true),
	})

	// send request to AWS
	resp, err := req.Send(context.Background())
	if err != nil {
		Logger.Fatal("unable to send request, " + err.Error())
	}
	if *resp.Parameter.Value != "" {
		err = yaml.Unmarshal([]byte(*resp.Parameter.Value), &Config)
		if err != nil {
			Logger.Fatal("unable to unmarshal yaml, " + err.Error())
		}
		if debug {
			Logger.Info("config reloaded from SSM Parameter Store")
		}
	} else {
		Logger.Warn("config empty from SSM Parameter Store")
	}
}

// InitSSMPSCfgLoop - get configuration from AWS SSM Parameter Store
func InitSSMPSCfgLoop(region string, key string, refresh int64, debug bool) {
	Logger, _ = zap.NewProduction()
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		Logger.Fatal("unable to load SDK config, " + err.Error())
	}
	cfg.Region = region

	for {
		svc := ssm.New(cfg)

		req := svc.GetParameterRequest(&ssm.GetParameterInput{
			Name:           aws.String(key),
			WithDecryption: aws.Bool(true),
		})
		resp, err := req.Send(context.Background())
		if err != nil {
			Logger.Fatal("unable to send request, " + err.Error())
		}
		if *resp.Parameter.Value != "" {
			err = yaml.Unmarshal([]byte(*resp.Parameter.Value), &Config)
			if err != nil {
				Logger.Fatal("unable to unmarshal yaml, " + err.Error())
			}
			if debug {
				Logger.Info("config reloaded from SSM Parameter Store")
			}
		} else {
			Logger.Warn("config empty from SSM Parameter Store")
		}

		time.Sleep(time.Duration(refresh) * time.Second)
	}
}
