package main

import "github.com/Ali-Assar/car-rental-system/types"

type GRPcAggregatorServer struct {
	types.UnimplementedAggregatorServer
	svc Aggregator
}

func NewGRPCAggregatorServer(svc Aggregator) *GRPcAggregatorServer {
	return &GRPcAggregatorServer{
		svc: svc,
	}
}

func (s *GRPcAggregatorServer) AggregateDistance(req *types.AggregateRequest) error {
	distance := types.Distance{
		OBUID: int(req.ObuID),
		Value: req.Value,
		Unix:  req.Unix,
	}
	return s.svc.AggregateDistance(distance)
}
