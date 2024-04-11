// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sql provides a generic interface around SQL (or SQL-like)
// databases.
//
// The sql package must be used in conjunction with a database driver.
// See https://golang.org/s/sqldrivers for a list of drivers.
//
// Drivers that do not support context cancellation will not return until
// after the query is completed.
//
// For usage examples, see the wiki page at
// https://golang.org/s/sqlwiki.
package sql

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/database/sql/driver"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/reflect"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/sync/atomic"
	"github.com/shogo82148/std/time"
)

// Register makes a database driver available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, driver driver.Driver)

// Drivers returns a sorted list of the names of the registered drivers.
func Drivers() []string

// A NamedArg is a named argument. NamedArg values may be used as
// arguments to [DB.Query] or [DB.Exec] and bind to the corresponding named
// parameter in the SQL statement.
//
// For a more concise way to create NamedArg values, see
// the [Named] function.
type NamedArg struct {
	_NamedFieldsRequired struct{}

	// Name is the name of the parameter placeholder.
	//
	// If empty, the ordinal position in the argument list will be
	// used.
	//
	// Name must omit any symbol prefix.
	Name string

	// Value is the value of the parameter.
	// It may be assigned the same value types as the query
	// arguments.
	Value any
}

// Named provides a more concise way to create [NamedArg] values.
//
// Example usage:
//
//	db.ExecContext(ctx, `
//	    delete from Invoice
//	    where
//	        TimeCreated < @end
//	        and TimeCreated >= @start;`,
//	    sql.Named("start", startTime),
//	    sql.Named("end", endTime),
//	)
func Named(name string, value any) NamedArg

// IsolationLevel is the transaction isolation level used in [TxOptions].
type IsolationLevel int

// Various isolation levels that drivers may support in [DB.BeginTx].
// If a driver does not support a given isolation level an error may be returned.
//
// See https://en.wikipedia.org/wiki/Isolation_(database_systems)#Isolation_levels.
const (
	LevelDefault IsolationLevel = iota
	LevelReadUncommitted
	LevelReadCommitted
	LevelWriteCommitted
	LevelRepeatableRead
	LevelSnapshot
	LevelSerializable
	LevelLinearizable
)

// String returns the name of the transaction isolation level.
func (i IsolationLevel) String() string

var _ fmt.Stringer = LevelDefault

// TxOptions holds the transaction options to be used in [DB.BeginTx].
type TxOptions struct {
	// Isolation is the transaction isolation level.
	// If zero, the driver or database's default level is used.
	Isolation IsolationLevel
	ReadOnly  bool
}

// RawBytes is a byte slice that holds a reference to memory owned by
// the database itself. After a [Rows.Scan] into a RawBytes, the slice is only
// valid until the next call to [Rows.Next], [Rows.Scan], or [Rows.Close].
type RawBytes []byte

// NullString represents a string that may be null.
// NullString implements the [Scanner] interface so
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

// Scan implements the [Scanner] interface.
func (ns *NullString) Scan(value any) error

// Value implements the [driver.Valuer] interface.
func (ns NullString) Value() (driver.Value, error)

// NullInt64 represents an int64 that may be null.
// NullInt64 implements the [Scanner] interface so
// it can be used as a scan destination, similar to [NullString].
type NullInt64 struct {
	Int64 int64
	Valid bool
}

// Scan implements the [Scanner] interface.
func (n *NullInt64) Scan(value any) error

// Value implements the [driver.Valuer] interface.
func (n NullInt64) Value() (driver.Value, error)

// NullInt32 represents an int32 that may be null.
// NullInt32 implements the [Scanner] interface so
// it can be used as a scan destination, similar to [NullString].
type NullInt32 struct {
	Int32 int32
	Valid bool
}

// Scan implements the [Scanner] interface.
func (n *NullInt32) Scan(value any) error

// Value implements the [driver.Valuer] interface.
func (n NullInt32) Value() (driver.Value, error)

// NullInt16 represents an int16 that may be null.
// NullInt16 implements the [Scanner] interface so
// it can be used as a scan destination, similar to [NullString].
type NullInt16 struct {
	Int16 int16
	Valid bool
}

// Scan implements the [Scanner] interface.
func (n *NullInt16) Scan(value any) error

// Value implements the [driver.Valuer] interface.
func (n NullInt16) Value() (driver.Value, error)

// NullByte represents a byte that may be null.
// NullByte implements the [Scanner] interface so
// it can be used as a scan destination, similar to [NullString].
type NullByte struct {
	Byte  byte
	Valid bool
}

