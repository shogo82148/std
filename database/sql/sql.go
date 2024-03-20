// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// sqlパッケージは、SQL（またはSQLライク）データベースを取り巻く汎用インターフェースを提供します。
//
// sqlパッケージは、データベースドライバと共に使用する必要があります。
// ドライバのリストについては、https://golang.org/s/sqldrivers を参照してください。
//
// コンテキストのキャンセルをサポートしていないドライバは、クエリが完了するまで戻らないことに注意してください。
//
// 使用例については、https://golang.org/s/sqlwiki のウィキページを参照してください。
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

// Registerは、指定された名前でデータベースドライバを利用可能にします。
// Registerが同じ名前で2回呼び出された場合、またはdriverがnilの場合、
// panicが発生します。
func Register(name string, driver driver.Driver)

// Driversは、登録されたドライバーの名前のソートされたリストを返します。
func Drivers() []string

// NamedArgは名前付き引数です。NamedArg値は、 [DB.Query] または [DB.Exec] の引数として使用でき、
// SQLステートメントの対応する名前付きパラメータにバインドされます。
//
// NamedArg値をより簡潔に作成する方法については、 [Named] 関数を参照してください。
type NamedArg struct {
	_NamedFieldsRequired struct{}

	// Nameはパラメータプレースホルダーの名前です。
	//
	// 空の場合、引数リストの序数が使用されます。
	//
	// Nameには、シンボル接頭辞を省略する必要があります。
	Name string

	// Valueはパラメータの値です。
	// クエリ引数と同じ値型が割り当てられる可能性があります。
	Value any
}

// Namedは、 [NamedArg] 値をより簡潔に作成する方法を提供します。
//
// 使用例:
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

// IsolationLevelは、 [TxOptions] で使用されるトランザクション分離レベルです。
type IsolationLevel int

// [DB.BeginTx] でドライバーがサポートする可能性のあるさまざまな分離レベル。
// ドライバーが特定の分離レベルをサポートしていない場合、エラーが返される場合があります。
//
// https://en.wikipedia.org/wiki/Isolation_(database_systems)#Isolation_levels を参照してください。
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

// Stringはトランザクション分離レベルの名前を返します。
func (i IsolationLevel) String() string

var _ fmt.Stringer = LevelDefault

// TxOptionsは、 [DB.BeginTx] で使用されるトランザクションオプションを保持します。
type TxOptions struct {
	// Isolationはトランザクション分離レベルです。
	// ゼロの場合、ドライバーまたはデータベースのデフォルトレベルが使用されます。
	Isolation IsolationLevel
	ReadOnly  bool
}

// RawBytesは、データベース自体が所有するメモリへの参照を保持するバイトスライスです。
// RawBytesに対して [Rows.Scan] を実行した後、スライスは次の [Rows.Next]、[Rows.Scan]、または [Rows.Close] の呼び出しまでのみ有効です。
type RawBytes []byte

// NullStringは、nullである可能性がある文字列を表します。
// NullStringは [Scanner] インターフェースを実装するため、
// スキャン先として使用できます。
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

// Scanは [Scanner] インターフェースを実装します。
func (ns *NullString) Scan(value any) error

// Valueは、 [driver.Valuer] インターフェースを実装します。
func (ns NullString) Value() (driver.Value, error)

// NullInt64は、nullである可能性があるint64を表します。
// NullInt64はScannerインターフェースを実装するため、
// [NullString] と同様にスキャン先として使用できます。
type NullInt64 struct {
	Int64 int64
	Valid bool
}

// Scanは [Scanner] インターフェースを実装します。
func (n *NullInt64) Scan(value any) error

// Valueは、 [driver.Valuer] インターフェースを実装します。
func (n NullInt64) Value() (driver.Value, error)

// NullInt32は、nullである可能性があるint32を表します。
// NullInt32は [Scanner] インターフェースを実装するため、
// NullStringと同様にスキャン先として使用できます。
type NullInt32 struct {
	Int32 int32
	Valid bool
}

// Scanは [Scanner] インターフェースを実装します。
func (n *NullInt32) Scan(value any) error

// Valueは、 [driver.Valuer] インターフェースを実装します。
func (n NullInt32) Value() (driver.Value, error)

// NullInt16は、nullである可能性があるint16を表します。
// NullInt16は [Scanner] インターフェースを実装するため、
// [NullString] と同様にスキャン先として使用できます。
type NullInt16 struct {
	Int16 int16
	Valid bool
}

// Scanは [Scanner] インターフェースを実装します。
func (n *NullInt16) Scan(value any) error

// Valueは、 [driver.Valuer] インターフェースを実装します。
func (n NullInt16) Value() (driver.Value, error)

