package edtf

import (
	"io"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBounds(t *testing.T) {
	r, err := newFixtureReader("testdata/bounds.csv", []string{
		"LowerStrict", "UpperStrict",
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

		t.Run(row.input, func(t *testing.T) {
			require := require.New(t)

			d, err := ParseDate(row.input)
			require.NoError(err, "input")

			expected, err := ParseDate(row.vals[0])
			require.NoError(err, "LowerStrict")
			require.Equal(expected, d.LowerStrict(), "LowerStrict")

			expected, err = ParseDate(row.vals[1])
			require.NoError(err, "LowerStrict")
			require.Equal(expected, d.UpperStrict(), "UpperStrict")
		})
	}
}