// Scan implements the [Scanner] interface.
func (n *NullByte) Scan(value any) error

// Value implements the [driver.Valuer] interface.
func (n NullByte) Value() (driver.Value, error)

// NullFloat64 represents a float64 that may be null.
// NullFloat64 implements the [Scanner] interface so
// it can be used as a scan destination, similar to [NullString].
type NullFloat64 struct {
	Float64 float64
	Valid   bool
}

// Scan implements the [Scanner] interface.
func (n *NullFloat64) Scan(value any) error

// Value implements the [driver.Valuer] interface.
func (n NullFloat64) Value() (driver.Value, error)

// NullBool represents a bool that may be null.
// NullBool implements the [Scanner] interface so
// it can be used as a scan destination, similar to [NullString].
type NullBool struct {
	Bool  bool
	Valid bool
}

// Scan implements the [Scanner] interface.
func (n *NullBool) Scan(value any) error

// Value implements the [driver.Valuer] interface.
func (n NullBool) Value() (driver.Value, error)

// NullTime represents a [time.Time] that may be null.
// NullTime implements the [Scanner] interface so
// it can be used as a scan destination, similar to [NullString].
type NullTime struct {
	Time  time.Time
	Valid bool
}

// Scan implements the [Scanner] interface.
func (n *NullTime) Scan(value any) error

// Value implements the [driver.Valuer] interface.
func (n NullTime) Value() (driver.Value, error)

// Null represents a value that may be null.
// Null implements the [Scanner] interface so
// it can be used as a scan destination:
//
//	var s Null[string]
//	err := db.QueryRow("SELECT name FROM foo WHERE id=?", id).Scan(&s)
//	...
//	if s.Valid {
//	   // use s.V
//	} else {
//	   // NULL value
//	}
type Null[T any] struct {
	V     T
	Valid bool
}

func (n *Null[T]) Scan(value any) error

func (n Null[T]) Value() (driver.Value, error)

// Scanner is an interface used by [Rows.Scan].
type Scanner interface {
	Scan(src any) error
}

// Out may be used to retrieve OUTPUT value parameters from stored procedures.
//
// Not all drivers and databases support OUTPUT value parameters.
//
// Example usage:
//
//	var outArg string
//	_, err := db.ExecContext(ctx, "ProcName", sql.Named("Arg1", sql.Out{Dest: &outArg}))
type Out struct {
	_NamedFieldsRequired struct{}

	// Dest is a pointer to the value that will be set to the result of the
	// stored procedure's OUTPUT parameter.
	Dest any

	// In is whether the parameter is an INOUT parameter. If so, the input value to the stored
	// procedure is the dereferenced value of Dest's pointer, which is then replaced with
	// the output value.
	In bool
}

// ErrNoRows is returned by [Row.Scan] when [DB.QueryRow] doesn't return a
// row. In such a case, QueryRow returns a placeholder [*Row] value that
// defers this error until a Scan.
var ErrNoRows = errors.New("sql: no rows in result set")

// DB is a database handle representing a pool of zero or more
// underlying connections. It's safe for concurrent use by multiple
// goroutines.
//
// The sql package creates and frees connections automatically; it
// also maintains a free pool of idle connections. If the database has
// a concept of per-connection state, such state can be reliably observed
// within a transaction ([Tx]) or connection ([Conn]). Once [DB.Begin] is called, the
// returned [Tx] is bound to a single connection. Once [Tx.Commit] or
// [Tx.Rollback] is called on the transaction, that transaction's
// connection is returned to [DB]'s idle connection pool. The pool size
// can be controlled with [DB.SetMaxIdleConns].
type DB struct {
	// Total time waited for new connections.
	waitDuration atomic.Int64

	connector driver.Connector
	// numClosed is an atomic counter which represents a total number of
	// closed connections. Stmt.openStmt checks it before cleaning closed
	// connections in Stmt.css.
	numClosed atomic.Uint64

	mu           sync.Mutex
	freeConn     []*driverConn
	connRequests connRequestSet
	numOpen      int
	// Used to signal the need for new connections
	// a goroutine running connectionOpener() reads on this chan and
	// maybeOpenNewConnections sends on the chan (one send per needed connection)
	// It is closed during db.Close(). The close tells the connectionOpener
	// goroutine to exit.
	openerCh          chan struct{}
	closed            bool
	dep               map[finalCloser]depSet
	lastPut           map[*driverConn]string
	maxIdleCount      int
	maxOpen           int
	maxLifetime       time.Duration
	maxIdleTime       time.Duration
	cleanerCh         chan struct{}
	waitCount         int64
	maxIdleClosed     int64
	maxIdleTimeClosed int64
	maxLifetimeClosed int64

	stop func()
}

