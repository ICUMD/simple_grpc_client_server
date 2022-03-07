package api

import (
	"context"
	"log"
)

type Server struct {
	UnimplementedCommunicationServer
}

func (s *Server) SendData(ctx context.Context, msg *ExchangeData) (*ExchangeData, error) {
	log.Printf("Message to server: %s", msg.CurrData)
	return &ExchangeData{CurrData: "Received"}, nil
}
