package service

import "github.com/berdoezt/taxi-fare-go/app/model"

//go:generate mockgen -destination ./mockservice/mock_fare_service.go -package mockservice github.com/berdoezt/taxi-fare-go/app/service FareService
type FareService interface {
	GetFareMeter(distanceMeters []model.DistanceMeter) (float64, error)
}