// NullByteは、nullである可能性があるバイトを表します。
// NullByteは [Scanner] インターフェースを実装するため、
// [NullString] と同様にスキャン先として使用できます。
type NullByte struct {
	Byte  byte
	Valid bool
}

// Scanは [Scanner] インターフェースを実装します。
func (n *NullByte) Scan(value any) error

// Valueは、 [driver.Valuer] インターフェースを実装します。
func (n NullByte) Value() (driver.Value, error)

// NullFloat64は、nullである可能性があるfloat64を表します。
// NullFloat64は [Scanner] インターフェースを実装するため、
// [NullString] と同様にスキャン先として使用できます。
type NullFloat64 struct {
	Float64 float64
	Valid   bool
}

// Scanは [Scanner] インターフェースを実装します。
func (n *NullFloat64) Scan(value any) error

// Valueは、 [driver.Valuer] インターフェースを実装します。
func (n NullFloat64) Value() (driver.Value, error)

// NullBoolは、nullである可能性があるboolを表します。
// NullBoolは [Scanner] インターフェースを実装するため、
// [NullString] と同様にスキャン先として使用できます。
type NullBool struct {
	Bool  bool
	Valid bool
}

// Scanは [Scanner] インターフェースを実装します。
func (n *NullBool) Scan(value any) error

// Valueは、 [driver.Valuer] インターフェースを実装します。
func (n NullBool) Value() (driver.Value, error)

// NullTimeは、nullである可能性がある [time.Time] を表します。
// NullTimeは [Scanner] インターフェースを実装するため、
// [NullString] と同様にスキャン先として使用できます。
type NullTime struct {
	Time  time.Time
	Valid bool
}

// Scanは [Scanner] インターフェースを実装します。
func (n *NullTime) Scan(value any) error

// Valueは、 [driver.Valuer] インターフェースを実装します。
func (n NullTime) Value() (driver.Value, error)

// Nullは、nullである可能性がある値を表します。
// Nullは [Scanner] インターフェースを実装するため、
// スキャン先として使用できます。
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

// Outは、ストアドプロシージャからOUTPUT値パラメータを取得するために使用できます。
//
// すべてのドライバーとデータベースがOUTPUT値パラメータをサポートしているわけではありません。
//
// 使用例:
//
//	var outArg string
//	_, err := db.ExecContext(ctx, "ProcName", sql.Named("Arg1", sql.Out{Dest: &outArg}))
type Out struct {
	_NamedFieldsRequired struct{}

	// Destは、ストアドプロシージャのOUTPUTパラメータの結果に設定される値へのポインタです。
	Dest any

	// Inは、パラメータがINOUTパラメータであるかどうかを示します。
	// その場合、ストアドプロシージャへの入力値はDestのポインタの参照解除された値であり、
	// その後、出力値で置き換えられます。
	In bool
}

// ErrNoRowsは、 [DB.QueryRow] が行を返さない場合に [Row.Scan] によって返されます。
// そのような場合、QueryRowはプレースホルダー [*Row] 値を返し、
// このエラーはScanまで遅延されます。
var ErrNoRows = errors.New("sql: no rows in result set")

// DBは、ゼロ個以上の基礎接続を表すデータベースハンドルです。
// 複数のゴルーチンによる同時使用に対して安全です。
//
// sqlパッケージは、接続を自動的に作成および解放します。
// また、アイドル接続のフリープールを維持します。
// データベースが接続ごとの状態を持つ場合、そのような状態はトランザクション（[Tx]）または接続（[Conn]）内で信頼性が高く観察できます。
// [DB.Begin] が呼び出されると、返された [Tx] は単一の接続にバインドされます。
// [Tx.Commit] または [Tx.Rollback] が呼び出されると、そのトランザクションの接続が [DB] のアイドル接続プールに返されます。
// プールのサイズは [DB.SetMaxIdleConns] で制御できます。
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

// OpenDBは、コネクタを使用してデータベースを開き、ドライバーが文字列ベースのデータソース名をバイパスできるようにします。
//
// ほとんどのユーザーは、 [*DB] を返すドライバー固有の接続ヘルパー関数を介してデータベースを開きます。
// Go標準ライブラリにはデータベースドライバーは含まれていません。サードパーティのドライバーのリストについては、https://golang.org/s/sqldrivers を参照してください。
//
// OpenDBは、データベースへの接続を作成せずに引数を検証する場合があります。
// データソース名が有効であることを確認するには、[DB.Ping] を呼び出します。
//
// 返された [DB] は、複数のゴルーチンによる同時使用に対して安全であり、アイドル接続のプールを維持します。
// したがって、OpenDB関数は1回だけ呼び出す必要があります。 [DB] を閉じる必要はほとんどありません。
func OpenDB(c driver.Connector) *DB

