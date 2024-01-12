package main

import (
	"time"

	"github.com/Ali-Assar/car-rental-system/types"
	"github.com/sirupsen/logrus"
)

type logMiddleware struct {
	next CalculatorServicer
}

func NewLogMiddleware(next CalculatorServicer) CalculatorServicer {
	return &logMiddleware{
		next: next,
	}
}

func (l *logMiddleware) CalculateDistance(data types.OBUData) (dist float64, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took":     time.Since(start),
			"error":    err,
			"distance": dist,
		}).Info("calculate distance")
	}(time.Now())
	dist, err = l.next.CalculateDistance(data)
	if err != nil {
		return 0.0, err
	}
	return dist, nil

}
