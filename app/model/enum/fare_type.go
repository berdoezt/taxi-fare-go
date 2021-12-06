package enum

// FareType is type of fare
type FareType int

// Below is the current known fare type
const (
	Base FareType = iota
	UpTo
	Over
)

// String convert FareType to string
func (f FareType) String() string {
	switch f {
	case Base:
		return "base"
	case UpTo:
		return "up_to"
	case Over:
		return "over"
	default:
		return "unknown"
	}
}
