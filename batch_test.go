package gockle

import (
	"fmt"
	"testing"
)

func TestBatchMock(t *testing.T) {
	var m = &BatchMock{}

	for _, test := range []error{nil, fmt.Errorf("a")} {
		t.Log("Test:", test)
		m.Reset()
		m.When("Execute").Return(test)

		if v := m.Execute(); v != test {
			t.Errorf("Actual %v, expected %v", v, test)
		}
	}

	for _, test := range []struct {
		s []interface{}
		b bool
		i Iterator
		e error
	}{
		{nil, false, (*iterator)(nil), nil},
		{[]interface{}{1}, true, iterator{}, fmt.Errorf("a")},
	} {
		t.Log("Test:", test)
		m.Reset()
		m.When("ExecuteTransaction", test.s).Return(test.b, test.i, test.e)

		if b, i, e := m.ExecuteTransaction(test.s...); b != test.b {
			t.Errorf("Actual %v, expected %v", b, test.b)
		} else if i != test.i {
			t.Errorf("Actual %v, expected %v", i, test.i)
		} else if e != test.e {
			t.Errorf("Actual %v, expected %v", e, test.e)
		}
	}

	for _, test := range []struct {
		m map[string]interface{}
		b bool
		i Iterator
		e error
	}{
		{nil, false, (*iterator)(nil), nil},
		{map[string]interface{}{"a": "b"}, true, iterator{}, fmt.Errorf("c")},
	} {
		t.Log("Test:", test)
		m.Reset()
		m.When("ExecuteTransactionMap", test.m).Return(test.b, test.i, test.e)

		if b, i, e := m.ExecuteTransactionMap(test.m); b != test.b {
			t.Errorf("Actual %v, expected %v", b, test.b)
		} else if i != test.i {
			t.Errorf("Actual %v, expected %v", i, test.i)
		} else if e != test.e {
			t.Errorf("Actual %v, expected %v", e, test.e)
		}
	}

	for _, test := range []struct {
		s string
		i []interface{}
	}{
		{"", nil},
		{"a", []interface{}{1}},
	} {
		t.Log("Test:", test)
		m.Reset()
		m.When("Query", test.s, test.i).Return()
		m.Query(test.s, test.i...)
	}
}
