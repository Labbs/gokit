package grpc

import {
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/Labbs/gokit/cfg"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
}

func InitGrpc(port string) (net.Listener, *grpc.Server) {
	cfg.Logger.Info("start server")
	listener, err := net.Listen("tcp", port)
	if err != nil {
		cfg.Logger.Fatal(err.Error())
	}

	opts := []grpc_zap.Option{
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
			return zap.Int64("grpc.time_ns", duration.Nanoseconds())
		}),
	}
	grpc_zap.ReplaceGrpcLogger(cfg.Logger)

	srv := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(cfg.Logger, opts...),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(cfg.Logger, opts...),
		)),
	)

	return (listener, srv)
}