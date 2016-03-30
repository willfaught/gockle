// Package gockle simplifies and mocks gocql. Use it if you don't need to customize batches or queries.
//
// Batch, Iterator, and Session wrap counterpart gocql types.
//
// The entry point is NewSession. Call it with a *gocql.Session to get a Session.
//
// A gocql.Iter method has a counterpart in Session:
//
// - gocql.Iter.SliceMap: Session.QuerySliceMap
//
// A gocql.Batch method has a counterpart in Batch:
//
// - gocql.Batch.Query: Batch.Query
//
// Some gocql.Session methods have counterparts in Batch:
//
// - gocql.Session.ExecuteBatch: Batch.Exec
//
// - gocql.Session.ExecuteBatchCAS: Batch.ExecTx
//
// - gocql.Session.MapExecuteBatchCAS: Batch.ExecTxMap
//
// Some gocql.Session methods have counterparts in Session:
//
// - gocql.Session.Close: Session.Close
//
// - gocql.Session.NewBatch: Session.QueryBatch
//
// Some gocql.Query methods have counterparts in Session:
//
// - gocql.Query.Exec: Session.QueryExec
//
// - gocql.Query.MapScan: Session.QueryScanMap
//
// - gocql.Query.MapScanCAS: Session.QueryScanMapTx
//
// - gocql.Query.Scan: Session.QueryScan
//
// Some gocql.Iter methods have counterparts in Iterator:
//
// - gocql.Iter.Close: Iterator.Close
//
// - gocql.Iter.Scan: Iterator.Scan
//
// - gocql.Iter.MapScan: Iterator.ScanMap
//
// The rest:
//
// - Session.QueryIterator returns an Iterator to iterate rows
//
// - Session.Tables returns the table names of a keyspace
//
// - Session.Columns returns the column names and types of a table
package gockle
