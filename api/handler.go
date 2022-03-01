package api

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) SendData(ctx context.Context, msg *StreamingData) (*StreamingData, error) {
	log.Printf("Message tp server: %s", msg.CurrData)
	return &StreamingData{CurrData: "Received"}, nil
}
