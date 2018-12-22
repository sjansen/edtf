package edtf

import (
	"fmt"
	"regexp"
	"strconv"
)

var collapseRE *regexp.Regexp
var dateRE *regexp.Regexp
var dateIdx map[string]int

func init() {
	collapseRE = regexp.MustCompile(`\s+`)
	dateRE = regexp.MustCompile(collapse(`
	  ^
	  (?P<simpleYear>-?[0-9]{4})
	  (?:
	    -
	    (?P<monthOrSeason>[0-9]{2})
	    (?:
	      -
	      (?P<day>[0-9]{2})
	    )?
	  )?
	  |
	  Y
	  (?P<complexYear>-?[0-9]+)
	  (?:
	    E
	    (?P<exponent>[0-9]+)
	  )?
	  (?:
	    S
	    (?P<sigdigits>[0-9]+)
	  )?
	  $
	`))
	names := dateRE.SubexpNames()
	dateIdx = make(map[string]int, len(names))
	for idx, name := range names {
		dateIdx[name] = idx
	}
}

func collapse(src string) string {
	return collapseRE.ReplaceAllString(src, "")
}

func ParseDate(s string) (d *Date, err error) {
	m := dateRE.FindStringSubmatch(s)
	if len(m) == 0 {
		err = fmt.Errorf("Invalid date: %q", s)
		return
	}

	d = &Date{
		Day: parseUint8(m[dateIdx["day"]]),
	}

	if s := m[dateIdx["simpleYear"]]; s != "" {
		d.Year = parseInt16(s)
	} else if s := m[dateIdx["complexYear"]]; s != "" {
		d.Year = parseInt16(s)
		d.Exponent = parseUint8(m[dateIdx["exponent"]])
		d.SigDigits = parseUint8(m[dateIdx["sigdigits"]])
	}

	monthOrSeason := parseUint8(m[dateIdx["monthOrSeason"]])
	switch {
	case monthOrSeason >= 21:
		d.Season = monthOrSeason
	default:
		d.Month = monthOrSeason
	}

	return
}

func parseInt16(s string) int16 {
	i, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		return 0
	}
	return int16(i)
}

func parseUint8(s string) uint8 {
	i, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return 0
	}
	return uint8(i)
}
