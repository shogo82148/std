// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// testingパッケージはGoパッケージの自動テストをサポートします。
// これは"go test"コマンドと一緒に使用することを意図しています。"go test"コマンドは以下の形式の関数を自動的に実行します。
//
//	func TestXxx(*testing.T)
//
// ただし、Xxxは小文字ではじまらないものとします。関数名はテストルーチンを識別するために使用されます。
//
// これらの関数内では、失敗を通知するために [T.Error]、[T.Fail] などの関連メソッドを使用してください。
//
// 新しいテストスイートを書くためには、次の要件に従っているファイルを作成し、"_test.go"で終わる名前を付けます。
// このファイルは通常のパッケージのビルドから除外されますが、"go test"コマンドの実行時には含まれます。
//
// テストファイルは、テスト対象のパッケージと同じパッケージにあるか、"_test"という接尾辞の付いた対応するパッケージに含めることができます。
//
// テストファイルが同じパッケージにある場合、パッケージ内の非公開の識別子を参照することができます。次の例のようになります。
//
//	package abs
//
//	import "testing"
//
//	func TestAbs(t *testing.T) {
//	    got := Abs(-1)
//	    if got != 1 {
//	        t.Errorf("Abs(-1) = %d; want 1", got)
//	    }
//	}
//
// ファイルが別の"_test"パッケージにある場合、テスト対象のパッケージを明示的にインポートする必要があり、公開された識別子のみ使用できます。これは「ブラックボックス」テストとして知られています。
//
//	package abs_test
//
//	import (
//		"testing"
//
//		"path_to_pkg/abs"
//	)
//
//	func TestAbs(t *testing.T) {
//	    got := abs.Abs(-1)
//	    if got != 1 {
//	        t.Errorf("Abs(-1) = %d; want 1", got)
//	    }
//	}
//
// 詳細については [go help test] および [go help testflag] を実行してください。
//
// # Benchmarks
//
// 次の形式の関数
//
//	func BenchmarkXxx(*testing.B)
//
// はベンチマークとみなされ、-benchフラグが指定されたときに"go test"コマンドによって
// 実行されます。ベンチマークは逐次的に実行されます。
//
// テストフラグの説明については、 [go help testflag] を参照してください。
//
// ベンチマーク関数の例は次のとおりです。
//
//	func BenchmarkRandInt(b *testing.B) {
//	    for b.Loop() {
//	        rand.Int()
//	    }
//	}
//
// 出力
//
//	BenchmarkRandInt-8   	68453040	        17.8 ns/op
//
// は、ループの本体が1ループあたり17.8 nsの速度で68453040回実行されたことを意味します。
//
// 時間が計測されるのはループの本体だけなので、ベンチマークはb.Loopを呼び出す前に
// コストの高いセットアップを行うことができ、それはベンチマークの計測には含まれません。
//
//	func BenchmarkBigLen(b *testing.B) {
//	    big := NewBig()
//	    for b.Loop() {
//	        big.Len()
//	    }
//	}
//
// ベンチマークが並列な環境でのパフォーマンスをテストする必要がある場合は、
// RunParallelヘルパー関数を使用できます。このようなベンチマークは、
// go test -cpuフラグと一緒に使用することを意図しています。
//
//	func BenchmarkTemplateParallel(b *testing.B) {
//	    templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
//	    b.RunParallel(func(pb *testing.PB) {
//	        var buf bytes.Buffer
//	        for pb.Next() {
//	            buf.Reset()
//	            templ.Execute(&buf, "World")
//	        }
//	    })
//	}
//
// ベンチマーク結果のフォーマットの詳細な仕様は、
// https://go.dev/design/14313-benchmark-format にあります。
//
// ベンチマーク結果を扱うための標準的なツールが [golang.org/x/perf/cmd] にあります。
// 特に、 [golang.org/x/perf/cmd/benchstat] は統計的に頑健なA/B比較を行います。
//
// # b.N-style benchmarks
//
// [B.Loop] が導入される以前は、ベンチマークはB.Nを使った別の形式で書かれていました。
// 例えば次のとおりです。
//
//	func BenchmarkRandInt(b *testing.B) {
//	    for range b.N {
//	        rand.Int()
//	    }
//	}
//
// この形式のベンチマークでは、ベンチマーク関数は対象のコードをb.N回実行しなければなりません。
// ベンチマーク関数は、確実に時間を計測できるほど長く続くようになるまで、
// b.Nを調整しながら複数回呼び出されます。これは、ループの前に行われるセットアップも
// 複数回実行される可能性があることを意味します。
//
// ベンチマークの実行前にコストの高いセットアップが必要な場合は、
// タイマーを明示的にリセットする必要があります。
//
//	func BenchmarkBigLen(b *testing.B) {
//	    big := NewBig()
//	    b.ResetTimer()
//	    for range b.N {
//	        big.Len()
//	    }
//	}
//
// 新しいベンチマークでは、より頑健で効率的な [B.Loop] を使用することを推奨します。
//
// # Examples
//
// このパッケージは、サンプルコードの実行と検証も行います。Example関数には、
// "Output:"で始まる結びの行コメントを含めることができ、テストの実行時に関数の
// 標準出力と比較されます。（比較では先頭と末尾の空白は無視されます。）
// 以下はExampleの例です。
//
//	func ExampleHello() {
//	    fmt.Println("hello")
//	    // Output: hello
//	}
//
//	func ExampleSalutations() {
//	    fmt.Println("hello, and")
//	    fmt.Println("goodbye")
//	    // Output:
//	    // hello, and
//	    // goodbye
//	}
//
// コメントの接頭辞"Unordered output:"は"Output:"に似ていますが、
// 行の順序を問わずにマッチします。
//
//	func ExamplePerm() {
//	    for _, value := range Perm(5) {
//	        fmt.Println(value)
//	    }
//	    // Unordered output: 4
//	    // 2
//	    // 1
//	    // 3
//	    // 0
//	}
//
// 出力コメントのないExample関数は、コンパイルされますが実行されません。
//
// パッケージ、関数F、型T、および型TのメソッドMに対するExampleを宣言するための
// 命名規則は次のとおりです。
//
//	func Example() { ... }
//	func ExampleF() { ... }
//	func ExampleT() { ... }
//	func ExampleT_M() { ... }
//
// パッケージ/型/関数/メソッドに対する複数のExample関数は、名前に異なる接尾辞を
// 付けることで提供できます。接尾辞は小文字で始まる必要があります。
//
//	func Example_suffix() { ... }
//	func ExampleF_suffix() { ... }
//	func ExampleT_suffix() { ... }
//	func ExampleT_M_suffix() { ... }
//
// テストファイルが、1つのExample関数と、少なくとも1つの他の関数、型、変数、
// または定数の宣言を含み、テスト関数やベンチマーク関数を含まない場合、
// テストファイル全体がExampleとして提示されます。
//
// # Fuzzing
//
// 'go test'とtestingパッケージはファジングをサポートしています。ファジングは、
// ユニットテストでは予期されていなかったバグを見つけるために、ランダムに生成された
// 入力で関数を呼び出すテスト手法です。
//
// 次の形式の関数
//
//	func FuzzXxx(*testing.F)
//
// はファズテストとみなされます。
//
// 例えば次のとおりです。
//
//	func FuzzHex(f *testing.F) {
//	  for _, seed := range [][]byte{{}, {0}, {9}, {0xa}, {0xf}, {1, 2, 3, 4}} {
//	    f.Add(seed)
//	  }
//	  f.Fuzz(func(t *testing.T, in []byte) {
//	    enc := hex.EncodeToString(in)
//	    out, err := hex.DecodeString(enc)
//	    if err != nil {
//	      t.Fatalf("%v: decode: %v", in, err)
//	    }
//	    if !bytes.Equal(in, out) {
//	      t.Fatalf("%v: not equal after round trip: %v", in, out)
//	    }
//	  })
//	}
//
// ファズテストはシードコーパス、つまりデフォルトで実行され、入力の生成の種となる
// 入力の集合を保持します。シード入力は、 [F.Add] を呼び出すか、ファズテストを含む
// パッケージ内のディレクトリtestdata/fuzz/<Name>（<Name>はファズテストの名前）に
// ファイルを格納することで登録できます。シード入力は省略可能ですが、
// コードカバレッジの良い小さなシード入力の集合が与えられると、ファジングエンジンは
// より効率的にバグを見つけられる場合があります。これらのシード入力は、ファジングによって
// 特定されたバグに対するリグレッションテストとしても機能します。
//
// ファズテスト内で [F.Fuzz] に渡される関数は、ファズターゲットとみなされます。
// ファズターゲットは、 [*T] パラメータと、それに続くランダムな入力のための1つ以上の
// パラメータを受け取らなければなりません。 [F.Add] に渡される引数の型は、
// これらのパラメータの型と同一でなければなりません。ファズターゲットは、テストと同じ方法、
// つまり [T.Fail] （または [T.Error] や [T.Fatal] のようにそれを呼び出すメソッド）を
// 呼び出すか、パニックすることで、問題を見つけたことを通知できます。
//
// ファジングが有効な場合（-fuzzフラグに特定のファズテストにマッチする正規表現を
// 設定した場合）、ファズターゲットは、シード入力にランダムな変更を繰り返し加えて生成された
// 引数で呼び出されます。サポートされているプラットフォームでは、'go test'は
// ファジング用のカバレッジ計装を施してテスト実行ファイルをコンパイルします。
// ファジングエンジンはその計装を使用して、カバレッジを拡大する入力を見つけて
// キャッシュし、バグを見つける可能性を高めます。ある入力に対してファズターゲットが
// 失敗した場合、ファジングエンジンは失敗の原因となった入力を、パッケージディレクトリ内の
// ディレクトリtestdata/fuzz/<Name>にあるファイルに書き込みます。このファイルは後で
// シード入力として機能します。その場所にファイルを書き込めない場合（例えば、
// ディレクトリが読み取り専用である場合）、ファジングエンジンは代わりにビルドキャッシュ内の
// ファズキャッシュディレクトリにファイルを書き込みます。
//
// ファジングが無効な場合、ファズターゲットは [F.Add] で登録されたシード入力と
// testdata/fuzz/<Name>のシード入力で呼び出されます。このモードでは、ファズテストは
// 通常のテストとほぼ同じように動作し、サブテストは [T.Run] の代わりに [F.Fuzz] で開始されます。
//
// ファジングについてのドキュメントは https://go.dev/doc/fuzz を参照してください。
//
// # Skipping
//
// テストやベンチマークは、 [T.Skip] または [B.Skip] を呼び出すことで、
// 実行時にスキップできます。
//
//	func TestTimeConsuming(t *testing.T) {
//	    if testing.Short() {
//	        t.Skip("skipping test in short mode.")
//	    }
//	    ...
//	}
//
// [T.Skip] メソッドは、入力が無効であるものの失敗する入力とみなすべきではない場合に、
// ファズターゲット内で使用できます。例えば次のとおりです。
//
//	func FuzzJSONMarshaling(f *testing.F) {
//	    f.Fuzz(func(t *testing.T, b []byte) {
//	        var v interface{}
//	        if err := json.Unmarshal(b, &v); err != nil {
//	            t.Skip()
//	        }
//	        if _, err := json.Marshal(v); err != nil {
//	            t.Errorf("Marshal: %v", err)
//	        }
//	    })
//	}
//
// # Subtests and Sub-benchmarks
//
// [T.Run] と [B.Run] メソッドを使うと、それぞれに別々の関数を定義することなく、
// サブテストとサブベンチマークを定義できます。これにより、テーブル駆動のベンチマークや
// 階層的なテストの作成といった使い方が可能になります。
// また、共通のセットアップとティアダウンのコードを共有する方法も提供します。
//
//	func TestFoo(t *testing.T) {
//	    // <setup code>
//	    t.Run("A=1", func(t *testing.T) { ... })
//	    t.Run("A=2", func(t *testing.T) { ... })
//	    t.Run("B=1", func(t *testing.T) { ... })
//	    // <tear-down code>
//	}
//
// 各サブテストとサブベンチマークは一意な名前を持ちます。それは、トップレベルのテストの
// 名前とRunに渡された名前の並びをスラッシュで区切って組み合わせたもので、
// 曖昧さをなくすために末尾に連番が付く場合があります。
//
// -run、-bench、-fuzzコマンドラインフラグの引数は、テストの名前にマッチする
// アンカーなしの正規表現です。サブテストのようにスラッシュで区切られた複数の要素を持つ
// テストの場合、引数自体もスラッシュで区切られ、各式が名前の各要素に順にマッチします。
// アンカーなしであるため、空の式は任意の文字列にマッチします。
// 例えば、「マッチする」を「名前に含む」という意味で使うと次のようになります。
//
//	go test -run ''        # Run all tests.
//	go test -run Foo       # Run top-level tests matching "Foo", such as "TestFooBar".
//	go test -run Foo/A=    # For top-level tests matching "Foo", run subtests matching "A=".
//	go test -run /A=1      # For all top-level tests, run subtests matching "A=1".
//	go test -fuzz FuzzFoo  # Fuzz the target matching "FuzzFoo"
//
// -run引数は、デバッグのためにシードコーパス内の特定の値を実行するためにも使用できます。
// 例えば次のとおりです。
//
//	go test -run=FuzzFoo/9ddb952d9814
//
// -fuzzフラグと-runフラグの両方を設定すると、ターゲットをファジングしつつ、
// 他のすべてのテストの実行をスキップできます。
//
// サブテストは並列性の制御にも使用できます。親テストは、そのすべてのサブテストが
// 完了したときにのみ完了します。この例では、定義されている可能性のある他の
// トップレベルのテストに関係なく、すべてのテストが互いに並列に、かつ互いの間でのみ
// 並列に実行されます。
//
//	func TestGroupedParallel(t *testing.T) {
//	    for _, tc := range tests {
//	        t.Run(tc.Name, func(t *testing.T) {
//	            t.Parallel()
//	            ...
//	        })
//	    }
//	}
//
// Runは並列のサブテストが完了するまで戻らないため、並列テストのグループの後で
// クリーンアップを行う方法を提供します。
//
//	func TestTeardownParallel(t *testing.T) {
//	    // This Run will not return until the parallel tests finish.
//	    t.Run("group", func(t *testing.T) {
//	        t.Run("Test1", parallelTest1)
//	        t.Run("Test2", parallelTest2)
//	        t.Run("Test3", parallelTest3)
//	    })
//	    // <tear-down code>
//	}
//
// # Main
//
// テストやベンチマークのプログラムで、実行の前後に追加のセットアップやティアダウンを
// 行う必要がある場合があります。また、どのコードをメインスレッドで実行するかを
// 制御する必要がある場合もあります。これらのケースやその他のケースをサポートするために、
// テストファイルに次の関数が含まれている場合、
//
//	func TestMain(m *testing.M)
//
// 生成されたテストは、テストやベンチマークを直接実行する代わりにTestMain(m)を呼び出します。
// TestMainはメインゴルーチンで実行され、m.Runの呼び出しの前後で必要なセットアップと
// ティアダウンを何でも行うことができます。m.Runは [os.Exit] に渡すことができる
// 終了コードを返します。TestMainが戻った場合、テストのラッパーがm.Runの結果を
// 自ら [os.Exit] に渡します。
//
// TestMainが呼び出された時点では、flag.Parseは実行されていません。TestMainが
// testingパッケージのフラグを含むコマンドラインフラグに依存する場合は、
// 明示的に [flag.Parse] を呼び出す必要があります。コマンドラインフラグは、
// テスト関数やベンチマーク関数が実行される時点では常に解析済みです。
//
// TestMainの簡単な実装は次のとおりです。
//
//	func TestMain(m *testing.M) {
//		// call flag.Parse() here if TestMain uses flags
//		m.Run()
//	}
//
// TestMainは低レベルのプリミティブであり、通常のテスト関数で十分な
// 一般的なテストの用途では必要ないはずです。
//
// [go help test]: https://pkg.go.dev/cmd/go#hdr-Test_packages
// [go help testflag]: https://pkg.go.dev/cmd/go#hdr-Testing_flags
package testing

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/time"
)

