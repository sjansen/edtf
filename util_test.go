package edtf

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type row struct {
	input string
	err   string
	vals  []string
}

type fixtureReader struct {
	cols []string
	csv  *csv.Reader
}

func newFixtureReader(filename string, cols []string) (*fixtureReader, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	r := &fixtureReader{
		cols: cols,
		csv:  csv.NewReader(f),
	}

	record, err := r.csv.Read()
	if err != nil {
		return nil, err
	}

	if len(record) != len(cols)+2 {
		err := fmt.Errorf("column count mismatch: %q", filename)
		return nil, err
	}
	if record[0] != "input" || record[1] != "error" {
		err := fmt.Errorf("column name mismatch: %q", filename)
		return nil, err
	}
	for i := range cols {
		if record[i+2] != cols[i] {
			err := fmt.Errorf(
				"expected=%q found=%q: %q",
				record[i+2], cols[i], filename,
			)
			return nil, err
		}
	}

	return r, nil
}

func (r *fixtureReader) next() (*row, error) {
	record, err := r.csv.Read()
	if err == io.EOF {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	n := len(r.cols)
	if m := len(record) - 2; n > m {
		n = m
	}
	row := &row{
		input: record[0],
		err:   record[1],
	}
	if n < 0 {
		return row, nil
	}

	row.vals = make([]string, n)
	for i := range row.vals {
		row.vals[i] = record[i+2]
	}
	return row, nil
}
