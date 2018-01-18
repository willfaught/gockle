package gockle

import (
	"github.com/gocql/gocql"
	"github.com/maraino/go-mock"
)

// Iterator iterates CQL query result rows.
type Iterator interface {
	// Close closes the Iterator.
	Close() error

	// Scan puts the current result row in results and returns whether there are
	// more result rows.
	Scan(results ...interface{}) bool

	// ScanMap puts the current result row in results and returns whether there are
	// more result rows.
	ScanMap(results map[string]interface{}) bool

	// WillSwitchPage detects if iterator reached end of current page and the
	// next page is available.
	WillSwitchPage() bool

	// PageState return the current paging state for a query which can be used
	// for subsequent quries to resume paging this point.
	PageState() []byte
}

var (
	_ Iterator = IteratorMock{}
	_ Iterator = iterator{}
)

// IteratorMock is a mock Iterator. See github.com/maraino/go-mock.
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

// WillSwitchPage implements Iterator.
func (m IteratorMock) WillSwitchPage() bool {
	return m.Called().Bool(0)
}

// PageState implements Iterator.
func (m IteratorMock) PageState() []byte {
	return m.Called().Bytes(0)
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

func (i iterator) WillSwitchPage() bool {
	return i.i.WillSwitchPage()
}

func (i iterator) PageState() []byte {
	return i.i.PageState()
}
