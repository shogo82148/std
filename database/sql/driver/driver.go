// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package driverは、package sqlによって使用されるデータベースドライバが実装するインターフェースを定義します。
//
// ほとんどのコードは、package sqlを使用するべきです。
//
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

// Valueは、ドライバが扱える必要がある値です。
// nil、データベースドライバのNamedValueCheckerインターフェースで扱われる型、または次のいずれかの型のインスタンスです：
//
//	int64
//	float64
//	bool
//	[]byte
//	string
//	time.Time
//
// ドライバがカーソルをサポートしている場合、返されたValueはこのパッケージのRowsインターフェースも実装する場合があります。
// これは、ユーザが "select cursor(select * from my_table) from dual" のようなカーソルを選択した場合に使用されます。
// セレクトのRowsがクローズされると、カーソルのRowsもクローズされます。
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
// 接続ごとに一度ずつではなく、DriverContextを実装することもできます。
type Driver interface {
	Open(name string) (Conn, error)
}

// もしDriverがDriverContextを実装している場合、sql.DBはOpenConnectorを呼び出してConnectorを取得し、
// そのConnectorのConnectメソッドを呼び出して必要な接続を取得します。
// これにより、接続ごとにDriverのOpenメソッドを呼び出すのではなく、名前を1回だけ解析することができ、
// またper-Connコンテキストにアクセスすることもできます。
type DriverContext interface {
	OpenConnector(name string) (Connector, error)
}

// コネクタは、固定の構成でドライバを表し、複数のゴルーチンで使用するための同等の接続を作成できます。
//
// コネクタはsql.OpenDBに渡すことができ、ドライバは独自のsql.DBコンストラクタを実装するため、また、DriverContextのOpenConnectorメソッドによって返されることができます。これにより、ドライバはコンテキストへのアクセスとドライバ構成の繰り返し解析を避けることができます。
//
// コネクタがio.Closerを実装している場合、sqlパッケージのDB.CloseメソッドはCloseを呼び出し、エラー（あれば）を返します。
type Connector interface {
	Connect(context.Context) (Conn, error)

	Driver() Driver
}

// ErrSkipは、一部のオプションのインタフェースメソッドによって、高速経路が利用できないことを実行時に示すために返される場合があります。sqlパッケージは、オプションのインタフェースが実装されていないかのように続行する必要があります。ErrSkipは、明示的に文書化されている場所でのみサポートされます。
var ErrSkip = errors.New("driver: skip fast-path; continue as if unimplemented")

// ErrBadConnは、ドライバがドライバ.Connが不良な状態であることを示すために、
// sqlパッケージに返すべきです（たとえば、サーバーが接続を早期に閉じたなど）。
// また、sqlパッケージは新しい接続で再試行する必要があります。
//
// 重複した操作を防ぐために、可能性がある場合には、ErrBadConnを返してはいけません。
// データベースサーバーが操作を実行した可能性があっても、ErrBadConnは返してはいけません。
//
// エラーはerrors.Isを使用してチェックされます。エラーは
// ErrBadConnをラップするか、Is(error) boolメソッドを実装することがあります。
var ErrBadConn = errors.New("driver: bad connection")

// PingerはConnによって実装される可能性のあるオプションのインターフェースです。
//
// ConnがPingerを実装していない場合、sqlパッケージのDB.PingおよびDB.PingContextは少なくとも1つのConnが利用可能かどうかを確認します。
//
// Conn.PingがErrBadConnを返す場合、DB.PingおよびDB.PingContextはConnをプールから削除します。
type Pinger interface {
	Ping(ctx context.Context) error
}

// ExecerはConnによって実装されるかもしれないオプションのインターフェースです。
//
// もしConnがExecerContextまたはExecerのどちらの実装も持っていない場合、
// sqlパッケージのDB.Execはまずクエリを準備し、ステートメントを実行し、そしてステートメントを閉じます。
//
// ExecはErrSkipを返す場合があります。
//
// 廃止予定: ドライバは代わりにExecerContextを実装するべきです。
type Execer interface {
	Exec(query string, args []Value) (Result, error)
}

