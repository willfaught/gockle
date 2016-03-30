// Package gockle simplifies and mocks github.com/gocql/gocql.
//
// Gockle provides the basic abilities to insert, query, and mutate data, as
// well as access to basic keyspace and table metadata.
//
// The entry point is NewSession or NewSimpleSession. Call it to get a Session.
//
// A Session is the connection to the database and the primary means to access
// it. It has all the query methods and the means to iterate result rows and
// batch together mutations. The Session implementation simply wraps
// gocql.Session and adapts a few things to have a simpler interface.
//
// Closing the Session closes the underlying gocql.Session, including the one
// passed in with NewSimpleSession.
//
// Mocks are provided for testing use of Batch, Iterator, and Session.
package gockle
