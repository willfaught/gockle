package gockle

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gocql/gocql"
	"github.com/maraino/go-mock"
)

func TestNewSession(t *testing.T) {
	for _, test := range []*gocql.Session{} {
		var a = NewSession(test)

		if e := (session{s: test}); a != e {
			t.Errorf("Actual %v, expected %v", a, e)
		}
	}
}

func TestSessionMock(t *testing.T) {
	var m, e = &SessionMock{}, fmt.Errorf("e")

	testMock(t, m, &m.Mock, []struct {
		method    string
		arguments []interface{}
		results   []interface{}
	}{
		{"Close", nil, nil},
		{"Columns", []interface{}{"", ""}, []interface{}{map[string]gocql.TypeInfo(nil), nil}},
		{"Columns", []interface{}{"a", "b"}, []interface{}{map[string]gocql.TypeInfo{"c": gocql.NativeType{}}, e}},
		{"QueryBatch", []interface{}{BatchKind(0)}, []interface{}{(*batch)(nil)}},
		{"QueryBatch", []interface{}{BatchKind(1)}, []interface{}{&batch{}}},
		{"QueryExecute", []interface{}{"", []interface{}(nil)}, []interface{}{nil}},
		{"QueryExecute", []interface{}{"a", []interface{}{1}}, []interface{}{e}},
		{"QueryIterate", []interface{}{"", []interface{}(nil)}, []interface{}{(*iterator)(nil)}},
		{"QueryIterate", []interface{}{"a", []interface{}{1}}, []interface{}{iterator{}}},
		{"QueryScan", []interface{}{"", []interface{}(nil), []interface{}(nil)}, []interface{}{nil}},
		{"QueryScan", []interface{}{"a", []interface{}{1}, []interface{}{1}}, []interface{}{e}},
		{"QueryScanMap", []interface{}{"", []interface{}(nil), map[string]interface{}(nil)}, []interface{}{nil}},
		{"QueryScanMap", []interface{}{"a", []interface{}{1}, map[string]interface{}{"b": 2}}, []interface{}{e}},
		{"QueryScanMapTransaction", []interface{}{"", []interface{}(nil), map[string]interface{}(nil)}, []interface{}{false, nil}},
		{"QueryScanMapTransaction", []interface{}{"a", []interface{}{1}, map[string]interface{}{"b": 2}}, []interface{}{true, e}},
		{"QuerySliceMap", []interface{}{"", []interface{}(nil)}, []interface{}{[]map[string]interface{}(nil), nil}},
		{"QuerySliceMap", []interface{}{"a", []interface{}{1}}, []interface{}{[]map[string]interface{}{{"b": 2}}, e}},
		{"Tables", []interface{}{""}, []interface{}{[]string(nil), nil}},
		{"Tables", []interface{}{"a"}, []interface{}{[]string{"b"}, e}},
	})
}

func testMock(t *testing.T, i interface{}, m *mock.Mock, tests []struct {
	method    string
	arguments []interface{}
	results   []interface{}
}) {
	var v = reflect.ValueOf(i)

	for _, test := range tests {
		t.Log("Test:", test)
		m.Reset()
		m.When(test.method, test.arguments...).Return(test.results...)

		var vs []reflect.Value

		for _, a := range test.arguments {
			vs = append(vs, reflect.ValueOf(a))
		}

		var method = v.MethodByName(test.method)

		if method.Type().IsVariadic() {
			vs = method.CallSlice(vs)
		} else {
			vs = method.Call(vs)
		}

		var is []interface{}

		for _, v := range vs {
			is = append(is, v.Interface())
		}

		if !reflect.DeepEqual(is, test.results) {
			t.Errorf("Actual %v, expected %v", is, test.results)
		}
	}
}
