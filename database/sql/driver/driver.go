// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package driverは、package sqlによって使用されるデータベースドライバが実装するインターフェースを定義します。
//
// ほとんどのコードは、package sqlを使用するべきです。
//
<<<<<<< HEAD
// ドライバのインターフェースは時間の経過とともに進化してきました。ドライバはConnectorとDriverContextのインターフェースを実装する必要があります。
// Connector.ConnectとDriver.Openメソッドは、決してErrBadConnを返してはいけません。
// ErrBadConnは、Validator、SessionResetter、または接続が既に無効な（閉じられたなど）状態にある場合にのみ返されるべきです。
//
// すべてのConnの実装は、以下のインターフェースを実装する必要があります：
// Pinger、SessionResetter、およびValidator。
//
// 名前付きパラメータやコンテキストがサポートされている場合、ドライバのConnは以下を実装する必要があります：
// ExecerContext、QueryerContext、ConnPrepareContext、およびConnBeginTx。
//
// カスタムデータ型をサポートするためには、NamedValueCheckerを実装します。NamedValueCheckerは、CheckNamedValueからErrRemoveArgumentを返すことで、クエリごとのオプションをパラメータとして受け入れることも可能にします。
//
// 複数の結果セットがサポートされている場合、RowsはRowsNextResultSetを実装する必要があります。
// ドライバが返された結果に含まれる型を説明する方法を知っている場合、以下のインターフェースを実装する必要があります：
// RowsColumnTypeScanType、RowsColumnTypeDatabaseTypeName、RowsColumnTypeLength、RowsColumnTypeNullable、およびRowsColumnTypePrecisionScale。
// ある行の値は、Rows型を返すこともあり、それはデータベースカーソル値を表すことができます。
=======
// The driver interface has evolved over time. Drivers should implement
// [Connector] and [DriverContext] interfaces.
// The Connector.Connect and Driver.Open methods should never return [ErrBadConn].
// [ErrBadConn] should only be returned from [Validator], [SessionResetter], or
// a query method if the connection is already in an invalid (e.g. closed) state.
//
// All [Conn] implementations should implement the following interfaces:
// [Pinger], [SessionResetter], and [Validator].
//
// If named parameters or context are supported, the driver's [Conn] should implement:
// [ExecerContext], [QueryerContext], [ConnPrepareContext], and [ConnBeginTx].
//
// To support custom data types, implement [NamedValueChecker]. [NamedValueChecker]
// also allows queries to accept per-query options as a parameter by returning
// [ErrRemoveArgument] from CheckNamedValue.
//
// If multiple result sets are supported, [Rows] should implement [RowsNextResultSet].
// If the driver knows how to describe the types present in the returned result
// it should implement the following interfaces: [RowsColumnTypeScanType],
// [RowsColumnTypeDatabaseTypeName], [RowsColumnTypeLength], [RowsColumnTypeNullable],
// and [RowsColumnTypePrecisionScale]. A given row value may also return a [Rows]
// type, which may represent a database cursor value.
>>>>>>> upstream/master
//
// 接続が使用後に接続プールに返される前に、実装されている場合にはIsValidが呼び出されます。
// 別のクエリに再利用される前に、実装されている場合にはResetSessionが呼び出されます。
// 接続が接続プールに返されないで直接再利用される場合は、再利用の前にResetSessionが呼び出されますが、IsValidは呼び出されません。
package driver

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/reflect"
)

<<<<<<< HEAD
// Valueは、ドライバが扱える必要がある値です。
// nil、データベースドライバのNamedValueCheckerインターフェースで扱われる型、または次のいずれかの型のインスタンスです：
=======
// Value is a value that drivers must be able to handle.
// It is either nil, a type handled by a database driver's [NamedValueChecker]
// interface, or an instance of one of these types:
>>>>>>> upstream/master
//
//	int64
//	float64
//	bool
//	[]byte
//	string
//	time.Time
//
<<<<<<< HEAD
// ドライバがカーソルをサポートしている場合、返されたValueはこのパッケージのRowsインターフェースも実装する場合があります。
// これは、ユーザが "select cursor(select * from my_table) from dual" のようなカーソルを選択した場合に使用されます。
// セレクトのRowsがクローズされると、カーソルのRowsもクローズされます。
=======
// If the driver supports cursors, a returned Value may also implement the [Rows] interface
// in this package. This is used, for example, when a user selects a cursor
// such as "select cursor(select * from my_table) from dual". If the [Rows]
// from the select is closed, the cursor [Rows] will also be closed.
>>>>>>> upstream/master
type Value any

