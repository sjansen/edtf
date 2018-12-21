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

type Season uint8

const (
	Spring Season = 21 + iota
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
	Season    Season
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