// Initはテストフラグを登録します。これらのフラグは、テスト関数を実行する前に"go test"コマンドによって自動的に登録されるため、Initは、"go test"を使用せずにBenchmarkなどの関数を呼び出す場合にのみ必要です。
//
// Initは並行して呼び出すことは安全ではありません。 すでに呼び出されている場合は、効果がありません。
func Init()

// Short は -test.short フラグが設定されているかどうかを報告します。
func Short() bool

// Testingは現在のコードがテストで実行されているかどうかを報告します。
// "go test"で作られたプログラムではtrueを報告し、
// "go build"で作られたプログラムではfalseを報告します。
func Testing() bool

// CoverModeはテストカバレッジモードの設定を報告します。値は「set」、「count」または「atomic」です。テストカバレッジが有効でない場合、戻り値は空になります。
func CoverMode() string

// Verboseは、-test.vフラグが設定されているかどうかを報告します。
func Verbose() bool

// TBは [T], [B], [F] に共通するインターフェースです。
type TB interface {
	ArtifactDir() string
	Attr(key, value string)
	Cleanup(func())
	Error(args ...any)
	Errorf(format string, args ...any)
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	Helper()
	Log(args ...any)
	Logf(format string, args ...any)
	Name() string
	Setenv(key, value string)
	Chdir(dir string)
	Skip(args ...any)
	SkipNow()
	Skipf(format string, args ...any)
	Skipped() bool
	TempDir() string
	Context() context.Context
	Output() io.Writer

	private()
}