// NamedValueは値の名前と値を保持します。
type NamedValue struct {

	// もし Name が空でなければ、パラメータの識別子に使用されるべきであり、
	// 順序位置ではないです。
	//
	// Name にはシンボルのプレフィックスはつきません。
	Name string

	// パラメータの序数位置は常に1から始まります。
	Ordinal int

	// Valueはパラメーターの値です。
	Value Value
}

// Driverはデータベースドライバーによって実装される必要があるインターフェースです。
//
<<<<<<< HEAD
// データベースドライバーは、コンテキストへのアクセスと、接続プールの名前の解析を一度だけ行うために、
// 接続ごとに一度ずつではなく、DriverContextを実装することもできます。
=======
// Database drivers may implement [DriverContext] for access
// to contexts and to parse the name only once for a pool of connections,
// instead of once per connection.
>>>>>>> upstream/master
type Driver interface {
	Open(name string) (Conn, error)
}

<<<<<<< HEAD
// もしDriverがDriverContextを実装している場合、sql.DBはOpenConnectorを呼び出してConnectorを取得し、
// そのConnectorのConnectメソッドを呼び出して必要な接続を取得します。
// これにより、接続ごとにDriverのOpenメソッドを呼び出すのではなく、名前を1回だけ解析することができ、
// またper-Connコンテキストにアクセスすることもできます。
=======
// If a [Driver] implements [DriverContext], then sql.DB will call
// OpenConnector to obtain a [Connector] and then invoke
// that [Connector]'s Connect method to obtain each needed connection,
// instead of invoking the [Driver]'s Open method for each connection.
// The two-step sequence allows drivers to parse the name just once
// and also provides access to per-Conn contexts.
>>>>>>> upstream/master
type DriverContext interface {
	OpenConnector(name string) (Connector, error)
}

// コネクタは、固定の構成でドライバを表し、複数のゴルーチンで使用するための同等の接続を作成できます。
//
<<<<<<< HEAD
// コネクタはsql.OpenDBに渡すことができ、ドライバは独自のsql.DBコンストラクタを実装するため、また、DriverContextのOpenConnectorメソッドによって返されることができます。これにより、ドライバはコンテキストへのアクセスとドライバ構成の繰り返し解析を避けることができます。
//
// コネクタがio.Closerを実装している場合、sqlパッケージのDB.CloseメソッドはCloseを呼び出し、エラー（あれば）を返します。
=======
// A Connector can be passed to [database/sql.OpenDB], to allow drivers
// to implement their own sql.DB constructors, or returned by
// [DriverContext]'s OpenConnector method, to allow drivers
// access to context and to avoid repeated parsing of driver
// configuration.
//
// If a Connector implements [io.Closer], the [database/sql.DB.Close]
// method will call Close and return error (if any).
>>>>>>> upstream/master
type Connector interface {
	Connect(context.Context) (Conn, error)

	Driver() Driver
}

// ErrSkipは、一部のオプションのインタフェースメソッドによって、高速経路が利用できないことを実行時に示すために返される場合があります。sqlパッケージは、オプションのインタフェースが実装されていないかのように続行する必要があります。ErrSkipは、明示的に文書化されている場所でのみサポートされます。
var ErrSkip = errors.New("driver: skip fast-path; continue as if unimplemented")

<<<<<<< HEAD
// ErrBadConnは、ドライバがドライバ.Connが不良な状態であることを示すために、
// sqlパッケージに返すべきです（たとえば、サーバーが接続を早期に閉じたなど）。
// また、sqlパッケージは新しい接続で再試行する必要があります。
=======
// ErrBadConn should be returned by a driver to signal to the sql
// package that a [driver.Conn] is in a bad state (such as the server
// having earlier closed the connection) and the sql package should
// retry on a new connection.
>>>>>>> upstream/master
//
// 重複した操作を防ぐために、可能性がある場合には、ErrBadConnを返してはいけません。
// データベースサーバーが操作を実行した可能性があっても、ErrBadConnは返してはいけません。
//
// エラーはerrors.Isを使用してチェックされます。エラーは
// ErrBadConnをラップするか、Is(error) boolメソッドを実装することがあります。
var ErrBadConn = errors.New("driver: bad connection")

