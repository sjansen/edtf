package edtf

import (
	"io"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseDate(t *testing.T) {
	require := require.New(t)

	r, err := newFixtureReader("testdata/date.csv", []string{
		"Year", "Month", "Day",
		"Exponent", "SigDigits", "Season",
	})
	if err != nil {
		log.Fatal(err)
	}

	for {
		row, err := r.next()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		actual, err := ParseDate(row.input)
		if row.err != "" {
			require.EqualError(err, row.err, row.input)
		} else {
			require.NoError(err, row.input)

			expected := &Date{
				Year:      parseInt64(row.vals[0]),
				Month:     parseUint8(row.vals[1]),
				Day:       parseUint8(row.vals[2]),
				Exponent:  parseUint8(row.vals[3]),
				SigDigits: parseUint8(row.vals[4]),
				Season:    parseUint8(row.vals[5]),
			}
			require.Equal(expected, actual, row.input)
		}
	}
}