// Openは、データベースドライバー名とドライバー固有のデータソース名で指定されたデータベースを開きます。
// 通常、少なくともデータベース名と接続情報が含まれます。
//
// ほとんどのユーザーは、 [*DB] を返すドライバー固有の接続ヘルパー関数を介してデータベースを開きます。
// Go標準ライブラリにはデータベースドライバーは含まれていません。サードパーティのドライバーのリストについては、https://golang.org/s/sqldrivers を参照してください。
//
// Openは、データベースへの接続を作成せずに引数を検証する場合があります。
// データソース名が有効であることを確認するには、[DB.Ping] を呼び出します。
//
// 返された [DB] は、複数のゴルーチンによる同時使用に対して安全であり、アイドル接続のプールを維持します。
// したがって、Open関数は1回だけ呼び出す必要があります。 [DB] を閉じる必要はほとんどありません。
func Open(driverName, dataSourceName string) (*DB, error)

// PingContextは、データベースへの接続がまだ有効であることを確認し、必要に応じて接続を確立します。
func (db *DB) PingContext(ctx context.Context) error

// Pingは、データベースへの接続がまだ有効であることを確認し、必要に応じて接続を確立します。
//
// Pingは、内部的に [context.Background] を使用します。コンテキストを指定するには、 [DB.PingContext] を使用してください。
func (db *DB) Ping() error

// Closeはデータベースを閉じ、新しいクエリの開始を防止します。
// Closeは、サーバーで処理を開始したすべてのクエリが完了するのを待ってから終了します。
//
// [DB] ハンドルは長期間生存し、多くのゴルーチンで共有されることを意図しているため、 [DB] をCloseすることはまれです。
func (db *DB) Close() error

// SetMaxIdleConnsは、アイドル接続プール内の最大接続数を設定します。
//
// MaxOpenConnsが0より大きく、新しいMaxIdleConnsより小さい場合、
// 新しいMaxIdleConnsはMaxOpenConnsの制限に合わせて減少します。
//
// n <= 0の場合、アイドル接続は保持されません。
//
// デフォルトの最大アイドル接続数は現在2です。将来のリリースで変更される可能性があります。
func (db *DB) SetMaxIdleConns(n int)

// SetMaxOpenConnsは、データベースへの最大オープン接続数を設定します。
//
// MaxIdleConnsが0より大きく、新しいMaxOpenConnsより小さい場合、
// 新しいMaxIdleConnsはMaxOpenConnsの制限に合わせて減少します。
//
// n <= 0の場合、オープン接続数に制限はありません。
// デフォルトは0（無制限）です。
func (db *DB) SetMaxOpenConns(n int)

// SetConnMaxLifetimeは、接続が再利用される最大時間を設定します。
//
// 期限切れの接続は再利用前に遅延して閉じることができます。
//
// d <= 0の場合、接続の年齢により接続が閉じられることはありません。
func (db *DB) SetConnMaxLifetime(d time.Duration)

// SetConnMaxIdleTimeは、接続がアイドル状態になっている最大時間を設定します。
//
// 期限切れの接続は再利用前に遅延して閉じることができます。
//
// d <= 0の場合、接続のアイドル時間により接続が閉じられることはありません。
func (db *DB) SetConnMaxIdleTime(d time.Duration)

// DBStatsには、データベースの統計情報が含まれます。
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

// Statsは、データベースの統計情報を返します。
func (db *DB) Stats() DBStats

// PrepareContextは、後でのクエリまたは実行のためにプリペアドステートメントを作成します。
// 返されたステートメントから複数のクエリまたは実行を同時に実行できます。
// ステートメントが不要になったら、呼び出し元はステートメントの [*Stmt.Close] メソッドを呼び出す必要があります。
//
// 提供されたコンテキストは、ステートメントの実行ではなく、ステートメントの準備に使用されます。
func (db *DB) PrepareContext(ctx context.Context, query string) (*Stmt, error)

// Prepareは、後でのクエリまたは実行のためにプリペアドステートメントを作成します。
// 返されたステートメントから複数のクエリまたは実行を同時に実行できます。
// ステートメントが不要になったら、呼び出し元はステートメントの [*Stmt.Close] メソッドを呼び出す必要があります。
//
// Prepareは、内部的に [context.Background] を使用します。コンテキストを指定するには、 [DB.PrepareContext] を使用してください。
func (db *DB) Prepare(query string) (*Stmt, error)