// ExecerContextはConnによって実装されるかもしれないオプションのインターフェースです。
//
// ConnがExecerContextを実装していない場合、sqlパッケージのDB.ExecはExecerにフォールバックします。
// もしConnがExecerも実装していない場合、DB.Execはまずクエリを準備し、ステートメントを実行してからステートメントを閉じます。
//
// ExecContextはErrSkipを返すことがあります。
//
// ExecContextはコンテキストのタイムアウトを尊重し、コンテキストがキャンセルされたら返ります。
type ExecerContext interface {
	ExecContext(ctx context.Context, query string, args []NamedValue) (Result, error)
}

// QueryerはConnによって実装されるかもしれないオプションのインターフェースです。
//
// ConnがQueryerContextでもQueryerでも実装していない場合、sqlパッケージのDB.Queryはまずクエリを準備し、ステートメントを実行してからステートメントを閉じます。
//
// QueryはErrSkipを返すことがあります。
//
// 廃止予定です: ドライバは代わりにQueryerContextを実装するべきです。
type Queryer interface {
	Query(query string, args []Value) (Rows, error)
}

// QueryerContextは、Connによって実装されるかもしれないオプションのインターフェースです。
//
// ConnがQueryerContextを実装していない場合、sqlパッケージのDB.QueryはQueryerにフォールバックします。
// もし、ConnがQueryerを実装していない場合、DB.Queryはまずクエリを準備し、ステートメントを実行してからステートメントを閉じます。
//
// QueryContextはErrSkipを返す場合があります。
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

// ConnPrepareContextはコンテキストを使用してConnインターフェースを拡張します。
type ConnPrepareContext interface {
	PrepareContext(ctx context.Context, query string) (Stmt, error)
}

// IsolationLevelはTxOptionsに保存されるトランザクション分離レベルです。
//
// この型は、sql.IsolationLevelと一緒に定義された値と同じものと考えられるべきです。
type IsolationLevel int

// TxOptionsはトランザクションのオプションを保持します。
//
// この型はsql.TxOptionsと同一と見なされるべきです。
type TxOptions struct {
	Isolation IsolationLevel
	ReadOnly  bool
}

// ConnBeginTxは、コンテキストとTxOptionsを追加してConnインタフェースを拡張します。
type ConnBeginTx interface {
	BeginTx(ctx context.Context, opts TxOptions) (Tx, error)
}

// SessionResetterは、Connによって実装される可能性があります。これにより、ドライバは接続に関連付けられたセッション状態をリセットし、悪い接続を通知することができます。
type SessionResetter interface {
	ResetSession(ctx context.Context) error
}

// Validaterは、Connによって実装されることがあります。これにより、ドライバーは接続が有効であるか、破棄すべきかを示すことができます。
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

// Stmtはプリペアドステートメントです。これはConnにバインドされており、複数のゴルーチンで同時に使用されません。
type Stmt interface {
	Close() error

	NumInput() int

	Exec(args []Value) (Result, error)

	Query(args []Value) (Rows, error)
}

// StmtExecContextはコンテキストを提供することにより、Stmtインターフェースを拡張します。
type StmtExecContext interface {
	ExecContext(ctx context.Context, args []NamedValue) (Result, error)
}

// StmtQueryContextはコンテキストを持つQueryを提供することにより、Stmtインターフェースを強化します。
type StmtQueryContext interface {
	QueryContext(ctx context.Context, args []NamedValue) (Rows, error)
}

// ErrRemoveArgumentは、NamedValueCheckerから返されることがあります。
// これは、sqlパッケージに対して引数をドライバのクエリインターフェースに渡さないよう指示するためです。
// クエリ固有のオプションやSQLクエリ引数ではない構造体を受け入れる場合に返します。
var ErrRemoveArgument = errors.New("driver: remove argument from query")

