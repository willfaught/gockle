package gockle

import "github.com/gocql/gocql"

type Iterator interface {
	Close() error

	Scan(results ...interface{}) bool

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
