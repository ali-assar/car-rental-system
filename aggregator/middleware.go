package main

import (
	"time"

	"github.com/Ali-Assar/car-rental-system/types"
	"github.com/sirupsen/logrus"
)

type logMiddleware struct {
	next Aggregator
}

func NewLogMiddleware(next Aggregator) Aggregator {
	return &logMiddleware{
		next: next,
	}
}

func (l *logMiddleware) AggregateDistance(distance types.Distance) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took":  time.Since(start),
			"error": err,
		}).Info("aggregate distance")
	}(time.Now())
	err = l.next.AggregateDistance(distance)
	return
}

func (l *logMiddleware) CalculateInvoice(obuID int) (invoice *types.Invoice, err error) {
	defer func(start time.Time) {
		var (
			distance float64
			amount   float64
		)
		if invoice != nil {
			distance = invoice.TotalDistance
			amount = invoice.TotalAmount
		}
		logrus.WithFields(logrus.Fields{
			"took":          time.Since(start),
			"error":         err,
			"obuID":         obuID,
			"totalDistance": distance,
			"totalAmount":   amount,
		}).Info("Calculate invoice")
	}(time.Now())
	invoice, err = l.next.CalculateInvoice(obuID)
	return
}
