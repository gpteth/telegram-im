package grpc

import (
	"google.golang.org/grpc"

	"open.chat/app/bots/botfather/internal/service"
	"open.chat/app/bots/botpb"
	"open.chat/pkg/grpc_util/server"
)

func New(appId string, svc *service.Service) *server.RPCServer {
	return server.NewRpcServer(appId, func(s *grpc.Server) {
		botpb.RegisterRPCBotsServer(s, svc)
	})
}