<<<<<<< HEAD
// PingerはConnによって実装される可能性のあるオプションのインターフェースです。
//
// ConnがPingerを実装していない場合、sqlパッケージのDB.PingおよびDB.PingContextは少なくとも1つのConnが利用可能かどうかを確認します。
//
// Conn.PingがErrBadConnを返す場合、DB.PingおよびDB.PingContextはConnをプールから削除します。
=======
// Pinger is an optional interface that may be implemented by a [Conn].
//
// If a [Conn] does not implement Pinger, the [database/sql.DB.Ping] and
// [database/sql.DB.PingContext] will check if there is at least one [Conn] available.
//
// If Conn.Ping returns [ErrBadConn], [database/sql.DB.Ping] and [database/sql.DB.PingContext] will remove
// the [Conn] from pool.
>>>>>>> upstream/master
type Pinger interface {
	Ping(ctx context.Context) error
}

// ExecerはConnによって実装されるかもしれないオプションのインターフェースです。
//
<<<<<<< HEAD
// もしConnがExecerContextまたはExecerのどちらの実装も持っていない場合、
// sqlパッケージのDB.Execはまずクエリを準備し、ステートメントを実行し、そしてステートメントを閉じます。
//
// ExecはErrSkipを返す場合があります。
//
// 廃止予定: ドライバは代わりにExecerContextを実装するべきです。
=======
// If a [Conn] implements neither [ExecerContext] nor Execer,
// the [database/sql.DB.Exec] will first prepare a query, execute the statement,
// and then close the statement.
//
// Exec may return [ErrSkip].
//
// Deprecated: Drivers should implement [ExecerContext] instead.
>>>>>>> upstream/master
type Execer interface {
	Exec(query string, args []Value) (Result, error)
}

<<<<<<< HEAD
// ExecerContextはConnによって実装されるかもしれないオプションのインターフェースです。
//
// ConnがExecerContextを実装していない場合、sqlパッケージのDB.ExecはExecerにフォールバックします。
// もしConnがExecerも実装していない場合、DB.Execはまずクエリを準備し、ステートメントを実行してからステートメントを閉じます。
//
// ExecContextはErrSkipを返すことがあります。
=======
// ExecerContext is an optional interface that may be implemented by a [Conn].
//
// If a Conn does not implement ExecerContext, the [database/sql.DB.Exec]
// will fall back to [Execer]; if the Conn does not implement Execer either,
// [database/sql.DB.Exec] will first prepare a query, execute the statement, and then
// close the statement.
//
// ExecContext may return [ErrSkip].
>>>>>>> upstream/master
//
// ExecContextはコンテキストのタイムアウトを尊重し、コンテキストがキャンセルされたら返ります。
type ExecerContext interface {
	ExecContext(ctx context.Context, query string, args []NamedValue) (Result, error)
}

<<<<<<< HEAD
// QueryerはConnによって実装されるかもしれないオプションのインターフェースです。
//
// ConnがQueryerContextでもQueryerでも実装していない場合、sqlパッケージのDB.Queryはまずクエリを準備し、ステートメントを実行してからステートメントを閉じます。
//
// QueryはErrSkipを返すことがあります。
//
// 廃止予定です: ドライバは代わりにQueryerContextを実装するべきです。
=======
// Queryer is an optional interface that may be implemented by a [Conn].
//
// If a Conn implements neither [QueryerContext] nor Queryer,
// the [database/sql.DB.Query] will first prepare a query, execute the statement,
// and then close the statement.
//
// Query may return [ErrSkip].
//
// Deprecated: Drivers should implement [QueryerContext] instead.
>>>>>>> upstream/master
type Queryer interface {
	Query(query string, args []Value) (Rows, error)
}

// QueryerContextは、Connによって実装されるかもしれないオプションのインターフェースです。
//
<<<<<<< HEAD
// ConnがQueryerContextを実装していない場合、sqlパッケージのDB.QueryはQueryerにフォールバックします。
// もし、ConnがQueryerを実装していない場合、DB.Queryはまずクエリを準備し、ステートメントを実行してからステートメントを閉じます。
//
// QueryContextはErrSkipを返す場合があります。
=======
// If a [Conn] does not implement QueryerContext, the [database/sql.DB.Query]
// will fall back to [Queryer]; if the [Conn] does not implement [Queryer] either,
// [database/sql.DB.Query] will first prepare a query, execute the statement, and then
// close the statement.
//
// QueryContext may return [ErrSkip].
>>>>>>> upstream/master
//
// QueryContextはコンテキストのタイムアウトに従う必要があり、コンテキストがキャンセルされた場合にreturnします。
type QueryerContext interface {
	QueryContext(ctx context.Context, query string, args []NamedValue) (Rows, error)
}

