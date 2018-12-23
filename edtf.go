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
	SpringNorthernHemisphere
	SummerNorthernHemisphere
	AutumnNorthernHemisphere
	WinterNorthernHemisphere
	SpringSouthernHemisphere
	SummerSouthernHemisphere
	AutumnSouthernHemisphere
	WinterSouthernHemisphere
	Quarter1
	Quarter2
	Quarter3
	Quarter4
	Quadrimester1
	Quadrimester2
	Quadrimester3
	Semestral1
	Semestral2
)

type Date struct {
	UA        UncertainOrApproximate
	Year      int64
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
