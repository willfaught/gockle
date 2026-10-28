package gockle

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/gocql/gocql"
)

func TestQuery(t *testing.T) {
	var s = newSession(t)
	defer s.Close()
	var exec = func(q string) {
		if err := s.Exec(q); err != nil {
			t.Fatalf("Actual error %v, expected no error", err)
		}
	}
	exec(ksDropIf)
	exec(ksCreate)
	defer exec(ksDrop)
	exec(tabCreate)
	defer exec(tabDrop)
	exec(rowInsert)
	q := s.Query("select * from gockle_test.test").
		PageSize(3).
		WithContext(context.Background()).
		PageState(nil)
	defer q.Release()
	if err := q.Exec(); err != nil {
		t.Fatalf("Actual error %v, expected no error", err)
	}
	// Scan
	var id, n int
	if err := q.Scan(&id, &n); err == nil {
		if id != 1 {
			t.Errorf("Actual id %v, expected 1", id)
		}
		if n != 2 {
			t.Errorf("Actual n %v, expected 2", n)
		}
	} else {
		t.Errorf("Actual error %v, expected no error", err)
	}
	// MapScan
	var actual = map[string]any{}
	var expected = map[string]any{"id": 1, "n": 2}
	if err := q.MapScan(actual); err == nil {
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Actual map %v, expected %v", actual, expected)
		}
	} else {
		t.Errorf("Actual error %v, expected no error", err)
	}
}

func TestQueryMock(t *testing.T) {
	var m, e = &QueryMock{}, fmt.Errorf("e")
	ctx := context.Background()
	it := &IteratorMock{}
	rt := &gocql.SimpleRetryPolicy{NumRetries: 3}
	testMock(t, m, &m.Mock, []struct {
		method    string
		arguments []any
		results   []any
	}{
		{"PageSize", []any{1}, []any{m}},
		{"WithContext", []any{ctx}, []any{m}},
		{"PageState", []any{[]byte{1}}, []any{m}},
		{"Scan", []any{[]any(nil)}, []any{e}},
		{"Scan", []any{[]any{1}}, []any{nil}},
		{"Exec", nil, []any{e}},
		{"Exec", nil, []any{nil}},
		{"Iter", nil, []any{it}},
		{"MapScan", []any{map[string]any(nil)}, []any{nil}},
		{"MapScan", []any{map[string]any{"a": 1}}, []any{e}},
		{"MapScanCAS", []any{map[string]any{"a": 1}}, []any{true, e}},
		{"Release", nil, nil},
		{"GetConsistency", nil, []any{gocql.Quorum}},
		{"SetConsistency", []any{gocql.One}, nil},
		{"Consistency", []any{gocql.One}, []any{m}},
		{"SerialConsistency", []any{gocql.LocalSerial}, []any{m}},
		{"RetryPolicy", []any{rt}, []any{m}},
	})
}

func TestQueryConsistency(t *testing.T) {
	var s = newSession(t)
	defer s.Close()
	q := s.Query("select * from gockle_test.test")
	actual := q.GetConsistency()
	if gocql.Quorum != actual {
		t.Errorf("Actual consistency %s, expected %s", actual, gocql.Quorum)
	}

	q.SetConsistency(gocql.One)
	actual = q.GetConsistency()
	if gocql.One != actual {
		t.Errorf("Actual consistency %s, expected %s", actual, gocql.One)
	}
}
