package gockle

import (
	"github.com/gocql/gocql"
	"github.com/maraino/go-mock"
)

// Batch is an ordered collection of CQL queries.
type Batch interface {
	// Exec executes the queries in the order they were added.
	Exec() error

	// ExecTx executes the queries in the order they were added. It returns a slice
	// of maps from columns to values, the maps corresponding to all the conditional
	// queries, and ordered in the same relative order. The special column
	// "[applied]" has a bool that indicates whether the conditional statement was
	// applied. If a conditional statement was not applied, the current values for
	// the columns are put into the map.
	ExecTx() ([]map[string]interface{}, error)

	// Query adds the query for statement and arguments.
	Query(statement string, arguments ...interface{})
}

var (
	_ Batch = BatchMock{}
	_ Batch = batch{}
)

// BatchKind matches gocql.BatchType.
type BatchKind byte

// Types of batches. TODO: Reorder and add doc.
const (
	BatchCounter  BatchKind = 2
	BatchLogged   BatchKind = 0
	BatchUnlogged BatchKind = 1
)

// BatchMock is a mock Batch. See github.com/maraino/go-mock.
type BatchMock struct {
	mock.Mock
}

// Exec implements Batch.
func (m BatchMock) Exec() error {
	return m.Called().Error(0)
}

// ExecTx implements Batch.
func (m BatchMock) ExecTx() ([]map[string]interface{}, error) {
	var r = m.Called()

	return r.Get(0).([]map[string]interface{}), r.Error(1)
}

// Query implements Batch.
func (m BatchMock) Query(statement string, arguments ...interface{}) {
	m.Called(statement, arguments)
}

type batch struct {
	b *gocql.Batch

	s *gocql.Session
}

func (b batch) Exec() error {
	return b.s.ExecuteBatch(b.b)
}

func (b batch) ExecTx() ([]map[string]interface{}, error) {
	var m = map[string]interface{}{}
	var a, i, err = b.s.MapExecuteBatchCAS(b.b, m)

	if err != nil {
		return nil, err
	}

	s, err := i.SliceMap()

	if err != nil {
		return nil, err
	}

	if err := i.Close(); err != nil {
		return nil, err
	}

	m["[applied]"] = a
	s = append([]map[string]interface{}{m}, s...)

	return s, nil
}

func (b batch) Query(statement string, arguments ...interface{}) {
	b.b.Query(statement, arguments...)
}
