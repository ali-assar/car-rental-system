package main

import (
	"context"

	"github.com/Ali-Assar/car-rental-system/types"
)

type Middleware func(Service) Service

type LoggingMiddleware struct {
	next Service
}

func newLoggingMiddleware() Middleware {
	return func(next Service) Service {
		return &LoggingMiddleware{
			next: next,
		}
	}
}

func (mw *LoggingMiddleware) Aggregate(_ context.Context, dist types.Distance) error {
	return nil
}

func (mw *LoggingMiddleware) Calculate(_ context.Context, obuID int) (*types.Invoice, error) {
	return nil, nil
}

type InstrumentationMiddleware struct {
	next Service
}

func newInstrumentationMiddleware() Middleware {
	return func(next Service) Service {
		return &InstrumentationMiddleware{
			next: next,
		}
	}
}

func (mw *InstrumentationMiddleware) Aggregate(_ context.Context, dist types.Distance) error {
	return nil
}

func (mw *InstrumentationMiddleware) Calculate(_ context.Context, obuID int) (*types.Invoice, error) {
	return nil, nil
}