// OpenDB opens a database using a [driver.Connector], allowing drivers to
// bypass a string based data source name.
//
// Most users will open a database via a driver-specific connection
// helper function that returns a [*DB]. No database drivers are included
// in the Go standard library. See https://golang.org/s/sqldrivers for
// a list of third-party drivers.
//
// OpenDB may just validate its arguments without creating a connection
// to the database. To verify that the data source name is valid, call
// [DB.Ping].
//
// The returned [DB] is safe for concurrent use by multiple goroutines
// and maintains its own pool of idle connections. Thus, the OpenDB
// function should be called just once. It is rarely necessary to
// close a [DB].
func OpenDB(c driver.Connector) *DB

// Open opens a database specified by its database driver name and a
// driver-specific data source name, usually consisting of at least a
// database name and connection information.
//
// Most users will open a database via a driver-specific connection
// helper function that returns a [*DB]. No database drivers are included
// in the Go standard library. See https://golang.org/s/sqldrivers for
// a list of third-party drivers.
//
// Open may just validate its arguments without creating a connection
// to the database. To verify that the data source name is valid, call
// [DB.Ping].
//
// The returned [DB] is safe for concurrent use by multiple goroutines
// and maintains its own pool of idle connections. Thus, the Open
// function should be called just once. It is rarely necessary to
// close a [DB].
func Open(driverName, dataSourceName string) (*DB, error)

// PingContext verifies a connection to the database is still alive,
// establishing a connection if necessary.
func (db *DB) PingContext(ctx context.Context) error

// Ping verifies a connection to the database is still alive,
// establishing a connection if necessary.
//
// Ping uses [context.Background] internally; to specify the context, use
// [DB.PingContext].
func (db *DB) Ping() error

// Close closes the database and prevents new queries from starting.
// Close then waits for all queries that have started processing on the server
// to finish.
//
// It is rare to Close a [DB], as the [DB] handle is meant to be
// long-lived and shared between many goroutines.
func (db *DB) Close() error

// SetMaxIdleConns sets the maximum number of connections in the idle
// connection pool.
//
// If MaxOpenConns is greater than 0 but less than the new MaxIdleConns,
// then the new MaxIdleConns will be reduced to match the MaxOpenConns limit.
//
// If n <= 0, no idle connections are retained.
//
// The default max idle connections is currently 2. This may change in
// a future release.
func (db *DB) SetMaxIdleConns(n int)

// SetMaxOpenConns sets the maximum number of open connections to the database.
//
// If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than
// MaxIdleConns, then MaxIdleConns will be reduced to match the new
// MaxOpenConns limit.
//
// If n <= 0, then there is no limit on the number of open connections.
// The default is 0 (unlimited).
func (db *DB) SetMaxOpenConns(n int)

// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
//
// Expired connections may be closed lazily before reuse.
//
// If d <= 0, connections are not closed due to a connection's age.
func (db *DB) SetConnMaxLifetime(d time.Duration)

// SetConnMaxIdleTime sets the maximum amount of time a connection may be idle.
//
// Expired connections may be closed lazily before reuse.
//
// If d <= 0, connections are not closed due to a connection's idle time.
func (db *DB) SetConnMaxIdleTime(d time.Duration)

// DBStats contains database statistics.
type DBStats struct {
	MaxOpenConnections int

	// Pool Status
	OpenConnections int
	InUse           int
	Idle            int

	// Counters
	WaitCount         int64
	WaitDuration      time.Duration
	MaxIdleClosed     int64
	MaxIdleTimeClosed int64
	MaxLifetimeClosed int64
}

// Stats returns database statistics.
func (db *DB) Stats() DBStats

// PrepareContext creates a prepared statement for later queries or executions.
// Multiple queries or executions may be run concurrently from the
// returned statement.
// The caller must call the statement's [*Stmt.Close] method
// when the statement is no longer needed.
//
// The provided context is used for the preparation of the statement, not for the
// execution of the statement.
func (db *DB) PrepareContext(ctx context.Context, query string) (*Stmt, error)

// Prepare creates a prepared statement for later queries or executions.
// Multiple queries or executions may be run concurrently from the
// returned statement.
// The caller must call the statement's [*Stmt.Close] method
// when the statement is no longer needed.
//
// Prepare uses [context.Background] internally; to specify the context, use
// [DB.PrepareContext].
func (db *DB) Prepare(query string) (*Stmt, error)

