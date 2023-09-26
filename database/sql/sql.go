// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sql provides a generic interface around SQL (or SQL-like)
// databases.
package sql

import (
	"github.com/shogo82148/std/database/sql/driver"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/sync"
)

// Register makes a database driver available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, driver driver.Driver)

// RawBytes is a byte slice that holds a reference to memory owned by
// the database itself. After a Scan into a RawBytes, the slice is only
// valid until the next call to Next, Scan, or Close.
type RawBytes []byte

// NullString represents a string that may be null.
// NullString implements the Scanner interface so
// it can be used as a scan destination:
//
//	var s NullString
//	err := db.QueryRow("SELECT name FROM foo WHERE id=?", id).Scan(&s)
//	...
//	if s.Valid {
//	   // use s.String
//	} else {
//	   // NULL value
//	}
type NullString struct {
	String string
	Valid  bool
}

// Scan implements the Scanner interface.
func (ns *NullString) Scan(value interface{}) error

// Value implements the driver Valuer interface.
func (ns NullString) Value() (driver.Value, error)

// NullInt64 represents an int64 that may be null.
// NullInt64 implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullInt64 struct {
	Int64 int64
	Valid bool
}

// Scan implements the Scanner interface.
func (n *NullInt64) Scan(value interface{}) error

// Value implements the driver Valuer interface.
func (n NullInt64) Value() (driver.Value, error)

// NullFloat64 represents a float64 that may be null.
// NullFloat64 implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullFloat64 struct {
	Float64 float64
	Valid   bool
}

// Scan implements the Scanner interface.
func (n *NullFloat64) Scan(value interface{}) error

// Value implements the driver Valuer interface.
func (n NullFloat64) Value() (driver.Value, error)

// NullBool represents a bool that may be null.
// NullBool implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullBool struct {
	Bool  bool
	Valid bool
}

// Scan implements the Scanner interface.
func (n *NullBool) Scan(value interface{}) error

// Value implements the driver Valuer interface.
func (n NullBool) Value() (driver.Value, error)

// Scanner is an interface used by Scan.
type Scanner interface {
	Scan(src interface{}) error
}

// ErrNoRows is returned by Scan when QueryRow doesn't return a
// row. In such a case, QueryRow returns a placeholder *Row value that
// defers this error until a Scan.
var ErrNoRows = errors.New("sql: no rows in result set")

// DB is a database handle. It's safe for concurrent use by multiple
// goroutines.
//
// If the underlying database driver has the concept of a connection
// and per-connection session state, the sql package manages creating
// and freeing connections automatically, including maintaining a free
// pool of idle connections. If observing session state is required,
// either do not share a *DB between multiple concurrent goroutines or
// create and observe all state only within a transaction. Once
// DB.Open is called, the returned Tx is bound to a single isolated
// connection. Once Tx.Commit or Tx.Rollback is called, that
// connection is returned to DB's idle connection pool.
type DB struct {
	driver driver.Driver
	dsn    string

	mu       sync.Mutex
	freeConn []driver.Conn
	closed   bool
}

// Open opens a database specified by its database driver name and a
// driver-specific data source name, usually consisting of at least a
// database name and connection information.
//
// Most users will open a database via a driver-specific connection
// helper function that returns a *DB.
func Open(driverName, dataSourceName string) (*DB, error)

// Close closes the database, releasing any open resources.
func (db *DB) Close() error

// putConnHook is a hook for testing.

// Prepare creates a prepared statement for later execution.
func (db *DB) Prepare(query string) (*Stmt, error)

// Exec executes a query without returning any rows.
func (db *DB) Exec(query string, args ...interface{}) (Result, error)

// Query executes a query that returns rows, typically a SELECT.
func (db *DB) Query(query string, args ...interface{}) (*Rows, error)

// QueryRow executes a query that is expected to return at most one row.
// QueryRow always return a non-nil value. Errors are deferred until
// Row's Scan method is called.
func (db *DB) QueryRow(query string, args ...interface{}) *Row

// Begin starts a transaction. The isolation level is dependent on
// the driver.
func (db *DB) Begin() (*Tx, error)

// Driver returns the database's underlying driver.
func (db *DB) Driver() driver.Driver

// Tx is an in-progress database transaction.
//
// A transaction must end with a call to Commit or Rollback.
//
// After a call to Commit or Rollback, all operations on the
// transaction fail with ErrTxDone.
type Tx struct {
	db *DB

	ci  driver.Conn
	txi driver.Tx

	cimu sync.Mutex

	done bool
}

var ErrTxDone = errors.New("sql: Transaction has already been committed or rolled back")

// Commit commits the transaction.
func (tx *Tx) Commit() error

// Rollback aborts the transaction.
func (tx *Tx) Rollback() error

