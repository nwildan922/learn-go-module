package handler

import (
	"context"
	"log"
	"time"

	"github.com/nwildan922/learn-go-module/proto/counterpb"
	"github.com/nwildan922/learn-go-module/service"
)

type CounterHandler struct {
	counterpb.UnimplementedCounterServiceServer
	service *service.CounterService
}

func NewCounterHandler(svc *service.CounterService) *CounterHandler {
	return &CounterHandler{service: svc}
}

func (h *CounterHandler) SendCounter(
	ctx context.Context,
	req *counterpb.CounterRequest,
) (*counterpb.CounterResponse, error) {

	log.Println("counter :", req.Counter)

	currentTime := time.Now()
	currentTimeString := currentTime.Format(time.RFC3339)

	// call service
	err := h.service.SaveCounter(ctx, req.Counter, currentTime)
	if err != nil {
		return nil, err
	}

	// response
	return &counterpb.CounterResponse{
		Counter:   req.Counter,
		Timestamp: currentTimeString,
	}, nil
}