var (
	_ TB = (*T)(nil)
	_ TB = (*B)(nil)
)

// Tはテスト状態を管理し、フォーマットされたテストログをサポートするためにTest関数に渡される型です。
//
// テストは、そのTest関数がreturnするか、[T.FailNow]、[T.Fatal]、[T.Fatalf]、[T.SkipNow]、[T.Skip]、[T.Skipf] のいずれかのメソッドを呼び出すことで終了します。
// これらのメソッドおよび [T.Parallel] メソッドは、Test関数を実行しているゴルーチンからのみ呼び出す必要があります。
//
// その他の報告用メソッド（[T.Log] や [T.Error] のバリエーションなど）は、複数のゴルーチンから同時に呼び出すことができます。
type T struct {
	common
	denyParallel bool
	tstate       *testState
}

// Parallel は、このテストを他の並列テストとだけ並行して実行することを示し、
// すべての非並列テストが終了するまで待機します。
//
// -test.count または -test.cpu の使用によってテストが複数回実行される場合、
// 同一テストの複数のインスタンスが互いに並行して実行されることはありません。
func (t *T) Parallel()

// Setenvはos.Setenv(key, value)を呼び出し、Cleanupを使用してテスト後に環境変数を元の値に復元します。
//
// Setenvはプロセス全体に影響を与えるため、並列テストや並列祖先を持つテストで使用することはできません。
func (t *T) Setenv(key, value string)