// ExecContextは、行を返さないクエリを実行します。
// argsは、クエリ内のプレースホルダーパラメーター用です。
func (db *DB) ExecContext(ctx context.Context, query string, args ...any) (Result, error)

// Execは、行を返さないクエリを実行します。
// argsは、クエリ内のプレースホルダーパラメーター用です。
//
// Execは、内部的に [context.Background] を使用します。コンテキストを指定するには、 [DB.ExecContext] を使用してください。
func (db *DB) Exec(query string, args ...any) (Result, error)

// QueryContextは、通常はSELECTで返される行を返すクエリを実行します。
// argsは、クエリ内のプレースホルダーパラメーター用です。
func (db *DB) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)

// QueryContextは、通常はSELECTで返される行を返すクエリを実行します。
// argsは、クエリ内のプレースホルダーパラメーター用です。
//
// QueryContextは、内部的に [context.Background] を使用します。コンテキストを指定するには、 [DB.QueryContext] を使用してください。
func (db *DB) Query(query string, args ...any) (*Rows, error)

// QueryRowContextは、最大1行を返すと予想されるクエリを実行します。
// QueryRowContextは常にnil以外の値を返します。エラーは [Row] のScanメソッドが呼び出されるまで遅延されます。
// クエリが行を選択しない場合、[*Row.Scan] は [ErrNoRows] を返します。
// そうでない場合、[*Row.Scan] は最初に選択された行をスキャンし、残りを破棄します。
func (db *DB) QueryRowContext(ctx context.Context, query string, args ...any) *Row

// QueryRowは、最大1行を返すと予想されるクエリを実行します。
// QueryRowは常にnil以外の値を返します。エラーは [Row] のScanメソッドが呼び出されるまで遅延されます。
// クエリが行を選択しない場合、[*Row.Scan] は [ErrNoRows] を返します。
// そうでない場合、[*Row.Scan] は最初に選択された行をスキャンし、残りを破棄します。
//
// QueryRowは、内部的に [context.Background] を使用します。コンテキストを指定するには、 [DBQueryRowContext] を使用してください。
func (db *DB) QueryRow(query string, args ...any) *Row

// BeginTxはトランザクションを開始します。
//
// 提供されたコンテキストは、トランザクションがコミットまたはロールバックされるまで使用されます。
// コンテキストがキャンセルされると、sqlパッケージはトランザクションをロールバックします。
// BeginTxに提供されたコンテキストがキャンセルされた場合、 [Tx.Commit] はエラーを返します。
//
// 提供された [TxOptions] はオプションであり、デフォルトを使用する場合はnilにすることができます。
// ドライバーがサポートしていない非デフォルトの分離レベルが使用された場合、エラーが返されます。
func (db *DB) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)

// Beginはトランザクションを開始します。デフォルトの分離レベルはドライバーに依存します。
//
// Beginは、内部的に [context.Background] を使用します。コンテキストを指定するには、 [DB.BeginTx] を使用してください。
func (db *DB) Begin() (*Tx, error)

// Driverは、データベースの基礎となるドライバーを返します。
func (db *DB) Driver() driver.Driver

// ErrConnDoneは、接続プールに返された接続で実行された操作によって返されます。
var ErrConnDone = errors.New("sql: connection is already closed")

// Connは、新しい接続を開くか、接続プールから既存の接続を返して、単一の接続を返します。
// Connは、接続が返されるか、ctxがキャンセルされるまでブロックされます。
// 同じConnで実行されるクエリは、同じデータベースセッションで実行されます。
//
// 各Connは、 [Conn.Close] を呼び出して使用後にデータベースプールに返す必要があります。
func (db *DB) Conn(ctx context.Context) (*Conn, error)

// Connは、データベース接続プールではなく、単一のデータベース接続を表します。
// 特定の理由がない限り、クエリは [DB] から実行することをお勧めします。
//
// Connは、 [Conn.Close] を呼び出して接続をデータベースプールに返す必要があります。
// また、実行中のクエリと同時に呼び出すことができます。
//
// [Conn.Close] の呼び出し後、接続に対するすべての操作は [ErrConnDone] で失敗します。
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

// PingContextは、データベースへの接続がまだ有効であることを確認します。
func (c *Conn) PingContext(ctx context.Context) error

// ExecContextは、行を返さないクエリを実行します。
// argsは、クエリ内のプレースホルダーパラメーター用です。
func (c *Conn) ExecContext(ctx context.Context, query string, args ...any) (Result, error)

// QueryContextは、通常はSELECTで返される行を返すクエリを実行します。
// argsは、クエリ内のプレースホルダーパラメーター用です。
func (c *Conn) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)

