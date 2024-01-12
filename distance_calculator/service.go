package main

import (
	"math"

	"github.com/Ali-Assar/car-rental-system/types"
)

type Point struct {
	Lat  float64
	Long float64
}

type CalculatorServicer interface {
	CalculateDistance(types.OBUData) (float64, error)
}

type CalculatorService struct {
	prevPoint Point
}

func NewCalculatorService() CalculatorServicer {
	return &CalculatorService{
		prevPoint: Point{},
	}
}

func (s *CalculatorService) CalculateDistance(data types.OBUData) (float64, error) {
	distance := 0.0
	if s.prevPoint.Lat != 0 || s.prevPoint.Long != 0 {
		distance = calculateDistance(s.prevPoint.Lat, s.prevPoint.Long, data.Lat, data.Long)
	}
	s.prevPoint.Lat = data.Lat
	s.prevPoint.Long = data.Long
	return distance, nil
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