// Connはデータベースへの接続です。同時に複数のゴルーチンで使用されません。
//
// Connは状態を保持していると想定されています。
type Conn interface {
	Prepare(query string) (Stmt, error)

	Close() error

	Begin() (Tx, error)
}

<<<<<<< HEAD
// ConnPrepareContextはコンテキストを使用してConnインターフェースを拡張します。
=======
// ConnPrepareContext enhances the [Conn] interface with context.
>>>>>>> upstream/master
type ConnPrepareContext interface {
	PrepareContext(ctx context.Context, query string) (Stmt, error)
}

<<<<<<< HEAD
// IsolationLevelはTxOptionsに保存されるトランザクション分離レベルです。
=======
// IsolationLevel is the transaction isolation level stored in [TxOptions].
>>>>>>> upstream/master
//
// この型は、sql.IsolationLevelと一緒に定義された値と同じものと考えられるべきです。
type IsolationLevel int

// TxOptionsはトランザクションのオプションを保持します。
//
<<<<<<< HEAD
// この型はsql.TxOptionsと同一と見なされるべきです。
=======
// This type should be considered identical to [database/sql.TxOptions].
>>>>>>> upstream/master
type TxOptions struct {
	Isolation IsolationLevel
	ReadOnly  bool
}

<<<<<<< HEAD
// ConnBeginTxは、コンテキストとTxOptionsを追加してConnインタフェースを拡張します。
=======
// ConnBeginTx enhances the [Conn] interface with context and [TxOptions].
>>>>>>> upstream/master
type ConnBeginTx interface {
	BeginTx(ctx context.Context, opts TxOptions) (Tx, error)
}

<<<<<<< HEAD
// SessionResetterは、Connによって実装される可能性があります。これにより、ドライバは接続に関連付けられたセッション状態をリセットし、悪い接続を通知することができます。
=======
// SessionResetter may be implemented by [Conn] to allow drivers to reset the
// session state associated with the connection and to signal a bad connection.
>>>>>>> upstream/master
type SessionResetter interface {
	ResetSession(ctx context.Context) error
}

<<<<<<< HEAD
// Validaterは、Connによって実装されることがあります。これにより、ドライバーは接続が有効であるか、破棄すべきかを示すことができます。
=======
// Validator may be implemented by [Conn] to allow drivers to
// signal if a connection is valid or if it should be discarded.
>>>>>>> upstream/master
//
// 実装されている場合、ドライバーはクエリから基礎となるエラーを返すことができます。たとえ接続プールによって接続が破棄されるべきであってもです。
type Validator interface {
	IsValid() bool
}

// Resultはクエリの実行結果です。
type Result interface {
	LastInsertId() (int64, error)

	RowsAffected() (int64, error)
}

<<<<<<< HEAD
// Stmtはプリペアドステートメントです。これはConnにバインドされており、複数のゴルーチンで同時に使用されません。
=======
// Stmt is a prepared statement. It is bound to a [Conn] and not
// used by multiple goroutines concurrently.
>>>>>>> upstream/master
type Stmt interface {
	Close() error

	NumInput() int

	Exec(args []Value) (Result, error)

	Query(args []Value) (Rows, error)
}

<<<<<<< HEAD
// StmtExecContextはコンテキストを提供することにより、Stmtインターフェースを拡張します。
=======
// StmtExecContext enhances the [Stmt] interface by providing Exec with context.
>>>>>>> upstream/master
type StmtExecContext interface {
	ExecContext(ctx context.Context, args []NamedValue) (Result, error)
}

<<<<<<< HEAD
// StmtQueryContextはコンテキストを持つQueryを提供することにより、Stmtインターフェースを強化します。
=======
// StmtQueryContext enhances the [Stmt] interface by providing Query with context.
>>>>>>> upstream/master
type StmtQueryContext interface {
	QueryContext(ctx context.Context, args []NamedValue) (Rows, error)
}

<<<<<<< HEAD
// ErrRemoveArgumentは、NamedValueCheckerから返されることがあります。
// これは、sqlパッケージに対して引数をドライバのクエリインターフェースに渡さないよう指示するためです。
// クエリ固有のオプションやSQLクエリ引数ではない構造体を受け入れる場合に返します。
var ErrRemoveArgument = errors.New("driver: remove argument from query")