// QueryRowContextは、最大1行を返すと予想されるクエリを実行します。
// QueryRowContextは常にnil以外の値を返します。エラーは [*Row.Scan] メソッドが呼び出されるまで遅延されます。
// クエリが行を選択しない場合、[*Row.Scan] は [ErrNoRows] を返します。
// そうでない場合、[*Row.Scan] は最初に選択された行をスキャンし、残りを破棄します。
func (c *Conn) QueryRowContext(ctx context.Context, query string, args ...any) *Row

// PrepareContextは、後でのクエリまたは実行のためにプリペアドステートメントを作成します。
// 返されたステートメントから複数のクエリまたは実行を同時に実行できます。
// ステートメントが不要になったら、呼び出し元はステートメントの [*Stmt.Close] メソッドを呼び出す必要があります。
//
// 提供されたコンテキストは、ステートメントの実行ではなく、ステートメントの準備に使用されます。
func (c *Conn) PrepareContext(ctx context.Context, query string) (*Stmt, error)

// Rawは、fを実行し、fの実行中に基礎となるドライバー接続を公開します。
// driverConnは、fの外部で使用してはいけません。
//
// fが返り、errが [driver.ErrBadConn] でない場合、 [Conn] は [Conn.Close] が呼び出されるまで使用可能です。
func (c *Conn) Raw(f func(driverConn any) error) (err error)

// BeginTxはトランザクションを開始します。
//
// 提供されたコンテキストは、トランザクションがコミットまたはロールバックされるまで使用されます。
// コンテキストがキャンセルされると、sqlパッケージはトランザクションをロールバックします。
// BeginTxに提供されたコンテキストがキャンセルされた場合、 [Tx.Commit] はエラーを返します。
//
// 提供された [TxOptions] はオプションであり、デフォルトを使用する場合はnilにすることができます。
// ドライバーがサポートしていない非デフォルトの分離レベルが使用された場合、エラーが返されます。
func (c *Conn) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)

// Closeは、接続を接続プールに返します。
// Closeの後に行われたすべての操作は、 [ErrConnDone] で返されます。
// Closeは、他の操作と同時に安全に呼び出すことができ、
// すべての他の操作が完了するまでブロックされます。
// 使用されたコンテキストを最初にキャンセルしてから直接Closeを呼び出すことが役立つ場合があります。
func (c *Conn) Close() error

// Txは、進行中のデータベーストランザクションです。
//
// トランザクションは、 [Tx.Commit] または [Tx.Rollback] の呼び出しで終了する必要があります。
//
// [Tx.Commit] または [Tx.Rollback] の呼び出し後、トランザクション上のすべての操作は [ErrTxDone] で失敗します。
//
// トランザクションの [Tx.Prepare] または [Tx.Stmt] メソッドを呼び出して準備されたステートメントは、 [Tx.Commit] または [Tx.Rollback] の呼び出しで閉じられます。
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

// ErrTxDoneは、すでにコミットまたはロールバックされたトランザクションに対して実行された操作によって返されます。
var ErrTxDone = errors.New("sql: transaction has already been committed or rolled back")

// Commitはトランザクションをコミットします。
func (tx *Tx) Commit() error

// Rollbackはトランザクションを中止します。
func (tx *Tx) Rollback() error

// PrepareContextは、トランザクション内で使用するためのプリペアドステートメントを作成します。
//
// 返されたステートメントはトランザクション内で動作し、トランザクションがコミットまたはロールバックされたときに閉じられます。
//
// このトランザクションで既存のプリペアドステートメントを使用するには、 [Tx.Stmt] を参照してください。
//
// 提供されたコンテキストは、ステートメントの実行ではなく、ステートメントの準備に使用されます。
// 返されたステートメントはトランザクションコンテキストで実行されます。
func (tx *Tx) PrepareContext(ctx context.Context, query string) (*Stmt, error)

// Prepareは、トランザクション内で使用するためのプリペアドステートメントを作成します。
//
// 返されたステートメントはトランザクション内で動作し、トランザクションがコミットまたはロールバックされたときに閉じられます。
//
// このトランザクションで既存のプリペアドステートメントを使用するには、 [Tx.Stmt] を参照してください。
//
// Prepareは、内部的に [context.Background] を使用します。コンテキストを指定するには、 [Tx.PrepareContext] を使用してください。
func (tx *Tx) Prepare(query string) (*Stmt, error)

