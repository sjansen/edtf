package edtf

import "math"

func (d *Date) LowerStrict() *Date {
	year := d.Year
	if d.Exponent > 0 {
		year = year * int64(math.Pow10(int(d.Exponent)))
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
		year = year * int64(math.Pow10(int(d.Exponent)))
	}

	month := d.Month
	if month < 1 {
		month = 12
	}

	day := d.Day
	if day < 1 {
		switch month {
		case 2:
			day = 28
		case 1, 3, 5, 7, 8, 10, 12:
			day = 31
		default:
			day = 30
		}
	}

	return &Date{Year: year, Month: month, Day: day}
}