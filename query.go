package gockle

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/maraino/go-mock"
)

// Query represents a CQL query.
type Query interface {
	// PageSize will tell the iterator to fetch the result in pages of size n.
	PageSize(n int) Query

	// WithContext will set the context to use during a query, it will be used
	// to timeout when waiting for responses from Cassandra.
	WithContext(ctx context.Context) Query

	// PageState sets the paging state for the query to resume paging from a
	// specific point in time. Setting this will disable to query paging for
	// this query, and must be used for all subsequent pages.
	PageState(state []byte) Query

	// Exec executes the query without returning any rows.
	Exec() error

	// Iter executes the query and returns an iterator capable of iterating
	// over all results.
	Iter() Iterator

	// MapScan executes the query, copies the columns of the first selected row
	// into the map pointed at by m and discards the rest. If no rows
	// were selected, ErrNotFound is returned.
	MapScan(m map[string]interface{}) error

	// Scan executes the query, copies the columns of the first selected row
	// into the values pointed at by dest and discards the rest. If no rows
	// were selected, ErrNotFound is returned.
	Scan(dest ...interface{}) error

	// Release releases a query back into a pool of queries. Released Queries
	// cannot be reused.
	Release()
}

var (
	_ Query = QueryMock{}
	_ Query = query{}
)

// QueryMock is a mock Query. See github.com/maraino/go-mock.
type QueryMock struct {
	mock.Mock
}

// PageSize implements Query.
func (m QueryMock) PageSize(n int) Query {
	return m.Called(n).Get(0).(Query)
}

// WithContext implements Query.
func (m QueryMock) WithContext(ctx context.Context) Query {
	return m.Called(ctx).Get(0).(Query)
}

// PageState implements Query.
func (m QueryMock) PageState(state []byte) Query {
	return m.Called(state).Get(0).(Query)
}

// Exec implements Query.
func (m QueryMock) Exec() error {
	return m.Called().Error(0)
}

// Iter implements Query.
func (m QueryMock) Iter() Iterator {
	return m.Called().Get(0).(Iterator)
}

// MapScan implements Query.
func (m QueryMock) MapScan(mm map[string]interface{}) error {
	return m.Called(mm).Error(0)
}

// Scan implements Query.
func (m QueryMock) Scan(dest ...interface{}) error {
	return m.Called(dest).Error(0)
}

// Release implements Query.
func (m QueryMock) Release() {
	m.Called()
}

type query struct {
	q *gocql.Query
}

func (q query) PageSize(n int) Query {
	return &query{q: q.q.PageSize(n)}
}

func (q query) WithContext(ctx context.Context) Query {
	return &query{q: q.q.WithContext(ctx)}
}

func (q query) PageState(state []byte) Query {
	return &query{q: q.q.PageState(state)}
}

func (q query) Exec() error {
	return q.q.Exec()
}

func (q query) Iter() Iterator {
	return &iterator{i: q.q.Iter()}
}

func (q query) MapScan(m map[string]interface{}) error {
	return q.q.MapScan(m)
}

func (q query) Scan(dest ...interface{}) error {
	return q.q.Scan(dest...)
}

func (q query) Release() {
	q.q.Release()
}
