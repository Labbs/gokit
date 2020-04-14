package cfg

import "go.uber.org/zap"

var (
	// Config - var contains configuration from AWS SSM Parameter Store
	Config map[string]interface{}
	// Logger - zap logger
	Logger *zap.Logger
)
