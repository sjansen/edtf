package edtf

type UncertainOrApproximate byte

const (
	YearApproximate UncertainOrApproximate = 1 << iota
	YearUncertain
	MonthApproximate
	MonthUncertain
	DayApproximate
	DayUncertain
)

const (
	Spring uint8 = 21 + iota
	Summer
	Autumn
	Winter
)

type Date struct {
	UA        UncertainOrApproximate
	Year      int16
	Exponent  uint8
	SigDigits uint8
	Month     uint8
	Season    uint8
	Day       uint8
}

type Time struct {
	UTC     bool
	Hours   uint8
	Minutes uint8
	Seconds uint8
	Offset  int16 // minutes
}

type DateTime struct {
	Date
	Time
}
