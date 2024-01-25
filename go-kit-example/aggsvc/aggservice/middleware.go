package aggservice

import (
	"context"
	"time"

	"github.com/Ali-Assar/car-rental-system/types"
	"github.com/go-kit/log"
)

type Middleware func(Service) Service

type LoggingMiddleware struct {
	log  log.Logger
	next Service
}

func newLoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &LoggingMiddleware{
			next: next,
			log:  logger,
		}
	}
}

func (mw *LoggingMiddleware) Aggregate(ctx context.Context, dist types.Distance) (err error) {
	defer func(start time.Time) {
		mw.log.Log("took", time.Since(start), "obu", dist.OBUID, "distance", dist.Value, "err", err)
	}(time.Now())
	err = mw.next.Aggregate(ctx, dist)
	return
}

func (mw *LoggingMiddleware) Calculate(ctx context.Context, obuID int) (inv *types.Invoice, err error) {
	defer func(start time.Time) {
		mw.log.Log("took", time.Since(start), "obuID", obuID, "invoice", inv, "err", err)
	}(time.Now())
	inv, err = mw.next.Calculate(ctx, obuID)
	return
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

func (mw *InstrumentationMiddleware) Aggregate(ctx context.Context, dist types.Distance) error {
	return mw.next.Aggregate(ctx, dist)
}

func (mw *InstrumentationMiddleware) Calculate(ctx context.Context, obuID int) (*types.Invoice, error) {
	return mw.next.Calculate(ctx, obuID)
}