// Chdirは [os.Chdir] を呼び出し、テスト後に作業ディレクトリを元の値に復元するためにCleanupを使用します。Unixでは、テストの実行中のみPWD環境変数も設定されます。
//
// Chdirはプロセス全体に影響を与えるため、並列テストや並列祖先を持つテストでは使用できません。
func (t *T) Chdir(dir string)

// InternalTestは内部型ですが、クロスパッケージであるためにエクスポートされています。
// これは"go test"コマンドの実装の一部です。
type InternalTest struct {
	Name string
	F    func(*T)
}

// Runはfをtのサブテストとして実行します。nameという名前で実行されます。fは別のゴルーチンで実行され、
// fが返るか、t.Parallelを呼び出して並列テストになるまでブロックされます。
// Runはfが成功したか（またはt.Parallelを呼び出す前に失敗しなかったか）を報告します。
//
// Runは複数のゴルーチンから同時に呼び出すことができますが、そのようなすべての呼び出しは、
// tが外部テスト関数を返す前に返さなければなりません。
func (t *T) Run(name string, f func(t *T)) bool

// Deadlineは、-timeoutフラグで指定されたタイムアウトを超えるテストバイナリの時間を報告します。
//
// -timeoutフラグが「タイムアウトなし」（0）を示す場合、ok結果はfalseです。
func (t *T) Deadline() (deadline time.Time, ok bool)