// NamedValueCheckerはConnまたはStmtによってオプションで実装されることがあります。これにより、ドライバはデフォルトのValuesタイプを超えたGoおよびデータベースのタイプを処理するための制御を提供します。
// sqlパッケージは、値チェッカーを以下の順序でチェックし、最初に一致したもので停止します： Stmt.NamedValueChecker、Conn.NamedValueChecker、Stmt.ColumnConverter、DefaultParameterConverter。
// CheckNamedValueがErrRemoveArgumentを返す場合、NamedValueは最終的なクエリ引数に含まれません。これはクエリ自体に特殊なオプションを渡すために使用される場合があります。
// ErrSkipが返された場合、列コンバーターのエラーチェックパスが引数に使用されます。ドライバは、独自の特殊なケースを使い果たした後にErrSkipを返すことを望むかもしれません。
=======
// ErrRemoveArgument may be returned from [NamedValueChecker] to instruct the
// sql package to not pass the argument to the driver query interface.
// Return when accepting query specific options or structures that aren't
// SQL query arguments.
var ErrRemoveArgument = errors.New("driver: remove argument from query")

// NamedValueChecker may be optionally implemented by [Conn] or [Stmt]. It provides
// the driver more control to handle Go and database types beyond the default
// Values types allowed.
//
// The sql package checks for value checkers in the following order,
// stopping at the first found match: [database/sql.Stmt.NamedValueChecker],
// Conn.NamedValueChecker, [database/sql.Stmt.ColumnConverter], [DefaultParameterConverter].
//
// If CheckNamedValue returns [ErrRemoveArgument], the [NamedValue] will not be included in
// the final query arguments. This may be used to pass special options to
// the query itself.
//
// If [ErrSkip] is returned the column converter error checking
// path is used for the argument. Drivers may wish to return [ErrSkip] after
// they have exhausted their own special cases.
>>>>>>> upstream/master
type NamedValueChecker interface {
	CheckNamedValue(*NamedValue) error
}

<<<<<<< HEAD
// ColumnConverterは、ステートメントが自身の列の型を認識しており、任意の型からドライバーのValueに変換できる場合、Stmtによってオプションで実装されることがあります。
// 廃止予定 : ドライバーはNamedValueCheckerを実装する必要があります。
=======
// ColumnConverter may be optionally implemented by [Stmt] if the
// statement is aware of its own columns' types and can convert from
// any type to a driver [Value].
//
// Deprecated: Drivers should implement [NamedValueChecker].
>>>>>>> upstream/master
type ColumnConverter interface {
	ColumnConverter(idx int) ValueConverter
}

// Rowsは実行されたクエリの結果に対するイテレータです。
type Rows interface {
	Columns() []string

	Close() error

	Next(dest []Value) error
}

<<<<<<< HEAD
// RowsNextResultSetは、ドライバに次の結果セットに進むようにシグナルを送る方法を提供するためにRowsインターフェースを拡張しています。
=======
// RowsNextResultSet extends the [Rows] interface by providing a way to signal
// the driver to advance to the next result set.
>>>>>>> upstream/master
type RowsNextResultSet interface {
	Rows

	HasNextResultSet() bool

	NextResultSet() error
}

<<<<<<< HEAD
// RowsColumnTypeScanTypeは、Rowsによって実装されるかもしれません。スキャンに使用できる値の型を返す必要があります。例えば、データベースのカラムタイプが「bigint」の場合、これは「reflect.TypeOf(int64(0))」を返すべきです。
=======
// RowsColumnTypeScanType may be implemented by [Rows]. It should return
// the value type that can be used to scan types into. For example, the database
// column type "bigint" this should return "[reflect.TypeOf](int64(0))".
>>>>>>> upstream/master
type RowsColumnTypeScanType interface {
	Rows
	ColumnTypeScanType(index int) reflect.Type
}

