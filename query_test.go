package gockle

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
	"reflect"
	"testing"
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
	var actual = map[string]interface{}{}
	var expected = map[string]interface{}{"id": 1, "n": 2}
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
	testMock(t, m, &m.Mock, []struct {
		method    string
		arguments []interface{}
		results   []interface{}
	}{
		{"PageSize", []interface{}{1}, []interface{}{m}},
		{"WithContext", []interface{}{ctx}, []interface{}{m}},
		{"PageState", []interface{}{[]byte{1}}, []interface{}{m}},
		{"Scan", []interface{}{[]interface{}(nil)}, []interface{}{e}},
		{"Scan", []interface{}{[]interface{}{1}}, []interface{}{nil}},
		{"Exec", nil, []interface{}{e}},
		{"Exec", nil, []interface{}{nil}},
		{"Iter", nil, []interface{}{it}},
		{"MapScan", []interface{}{map[string]interface{}(nil)}, []interface{}{nil}},
		{"MapScan", []interface{}{map[string]interface{}{"a": 1}}, []interface{}{e}},
		{"Release", nil, nil},
		{"GetConsistency", nil, []interface{}{gocql.Quorum}},
		{"SetConsistency", []interface{}{gocql.One}, nil},
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
