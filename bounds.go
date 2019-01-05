package edtf

import "math"

func isLeapYear(year int64) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		}
		return true
	}
	return false
}

func (d *Date) LowerStrict() *Date {
	year := d.Year
	if d.Exponent > 0 {
		year *= int64(math.Pow10(int(d.Exponent)))
	}

	month := d.Month
	if month < 1 {
		month = 1
	}

	day := d.Day
	if day < 1 {
		day = 1
	}

	return &Date{Year: year, Month: month, Day: day}
}

func (d *Date) UpperStrict() *Date {
	year := d.Year
	if d.Exponent > 0 {
		year *= int64(math.Pow10(int(d.Exponent)))
	}

	month := d.Month
	if month < 1 {
		month = 12
	}

	day := d.Day
	if day < 1 {
		switch month {
		case 2:
			if isLeapYear(year) {
				day = 29
			} else {
				day = 28
			}
		case 1, 3, 5, 7, 8, 10, 12:
			day = 31
		default:
			day = 30
		}
	}

	return &Date{Year: year, Month: month, Day: day}
}
