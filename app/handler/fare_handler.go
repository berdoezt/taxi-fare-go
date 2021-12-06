package handler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/berdoezt/taxi-fare-go/app/model"
	"github.com/berdoezt/taxi-fare-go/app/service"
	"github.com/berdoezt/taxi-fare-go/app/validation"
	"github.com/sirupsen/logrus"
)

type FareHandler struct {
	fareService   service.FareService
	fareValidator validation.Validator
}

type FareHandlerOpt func(*FareHandler)

func WithService(service service.FareService) FareHandlerOpt {
	return func(fh *FareHandler) {
		fh.fareService = service
	}
}

func WithValidator(validator validation.Validator) FareHandlerOpt {
	return func(fh *FareHandler) {
		fh.fareValidator = validator
	}
}

func NewFareHandler(opts ...FareHandlerOpt) *FareHandler {
	handler := &FareHandler{}

	for _, opt := range opts {
		opt(handler)
	}

	return handler
}

func (c *FareHandler) Handle() {
loop:
	for {
		var distanceMeters []model.DistanceMeter

		reader := bufio.NewReader(os.Stdin)
		inputs := make([]string, 0)

		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				logrus.Error(err)
				fmt.Println(0)
				break loop
			}

			strLine := string(line)

			if strLine == "-1" {
				break
			}

			inputs = append(inputs, strLine)
		}

		err := c.fareValidator.Validate(inputs)
		if err != nil {
			logrus.Error(err)
			fmt.Println(0)
			continue
		}

		for _, input := range inputs {
			distanceMeter, err := c.toDistanceMeter(input)
			if err != nil {
				logrus.Error(err)
				fmt.Println(0)
				break loop
			}
			distanceMeters = append(distanceMeters, distanceMeter)
		}

		result, err := c.fareService.GetFareMeter(distanceMeters)
		if err != nil {
			logrus.Error(err)
			fmt.Println(0)
			continue
		}

		fmt.Println(int(result))
	}
}

func (c *FareHandler) toDistanceMeter(line string) (model.DistanceMeter, error) {
	var distanceMeter model.DistanceMeter

	splitLine := strings.Split(line, " ")
	mileage, err := strconv.ParseFloat(splitLine[1], 64)
	if err != nil {
		return distanceMeter, err
	}

	distanceMeter.
		WithElapsedTime(splitLine[0]).
		WithMileage(mileage)

	return distanceMeter, nil
}
