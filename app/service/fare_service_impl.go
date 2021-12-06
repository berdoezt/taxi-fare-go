package service

import (
	"github.com/berdoezt/taxi-fare-go/app/model"
	"github.com/berdoezt/taxi-fare-go/app/model/enum"
	"github.com/berdoezt/taxi-fare-go/config"
)

type FareServiceImpl struct {
	fareRules map[string]config.FareRule
}

type FareServiceImplOpt func(*FareServiceImpl)

func WithFareRules(fareRules map[string]config.FareRule) FareServiceImplOpt {
	return func(fsi *FareServiceImpl) {
		fsi.fareRules = fareRules
	}
}

func NewFareServiceImpl(opts ...FareServiceImplOpt) *FareServiceImpl {
	service := &FareServiceImpl{}

	for _, opt := range opts {
		opt(service)
	}

	return service
}

var _ FareService = (*FareServiceImpl)(nil)

func (f *FareServiceImpl) GetFareMeter(distanceMeters []model.DistanceMeter) (float64, error) {
	fare := f.fareRules[enum.Base.String()].Price
	mileage := (distanceMeters[len(distanceMeters)-1].GetMileage() - distanceMeters[0].GetMileage())

	if mileage <= f.fareRules[enum.Base.String()].Distance {
		return fare, nil
	}

	mileage = mileage - f.fareRules[enum.Base.String()].Distance
	remainderDistanceRule := f.fareRules[enum.UpTo.String()].DistanceThreshold - f.fareRules[enum.Base.String()].Distance

	if mileage <= remainderDistanceRule {
		fare = fare + (mileage/f.fareRules[enum.UpTo.String()].Distance)*f.fareRules[enum.UpTo.String()].Price
		return fare, nil
	}

	fare = fare + (remainderDistanceRule/f.fareRules[enum.UpTo.String()].Distance)*f.fareRules[enum.UpTo.String()].Price
	mileage = mileage - remainderDistanceRule

	fare = fare + (mileage/f.fareRules[enum.Over.String()].Distance)*f.fareRules[enum.Over.String()].Price
	return fare, nil
}
