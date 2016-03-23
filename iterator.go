package gockle

import "github.com/gocql/gocql"

// Iterator wraps *gocql.Iter.
type Iterator interface {
	// Close closes the iterator.
	Close() error

	// Scan puts the current row in results and returns whether there are more rows.
	Scan(results ...interface{}) bool

	// ScanMap puts the current row in results and returns whether there are more rows.
	ScanMap(results map[string]interface{}) bool
}

var _ Iterator = iterator{}

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
