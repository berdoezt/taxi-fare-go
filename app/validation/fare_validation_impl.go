package validation

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	e "github.com/berdoezt/taxi-fare-go/app/err"
)

var (
	inputPattern = `^\d\d:\d\d:\d\d.\d\d\d\s\d+\.?\d*$`
	timeFormat   = "15:04:05.000"
	timeInterval = 300
)

func inputLengthValid(inputs []string) bool {
	return len(inputs) >= 2
}

func intervalValid(current time.Time, prev time.Time) bool {
	if current.Before(prev) {
		return false
	}

	return current.Sub(prev) <= time.Duration(timeInterval)*time.Second
}

func mileageValid(current float64, prev float64) bool {
	if current == 0 {
		return false
	}

	return current >= prev
}

func inputFormatValid(input string) bool {
	r := regexp.MustCompile(inputPattern)
	return r.Match([]byte(input))
}

type FareValidation struct {
}

var _ Validator = (*FareValidation)(nil)

func NewFareValidation() *FareValidation {
	return &FareValidation{}
}

func (f *FareValidation) Validate(data interface{}) error {
	inputs, ok := data.([]string)
	if !ok {
		return e.ErrInvalidDataType
	}
	if !inputLengthValid(inputs) {
		return e.ErrInvalidDataAmount
	}

	var prevTime time.Time
	var prevMileage float64

	for idx, input := range inputs {
		if !inputFormatValid(input) {
			return e.ErrInvalidFormat
		}

		splitString := strings.Split(input, " ")
		t, err := time.Parse(timeFormat, splitString[0])
		if err != nil {
			return err
		}

		m, err := strconv.ParseFloat(splitString[1], 64)
		if err != nil {
			return err
		}

		if idx == 0 {
			prevTime = t
			prevMileage = m
			continue
		}

		if !intervalValid(t, prevTime) {
			return e.ErrInvalidTimeInterval
		}

		if !mileageValid(m, prevMileage) {
			return e.ErrInvalidMileage
		}

		prevTime = t
		prevMileage = m
	}
	return nil

}