// ExecContext executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (db *DB) ExecContext(ctx context.Context, query string, args ...any) (Result, error)

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
//
// Exec uses [context.Background] internally; to specify the context, use
// [DB.ExecContext].
func (db *DB) Exec(query string, args ...any) (Result, error)

// QueryContext executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (db *DB) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)

// Query executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
//
// Query uses [context.Background] internally; to specify the context, use
// [DB.QueryContext].
func (db *DB) Query(query string, args ...any) (*Rows, error)

// QueryRowContext executes a query that is expected to return at most one row.
// QueryRowContext always returns a non-nil value. Errors are deferred until
// [Row]'s Scan method is called.
// If the query selects no rows, the [*Row.Scan] will return [ErrNoRows].
// Otherwise, [*Row.Scan] scans the first selected row and discards
// the rest.
func (db *DB) QueryRowContext(ctx context.Context, query string, args ...any) *Row

// QueryRow executes a query that is expected to return at most one row.
// QueryRow always returns a non-nil value. Errors are deferred until
// [Row]'s Scan method is called.
// If the query selects no rows, the [*Row.Scan] will return [ErrNoRows].
// Otherwise, [*Row.Scan] scans the first selected row and discards
// the rest.
//
// QueryRow uses [context.Background] internally; to specify the context, use
// [DB.QueryRowContext].
func (db *DB) QueryRow(query string, args ...any) *Row

// BeginTx starts a transaction.
//
// The provided context is used until the transaction is committed or rolled back.
// If the context is canceled, the sql package will roll back
// the transaction. [Tx.Commit] will return an error if the context provided to
// BeginTx is canceled.
//
// The provided [TxOptions] is optional and may be nil if defaults should be used.
// If a non-default isolation level is used that the driver doesn't support,
// an error will be returned.
func (db *DB) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)

// Begin starts a transaction. The default isolation level is dependent on
// the driver.
//
// Begin uses [context.Background] internally; to specify the context, use
// [DB.BeginTx].
func (db *DB) Begin() (*Tx, error)

// Driver returns the database's underlying driver.
func (db *DB) Driver() driver.Driver

// ErrConnDone is returned by any operation that is performed on a connection
// that has already been returned to the connection pool.
var ErrConnDone = errors.New("sql: connection is already closed")

// Conn returns a single connection by either opening a new connection
// or returning an existing connection from the connection pool. Conn will
// block until either a connection is returned or ctx is canceled.
// Queries run on the same Conn will be run in the same database session.
//
// Every Conn must be returned to the database pool after use by
// calling [Conn.Close].
func (db *DB) Conn(ctx context.Context) (*Conn, error)

// Conn represents a single database connection rather than a pool of database
// connections. Prefer running queries from [DB] unless there is a specific
// need for a continuous single database connection.
//
// A Conn must call [Conn.Close] to return the connection to the database pool
// and may do so concurrently with a running query.
//
// After a call to [Conn.Close], all operations on the
// connection fail with [ErrConnDone].
type Conn struct {
	db *DB

	// closemu prevents the connection from closing while there
	// is an active query. It is held for read during queries
	// and exclusively during close.
	closemu sync.RWMutex

	// dc is owned until close, at which point
	// it's returned to the connection pool.
	dc *driverConn

	// done transitions from false to true exactly once, on close.
	// Once done, all operations fail with ErrConnDone.
	done atomic.Bool

	releaseConnOnce sync.Once
	// releaseConnCache is a cache of c.closemuRUnlockCondReleaseConn
	// to save allocations in a call to grabConn.
	releaseConnCache releaseConn
}

// PingContext verifies the connection to the database is still alive.
func (c *Conn) PingContext(ctx context.Context) error

// ExecContext executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (c *Conn) ExecContext(ctx context.Context, query string, args ...any) (Result, error)

// QueryContext executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (c *Conn) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)

// QueryRowContext executes a query that is expected to return at most one row.
// QueryRowContext always returns a non-nil value. Errors are deferred until
// the [*Row.Scan] method is called.
// If the query selects no rows, the [*Row.Scan] will return [ErrNoRows].
// Otherwise, the [*Row.Scan] scans the first selected row and discards
// the rest.
func (c *Conn) QueryRowContext(ctx context.Context, query string, args ...any) *Row

// PrepareContext creates a prepared statement for later queries or executions.
// Multiple queries or executions may be run concurrently from the
// returned statement.
// The caller must call the statement's [*Stmt.Close] method
// when the statement is no longer needed.
//
// The provided context is used for the preparation of the statement, not for the
// execution of the statement.
func (c *Conn) PrepareContext(ctx context.Context, query string) (*Stmt, error)