<<<<<<< HEAD
// RowsColumnTypeDatabaseTypeNameはRowsによって実装されるかもしれません。長さを除いたデータベースシステムのタイプ名を返す必要があります。タイプ名は大文字であるべきです。
// 返される型の例: "VARCHAR", "NVARCHAR", "VARCHAR2", "CHAR", "TEXT",
=======
// RowsColumnTypeDatabaseTypeName may be implemented by [Rows]. It should return the
// database system type name without the length. Type names should be uppercase.
// Examples of returned types: "VARCHAR", "NVARCHAR", "VARCHAR2", "CHAR", "TEXT",
>>>>>>> upstream/master
// "DECIMAL", "SMALLINT", "INT", "BIGINT", "BOOL", "[]BIGINT", "JSONB", "XML",
// "TIMESTAMP"。
type RowsColumnTypeDatabaseTypeName interface {
	Rows
	ColumnTypeDatabaseTypeName(index int) string
}

<<<<<<< HEAD
// RowsColumnTypeLengthは、Rowsによって実装されるかもしれません。カラムが可変長の場合、カラムタイプの長さを返す必要があります。カラムが可変長のタイプでない場合、okはfalseを返す必要があります。システムの制限以外で長さが制限されていない場合、math.MaxInt64を返す必要があります。以下は、さまざまなタイプの戻り値の例です：
// TEXT（math.MaxInt64、true）
// varchar(10)（10、true）
// nvarchar(10)（10、true）
// decimal（0、false）
// int（0、false）
// bytea(30)（30、true）
=======
// RowsColumnTypeLength may be implemented by [Rows]. It should return the length
// of the column type if the column is a variable length type. If the column is
// not a variable length type ok should return false.
// If length is not limited other than system limits, it should return math.MaxInt64.
// The following are examples of returned values for various types:
//
//	TEXT          (math.MaxInt64, true)
//	varchar(10)   (10, true)
//	nvarchar(10)  (10, true)
//	decimal       (0, false)
//	int           (0, false)
//	bytea(30)     (30, true)
>>>>>>> upstream/master
type RowsColumnTypeLength interface {
	Rows
	ColumnTypeLength(index int) (length int64, ok bool)
}

<<<<<<< HEAD
// RowsColumnTypeNullableは、Rowsによって実装される可能性があります。カラムがnullである可能性がある場合は、nullableの値をtrueにする必要があります。カラムがnullでないことが確認されている場合は、falseにする必要があります。
// カラムのヌラビリティが不明な場合は、okをfalseにしてください。
=======
// RowsColumnTypeNullable may be implemented by [Rows]. The nullable value should
// be true if it is known the column may be null, or false if the column is known
// to be not nullable.
// If the column nullability is unknown, ok should be false.
>>>>>>> upstream/master
type RowsColumnTypeNullable interface {
	Rows
	ColumnTypeNullable(index int) (nullable, ok bool)
}

<<<<<<< HEAD
// RowsColumnTypePrecisionScaleはRowsによって実装されるかもしれません。それはデシマル型の精度とスケールを返すべきです。該当しない場合、okはfalseであるべきです。
// 以下に、さまざまな型の戻り値の例を示します:
=======
// RowsColumnTypePrecisionScale may be implemented by [Rows]. It should return
// the precision and scale for decimal types. If not applicable, ok should be false.
// The following are examples of returned values for various types:
>>>>>>> upstream/master
//
//	decimal(38, 4)    (38, 4, true)
//	int               (0, 0, false)
//	decimal           (math.MaxInt64, math.MaxInt64, true)
type RowsColumnTypePrecisionScale interface {
	Rows
	ColumnTypePrecisionScale(index int) (precision, scale int64, ok bool)
}

// Txはトランザクションです。
type Tx interface {
	Commit() error
	Rollback() error
}

<<<<<<< HEAD
// RowsAffectedは、行数を変更するINSERTまたはUPDATE操作に対する Result を実装します。
=======
// RowsAffected implements [Result] for an INSERT or UPDATE operation
// which mutates a number of rows.
>>>>>>> upstream/master
type RowsAffected int64

var _ Result = RowsAffected(0)

func (RowsAffected) LastInsertId() (int64, error)

func (v RowsAffected) RowsAffected() (int64, error)

<<<<<<< HEAD
// ResultNoRowsは、DDLコマンド（CREATE TABLEなど）が成功した場合にドライバーが返すための事前定義された結果です。これは、LastInsertIdとRowsAffectedの両方に対してエラーを返します。
=======
// ResultNoRows is a pre-defined [Result] for drivers to return when a DDL
// command (such as a CREATE TABLE) succeeds. It returns an error for both
// LastInsertId and [RowsAffected].
>>>>>>> upstream/master
var ResultNoRows noRows

var _ Result = noRows{}
