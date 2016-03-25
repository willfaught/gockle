# gockle

*gocql: GOC-QuL*

**Gockle simplifies and mocks [gocql](https://github.com/gocql/gocql).** Use it if you don't need to customize batches or queries.

[![GoDoc Widget]][GoDoc] [![Travis Widget]][Travis] [![GoReportCard Widget]][GoReportCard] [![Coverage Status Widget]][Coverage Status]

[GoDoc]: https://godoc.org/github.com/willfaught/gockle
[GoDoc Widget]: https://godoc.org/github.com/willfaught/gockle?status.svg
[Travis]: https://travis-ci.org/willfaught/gockle
[Travis Widget]: https://travis-ci.org/willfaught/gockle.svg?branch=master
[GoReportCard]: https://goreportcard.com/report/github.com/willfaught/gockle
[GoReportCard Widget]: https://goreportcard.com/badge/github.com/willfaught/gockle
[Coverage Status]: https://coveralls.io/r/willfaught/gockle?branch=master
[Coverage Status Widget]: https://coveralls.io/repos/willfaught/gockle/badge.svg?branch=master

## Overview

`Batch`, `Iterator`, and `Session` wrap counterpart `gocql` types.

The entry point is `NewSession`. Call it with a `*gocql.Session` to get a `Session`.

A `gocql.Iter` method has a counterpart in `Session`:

- `gocql.Iter.SliceMap`: `Session.QuerySliceMap`

A `gocql.Batch` method has a counterpart in `Batch`:

- `gocql.Batch.Query`: `Batch.Query`

Some `gocql.Session` methods have counterparts in `Batch`:

- `gocql.Session.ExecuteBatch`: `Batch.Execute`
- `gocql.Session.ExecuteBatchCAS`: `Batch.ExecuteTransaction`
- `gocql.Session.MapExecuteBatchCAS`: `Batch.ExecuteTransactionMap`

Some `gocql.Session` methods have counterparts in `Session`:

- `gocql.Session.Close`: `Session.Close`
- `gocql.Session.NewBatch`: `Session.QueryBatch`

Some `gocql.Query` methods have counterparts in `Session`:

- `gocql.Query.Exec`: `Session.QueryExecute`
- `gocql.Query.MapScan`: `Session.QueryScanMap`
- `gocql.Query.MapScanCAS`: `Session.QueryScanMapTransaction`
- `gocql.Query.Scan`: `Session.QueryScan`

Some `gocql.Iter` methods have counterparts in `Iterator`:

- `gocql.Iter.Close`: `Iterator.Close`
- `gocql.Iter.Scan`: `Iterator.Scan`
- `gocql.Iter.MapScan`: `Iterator.ScanMap`

The rest:

- `Session.QueryIterate` returns an `Iterator` to iterate rows
- `Session.Tables` returns the table names of a keyspace
- `Session.Columns` returns the column names and types of a table

##  Examples

Insert a row:

    var s = NewSession(...)

    if err := s.QueryExecute("insert into users (id, name) values (123, 'me')"); err != nil {
        return err
    }

Print all rows:

    var s = NewSession(...)
    var i = s.QueryIterate("select * from users")

    for done := false; !done; {
        var m = map[string]interface{}{}

        done = i.ScanMap(m)

        fmt.Println(m)
    }

    if err := i.Close(); err != nil {
        return err
    }

Dump all rows:

    var s = NewSession(...)
    var m, err = s.QueryScanMap("select * from users")

    if err != nil {
        return err
    }

    fmt.Println(m)
