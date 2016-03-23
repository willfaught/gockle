package gockle

import "github.com/gocql/gocql"

// Batch is a batch of queries.
type Batch interface {
	// Execute executes each query in order.
	Execute() error

	// ExecuteTransaction executes each query in order. It puts the first result row in results.
	// If successful, it returns true and an Iterator that ranges over the conditional statement results.
	ExecuteTransaction(results ...interface{}) (bool, Iterator, error)

	// ExecuteTransactionMap executes each query in order. It puts the first result row in results.
	// If successful, it returns true and an Iterator that ranges over the conditional statement results.
	ExecuteTransactionMap(results map[string]interface{}) (bool, Iterator, error)

	// Query adds the query for statement and arguments.
	Query(statement string, arguments ...interface{})
}

var _ Batch = &batch{}

// BatchKind matches gocql.BatchType.
type BatchKind byte

// Types of batches.
const (
	BatchCounter  BatchKind = 2
	BatchLogged   BatchKind = 0
	BatchUnlogged BatchKind = 1
)

type batch struct {
	b *gocql.Batch

	s *gocql.Session
}

func (b *batch) Execute() error {
	var gb, gs = b.clear()

	return gs.ExecuteBatch(gb)
}

func (b *batch) ExecuteTransaction(results ...interface{}) (bool, Iterator, error) {
	var gb, gs = b.clear()
	var a, i, err = gs.ExecuteBatchCAS(gb, results...)

	if err != nil {
		return false, nil, err
	}

	return a, iterator{i: i}, nil
}

func (b *batch) ExecuteTransactionMap(results map[string]interface{}) (bool, Iterator, error) {
	var gb, gs = b.clear()
	var a, i, err = gs.MapExecuteBatchCAS(gb, results)

	if err != nil {
		return false, nil, err
	}

	return a, iterator{i: i}, nil
}

func (b batch) Query(statement string, arguments ...interface{}) {
	b.b.Query(statement, arguments...)
}

func (b *batch) clear() (*gocql.Batch, *gocql.Session) {
	var gb, gs = b.b, b.s

	b.b = nil
	b.s = nil

	return gb, gs
}
