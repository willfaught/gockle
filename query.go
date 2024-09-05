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

	// MapScanCAS executes a lightweight transaction (i.e. an UPDATE or INSERT
	// statement containing an IF clause). If the transaction fails because
	// the existing values did not match, the previous values will be stored
	// in dest map.
	//
	// As for INSERT .. IF NOT EXISTS, previous values will be returned as if
	// SELECT * FROM. So using ScanCAS with INSERT is inherently prone to
	// column mismatching. MapScanCAS is added to capture them safely.
	MapScanCAS(dest map[string]interface{}) (applied bool, err error)

	// Scan executes the query, copies the columns of the first selected row
	// into the values pointed at by dest and discards the rest. If no rows
	// were selected, ErrNotFound is returned.
	Scan(dest ...interface{}) error

	// Release releases a query back into a pool of queries. Released queries
	// cannot be reused.
	Release()

	// GetConsistency returns the currently configured consistency level for
	// the query.
	GetConsistency() gocql.Consistency

	// SetConsistency sets the consistency level for this query.
	SetConsistency(c gocql.Consistency)

	// Consistency sets the consistency level for this query. If no consistency
	// level has been set, the default consistency level of the cluster
	// is used.
	Consistency(c gocql.Consistency) Query

	// SerialConsistency sets the consistency level for the
	// serial phase of conditional updates. That consistency can only be
	// either SERIAL or LOCAL_SERIAL and if not present, it defaults to
	// SERIAL. This option will be ignored for anything else that a
	// conditional update/insert.
	SerialConsistency(cons gocql.SerialConsistency) Query

	// RetryPolicy sets the policy to use when retrying the query.
	RetryPolicy(r gocql.RetryPolicy) Query
}

var (
	_ Query = QueryMock{}
	_ Query = query{}
)

// QueryMock is a mock Query.
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

// MapScan implements Query.
func (m QueryMock) MapScanCAS(mm map[string]interface{}) (bool, error) {
	call := m.Called(mm)
	return call.Bool(0), call.Error(1)
}

// Scan implements Query.
func (m QueryMock) Scan(dest ...interface{}) error {
	return m.Called(dest).Error(0)
}

// Release implements Query.
func (m QueryMock) Release() {
	m.Called()
}

func (m QueryMock) GetConsistency() gocql.Consistency {
	return m.Called().Get(0).(gocql.Consistency)
}

func (m QueryMock) SetConsistency(c gocql.Consistency) {
	m.Called(c)
}

func (m QueryMock) Consistency(c gocql.Consistency) Query {
	return m.Called(c).Get(0).(Query)
}

func (m QueryMock) RetryPolicy(r gocql.RetryPolicy) Query {
	return m.Called(r).Get(0).(Query)
}

func (m QueryMock) SerialConsistency(c gocql.SerialConsistency) Query {
	return m.Called(c).Get(0).(Query)
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

func (q query) GetConsistency() gocql.Consistency {
	return q.q.GetConsistency()
}

func (q query) SetConsistency(c gocql.Consistency) {
	q.q.SetConsistency(c)
}

func (q query) Consistency(c gocql.Consistency) Query {
	return &query{q: q.q.Consistency(c)}
}

func (q query) SerialConsistency(cons gocql.SerialConsistency) Query {
	return &query{q: q.q.SerialConsistency(cons)}
}

func (q query) MapScanCAS(dest map[string]interface{}) (applied bool, err error) {
	return q.q.MapScanCAS(dest)
}

func (q query) RetryPolicy(r gocql.RetryPolicy) Query {
	return &query{q: q.q.RetryPolicy(r)}
}