// Raw executes f exposing the underlying driver connection for the
// duration of f. The driverConn must not be used outside of f.
//
// Once f returns and err is not [driver.ErrBadConn], the [Conn] will continue to be usable
// until [Conn.Close] is called.
func (c *Conn) Raw(f func(driverConn any) error) (err error)

// BeginTx starts a transaction.
//
// The provided context is used until the transaction is committed or rolled back.
// If the context is canceled, the sql package will roll back
// the transaction. [Tx.Commit] will return an error if the context provided to
// BeginTx is canceled.
//
// The provided [TxOptions] is optional and may be nil if defaults should be used.
// If a non-default isolation level is used that the driver doesn't support,
// an error will be returned.
func (c *Conn) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)

// Close returns the connection to the connection pool.
// All operations after a Close will return with [ErrConnDone].
// Close is safe to call concurrently with other operations and will
// block until all other operations finish. It may be useful to first
// cancel any used context and then call close directly after.
func (c *Conn) Close() error

// Tx is an in-progress database transaction.
//
// A transaction must end with a call to [Tx.Commit] or [Tx.Rollback].
//
// After a call to [Tx.Commit] or [Tx.Rollback], all operations on the
// transaction fail with [ErrTxDone].
//
// The statements prepared for a transaction by calling
// the transaction's [Tx.Prepare] or [Tx.Stmt] methods are closed
// by the call to [Tx.Commit] or [Tx.Rollback].
type Tx struct {
	db *DB

	// closemu prevents the transaction from closing while there
	// is an active query. It is held for read during queries
	// and exclusively during close.
	closemu sync.RWMutex

	// dc is owned exclusively until Commit or Rollback, at which point
	// it's returned with putConn.
	dc  *driverConn
	txi driver.Tx

	// releaseConn is called once the Tx is closed to release
	// any held driverConn back to the pool.
	releaseConn func(error)

	// done transitions from false to true exactly once, on Commit
	// or Rollback. once done, all operations fail with
	// ErrTxDone.
	done atomic.Bool

	// keepConnOnRollback is true if the driver knows
	// how to reset the connection's session and if need be discard
	// the connection.
	keepConnOnRollback bool

	// All Stmts prepared for this transaction. These will be closed after the
	// transaction has been committed or rolled back.
	stmts struct {
		sync.Mutex
		v []*Stmt
	}

	// cancel is called after done transitions from 0 to 1.
	cancel func()

	// ctx lives for the life of the transaction.
	ctx context.Context
}

// ErrTxDone is returned by any operation that is performed on a transaction
// that has already been committed or rolled back.
var ErrTxDone = errors.New("sql: transaction has already been committed or rolled back")

// Commit commits the transaction.
func (tx *Tx) Commit() error

// Rollback aborts the transaction.
func (tx *Tx) Rollback() error

// PrepareContext creates a prepared statement for use within a transaction.
//
// The returned statement operates within the transaction and will be closed
// when the transaction has been committed or rolled back.
//
// To use an existing prepared statement on this transaction, see [Tx.Stmt].
//
// The provided context will be used for the preparation of the context, not
// for the execution of the returned statement. The returned statement
// will run in the transaction context.
func (tx *Tx) PrepareContext(ctx context.Context, query string) (*Stmt, error)

// Prepare creates a prepared statement for use within a transaction.
//
// The returned statement operates within the transaction and will be closed
// when the transaction has been committed or rolled back.
//
// To use an existing prepared statement on this transaction, see [Tx.Stmt].
//
// Prepare uses [context.Background] internally; to specify the context, use
// [Tx.PrepareContext].
func (tx *Tx) Prepare(query string) (*Stmt, error)

// StmtContext returns a transaction-specific prepared statement from
// an existing statement.
//
// Example:
//
//	updateMoney, err := db.Prepare("UPDATE balance SET money=money+? WHERE id=?")
//	...
//	tx, err := db.Begin()
//	...
//	res, err := tx.StmtContext(ctx, updateMoney).Exec(123.45, 98293203)
//
// The provided context is used for the preparation of the statement, not for the
// execution of the statement.
//
// The returned statement operates within the transaction and will be closed
// when the transaction has been committed or rolled back.
func (tx *Tx) StmtContext(ctx context.Context, stmt *Stmt) *Stmt

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
//
// The returned statement operates within the transaction and will be closed
// when the transaction has been committed or rolled back.
//
// Stmt uses [context.Background] internally; to specify the context, use
// [Tx.StmtContext].
func (tx *Tx) Stmt(stmt *Stmt) *Stmt

