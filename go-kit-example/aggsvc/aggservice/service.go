package aggservice

import (
	"context"
	"os"

	"github.com/Ali-Assar/car-rental-system/types"
	"github.com/go-kit/kit/log"
)

const basePrice = 3.5

type Service interface {
	Aggregate(context.Context, types.Distance) error
	Calculate(context.Context, int) (*types.Invoice, error)
}

type Storer interface {
	Insert(types.Distance) error
	Get(int) (float64, error)
}

type BasicService struct {
	store Storer
}

func newBasicService(store Storer) Service {
	return &BasicService{
		store: store,
	}
}

func (svc *BasicService) Aggregate(_ context.Context, dist types.Distance) error {
	return svc.store.Insert(dist)
}
func (svc *BasicService) Calculate(_ context.Context, obuID int) (*types.Invoice, error) {
	dist, err := svc.store.Get(obuID)
	if err != nil {
		return nil, err
	}
	invoice := &types.Invoice{
		OBUID:         obuID,
		TotalDistance: dist,
		TotalAmount:   basePrice * dist,
	}
	return invoice, nil
}

// New will construct a complete micro service
// with logging and instrumentation middleware
func New() Service {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "service", "aggregator")
	}
	var svc Service
	{
		svc = newBasicService(NewMemoryStore())
		svc = newLoggingMiddleware(logger)(svc)
		svc = newInstrumentationMiddleware()(svc)
	}
	return svc
}