// StmtContextは、既存のステートメントからトランザクション固有のプリペアドステートメントを返します。
//
// 例：
//
//	updateMoney, err := db.Prepare("UPDATE balance SET money=money+? WHERE id=?")
//	...
//	tx, err := db.Begin()
//	...
//	res, err := tx.StmtContext(ctx, updateMoney).Exec(123.45, 98293203)
//
// 提供されたコンテキストはステートメントの実行ではなく、ステートメントの準備に使用されます。
//
// 返されたステートメントはトランザクション内で動作し、トランザクションがコミットまたはロールバックされたときに閉じられます。
func (tx *Tx) StmtContext(ctx context.Context, stmt *Stmt) *Stmt

// Stmtは、既存のステートメントからトランザクション固有のプリペアドステートメントを返します。
//
// 例：
//
//	updateMoney, err := db.Prepare("UPDATE balance SET money=money+? WHERE id=?")
//	...
//	tx, err := db.Begin()
//	...
//	res, err := tx.Stmt(updateMoney).Exec(123.45, 98293203)
//
// 返されたステートメントはトランザクション内で動作し、トランザクションがコミットまたはロールバックされたときに閉じられます。
//
// Stmtは、内部的に [context.Background] を使用します。コンテキストを指定するには、 [Tx.StmtContext] を使用してください。
func (tx *Tx) Stmt(stmt *Stmt) *Stmt

// ExecContextは、行を返さないクエリを実行します。
// 例えば、INSERTやUPDATEです。
func (tx *Tx) ExecContext(ctx context.Context, query string, args ...any) (Result, error)

// Execは、行を返さないクエリを実行します。
// 例えば、INSERTやUPDATEです。
//
// Execは、内部的に [context.Background] を使用します。コンテキストを指定するには、 [Tx.ExecContext] を使用してください。
func (tx *Tx) Exec(query string, args ...any) (Result, error)

// QueryContext executes a query that returns rows, typically a SELECT.
func (tx *Tx) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)

// Queryは、通常はSELECTで返される行を返すクエリを実行します。
//
// Queryは、内部的に [context.Background] を使用します。コンテキストを指定するには、 [Tx.QueryContext] を使用してください。
func (tx *Tx) Query(query string, args ...any) (*Rows, error)

// QueryRowContextは、最大1行を返すと予想されるクエリを実行します。
// QueryRowContextは常にnil以外の値を返します。エラーは [Row] のScanメソッドが呼び出されるまで遅延されます。
// クエリが行を選択しない場合、[*Row.Scan] は [ErrNoRows] を返します。
// そうでない場合、[*Row.Scan] は最初に選択された行をスキャンし、残りを破棄します。
func (tx *Tx) QueryRowContext(ctx context.Context, query string, args ...any) *Row

// QueryRowは、最大1行を返すと予想されるクエリを実行します。
// QueryRowは常にnil以外の値を返します。エラーは [Row] のScanメソッドが呼び出されるまで遅延されます。
// クエリが行を選択しない場合、[*Row.Scan] は [ErrNoRows] を返します。
// そうでない場合、[*Row.Scan] は最初に選択された行をスキャンし、残りを破棄します。
//
// QueryRowは、内部的に [context.Background] を使用します。コンテキストを指定するには、 [Tx.QueryRowContext] を使用してください。
func (tx *Tx) QueryRow(query string, args ...any) *Row

var (
	_ stmtConnGrabber = &Tx{}
	_ stmtConnGrabber = &Conn{}
)

// Stmtは、プリペアドステートメントです。
// Stmtは、複数のゴルーチンによる同時使用に対して安全です。
//
// Stmtが [Tx] または [Conn] で準備された場合、それは1つの基礎となる接続に永久にバインドされます。
// [Tx] または [Conn] が閉じられると、Stmtは使用できなくなり、すべての操作がエラーを返します。
// Stmtが [DB] で準備された場合、それは [DB] の寿命の間使用可能です。
// Stmtが新しい基礎となる接続で実行する必要がある場合、自動的に新しい接続で自己準備します。
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

// ExecContextは、指定された引数を使用してプリペアドステートメントを実行し、
// ステートメントの影響を要約する [Result] を返します。
func (s *Stmt) ExecContext(ctx context.Context, args ...any) (Result, error)

// Execは、指定された引数を使用してプリペアドステートメントを実行し、
// ステートメントの影響を要約する [Result] を返します。
//
// Execは、内部的に [context.Background] を使用します。コンテキストを指定するには、 [Stmt.ExecContext] を使用してください。
func (s *Stmt) Exec(args ...any) (Result, error)

// QueryContextは、指定された引数を使用してプリペアドクエリステートメントを実行し、
// クエリ結果を [*Rows] として返します。
func (s *Stmt) QueryContext(ctx context.Context, args ...any) (*Rows, error)

// Queryは、指定された引数を使用してプリペアドクエリステートメントを実行し、
// クエリ結果を*Rowsとして返します。
//
// Queryは、内部的に [context.Background] を使用します。コンテキストを指定するには、 [Stmt.QueryContext] を使用してください。
func (s *Stmt) Query(args ...any) (*Rows, error)