// ExecContext executes a query that doesn't return rows.
// For example: an INSERT and UPDATE.
func (tx *Tx) ExecContext(ctx context.Context, query string, args ...any) (Result, error)

// Exec executes a query that doesn't return rows.
// For example: an INSERT and UPDATE.
//
// Exec uses [context.Background] internally; to specify the context, use
// [Tx.ExecContext].
func (tx *Tx) Exec(query string, args ...any) (Result, error)

// QueryContext executes a query that returns rows, typically a SELECT.
func (tx *Tx) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)

// Query executes a query that returns rows, typically a SELECT.
//
// Query uses [context.Background] internally; to specify the context, use
// [Tx.QueryContext].
func (tx *Tx) Query(query string, args ...any) (*Rows, error)

// QueryRowContext executes a query that is expected to return at most one row.
// QueryRowContext always returns a non-nil value. Errors are deferred until
// [Row]'s Scan method is called.
// If the query selects no rows, the [*Row.Scan] will return [ErrNoRows].
// Otherwise, the [*Row.Scan] scans the first selected row and discards
// the rest.
func (tx *Tx) QueryRowContext(ctx context.Context, query string, args ...any) *Row

// QueryRow executes a query that is expected to return at most one row.
// QueryRow always returns a non-nil value. Errors are deferred until
// [Row]'s Scan method is called.
// If the query selects no rows, the [*Row.Scan] will return [ErrNoRows].
// Otherwise, the [*Row.Scan] scans the first selected row and discards
// the rest.
//
// QueryRow uses [context.Background] internally; to specify the context, use
// [Tx.QueryRowContext].
func (tx *Tx) QueryRow(query string, args ...any) *Row

var (
	_ stmtConnGrabber = &Tx{}
	_ stmtConnGrabber = &Conn{}
)

// Stmt is a prepared statement.
// A Stmt is safe for concurrent use by multiple goroutines.
//
// If a Stmt is prepared on a [Tx] or [Conn], it will be bound to a single
// underlying connection forever. If the [Tx] or [Conn] closes, the Stmt will
// become unusable and all operations will return an error.
// If a Stmt is prepared on a [DB], it will remain usable for the lifetime of the
// [DB]. When the Stmt needs to execute on a new underlying connection, it will
// prepare itself on the new connection automatically.
type Stmt struct {
	// Immutable:
	db        *DB
	query     string
	stickyErr error

	closemu sync.RWMutex

	// If Stmt is prepared on a Tx or Conn then cg is present and will
	// only ever grab a connection from cg.
	// If cg is nil then the Stmt must grab an arbitrary connection
	// from db and determine if it must prepare the stmt again by
	// inspecting css.
	cg   stmtConnGrabber
	cgds *driverStmt

	// parentStmt is set when a transaction-specific statement
	// is requested from an identical statement prepared on the same
	// conn. parentStmt is used to track the dependency of this statement
	// on its originating ("parent") statement so that parentStmt may
	// be closed by the user without them having to know whether or not
	// any transactions are still using it.
	parentStmt *Stmt

	mu     sync.Mutex
	closed bool

	// css is a list of underlying driver statement interfaces
	// that are valid on particular connections. This is only
	// used if cg == nil and one is found that has idle
	// connections. If cg != nil, cgds is always used.
	css []connStmt

	// lastNumClosed is copied from db.numClosed when Stmt is created
	// without tx and closed connections in css are removed.
	lastNumClosed uint64
}

// ExecContext executes a prepared statement with the given arguments and
// returns a [Result] summarizing the effect of the statement.
func (s *Stmt) ExecContext(ctx context.Context, args ...any) (Result, error)

// Exec executes a prepared statement with the given arguments and
// returns a [Result] summarizing the effect of the statement.
//
// Exec uses [context.Background] internally; to specify the context, use
// [Stmt.ExecContext].
func (s *Stmt) Exec(args ...any) (Result, error)

// QueryContext executes a prepared query statement with the given arguments
// and returns the query results as a [*Rows].
func (s *Stmt) QueryContext(ctx context.Context, args ...any) (*Rows, error)

// Query executes a prepared query statement with the given arguments
// and returns the query results as a *Rows.
//
// Query uses [context.Background] internally; to specify the context, use
// [Stmt.QueryContext].
func (s *Stmt) Query(args ...any) (*Rows, error)

