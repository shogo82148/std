// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package driverは、package sqlによって使用されるデータベースドライバが実装するインターフェースを定義します。
//
// ほとんどのコードは、[database/sql] パッケージを使用するべきです。
//
// ドライバのインターフェースは時間の経過とともに進化してきました。ドライバは [Connector] と [DriverContext] のインターフェースを実装する必要があります。
// Connector.ConnectとDriver.Openメソッドは、決して [ErrBadConn] を返してはいけません。
// [ErrBadConn] は、 [Validator] 、 [SessionResetter] 、または接続が既に無効な（閉じられたなど）状態にある場合にのみ返されるべきです。
//
// すべての [Conn] の実装は、以下のインターフェースを実装する必要があります：
// [Pinger] 、 [SessionResetter] 、および [Validator] 。
//
// 名前付きパラメータやコンテキストがサポートされている場合、ドライバの [Conn] は以下を実装する必要があります：
// [ExecerContext] 、 [QueryerContext] 、 [ConnPrepareContext] 、および [ConnBeginTx] 。
//
// カスタムデータ型をサポートするためには、 [NamedValueChecker] を実装します。 [NamedValueChecker] は、CheckNamedValueから [ErrRemoveArgument] を返すことで、クエリごとのオプションをパラメータとして受け入れることも可能にします。
//
// 複数の結果セットがサポートされている場合、 [Rows] は [RowsNextResultSet] を実装する必要があります。
// ドライバが返された結果に含まれる型を説明する方法を知っている場合、以下のインターフェースを実装する必要があります：
// [RowsColumnTypeScanType] 、 [RowsColumnTypeDatabaseTypeName] 、 [RowsColumnTypeLength] 、 [RowsColumnTypeNullable] 、および [RowsColumnTypePrecisionScale] 。
// ある行の値は、Rows型を返すこともあり、それはデータベースカーソル値を表すことができます。
//
// [Conn] が [Validator] を実装している場合には、接続が使用後に接続プールに返される前にIsValidメソッドが呼び出されます。
// コネクションプールのエントリーが [SessionResetter] を実装している場合には、別のクエリに再利用される前にResetSessionが呼び出されます。
// 接続が接続プールに返されないで直接再利用される場合は、再利用の前にResetSessionが呼び出されますが、IsValidは呼び出されません。
package driver

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/reflect"
)

// Valueは、ドライバが扱える必要がある値です。
// nil、データベースドライバの [NamedValueChecker] インターフェースで扱われる型、または次のいずれかの型のインスタンスです：
//
//	int64
//	float64
//	bool
//	[]byte
//	string
//	time.Time
//
// ドライバがカーソルをサポートしている場合、返されたValueはこのパッケージの [Rows] インターフェースも実装する場合があります。
// これは、ユーザが "select cursor(select * from my_table) from dual" のようなカーソルを選択した場合に使用されます。
// セレクトの [Rows] がクローズされると、カーソルの [Rows] もクローズされます。
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
// データベースドライバーは、コンテキストへのアクセスと、接続プールの名前の解析を一度だけ行うために、
// 接続ごとに一度ずつではなく、 [DriverContext] を実装することもできます。
type Driver interface {
	Open(name string) (Conn, error)
}

// もし [Driver] が DriverContext を実装している場合、[database/sql.DB] はOpenConnectorを呼び出して [Connector] を取得し、
// その [Connector] のConnectメソッドを呼び出して必要な接続を取得します。
// これにより、接続ごとに [Driver] のOpenメソッドを呼び出すのではなく、名前を1回だけ解析することができ、
// またper-[Conn] コンテキストにアクセスすることもできます。
type DriverContext interface {
	OpenConnector(name string) (Connector, error)
}

// コネクタは、固定の構成でドライバを表し、複数のゴルーチンで使用するための同等の接続を作成できます。
//
// コネクタは [database/sql.OpenDB] に渡すことができ、ドライバは独自の [database/sql.DB] コンストラクタを実装するため、また、 [DriverContext] のOpenConnectorメソッドによって返されることができます。
// これにより、ドライバはコンテキストへのアクセスとドライバ構成の繰り返し解析を避けることができます。
//
// コネクタが [io.Closer] を実装している場合、sqlパッケージの [database/sql.DB.Close] メソッドはCloseを呼び出し、エラー（あれば）を返します。
type Connector interface {
	Connect(context.Context) (Conn, error)

	Driver() Driver
}

// ErrSkipは、一部のオプションのインタフェースメソッドによって、高速経路が利用できないことを実行時に示すために返される場合があります。sqlパッケージは、オプションのインタフェースが実装されていないかのように続行する必要があります。ErrSkipは、明示的に文書化されている場所でのみサポートされます。
var ErrSkip = errors.New("driver: skip fast-path; continue as if unimplemented")

