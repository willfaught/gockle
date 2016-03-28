package gockle

import (
	"github.com/gocql/gocql"
	"github.com/maraino/go-mock"
)

// Batch is a batch of queries.
type Batch interface {
	// Execute executes each query in order.
	Execute() error

	// ExecuteTx executes each query in order. It puts the first
	// result row in results. If successful, it returns true and an Iterator
	// that ranges over the conditional statement results.
	ExecuteTx(results ...interface{}) (bool, Iterator, error)

	// ExecuteTxMap executes each query in order. It puts the first
	// result row in results. If successful, it returns true and an Iterator
	// that ranges over the conditional statement results.
	ExecuteTxMap(results map[string]interface{}) (bool, Iterator, error)

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

// Execute implements Batch.
func (m BatchMock) Execute() error {
	return m.Called().Error(0)
}

// ExecuteTx implements Batch.
func (m BatchMock) ExecuteTx(results ...interface{}) (bool, Iterator, error) {
	var r = m.Called(results)

	return r.Bool(0), r.Get(1).(Iterator), r.Error(2)
}

// ExecuteTxMap implements Batch.
func (m BatchMock) ExecuteTxMap(results map[string]interface{}) (bool, Iterator, error) {
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

func (b batch) Execute() error {
	return b.s.ExecuteBatch(b.b)
}

func (b batch) ExecuteTx(results ...interface{}) (bool, Iterator, error) {
	var a, i, err = b.s.ExecuteBatchCAS(b.b, results...)

	if err != nil {
		return false, nil, err
	}

	return a, iterator{i: i}, nil
}

func (b batch) ExecuteTxMap(results map[string]interface{}) (bool, Iterator, error) {
	var a, i, err = b.s.MapExecuteBatchCAS(b.b, results)

	if err != nil {
		return false, nil, err
	}

	return a, iterator{i: i}, nil
}

func (b batch) Query(statement string, arguments ...interface{}) {
	b.b.Query(statement, arguments...)
}
