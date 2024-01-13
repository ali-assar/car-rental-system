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
		}).Info("")
	}(time.Now())
	err = l.next.AggregateDistance(distance)
	return
}
