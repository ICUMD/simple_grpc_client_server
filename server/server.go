package serverPkg

import (
	"context"
	"log"

	__ "github.com/ICUMD/simple_grpc_client_server/server_stub"
)

type Server struct {
}

func (s *Server) SendData(ctx context.Context, msg *__.StreamingData) {
	log.Printf("Message tp server: %s", msg.CurrData)
}