// QueryRowContextは、指定された引数を使用してプリペアドクエリステートメントを実行し、
// ステートメントの実行中にエラーが発生した場合、そのエラーは常にnil以外の [*Row] のScan呼び出しによって返されます。
// クエリが行を選択しない場合、[*Row.Scan] は [ErrNoRows] を返します。
// そうでない場合、[*Row.Scan] は最初に選択された行をスキャンし、残りを破棄します。
func (s *Stmt) QueryRowContext(ctx context.Context, args ...any) *Row

// QueryRowは、指定された引数を使用してプリペアドクエリステートメントを実行し、
// ステートメントの実行中にエラーが発生した場合、そのエラーは常にnil以外の [*Row] のScan呼び出しによって返されます。
// クエリが行を選択しない場合、[*Row.Scan] は [ErrNoRows] を返します。
// そうでない場合、[*Row.Scan] のScanは最初に選択された行をスキャンし、残りを破棄します。
//
// 使用例:
//
//	var name string
//	err := nameByUseridStmt.QueryRow(id).Scan(&name)
//
// QueryRowは、内部的に [context.Background] を使用します。コンテキストを指定するには、 [Stmt.QueryRowContext] を使用してください。
func (s *Stmt) QueryRow(args ...any) *Row

// Closeはステートメントを閉じます。
func (s *Stmt) Close() error

// Rowsはクエリの結果です。そのカーソルは、結果セットの最初の行の前に開始します。
// 行から行に進むには、 [Rows.Next] を使用してください。
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

// Nextは、 [Rows.Scan] メソッドで読み取る次の結果行を準備します。
// 成功した場合はtrue、次の結果行がない場合や準備中にエラーが発生した場合はfalseを返します。
// 2つの場合を区別するには、 [Rows.Err] を参照する必要があります。
//
// 最初の呼び出しを含め、すべての [Rows.Scan] 呼び出しは、 [Rows.Next] の呼び出しに先立っている必要があります。
func (rs *Rows) Next() bool

// NextResultSetは、次の結果セットの読み取りの準備をします。
// さらに結果セットがある場合はtrue、それ以外の場合はfalseを報告します。
// または、それに進む際にエラーが発生した場合はfalseを報告します。
// 2つの場合を区別するには、 [Rows.Err] を参照する必要があります。
//
// NextResultSetを呼び出した後、スキャンする前に常に [Rows.Next] メソッドを呼び出す必要があります。
// さらに結果セットがある場合、結果セットに行がない場合があります。
func (rs *Rows) NextResultSet() bool

// Errは、反復中に遭遇したエラー（ある場合）を返します。
// 明示的または暗黙的な [Rows.Close] の後にErrを呼び出すことができます。
func (rs *Rows) Err() error

// Columnsは列名を返します。
// Rowsが閉じられている場合、Columnsはエラーを返します。
func (rs *Rows) Columns() ([]string, error)

// ColumnTypesは、列の型、長さ、null可能性などの列情報を返します。
// 一部の情報は、一部のドライバから利用できない場合があります。
func (rs *Rows) ColumnTypes() ([]*ColumnType, error)

// ColumnTypeは、列の名前と型を含みます。
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

// Nameは、列の名前またはエイリアスを返します。
func (ci *ColumnType) Name() string

// Lengthは、テキストやバイナリフィールドタイプなどの可変長カラムタイプのためのカラムタイプ長を返します。
// タイプ長が無制限の場合、値は [math.MaxInt64] になります（データベースの制限は引き続き適用されます）。
// カラムタイプが可変長でない場合、例えばintの場合、またはドライバでサポートされていない場合、okはfalseになります。
func (ci *ColumnType) Length() (length int64, ok bool)

// DecimalSizeは、10進数型のスケールと精度を返します。
// 適用できない場合やサポートされていない場合は、okがfalseになります。
func (ci *ColumnType) DecimalSize() (precision, scale int64, ok bool)

// ScanTypeは、 [Rows.Scan] を使用してスキャンするために適したGo型を返します。
// ドライバがこのプロパティをサポートしていない場合、ScanTypeは空のインターフェースの型を返します。
func (ci *ColumnType) ScanType() reflect.Type

// Nullableは、列がnullである可能性があるかどうかを報告します。
// ドライバがこのプロパティをサポートしていない場合、okはfalseになります。
func (ci *ColumnType) Nullable() (nullable, ok bool)

