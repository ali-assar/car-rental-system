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
	points []Point
}

func NewCalculatorService() *CalculatorService {
	return &CalculatorService{
		points: make([]Point, 0),
	}
}

func (s *CalculatorService) CalculateDistance(data types.OBUData) (float64, error) {
	distance := 0.0
	if len(s.points) > 0 {
		prevPoint := s.points[len(s.points)-1]
		distance = calculateDistance(prevPoint.Lat, prevPoint.Long, data.Lat, data.Long)
	}
	s.points = append(s.points, Point{Lat: data.Lat, Long: data.Long})
	return distance, nil
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