// Prepare creates a prepared statement for use within a transaction.
//
// The returned statement operates within the transaction and can no longer
// be used once the transaction has been committed or rolled back.
//
// To use an existing prepared statement on this transaction, see Tx.Stmt.
func (tx *Tx) Prepare(query string) (*Stmt, error)

// Stmt returns a transaction-specific prepared statement from
// an existing statement.
//
// Example:
//
//	updateMoney, err := db.Prepare("UPDATE balance SET money=money+? WHERE id=?")
//	...
//	tx, err := db.Begin()
//	...
//	res, err := tx.Stmt(updateMoney).Exec(123.45, 98293203)
func (tx *Tx) Stmt(stmt *Stmt) *Stmt

// Exec executes a query that doesn't return rows.
// For example: an INSERT and UPDATE.
func (tx *Tx) Exec(query string, args ...interface{}) (Result, error)

// Query executes a query that returns rows, typically a SELECT.
func (tx *Tx) Query(query string, args ...interface{}) (*Rows, error)

// QueryRow executes a query that is expected to return at most one row.
// QueryRow always return a non-nil value. Errors are deferred until
// Row's Scan method is called.
func (tx *Tx) QueryRow(query string, args ...interface{}) *Row

// connStmt is a prepared statement on a particular connection.

// Stmt is a prepared statement. Stmt is safe for concurrent use by multiple goroutines.
type Stmt struct {
	db        *DB
	query     string
	stickyErr error

	tx   *Tx
	txsi driver.Stmt

	mu     sync.Mutex
	closed bool

	css []connStmt
}

// Exec executes a prepared statement with the given arguments and
// returns a Result summarizing the effect of the statement.
func (s *Stmt) Exec(args ...interface{}) (Result, error)

// Query executes a prepared query statement with the given arguments
// and returns the query results as a *Rows.
func (s *Stmt) Query(args ...interface{}) (*Rows, error)

// QueryRow executes a prepared query statement with the given arguments.
// If an error occurs during the execution of the statement, that error will
// be returned by a call to Scan on the returned *Row, which is always non-nil.
// If the query selects no rows, the *Row's Scan will return ErrNoRows.
// Otherwise, the *Row's Scan scans the first selected row and discards
// the rest.
//
// Example usage:
//
//	var name string
//	err := nameByUseridStmt.QueryRow(id).Scan(&name)
func (s *Stmt) QueryRow(args ...interface{}) *Row

// Close closes the statement.
func (s *Stmt) Close() error

// Rows is the result of a query. Its cursor starts before the first row
// of the result set. Use Next to advance through the rows:
//
//	rows, err := db.Query("SELECT ...")
//	...
//	for rows.Next() {
//	    var id int
//	    var name string
//	    err = rows.Scan(&id, &name)
//	    ...
//	}
//	err = rows.Err() // get any error encountered during iteration
//	...
type Rows struct {
	db          *DB
	ci          driver.Conn
	releaseConn func(error)
	rowsi       driver.Rows

	closed    bool
	lastcols  []driver.Value
	lasterr   error
	closeStmt *Stmt
}

// Next prepares the next result row for reading with the Scan method.
// It returns true on success, false if there is no next result row.
// Every call to Scan, even the first one, must be preceded by a call
// to Next.
func (rs *Rows) Next() bool

// Err returns the error, if any, that was encountered during iteration.
func (rs *Rows) Err() error

// Columns returns the column names.
// Columns returns an error if the rows are closed, or if the rows
// are from QueryRow and there was a deferred error.
func (rs *Rows) Columns() ([]string, error)

// Scan copies the columns in the current row into the values pointed
// at by dest.
//
// If an argument has type *[]byte, Scan saves in that argument a copy
// of the corresponding data. The copy is owned by the caller and can
// be modified and held indefinitely. The copy can be avoided by using
// an argument of type *RawBytes instead; see the documentation for
// RawBytes for restrictions on its use.
//
// If an argument has type *interface{}, Scan copies the value
// provided by the underlying driver without conversion. If the value
// is of type []byte, a copy is made and the caller owns the result.
func (rs *Rows) Scan(dest ...interface{}) error

// Close closes the Rows, preventing further enumeration. If the
// end is encountered, the Rows are closed automatically. Close
// is idempotent.
func (rs *Rows) Close() error

// Row is the result of calling QueryRow to select a single row.
type Row struct {
	err  error
	rows *Rows
}

// Scan copies the columns from the matched row into the values
// pointed at by dest.  If more than one row matches the query,
// Scan uses the first row and discards the rest.  If no row matches
// the query, Scan returns ErrNoRows.
func (r *Row) Scan(dest ...interface{}) error

// A Result summarizes an executed SQL command.
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}
