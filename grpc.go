package server

import (
	"net"

	"github.com/sirupsen/logrus"
	"github.com/bblfsh/sdk/protocol"
	"github.com/bblfsh/server/runtime"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	*Server
	opts []grpc.ServerOption
}

func NewGRPCServer(r *runtime.Runtime, overrides map[string]string, transport string, maxMessageSize int) *GRPCServer {
	server := NewServer(r, overrides)
	server.Transport = transport

	opts := []grpc.ServerOption{}
	if maxMessageSize != 0 {
		logrus.Infof("setting maximum size for sending and receiving messages to %d", maxMessageSize)
		opts = append(opts, grpc.MaxRecvMsgSize(maxMessageSize))
		opts = append(opts, grpc.MaxSendMsgSize(maxMessageSize))
	}

	return &GRPCServer{server, opts}
}

func (s *GRPCServer) Serve(listener net.Listener) error {
	grpcServer := grpc.NewServer(s.opts...)

	logrus.Debug("registering gRPC service")
	protocol.RegisterProtocolServiceServer(
		grpcServer,
		protocol.NewProtocolServiceServer(),
	)

	protocol.DefaultParser = s.Server

	logrus.Info("starting gRPC server")
	return grpcServer.Serve(listener)
}
