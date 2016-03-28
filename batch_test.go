package gockle

import (
	"fmt"
	"testing"
)

func TestBatch(t *testing.T) {
	var s = newSession(t)

	var exec = func(q string) {
		if err := s.QueryExecute(q); err != nil {
			t.Errorf("Actual error %v, expected no error", err)
		}
	}

	exec(ksDropIf)
	exec(ksCreate)
	exec(tabCreate)
	exec(rowInsert)

	// Execute
	var b = s.QueryBatch(BatchKind(0))

	if b == nil {
		t.Error("Actual batch nil, expected not nil")
	}

	b.Query("update gockle_test.test set n = 3 where id = 1 if n = 2")

	if err := b.Execute(); err != nil {
		t.Errorf("Actual error %v, expected no error", err)
	}

	// ExecuteTransaction
	b = s.QueryBatch(BatchKind(0))
	b.Query("update gockle_test.test set n = 4 where id = 1 if n = 3")

	var id, n int

	if b, i, err := b.ExecuteTransaction(&id, &n); err == nil {
		if id != 0 {
			t.Errorf("Actual id %v, expected 0", id)
		}

		if n != 0 {
			t.Errorf("Actual n %v, expected 0", n)
		}

		if !b {
			t.Error("Actual applied false, expected true")
		}

		if i.Scan() {
			t.Error("Actual scan true, expected false")
		}

		if err := i.Close(); err != nil {
			t.Errorf("Actual error %v, expected no error", err)
		}
	} else {
		t.Errorf("Actual error %v, expected no error", err)
	}

	// ExecuteTransactionMap
	b = s.QueryBatch(BatchKind(0))
	b.Query("update gockle_test.test set n = 5 where id = 1 if n = 4")

	var m = map[string]interface{}{}

	if b, i, err := b.ExecuteTransactionMap(m); err == nil {
		if l := len(m); l > 0 {
			t.Errorf("Actual length %v, expected 0", l)
		}

		if !b {
			t.Error("Actual applied false, expected true")
		}

		if i.Scan() {
			t.Error("Actual scan true, expected false")
		}

		if err := i.Close(); err != nil {
			t.Errorf("Actual error %v, expected no error", err)
		}
	} else {
		t.Errorf("Actual error %v, expected no error", err)
	}

	exec(tabDrop)
	exec(ksDrop)

	s.Close()
}

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