// ErrBadConnは、ドライバが driver.[Conn] が不良な状態であることを示すために、
// [database/sql] パッケージに返すべきです（たとえば、サーバーが接続を早期に閉じたなど）。
// また、[database/sql] パッケージは新しい接続で再試行する必要があります。
//
// 重複した操作を防ぐために、可能性がある場合には、ErrBadConnを返してはいけません。
// データベースサーバーが操作を実行した可能性があっても、ErrBadConnは返してはいけません。
//
// エラーは [errors.Is] を使用してチェックされます。エラーは
// ErrBadConnをラップするか、Is(error) boolメソッドを実装することがあります。
var ErrBadConn = errors.New("driver: bad connection")

// Pingerは [Conn] によって実装される可能性のあるオプションのインターフェースです。
//
// [Conn] がPingerを実装していない場合、 [database/sql.DB.Ping] および [database/sql.DB.PingContext] は少なくとも1つの [Conn] が利用可能かどうかを確認します。
//
// Conn.Pingが [ErrBadConn] を返す場合、 [database/sql.DB.Ping] および [database/sql.DB.PingContext] は [Conn] をプールから削除します。
type Pinger interface {
	Ping(ctx context.Context) error
}

// Execerは [Conn] によって実装されるかもしれないオプションのインターフェースです。
//
// もし [Conn] が [ExecerContext] または [Execer] のどちらの実装も持っていない場合、
// [database/sql.DB.Exec] はまずクエリを準備し、ステートメントを実行し、そしてステートメントを閉じます。
//
// Execは [ErrSkip] を返す場合があります。
//
// Deprecated: ドライバは代わりに [ExecerContext] を実装するべきです。
type Execer interface {
	Exec(query string, args []Value) (Result, error)
}

// ExecerContextは [Conn] によって実装されるかもしれないオプションのインターフェースです。
//
// [Conn] が [ExecerContext] を実装していない場合、[database/sql.DB.Exec] は [Execer] にフォールバックします。
// もしConnがExecerも実装していない場合、[database/sql.DB.Exec] はまずクエリを準備し、ステートメントを実行してからステートメントを閉じます。
//
// ExecContextは [ErrSkip] を返すことがあります。
//
// ExecContextはコンテキストのタイムアウトを尊重し、コンテキストがキャンセルされたら返ります。
type ExecerContext interface {
	ExecContext(ctx context.Context, query string, args []NamedValue) (Result, error)
}

// Queryerは [Conn] によって実装されるかもしれないオプションのインターフェースです。
//
// [Conn] が [QueryerContext] も [Queryer] でも実装していない場合、 [database/sql.DB.Query] はまずクエリを準備し、ステートメントを実行してからステートメントを閉じます。
//
// Queryは [ErrSkip] を返すことがあります。
//
// Deprecated: ドライバは代わりに [QueryerContext] を実装するべきです。
type Queryer interface {
	Query(query string, args []Value) (Rows, error)
}

// QueryerContextは、[Conn] によって実装されるかもしれないオプションのインターフェースです。
//
// [Conn] がQueryerContextを実装していない場合、 [database/sql.DB.Query] は [Queryer] にフォールバックします。
// もし、Connが [Queryer] を実装していない場合、 [database/sql.DB.Query] はまずクエリを準備し、ステートメントを実行してからステートメントを閉じます。
//
// QueryContextは [ErrSkip] を返す場合があります。
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

// ConnPrepareContextはコンテキストを使用して [Conn] インターフェースを拡張します。
type ConnPrepareContext interface {
	PrepareContext(ctx context.Context, query string) (Stmt, error)
}

// IsolationLevelは [TxOptions] に保存されるトランザクション分離レベルです。
//
// この型は、[database/sql.IsolationLevel] と一緒に定義された値と同じものと考えられるべきです。
type IsolationLevel int

// TxOptionsはトランザクションのオプションを保持します。
//
// この型は [database/sql.TxOptions] と同一と見なされるべきです。
type TxOptions struct {
	Isolation IsolationLevel
	ReadOnly  bool
}

// ConnBeginTxは、コンテキストと [TxOptions] を追加して [Conn] インタフェースを拡張します。
type ConnBeginTx interface {
	BeginTx(ctx context.Context, opts TxOptions) (Tx, error)
}

// SessionResetterは、 [Conn] によって実装される可能性があります。これにより、ドライバは接続に関連付けられたセッション状態をリセットし、悪い接続を通知することができます。
type SessionResetter interface {
	ResetSession(ctx context.Context) error
}

// Validatorは、 [Conn] によって実装されることがあります。これにより、ドライバーは接続が有効であるか、破棄すべきかを示すことができます。
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

// Stmtはプリペアドステートメントです。これは [Conn] にバインドされており、複数のゴルーチンで同時に使用されません。
type Stmt interface {
	Close() error

	NumInput() int

	Exec(args []Value) (Result, error)

	Query(args []Value) (Rows, error)
}

// StmtExecContextはコンテキストを提供することにより、 [Stmt] インターフェースを拡張します。
type StmtExecContext interface {
	ExecContext(ctx context.Context, args []NamedValue) (Result, error)
}

// StmtQueryContextはコンテキストを持つQueryを提供することにより、 [Stmt] インターフェースを強化します。
type StmtQueryContext interface {
	QueryContext(ctx context.Context, args []NamedValue) (Rows, error)
}

