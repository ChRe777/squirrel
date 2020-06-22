package generator

import (
	"testing"
)

func TestError(t *testing.T) {

	specs := []struct {
		msg    string
		params []interface{}
		want   string
	}{
		{"Not found: %v", []interface{}{"b"}, "Error: \"Not found: b\""},
		{"Not found: %v %v", []interface{}{"b", "c"}, "Error: \"Not found: b c\""},
		{"Not found", []interface{}{}, "Error: \"Not found\""},
	}

	for _, spec := range specs {

		got := Err(spec.msg, spec.params...).Val
		if got != spec.want {
			t.Errorf("TestError got: %v, want: %v \n", got, spec.want)
		}

	}
}
