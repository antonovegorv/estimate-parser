package estimateparser

import (
	"os"
	"testing"
)

func TestParseFromReader(t *testing.T) {
	file, err := os.Open("./testdata/estimate.xlsx")
	if err != nil {
		t.Fatal(err)
	}

	ep := NewEstimateParser()

	e, err := ep.ParseFromReader(file)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(e)
}
