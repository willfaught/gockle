package gockle

import (
	"fmt"
	"testing"
)

func TestIteratorMock(t *testing.T) {
	var m, e = &IteratorMock{}, fmt.Errorf("e")

	testMock(t, m, &m.Mock, []struct {
		method    string
		arguments []interface{}
		results   []interface{}
	}{
		{"Close", nil, []interface{}{nil}},
		{"Close", nil, []interface{}{e}},
		{"Scan", []interface{}{[]interface{}(nil)}, []interface{}{false}},
		{"Scan", []interface{}{[]interface{}{1}}, []interface{}{true}},
		{"ScanMap", []interface{}{map[string]interface{}(nil)}, []interface{}{false}},
		{"ScanMap", []interface{}{map[string]interface{}{"a": 1}}, []interface{}{true}},
	})
}
