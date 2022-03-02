package api

import (
	"context"
	"log"
)

type Server struct {
	UnimplementedStreamServer
}

func (s *Server) SendData(ctx context.Context, msg *StreamingData) (*StreamingData, error) {
	log.Printf("Message to server: %s", msg.CurrData)
	return &StreamingData{CurrData: "Received"}, nil
}
