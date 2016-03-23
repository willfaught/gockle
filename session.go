package gockle

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/maraino/go-mock"
)

// Session wraps *gocql.Session.
type Session interface {
	// Close closes the session.
	Close()

	// Columns returns a map from column names to types for keyspace and table.
	Columns(keyspace, table string) (map[string]gocql.TypeInfo, error)

	// QueryBatch returns a Batch with type kind for batched queries.
	QueryBatch(kind BatchKind) Batch

	// QueryExecute runs statement with arguments.
	QueryExecute(statement string, arguments ...interface{}) error

	// QueryIterate runs statement with arguments and returns an Iterator for
	// the results.
	QueryIterate(statement string, arguments ...interface{}) Iterator

	// QueryScan runs statement with arguments.
	QueryScan(statement string, arguments, results []interface{}) error

	// QueryScanMap runs statement with arguments and puts the first result row
	// in results.
	QueryScanMap(statement string, arguments []interface{}, results map[string]interface{}) error

	// QueryScanMapTransaction runs statement with arguments as a lightweight
	// transaction and puts the first result row in results. It returns whether
	// the transaction succeeded. If not, it puts the old values in results.
	QueryScanMapTransaction(statement string, arguments []interface{}, results map[string]interface{}) (bool, error)

	// QuerySliceMap runs statement with arguments and returns all result rows.
	QuerySliceMap(statement string, arguments ...interface{}) ([]map[string]interface{}, error)

	// Tables returns the table names for keyspace.
	Tables(keyspace string) ([]string, error)
}

var _ Session = session{}

// NewSession returns a new Session for s.
func NewSession(s *gocql.Session) Session {
	return session{s: s}
}

// SessionMock is a mock Session.
type SessionMock struct {
	mock.Mock
}

// Close implements Session.
func (m SessionMock) Close() {
	m.Called()
}

// Columns implements Session.
func (m SessionMock) Columns(keyspace, table string) (map[string]gocql.TypeInfo, error) {
	var r = m.Called(keyspace, table)

	return r.Get(0).(map[string]gocql.TypeInfo), r.Error(1)
}

// QueryBatch implements Session.
func (m SessionMock) QueryBatch(kind BatchKind) Batch {
	return m.Called(kind).Get(0).(Batch)
}

// QueryExecute implements Session.
func (m SessionMock) QueryExecute(statement string, arguments ...interface{}) error {
	return m.Called(statement, arguments).Error(0)
}

// QueryIterate implements Session.
func (m SessionMock) QueryIterate(statement string, arguments ...interface{}) Iterator {
	return m.Called(statement, arguments).Get(0).(Iterator)
}

// QueryScan implements Session.
func (m SessionMock) QueryScan(statement string, arguments, results []interface{}) error {
	return m.Called(statement, arguments, results).Error(0)
}

// QueryScanMap implements Session.
func (m SessionMock) QueryScanMap(statement string, arguments []interface{}, results map[string]interface{}) error {
	return m.Called(statement, arguments, results).Error(0)
}

// QueryScanMapTransaction implements Session.
func (m SessionMock) QueryScanMapTransaction(statement string, arguments []interface{}, results map[string]interface{}) (bool, error) {
	var r = m.Called(statement, arguments)

	return r.Bool(0), r.Error(1)
}

// QuerySliceMap implements Session.
func (m SessionMock) QuerySliceMap(statement string, arguments ...interface{}) ([]map[string]interface{}, error) {
	var r = m.Called(statement, arguments)

	return r.Get(0).([]map[string]interface{}), r.Error(1)
}

// Tables implements Session.
func (m SessionMock) Tables(keyspace string) ([]string, error) {
	var r = m.Called(keyspace)

	return r.Get(0).([]string), r.Error(1)
}

type session struct {
	s *gocql.Session
}

func (s session) Close() {
	s.Close()
}

func (s session) Columns(keyspace, table string) (map[string]gocql.TypeInfo, error) {
	var m, err = s.s.KeyspaceMetadata(keyspace)

	if err != nil {
		return nil, err
	}

	var t, ok = m.Tables[table]

	if !ok {
		return nil, fmt.Errorf("table %v.%v invalid", keyspace, table)
	}

	var types = map[string]gocql.TypeInfo{}

	for n, c := range t.Columns {
		types[n] = c.Type
	}

	return types, nil
}

func (s session) QueryBatch(kind BatchKind) Batch {
	return &batch{b: s.s.NewBatch(gocql.BatchType(kind)), s: s.s}
}

func (s session) QueryExecute(statement string, arguments ...interface{}) error {
	return s.s.Query(statement, arguments...).Exec()
}

func (s session) QueryIterate(statement string, arguments ...interface{}) Iterator {
	return iterator{i: s.s.Query(statement, arguments...).Iter()}
}

func (s session) QueryScan(statement string, arguments, results []interface{}) error {
	return s.s.Query(statement, arguments...).Scan(results...)
}

func (s session) QueryScanMap(statement string, arguments []interface{}, results map[string]interface{}) error {
	return s.s.Query(statement, arguments...).MapScan(results)
}

func (s session) QueryScanMapTransaction(statement string, arguments []interface{}, results map[string]interface{}) (bool, error) {
	return s.s.Query(statement, arguments...).MapScanCAS(results)
}

func (s session) QuerySliceMap(statement string, arguments ...interface{}) ([]map[string]interface{}, error) {
	var i = s.s.Query(statement, arguments...).Iter()
	var m, err = i.SliceMap()

	if err != nil {
		return nil, err
	}

	if err := i.Close(); err != nil {
		return nil, err
	}

	return m, nil
}

func (s session) Tables(keyspace string) ([]string, error) {
	var m, err = s.s.KeyspaceMetadata(keyspace)

	if err != nil {
		return nil, err
	}

	var ts []string

	for t := range m.Tables {
		ts = append(ts, t)
	}

	return ts, nil
}