// Mainは"go test"コマンドの実装の一部である内部関数です。
// クロスパッケージであり "internal" パッケージ以前から存在するためエクスポートされています。
// 現在は"go test"では使用されていませんが、可能な限り他の"go test"を模倣するシステムのために保持されています。
// ただし、testingパッケージに新機能が追加されるとMainを更新できない場合があります。
// "go test"を模倣するシステムは [MainStart] の使用に切り替えるべきです。
func Main(matchString func(pat, str string) (bool, error), tests []InternalTest, benchmarks []InternalBenchmark, examples []InternalExample)

// MはTestMain関数に渡される型で、実際のテストを実行するために使用されます。
type M struct {
	deps        testDeps
	tests       []InternalTest
	benchmarks  []InternalBenchmark
	fuzzTargets []InternalFuzzTarget
	examples    []InternalExample

	timer     *time.Timer
	afterOnce sync.Once

	numRun int

	// os.Exitに渡す値、外部のtest func mainに
	// harnessはこのコードでos.Exitを呼び出します。#34129を参照してください。
	exitCode int
}

// MainStartは「go test」によって生成されたテストで使用することを意図しています。
// 直接呼び出すことは意図されておらず、Go 1の互換性ドキュメントの対象外です。
// リリースごとにシグネチャが変更される可能性があります。
func MainStart(deps testDeps, tests []InternalTest, benchmarks []InternalBenchmark, fuzzTargets []InternalFuzzTarget, examples []InternalExample) *M

// Runはテストを実行します。os.Exitに渡す終了コードを返します。
// 終了コードはすべてのテストが成功した場合はゼロ、何らかの失敗があった場合は非ゼロです。
// 機械可読なテスト結果が必要な場合は、'go test -json'の出力を解析してください。
func (m *M) Run() (code int)

// RunTestsは内部関数ですが、クロスパッケージであるためにエクスポートされています。
// これは"go test"コマンドの実装の一部です。
func RunTests(matchString func(pat, str string) (bool, error), tests []InternalTest) (ok bool)
