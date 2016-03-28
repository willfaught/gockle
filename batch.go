package gockle

import (
	"github.com/gocql/gocql"
	"github.com/maraino/go-mock"
)

// Batch is a batch of queries.
type Batch interface {
	// Exec executes each query in order.
	Exec() error

	// ExecTx executes each query in order. It puts the first result row in results.
	// If successful, it returns true and an Iterator that ranges over the
	// conditional statement results.
	ExecTx(results ...interface{}) (bool, Iterator, error)

	// ExecTxMap executes each query in order. It puts the first result row in
	// results. If successful, it returns true and an Iterator that ranges over the
	// conditional statement results.
	ExecTxMap(results map[string]interface{}) (bool, Iterator, error)

	// Query adds the query for statement and arguments.
	Query(statement string, arguments ...interface{})
}

var (
	_ Batch = BatchMock{}
	_ Batch = batch{}
)

// BatchKind matches gocql.BatchType.
type BatchKind byte

// Types of batches.
const (
	BatchCounter  BatchKind = 2
	BatchLogged   BatchKind = 0
	BatchUnlogged BatchKind = 1
)

// BatchMock is a mock Batch.
type BatchMock struct {
	mock.Mock
}

// Exec implements Batch.
func (m BatchMock) Exec() error {
	return m.Called().Error(0)
}

// ExecTx implements Batch.
func (m BatchMock) ExecTx(results ...interface{}) (bool, Iterator, error) {
	var r = m.Called(results)

	return r.Bool(0), r.Get(1).(Iterator), r.Error(2)
}

// ExecTxMap implements Batch.
func (m BatchMock) ExecTxMap(results map[string]interface{}) (bool, Iterator, error) {
	var r = m.Called(results)

	return r.Bool(0), r.Get(1).(Iterator), r.Error(2)
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

func (b batch) ExecTx(results ...interface{}) (bool, Iterator, error) {
	var a, i, err = b.s.ExecuteBatchCAS(b.b, results...)

	if err != nil {
		return false, nil, err
	}

	return a, iterator{i: i}, nil
}

func (b batch) ExecTxMap(results map[string]interface{}) (bool, Iterator, error) {
	var a, i, err = b.s.MapExecuteBatchCAS(b.b, results)

	if err != nil {
		return false, nil, err
	}

	return a, iterator{i: i}, nil
}

func (b batch) Query(statement string, arguments ...interface{}) {
	b.b.Query(statement, arguments...)
}
