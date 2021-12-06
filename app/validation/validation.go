package validation

//go:generate mockgen -destination ./mockvalidation/mock_validation.go -package mockvalidation github.com/berdoezt/taxi-fare-go/app/validation Validator
type Validator interface {
	Validate(data interface{}) error
}
