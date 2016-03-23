package gockle

import (
	"fmt"

	"github.com/gocql/gocql"
)

type Session interface {
	Close()

	Columns(keyspace, table string) (map[string]gocql.TypeInfo, error)

	QueryExecute(statement string, arguments ...interface{}) error

	QueryIterate(statement string, arguments ...interface{}) Iterator

	QueryScan(statement string, arguments, results []interface{}) error

	QueryScanMap(statement string, arguments []interface{}, results map[string]interface{}) error

	QueryScanMapTransaction(statement string, arguments []interface{}, results map[string]interface{}) (bool, error)

	QuerySliceMap(statement string, arguments ...interface{}) ([]map[string]interface{}, error)

	Tables(keyspace string) ([]string, error)
}

var _ Session = session{}

func NewSession(s *gocql.Session) Session {
	return session{s: s}
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
