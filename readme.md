# gockle

*gocql: GOC-QuL*

**Gockle simplifies and mocks [gocql](https://github.com/gocql/gocql).** Use it if you don't need to customize batches or queries.

[![GoDoc Widget]][GoDoc] [![Travis Widget]][Travis] [![GoReportCard Widget]][GoReportCard]

[GoDoc]: https://godoc.org/github.com/willfaught/gockle
[GoDoc Widget]: https://godoc.org/github.com/willfaught/gockle?status.svg
[Travis]: https://travis-ci.org/willfaught/gockle
[Travis Widget]: https://travis-ci.org/willfaught/gockle.svg?branch=master
[GoReportCard]: https://goreportcard.com/report/github.com/willfaught/gockle
[GoReportCard Widget]: https://goreportcard.com/badge/github.com/willfaught/gockle

The entry point is `NewSession`. Call it with a `*gocql.Session` to get a `Session`.

A `gocql.Session` method has a counterpart in `Session`:

- `gocql.Session.Close`: `Session.Close`

A `gocql.Iter` method has a counterpart in `Session`:

- `gocql.Iter.SliceMap`: `Session.QuerySliceMap`

Some of the `gocql.Query` methods have counterparts in `Session`:

- `gocql.Query.Exec`: `Session.QueryExecute`
- `gocql.Query.MapScan`: `Session.QueryScanMap`
- `gocql.Query.MapScanCAS`: `Session.QueryScanMapTransaction`
- `gocql.Query.Scan`: `Session.QueryScan`

Some of the `gocql.Iter` methods have counterparts in `Iterator`:

- `gocql.Iter.Close`: `Iterator.Close`
- `gocql.Iter.Scan`: `Iterator.Scan`
- `gocql.Iter.MapScan`: `Iterator.ScanMap`

The rest:

- `Session.QueryIterate` returns an `Iterator` to iterate rows
- `Session.Tables` returns the table names of a keyspace
- `Session.Columns` returns the column names and types of a table