// DatabaseTypeNameは、列のデータベースシステム名を返します。
// 空の文字列が返された場合、ドライバの型名はサポートされていません。
// ドライバのデータ型のリストについては、ドライバのドキュメントを参照してください。
// [ColumnType.Length] 指定子は含まれません。
// 一般的な型名には、"VARCHAR"、"TEXT"、"NVARCHAR"、"DECIMAL"、"BOOL"、"INT"、"BIGINT"があります。
func (ci *ColumnType) DatabaseTypeName() string

// Scanは、現在の行の列をdestが指す値にコピーします。
// destの数は [Rows] の列数と同じでなければなりません。
//
// Scanは、データベースから読み取った列を、以下の共通のGoの型およびsqlパッケージで提供される特殊な型に変換します。
//
//	*string
//	*[]byte
//	*int, *int8, *int16, *int32, *int64
//	*uint, *uint8, *uint16, *uint32, *uint64
//	*bool
//	*float32, *float64
//	*interface{}
//	*RawBytes
//	*Rows (カーソル値)
//	Scanner を実装する任意の型（Scannerドキュメントを参照）
//
// 最も単純な場合、ソース列の値の型が整数、ブール、または文字列型Tで、destが型*Tの場合、Scanは単にポインタを介して値を割り当てます。
//
// Scanは、文字列と数値型の間でも変換しますが、情報が失われない場合に限ります。Scanは、数値データベース列からスキャンされたすべての数値を*stringに文字列化しますが、数値型へのスキャンはオーバーフローのチェックが行われます。例えば、値が300のfloat64または値が"300"の文字列はuint16にスキャンできますが、uint8にはスキャンできません。ただし、float64(255)または"255"はuint8にスキャンできます。一部のfloat64数値を文字列に変換するスキャンは、文字列化すると情報が失われる場合があります。一般的には、浮動小数点列を*float64にスキャンします。
//
// dest引数の型が*[]byteの場合、Scanは対応するデータのコピーをその引数に保存します。コピーは呼び出し元が所有し、修正して無期限に保持できます。コピーを回避するには、代わりに [*RawBytes] の引数を使用します。[RawBytes] の使用制限については、 [RawBytes] のドキュメントを参照してください。
//
// 引数の型が*interface{}の場合、Scanは変換せずに基礎ドライバが提供する値をコピーします。[]byte型のソース値から*interface{}にスキャンする場合、スライスのコピーが作成され、呼び出し元が結果を所有します。
//
// [time.Time] 型のソース値は、*time.Time、*interface{}、*string、または*[]byte型の値にスキャンできます。後者2つに変換する場合、[time.RFC3339Nano] が使用されます。
//
// bool型のソース値は、*bool、*interface{}、*string、*[]byte、または [*RawBytes] 型にスキャンできます。
//
// *boolにスキャンする場合、ソースはtrue、false、1、0、または [strconv.ParseBool] で解析可能な文字列入力である必要があります。
//
// Scanは、クエリから返されたカーソル（例："select cursor(select * from my_table) from dual"）を、自体からスキャンできる [*Rows] 値に変換できます。親の [*Rows] が閉じられると、親の選択クエリはカーソル [*Rows] を閉じます。
//
// 最初の引数のいずれかがエラーを返す [Scanner] を実装している場合、そのエラーは返されたエラーにラップされます。
func (rs *Rows) Scan(dest ...any) error

// Closeは、 [Rows] を閉じ、以降の列挙を防止します。
// [Rows.Next] がfalseを返し、さらに結果セットがない場合、 [Rows] は自動的に閉じられ、 [Rows.Err] の結果を確認するだけで十分です。
// Closeは冪等性があり、 [Rows.Err] の結果に影響を与えません。
func (rs *Rows) Close() error

// Rowは、単一の行を選択するために [DB.QueryRow] を呼び出した結果です。
type Row struct {
	// One of these two will be non-nil:
	err  error
	rows *Rows
}

// Scanは、一致する行から列をdestが指す値にコピーします。
// 詳細については、 [Rows.Scan] のドキュメントを参照してください。
// クエリに複数の行が一致する場合、Scanは最初の行を使用し、残りを破棄します。
// クエリに一致する行がない場合、Scanは [ErrNoRows] を返します。
func (r *Row) Scan(dest ...any) error

// Errは、 [Row.Scan] を呼び出さずにクエリエラーをチェックするための方法を提供します。
// Errは、クエリを実行する際に遭遇したエラー（ある場合）を返します。
// このエラーがnilでない場合、このエラーは [Row.Scan] からも返されます。
func (r *Row) Err() error

// Resultは、実行されたSQLコマンドを要約します。
type Result interface {
	LastInsertId() (int64, error)

	RowsAffected() (int64, error)
}
