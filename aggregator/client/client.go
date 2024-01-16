package client

import (
	"context"

	"github.com/Ali-Assar/car-rental-system/types"
)

type Client interface {
	Aggregate(context.Context, *types.AggregateRequest) error
}
