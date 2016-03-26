package gockle

import (
	"fmt"
	"testing"
)

func TestBatchMock(t *testing.T) {
	var m, e = &BatchMock{}, fmt.Errorf("e")

	testMock(t, m, &m.Mock, []struct {
		method    string
		arguments []interface{}
		results   []interface{}
	}{
		{"Execute", nil, []interface{}{nil}},
		{"Execute", nil, []interface{}{e}},
		{"ExecuteTransaction", []interface{}{[]interface{}(nil)}, []interface{}{false, (*iterator)(nil), nil}},
		{"ExecuteTransaction", []interface{}{[]interface{}{1}}, []interface{}{true, &iterator{}, e}},
		{"ExecuteTransactionMap", []interface{}{map[string]interface{}(nil)}, []interface{}{false, (*iterator)(nil), nil}},
		{"ExecuteTransactionMap", []interface{}{map[string]interface{}{"a": 1}}, []interface{}{true, &iterator{}, e}},
	})
}
