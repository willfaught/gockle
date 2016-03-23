package gockle

import (
	"github.com/gocql/gocql"
	"github.com/maraino/go-mock"
)

// Iterator wraps *gocql.Iter.
type Iterator interface {
	// Close closes the iterator.
	Close() error

	// Scan puts the current row in results and returns whether there are more
	// rows.
	Scan(results ...interface{}) bool

	// ScanMap puts the current row in results and returns whether there are
	// more rows.
	ScanMap(results map[string]interface{}) bool
}

var (
	_ Iterator = IteratorMock{}
	_ Iterator = iterator{}
)

// IteratorMock is a mock Iterator.
type IteratorMock struct {
	mock.Mock
}

// Close implements Iterator.
func (m IteratorMock) Close() error {
	return m.Called().Error(0)
}

// Scan implements Iterator.
func (m IteratorMock) Scan(results ...interface{}) bool {
	return m.Called(results).Bool(0)
}

// ScanMap implements Iterator.
func (m IteratorMock) ScanMap(results map[string]interface{}) bool {
	return m.Called(results).Bool(0)
}

type iterator struct {
	i *gocql.Iter
}

func (i iterator) Close() error {
	return i.i.Close()
}

func (i iterator) Scan(results ...interface{}) bool {
	return i.i.Scan(results...)
}

func (i iterator) ScanMap(results map[string]interface{}) bool {
	return i.i.MapScan(results)
}
