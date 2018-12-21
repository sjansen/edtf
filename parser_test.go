package edtf

import (
	"io"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	assert := assert.New(t)

	r, err := newFixtureReader("testdata/date.csv", []string{
		"year", "month", "day",
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

		actual, err := parse(row.input)
		if row.err != "" {
			assert.EqualError(err, row.err)
		} else {
			expected := &Date{
				Year:  int16(row.vals[0]),
				Month: uint8(row.vals[1]),
				Day:   uint8(row.vals[2]),
			}
			assert.NoError(err)
			assert.Equal(expected, actual, row.input)
		}
	}
}
