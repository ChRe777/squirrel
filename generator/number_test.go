package generator

import (
	"fmt"
	"testing"
)

func TestNumber(t *testing.T) {

	specs := []struct{
		numStr	string
		want 	string
	} {
		{"123.34",  "123.34"},
		{"abc"	 ,  "Error: \"can't convert abc to decimal\""},
		{""	 ,  "Error: \"can't convert  to decimal\""},

	}

	for _, spec := range specs {

		got := fmt.Sprintf("%v", Num(spec.numStr))

		if got != spec.want {
			t.Errorf("TestNumber got: %v, want: %v \n", got, spec.want)
		}

	}
}