// QueryRowContext executes a prepared query statement with the given arguments.
// If an error occurs during the execution of the statement, that error will
// be returned by a call to Scan on the returned [*Row], which is always non-nil.
// If the query selects no rows, the [*Row.Scan] will return [ErrNoRows].
// Otherwise, the [*Row.Scan] scans the first selected row and discards
// the rest.
func (s *Stmt) QueryRowContext(ctx context.Context, args ...any) *Row

// QueryRow executes a prepared query statement with the given arguments.
// If an error occurs during the execution of the statement, that error will
// be returned by a call to Scan on the returned [*Row], which is always non-nil.
// If the query selects no rows, the [*Row.Scan] will return [ErrNoRows].
// Otherwise, the [*Row.Scan] scans the first selected row and discards
// the rest.
//
// Example usage:
//
//	var name string
//	err := nameByUseridStmt.QueryRow(id).Scan(&name)
//
// QueryRow uses [context.Background] internally; to specify the context, use
// [Stmt.QueryRowContext].
func (s *Stmt) QueryRow(args ...any) *Row

// Close closes the statement.
func (s *Stmt) Close() error

// Rows is the result of a query. Its cursor starts before the first row
// of the result set. Use [Rows.Next] to advance from row to row.
type Rows struct {
	dc          *driverConn
	releaseConn func(error)
	rowsi       driver.Rows
	cancel      func()
	closeStmt   *driverStmt

	contextDone atomic.Pointer[error]

	// closemu prevents Rows from closing while there
	// is an active streaming result. It is held for read during non-close operations
	// and exclusively during close.
	//
	// closemu guards lasterr and closed.
	closemu sync.RWMutex
	closed  bool
	lasterr error

	// lastcols is only used in Scan, Next, and NextResultSet which are expected
	// not to be called concurrently.
	lastcols []driver.Value

	// raw is a buffer for RawBytes that persists between Scan calls.
	// This is used when the driver returns a mismatched type that requires
	// a cloning allocation. For example, if the driver returns a *string and
	// the user is scanning into a *RawBytes, we need to copy the string.
	// The raw buffer here lets us reuse the memory for that copy across Scan calls.
	raw []byte

	// closemuScanHold is whether the previous call to Scan kept closemu RLock'ed
	// without unlocking it. It does that when the user passes a *RawBytes scan
	// target. In that case, we need to prevent awaitDone from closing the Rows
	// while the user's still using the memory. See go.dev/issue/60304.
	//
	// It is only used by Scan, Next, and NextResultSet which are expected
	// not to be called concurrently.
	closemuScanHold bool

	// hitEOF is whether Next hit the end of the rows without
	// encountering an error. It's set in Next before
	// returning. It's only used by Next and Err which are
	// expected not to be called concurrently.
	hitEOF bool
}

// Next prepares the next result row for reading with the [Rows.Scan] method. It
// returns true on success, or false if there is no next result row or an error
// happened while preparing it. [Rows.Err] should be consulted to distinguish between
// the two cases.
//
// Every call to [Rows.Scan], even the first one, must be preceded by a call to [Rows.Next].
func (rs *Rows) Next() bool

// NextResultSet prepares the next result set for reading. It reports whether
// there is further result sets, or false if there is no further result set
// or if there is an error advancing to it. The [Rows.Err] method should be consulted
// to distinguish between the two cases.
//
// After calling NextResultSet, the [Rows.Next] method should always be called before
// scanning. If there are further result sets they may not have rows in the result
// set.
func (rs *Rows) NextResultSet() bool

// Err returns the error, if any, that was encountered during iteration.
// Err may be called after an explicit or implicit [Rows.Close].
func (rs *Rows) Err() error

// Columns returns the column names.
// Columns returns an error if the rows are closed.
func (rs *Rows) Columns() ([]string, error)

// ColumnTypes returns column information such as column type, length,
// and nullable. Some information may not be available from some drivers.
func (rs *Rows) ColumnTypes() ([]*ColumnType, error)

// ColumnType contains the name and type of a column.
type ColumnType struct {
	name string

	hasNullable       bool
	hasLength         bool
	hasPrecisionScale bool

	nullable     bool
	length       int64
	databaseType string
	precision    int64
	scale        int64
	scanType     reflect.Type
}

// Name returns the name or alias of the column.
func (ci *ColumnType) Name() string

// Length returns the column type length for variable length column types such
// as text and binary field types. If the type length is unbounded the value will
// be [math.MaxInt64] (any database limits will still apply).
// If the column type is not variable length, such as an int, or if not supported
// by the driver ok is false.
func (ci *ColumnType) Length() (length int64, ok bool)

