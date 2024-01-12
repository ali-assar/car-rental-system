package main

import (
	"time"

	"github.com/Ali-Assar/car-rental-system/types"
	"github.com/sirupsen/logrus"
)

type logMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) *logMiddleware {
	return &logMiddleware{
		next: next,
	}
}

func (l *logMiddleware) ProduceData(data types.OBUData) error {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"obuID": data.OBUID,
			"lat":   data.Lat,
			"long":  data.Long,
			"took":  time.Since(start),
		}).Info("producing to kafka")
	}(time.Now())
	return l.next.ProduceData(data)
}
