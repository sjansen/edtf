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

func TestIsLeapYear(t *testing.T) {
	require := require.New(t)

	require.Equal(false, isLeapYear(1900))
	require.Equal(true, isLeapYear(1996))
	require.Equal(false, isLeapYear(1997))
	require.Equal(false, isLeapYear(1998))
	require.Equal(false, isLeapYear(1999))
	require.Equal(true, isLeapYear(2000))
	require.Equal(false, isLeapYear(2100))
	require.Equal(true, isLeapYear(2104))
	require.Equal(true, isLeapYear(2400))
}
