package main

import (
	"context"
	"log"
	"time"

	"github.com/Ali-Assar/car-rental-system/aggregator/client"
	"github.com/Ali-Assar/car-rental-system/types"
)

func main() {
	c, err := client.NewGRPCClient(":3001")
	if err != nil {
		log.Fatal(err)
	}
	if err := c.Aggregate(context.Background(), &types.AggregateRequest{
		ObuID: 1,
		Value: 52.4,
		Unix:  time.Now().UnixNano(),
	}); err != nil {
		log.Fatal(err)
	}
}
