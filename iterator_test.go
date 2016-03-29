package gockle

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	var s = newSession(t)

	defer s.Close()

	var exec = func(q string) {
		if err := s.QueryExec(q); err != nil {
			t.Fatalf("Actual error %v, expected no error", err)
		}
	}

	exec(ksDropIf)
	exec(ksCreate)

	defer exec(ksDrop)

	exec(tabCreate)

	defer exec(tabDrop)

	exec(rowInsert)

	var i = s.QueryIterator("select * from gockle_test.test")

	if i == nil {
		t.Fatal("Actual iterator nil, expected not nil")
	}

	var id, n int

	if !i.Scan(&id, &n) {
		t.Fatal("Actual more false, expected true")
	}

	if id != 1 {
		t.Fatalf("Actual id %v, expected 1", id)
	}

	if n != 2 {
		t.Fatalf("Actual n %v, expected 2", n)
	}

	if err := i.Close(); err != nil {
		t.Fatalf("Actual error %v, expected no error", err)
	}
}

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
