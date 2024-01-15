package main

import (
	"context"

	"github.com/Ali-Assar/car-rental-system/types"
)

type GRPcAggregatorServer struct {
	types.UnimplementedAggregatorServer
	svc Aggregator
}

func NewGRPCAggregatorServer(svc Aggregator) *GRPcAggregatorServer {
	return &GRPcAggregatorServer{
		svc: svc,
	}
}

func (s *GRPcAggregatorServer) AggregateD(ctx context.Context, req *types.AggregateRequest) (*types.None, error) {
	distance := types.Distance{
		OBUID: int(req.ObuID),
		Value: req.Value,
		Unix:  req.Unix,
	}
	return &types.None{}, s.svc.AggregateDistance(distance)
}
