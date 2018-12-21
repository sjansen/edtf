package edtf

import (
	"fmt"
	"regexp"
	"strconv"
)

var collapseRE *regexp.Regexp
var dateRE *regexp.Regexp

func init() {
	collapseRE = regexp.MustCompile(`\s+`)
	dateRE = regexp.MustCompile(collapse(`
	  ^
	  (?P<year>[0-9]{4})
	  (?:
	    -
	    (?P<month>[0-9]{2})
	    (?:
	      -
	      (?P<day>[0-9]{2})
	    )?
	  )?
	  $
	`))
}

func collapse(src string) string {
	return collapseRE.ReplaceAllString(src, "")
}

func parse(s string) (d *Date, err error) {
	m := dateRE.FindStringSubmatch(s)
	if len(m) == 0 {
		err = fmt.Errorf("Invalid date: %q", s)
		return
	}
	d = &Date{
		Year:  int16(parseInt(m[1])),
		Month: uint8(parseInt(m[2])),
		Day:   uint8(parseInt(m[3])),
	}
	return
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
