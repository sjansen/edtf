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
		"year", "month", "season", "day",
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
			require.EqualError(err, row.err)
		} else {
			expected := &Date{
				Year:   parseInt16(row.vals[0]),
				Month:  parseUint8(row.vals[1]),
				Season: parseUint8(row.vals[2]),
				Day:    parseUint8(row.vals[3]),
			}
			require.NoError(err)
			require.Equal(expected, actual, row.input)
		}
	}
}
