package err

import "fmt"

type Err struct {
	code    string
	message string
}

func (c Err) Error() string {
	return fmt.Sprintf("%s: %s", c.code, c.message)
}

var (
	ErrInvalidFormat       Err = Err{"100", "invalid format"}
	ErrInvalidTimeInterval Err = Err{"101", "invalid time interval"}
	ErrInvalidDataAmount   Err = Err{"102", "invalid data amount"}
	ErrInvalidMileage      Err = Err{"103", "invalid mileage"}
	ErrInvalidDataType     Err = Err{"104", "invalid data type"}
)
