// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package driver defines interfaces to be implemented by database
// drivers as used by package sql.
//
// Most code should use package sql.
package driver

import "github.com/shogo82148/std/errors"

// A driver Value is a value that drivers must be able to handle.
// A Value is either nil or an instance of one of these types:
//
//	int64
//	float64
//	bool
//	[]byte
//	string   [*] everywhere except from Rows.Next.
//	time.Time
type Value interface{}

// Driver is the interface that must be implemented by a database
// driver.
type Driver interface {
	Open(name string) (Conn, error)
}

// ErrSkip may be returned by some optional interfaces' methods to
// indicate at runtime that the fast path is unavailable and the sql
// package should continue as if the optional interface was not
// implemented. ErrSkip is only supported where explicitly
// documented.
var ErrSkip = errors.New("driver: skip fast-path; continue as if unimplemented")

// ErrBadConn should be returned by a driver to signal to the sql
// package that a driver.Conn is in a bad state (such as the server
// having earlier closed the connection) and the sql package should
// retry on a new connection.
//
// To prevent duplicate operations, ErrBadConn should NOT be returned
// if there's a possibility that the database server might have
// performed the operation. Even if the server sends back an error,
// you shouldn't return ErrBadConn.
var ErrBadConn = errors.New("driver: bad connection")

// Execer is an optional interface that may be implemented by a Conn.
//
// If a Conn does not implement Execer, the db package's DB.Exec will
// first prepare a query, execute the statement, and then close the
// statement.
//
// Exec may return ErrSkip.
type Execer interface {
	Exec(query string, args []Value) (Result, error)
}

// Conn is a connection to a database. It is not used concurrently
// by multiple goroutines.
//
// Conn is assumed to be stateful.
type Conn interface {
	Prepare(query string) (Stmt, error)

	Close() error

	Begin() (Tx, error)
}

// Result is the result of a query execution.
type Result interface {
	LastInsertId() (int64, error)

	RowsAffected() (int64, error)
}

// Stmt is a prepared statement. It is bound to a Conn and not
// used by multiple goroutines concurrently.
type Stmt interface {
	Close() error

	NumInput() int

	Exec(args []Value) (Result, error)

	Query(args []Value) (Rows, error)
}

// ColumnConverter may be optionally implemented by Stmt if the
// the statement is aware of its own columns' types and can
// convert from any type to a driver Value.
type ColumnConverter interface {
	ColumnConverter(idx int) ValueConverter
}

// Rows is an iterator over an executed query's results.
type Rows interface {
	Columns() []string

	Close() error

	Next(dest []Value) error
}

// Tx is a transaction.
type Tx interface {
	Commit() error
	Rollback() error
}

// RowsAffected implements Result for an INSERT or UPDATE operation
// which mutates a number of rows.
type RowsAffected int64

var _ Result = RowsAffected(0)

func (RowsAffected) LastInsertId() (int64, error)

func (v RowsAffected) RowsAffected() (int64, error)

// ResultNoRows is a pre-defined Result for drivers to return when a DDL
// command (such as a CREATE TABLE) succeeds. It returns an error for both
// LastInsertId and RowsAffected.
var ResultNoRows noRows

var _ Result = noRows{}