// NamedValueCheckerはConnまたはStmtによってオプションで実装されることがあります。これにより、ドライバはデフォルトのValuesタイプを超えたGoおよびデータベースのタイプを処理するための制御を提供します。
// sqlパッケージは、値チェッカーを以下の順序でチェックし、最初に一致したもので停止します： Stmt.NamedValueChecker、Conn.NamedValueChecker、Stmt.ColumnConverter、DefaultParameterConverter。
// CheckNamedValueがErrRemoveArgumentを返す場合、NamedValueは最終的なクエリ引数に含まれません。これはクエリ自体に特殊なオプションを渡すために使用される場合があります。
// ErrSkipが返された場合、列コンバーターのエラーチェックパスが引数に使用されます。ドライバは、独自の特殊なケースを使い果たした後にErrSkipを返すことを望むかもしれません。
type NamedValueChecker interface {
	CheckNamedValue(*NamedValue) error
}

// ColumnConverterは、ステートメントが自身の列の型を認識しており、任意の型からドライバーのValueに変換できる場合、Stmtによってオプションで実装されることがあります。
// 廃止予定 : ドライバーはNamedValueCheckerを実装する必要があります。
type ColumnConverter interface {
	ColumnConverter(idx int) ValueConverter
}

// Rowsは実行されたクエリの結果に対するイテレータです。
type Rows interface {
	Columns() []string

	Close() error

	Next(dest []Value) error
}

// RowsNextResultSetは、ドライバに次の結果セットに進むようにシグナルを送る方法を提供するためにRowsインターフェースを拡張しています。
type RowsNextResultSet interface {
	Rows

	HasNextResultSet() bool

	NextResultSet() error
}

// RowsColumnTypeScanTypeは、Rowsによって実装されるかもしれません。スキャンに使用できる値の型を返す必要があります。例えば、データベースのカラムタイプが「bigint」の場合、これは「reflect.TypeOf(int64(0))」を返すべきです。
type RowsColumnTypeScanType interface {
	Rows
	ColumnTypeScanType(index int) reflect.Type
}

// RowsColumnTypeDatabaseTypeNameはRowsによって実装されるかもしれません。長さを除いたデータベースシステムのタイプ名を返す必要があります。タイプ名は大文字であるべきです。
// 返される型の例: "VARCHAR", "NVARCHAR", "VARCHAR2", "CHAR", "TEXT",
// "DECIMAL", "SMALLINT", "INT", "BIGINT", "BOOL", "[]BIGINT", "JSONB", "XML",
// "TIMESTAMP"。
type RowsColumnTypeDatabaseTypeName interface {
	Rows
	ColumnTypeDatabaseTypeName(index int) string
}

// RowsColumnTypeLengthは、Rowsによって実装されるかもしれません。カラムが可変長の場合、カラムタイプの長さを返す必要があります。カラムが可変長のタイプでない場合、okはfalseを返す必要があります。システムの制限以外で長さが制限されていない場合、math.MaxInt64を返す必要があります。以下は、さまざまなタイプの戻り値の例です：
// TEXT（math.MaxInt64、true）
// varchar(10)（10、true）
// nvarchar(10)（10、true）
// decimal（0、false）
// int（0、false）
// bytea(30)（30、true）
type RowsColumnTypeLength interface {
	Rows
	ColumnTypeLength(index int) (length int64, ok bool)
}

// RowsColumnTypeNullableは、Rowsによって実装される可能性があります。カラムがnullである可能性がある場合は、nullableの値をtrueにする必要があります。カラムがnullでないことが確認されている場合は、falseにする必要があります。
// カラムのヌラビリティが不明な場合は、okをfalseにしてください。
type RowsColumnTypeNullable interface {
	Rows
	ColumnTypeNullable(index int) (nullable, ok bool)
}

// RowsColumnTypePrecisionScaleはRowsによって実装されるかもしれません。それはデシマル型の精度とスケールを返すべきです。該当しない場合、okはfalseであるべきです。
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

// RowsAffectedは、行数を変更するINSERTまたはUPDATE操作に対する Result を実装します。
type RowsAffected int64

var _ Result = RowsAffected(0)

func (RowsAffected) LastInsertId() (int64, error)

func (v RowsAffected) RowsAffected() (int64, error)

// ResultNoRowsは、DDLコマンド（CREATE TABLEなど）が成功した場合にドライバーが返すための事前定義された結果です。これは、LastInsertIdとRowsAffectedの両方に対してエラーを返します。
var ResultNoRows noRows

var _ Result = noRows{}