// ErrRemoveArgumentは、 [NamedValueChecker] から返されることがあります。
// これは、[database/sql] パッケージに対して引数をドライバのクエリインターフェースに渡さないよう指示するためです。
// クエリ固有のオプションやSQLクエリ引数ではない構造体を受け入れる場合に返します。
var ErrRemoveArgument = errors.New("driver: remove argument from query")

// NamedValueCheckerは [Conn] または [Stmt] によってオプションで実装されることがあります。これにより、ドライバはデフォルトの [Value] タイプを超えたGoおよびデータベースのタイプを処理するための制御を提供します。
//
// [database/sql] パッケージは、値チェッカーを以下の順序でチェックし、最初に一致したもので停止します： Stmt.NamedValueChecker、Conn.NamedValueChecker、Stmt.ColumnConverter、 [DefaultParameterConverter] 。
//
// CheckNamedValueが [ErrRemoveArgument] を返す場合、 [NamedValue] は最終的なクエリ引数に含まれません。これはクエリ自体に特殊なオプションを渡すために使用される場合があります。
//
// [ErrSkip] が返された場合、列コンバーターのエラーチェックパスが引数に使用されます。ドライバは、独自の特殊なケースを使い果たした後に [ErrSkip] を返すことを望むかもしれません。
type NamedValueChecker interface {
	CheckNamedValue(*NamedValue) error
}

// ColumnConverterは、ステートメントが自身の列の型を認識しており、任意の型からドライバーの [Value] に変換できる場合、Stmtによってオプションで実装されることがあります。
// Deprecated: ドライバーは [NamedValueChecker] を実装する必要があります。
type ColumnConverter interface {
	ColumnConverter(idx int) ValueConverter
}

// Rowsは実行されたクエリの結果に対するイテレータです。
type Rows interface {
	Columns() []string

	Close() error

	Next(dest []Value) error
}

// RowsNextResultSetは、ドライバに次の結果セットに進むようにシグナルを送る方法を提供するために [Rows] インターフェースを拡張しています。
type RowsNextResultSet interface {
	Rows

	HasNextResultSet() bool

	NextResultSet() error
}

// RowsColumnTypeScanTypeは、 [Rows] によって実装されるかもしれません。スキャンに使用できる値の型を返す必要があります。例えば、データベースのカラムタイプが「bigint」の場合、これは「 [reflect.TypeOf](int64(0)) 」を返すべきです。
type RowsColumnTypeScanType interface {
	Rows
	ColumnTypeScanType(index int) reflect.Type
}

// RowsColumnTypeDatabaseTypeNameは [Rows] によって実装されるかもしれません。長さを除いたデータベースシステムのタイプ名を返す必要があります。タイプ名は大文字であるべきです。
// 返される型の例: "VARCHAR", "NVARCHAR", "VARCHAR2", "CHAR", "TEXT",
// "DECIMAL", "SMALLINT", "INT", "BIGINT", "BOOL", "[]BIGINT", "JSONB", "XML",
// "TIMESTAMP"。
type RowsColumnTypeDatabaseTypeName interface {
	Rows
	ColumnTypeDatabaseTypeName(index int) string
}

// RowsColumnTypeLengthは、 [Rows] によって実装されるかもしれません。カラムが可変長の場合、カラムタイプの長さを返す必要があります。カラムが可変長のタイプでない場合、okはfalseを返す必要があります。
// システムの制限以外で長さが制限されていない場合、[math.MaxInt64] を返す必要があります。以下は、さまざまなタイプの戻り値の例です：
//
//	TEXT          (math.MaxInt64, true)
//	varchar(10)   (10, true)
//	nvarchar(10)  (10, true)
//	decimal       (0, false)
//	int           (0, false)
//	bytea(30)     (30, true)
type RowsColumnTypeLength interface {
	Rows
	ColumnTypeLength(index int) (length int64, ok bool)
}

// RowsColumnTypeNullableは、 [Rows] によって実装される可能性があります。カラムがnullである可能性がある場合は、nullableの値をtrueにする必要があります。カラムがnullでないことが確認されている場合は、falseにする必要があります。
// カラムのヌラビリティが不明な場合は、okをfalseにしてください。
type RowsColumnTypeNullable interface {
	Rows
	ColumnTypeNullable(index int) (nullable, ok bool)
}

// RowsColumnTypePrecisionScaleは [Rows] によって実装されるかもしれません。それはデシマル型の精度とスケールを返すべきです。該当しない場合、okはfalseであるべきです。
// 以下に、さまざまな型の戻り値の例を示します:
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

// RowsAffectedは、行数を変更するINSERTまたはUPDATE操作に対する [Result] を実装します。
type RowsAffected int64

var _ Result = RowsAffected(0)

func (RowsAffected) LastInsertId() (int64, error)

func (v RowsAffected) RowsAffected() (int64, error)

// ResultNoRowsは、DDLコマンド（CREATE TABLEなど）が成功した場合にドライバーが返すための事前定義された [Result] です。これは、LastInsertIdと [RowsAffected] の両方に対してエラーを返します。
var ResultNoRows noRows

var _ Result = noRows{}
