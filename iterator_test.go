package gockle

import (
	"fmt"
	"testing"
)

func TestIteratorMock(t *testing.T) {
	var m = &IteratorMock{}

	for _, test := range []error{nil, fmt.Errorf("a")} {
		t.Log("Test:", test)
		m.Reset()
		m.When("Close").Return(test)

		if v := m.Close(); v != test {
			t.Errorf("Actual %v, expected %v", v, test)
		}
	}

	for _, test := range []struct {
		in  []interface{}
		out bool
	}{
		{nil, false},
		{[]interface{}{1}, true},
	} {
		t.Log("Test:", test)
		m.Reset()
		m.When("Scan", test.in).Return(test.out)

		if v := m.Scan(test.in...); v != test.out {
			t.Errorf("Actual %v, expected %v", v, test.out)
		}
	}

	for _, test := range []struct {
		in  map[string]interface{}
		out bool
	}{
		{nil, false},
		{map[string]interface{}{"a": "b"}, true},
	} {
		t.Log("Test:", test)
		m.Reset()
		m.When("ScanMap", test.in).Return(test.out)

		if v := m.ScanMap(test.in); v != test.out {
			t.Errorf("Actual %v, expected %v", v, test.out)
		}
	}
}