// DecimalSize returns the scale and precision of a decimal type.
// If not applicable or if not supported ok is false.
func (ci *ColumnType) DecimalSize() (precision, scale int64, ok bool)

// ScanType returns a Go type suitable for scanning into using [Rows.Scan].
// If a driver does not support this property ScanType will return
// the type of an empty interface.
func (ci *ColumnType) ScanType() reflect.Type

// Nullable reports whether the column may be null.
// If a driver does not support this property ok will be false.
func (ci *ColumnType) Nullable() (nullable, ok bool)

// DatabaseTypeName returns the database system name of the column type. If an empty
// string is returned, then the driver type name is not supported.
// Consult your driver documentation for a list of driver data types. [ColumnType.Length] specifiers
// are not included.
// Common type names include "VARCHAR", "TEXT", "NVARCHAR", "DECIMAL", "BOOL",
// "INT", and "BIGINT".
func (ci *ColumnType) DatabaseTypeName() string

// Scan copies the columns in the current row into the values pointed
// at by dest. The number of values in dest must be the same as the
// number of columns in [Rows].
//
// Scan converts columns read from the database into the following
// common Go types and special types provided by the sql package:
//
//	*string
//	*[]byte
//	*int, *int8, *int16, *int32, *int64
//	*uint, *uint8, *uint16, *uint32, *uint64
//	*bool
//	*float32, *float64
//	*interface{}
//	*RawBytes
//	*Rows (cursor value)
//	any type implementing Scanner (see Scanner docs)
//
// In the most simple case, if the type of the value from the source
// column is an integer, bool or string type T and dest is of type *T,
// Scan simply assigns the value through the pointer.
//
// Scan also converts between string and numeric types, as long as no
// information would be lost. While Scan stringifies all numbers
// scanned from numeric database columns into *string, scans into
// numeric types are checked for overflow. For example, a float64 with
// value 300 or a string with value "300" can scan into a uint16, but
// not into a uint8, though float64(255) or "255" can scan into a
// uint8. One exception is that scans of some float64 numbers to
// strings may lose information when stringifying. In general, scan
// floating point columns into *float64.
//
// If a dest argument has type *[]byte, Scan saves in that argument a
// copy of the corresponding data. The copy is owned by the caller and
// can be modified and held indefinitely. The copy can be avoided by
// using an argument of type [*RawBytes] instead; see the documentation
// for [RawBytes] for restrictions on its use.
//
// If an argument has type *interface{}, Scan copies the value
// provided by the underlying driver without conversion. When scanning
// from a source value of type []byte to *interface{}, a copy of the
// slice is made and the caller owns the result.
//
// Source values of type [time.Time] may be scanned into values of type
// *time.Time, *interface{}, *string, or *[]byte. When converting to
// the latter two, [time.RFC3339Nano] is used.
//
// Source values of type bool may be scanned into types *bool,
// *interface{}, *string, *[]byte, or [*RawBytes].
//
// For scanning into *bool, the source may be true, false, 1, 0, or
// string inputs parseable by [strconv.ParseBool].
//
// Scan can also convert a cursor returned from a query, such as
// "select cursor(select * from my_table) from dual", into a
// [*Rows] value that can itself be scanned from. The parent
// select query will close any cursor [*Rows] if the parent [*Rows] is closed.
//
// If any of the first arguments implementing [Scanner] returns an error,
// that error will be wrapped in the returned error.
func (rs *Rows) Scan(dest ...any) error

// Close closes the [Rows], preventing further enumeration. If [Rows.Next] is called
// and returns false and there are no further result sets,
// the [Rows] are closed automatically and it will suffice to check the
// result of [Rows.Err]. Close is idempotent and does not affect the result of [Rows.Err].
func (rs *Rows) Close() error

// Row is the result of calling [DB.QueryRow] to select a single row.
type Row struct {
	// One of these two will be non-nil:
	err  error
	rows *Rows
}

// Scan copies the columns from the matched row into the values
// pointed at by dest. See the documentation on [Rows.Scan] for details.
// If more than one row matches the query,
// Scan uses the first row and discards the rest. If no row matches
// the query, Scan returns [ErrNoRows].
func (r *Row) Scan(dest ...any) error

// Err provides a way for wrapping packages to check for
// query errors without calling [Row.Scan].
// Err returns the error, if any, that was encountered while running the query.
// If this error is not nil, this error will also be returned from [Row.Scan].
func (r *Row) Err() error

// A Result summarizes an executed SQL command.
type Result interface {
	LastInsertId() (int64, error)

	RowsAffected() (int64, error)
}
