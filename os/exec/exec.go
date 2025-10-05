// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// execパッケージは外部コマンドを実行します。これはos.StartProcessをラップして、
// stdinとstdoutのリマップ、パイプを使用したI/Oの接続、その他の調整を
// 簡単に行うことができます。
//
// Cや他の言語からの"system"ライブラリ呼び出しとは異なり、
// os/execパッケージは意図的にシステムシェルを呼び出さず、
// グロブパターンを展開したり、シェルが通常行う他の展開、
// パイプライン、リダイレクションを処理しません。このパッケージは
// Cの"exec"関数群のように振る舞います。グロブパターンを展開するには、
// シェルを直接呼び出し、危険な入力をエスケープするか、
// [path/filepath] パッケージのGlob関数を使用します。
// 環境変数を展開するには、osパッケージのExpandEnvを使用します。
//
// このパッケージの例はUnixシステムを前提としています。
// Windowsでは動作しない場合があり、go.dev や pkg.go.dev で使われているGo Playgroundでは実行できません。
//
// # Executables in the current directory
//
// 関数 [Command] と [LookPath] は、ホストオペレーティングシステムの規則に従って、
// 現在のパスにリストされたディレクトリでプログラムを探します。
// オペレーティングシステムは何十年もの間、この検索に現在の
// ディレクトリを含めてきました。これは時々暗黙的に、時々
// デフォルトで明示的にそのように設定されています。
// 現代の慣行では、現在のディレクトリを含めることは通常予期しないもので、
// しばしばセキュリティ問題につながります。
//
// これらのセキュリティ問題を避けるために、Go 1.19から、このパッケージはプログラムを
// 現在のディレクトリに対する暗黙的または明示的なパスエントリを使用して解決しません。
// つまり、[LookPath]("go")を実行すると、パスがどのように設定されていても、
// Unixでは./go、Windowsでは.\go.exeを正常に返すことはありません。
// 代わりに、通常のパスアルゴリズムがその答えをもたらす場合、
// これらの関数はエラーerrを返し、[errors.Is](err, [ErrDot])を満たします。
//
// 例えば、以下の2つのプログラムスニペットを考えてみてください：
//
//	path, err := exec.LookPath("prog")
//	if err != nil {
//		log.Fatal(err)
//	}
//	use(path)
//
// そして
//
//	cmd := exec.Command("prog")
//	if err := cmd.Run(); err != nil {
//		log.Fatal(err)
//	}
//
// これらは、現在のパスの設定に関係なく、./progや.\prog.exeを見つけて実行することはありません。
//
// 常に現在のディレクトリからプログラムを実行したいコードは、"prog"の代わりに"./prog"と指定することで書き換えることができます。
//
// 相対パスエントリからの結果を含めることに固執するコードは、代わりに errors.Is チェックを使用してエラーをオーバーライドできます：
//
//	path, err := exec.LookPath("prog")
//	if errors.Is(err, exec.ErrDot) {
//		err = nil
//	}
//	if err != nil {
//		log.Fatal(err)
//	}
//	use(path)
//
// そして
//
//	cmd := exec.Command("prog")
//	if errors.Is(cmd.Err, exec.ErrDot) {
//		cmd.Err = nil
//	}
//	if err := cmd.Run(); err != nil {
//		log.Fatal(err)
//	}
//
// 環境変数GODEBUG=execerrdot=0を設定すると、
// ErrDotの生成が完全に無効になり、よりターゲット指向の修正を適用できないプログラムに対して、
// 一時的にGo 1.19以前の動作が復元されます。
// Goの将来のバージョンでは、この変数のサポートが削除される可能性があります。
//
// そのようなオーバーライドを追加する前に、
// それを行うことのセキュリティ上の意味を理解しておいてください。
// 詳細は https://go.dev/blog/path-security を参照してください。
package exec

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/syscall"
	"github.com/shogo82148/std/time"
)

// Errorは、[LookPath] がファイルを実行可能なものとして分類できなかったときに返されます。
type Error struct {
	// Nameは、エラーが発生したファイル名です。
	Name string
	// Errは、基になるエラーです。
	Err error
}

func (e *Error) Error() string

func (e *Error) Unwrap() error

// ErrWaitDelayは、プロセスが成功したステータスコードで終了するが、
// コマンドのWaitDelayが期限切れになる前にその出力パイプが閉じられない場合、
// [Cmd.Wait] によって返されます。
var ErrWaitDelay = errors.New("exec: WaitDelay expired before I/O complete")

// Cmdは、準備中または実行中の外部コマンドを表します。
//
// Cmdは、[Cmd.Run]、[Cmd.Output]、または [Cmd.CombinedOutput] メソッドを呼び出した後では再利用できません。
type Cmd struct {
	// Pathは、実行するコマンドのパスです。
	//
	// これは唯一、ゼロ以外の値に設定しなければならないフィールドです。
	// Pathが相対パスの場合、Dirに対して相対的に評価されます。
	Path string

	// Argsはコマンドライン引数を保持します。コマンド自体はArgs[0]として含まれます。
	// Argsフィールドが空またはnilの場合、Runは{Path}を使用します。
	//
	// 典型的な使用では、PathとArgsの両方はCommandを呼び出すことで設定されます。
	Args []string

	// Envはプロセスの環境変数を指定します。
	// 各要素は "key=value" の形式です。
	// Envがnilの場合、新しいプロセスは現在のプロセスの環境を使用します。
	// Envに重複するキーが含まれている場合、重複したキーごとに最後の値のみが使用されます。
	// Windowsでは特別なケースとして、SYSTEMROOTが存在しない場合は必ず追加され、空文字列に明示的に設定されていない限り追加されます。
	//
	// また、Dirフィールドも参照してください。これは環境変数PWDを設定する場合があります。
	Env []string

	// Dirはコマンドの作業ディレクトリを指定します。
	// Dirが空文字列の場合、Runは呼び出し元プロセスの現在のディレクトリでコマンドを実行します。
	//
	// Unixシステムでは、Dirの値は他に指定がなければ子プロセスのPWD環境変数も決定します。
	// Unixプロセスは作業ディレクトリを名前ではなく、ファイルツリー内のノードへの暗黙的な参照として表現します。
	// そのため、子プロセスがCのgetcwdなどの関数を呼び出してファイルツリーを辿って正規名を計算した場合、
	// Dirの値がシンボリックリンクを含むエイリアスだった場合は元の値を復元できません。
	// しかし、子プロセスがGoの [os.Getwd] やGNU Cのget_current_dir_nameを呼び出し、PWDの値が現在のディレクトリのエイリアスであれば、
	// それらの関数はPWDの値、つまりDirの値を返します。
	Dir string

	// Stdinはプロセスの標準入力を指定します。
	//
	// Stdinがnilの場合、プロセスはnullデバイス(os.DevNull)から読み取ります。
	//
	// Stdinが*os.Fileの場合、プロセスの標準入力はそのファイルに直接接続されます。
	//
	// それ以外の場合、コマンドの実行中に別のgoroutineがStdinから読み取り、
	// そのデータをパイプ経由でコマンドに送信します。この場合、Waitはgoroutineが
	// コピーを停止するまで完了しません。これは、Stdinの終わりに達したため（EOFまたは読み取りエラー）、
	// パイプへの書き込みがエラーを返したため、または非ゼロのWaitDelayが設定されて期限切れになったためです。
	Stdin io.Reader

	// StdoutとStderrは、プロセスの標準出力とエラーを指定します。
	//
	// どちらかがnilの場合、Runは対応するファイルディスクリプタを
	// nullデバイス(os.DevNull)に接続します。
	//
	// どちらかが*os.Fileの場合、プロセスからの対応する出力は
	// そのファイルに直接接続されます。
	//
	// それ以外の場合、コマンドの実行中に別のgoroutineがプロセスからパイプ経由で読み取り、
	// そのデータを対応するWriterに送信します。この場合、Waitはgoroutineが
	// EOFに達するか、エラーに遭遇するか、非ゼロのWaitDelayが期限切れになるまで完了しません。
	//
	// StdoutとStderrが同じWriterで、==で比較できる型を持っている場合、
	// 同時に最大1つのgoroutineだけがWriteを呼び出します。
	Stdout io.Writer
	Stderr io.Writer

	// ExtraFilesは、新しいプロセスに継承される追加のオープンファイルを指定します。
	// 標準入力、標準出力、または標準エラーは含まれません。非nilの場合、エントリiは
	// ファイルディスクリプタ3+iになります。
	//
	// ExtraFilesはWindowsではサポートされていません。
	ExtraFiles []*os.File

	// SysProcAttrは、オプションのオペレーティングシステム固有の属性を保持します。
	// Runは、os.ProcAttrのSysフィールドとしてos.StartProcessに渡します。
	SysProcAttr *syscall.SysProcAttr

	// Processは、開始された後の基本的なプロセスです。
	Process *os.Process

	// ProcessStateは、終了したプロセスに関する情報を含みます。
	// プロセスが正常に開始された場合、コマンドが完了するとWaitまたはRunが
	// そのProcessStateを設定します。
	ProcessState *os.ProcessState

	// ctx is the context passed to CommandContext, if any.
	ctx context.Context

	Err error

	// Cancelがnilでない場合、コマンドはCommandContextで作成されていなければならず、
	// コマンドのContextが完了したときにCancelが呼び出されます。デフォルトでは、
	// CommandContextはCancelをコマンドのProcessのKillメソッドを呼び出すように設定します。
	//
	// 通常、カスタムCancelはコマンドのProcessにシグナルを送信しますが、
	// 代わりにキャンセルを開始するための他のアクションを取ることもあります。
	// 例えば、stdinやstdoutのパイプを閉じる、またはネットワークソケットにシャットダウンリクエストを送信するなどです。
	//
	// Cancelが呼び出された後にコマンドが成功ステータスで終了し、
	// そしてCancelがos.ErrProcessDoneと等価のエラーを返さない場合、
	// Waitや類似のメソッドは非nilのエラーを返します：Cancelによって返されたエラーをラップするエラー、
	// またはContextからのエラーです。
	// (コマンドが非成功ステータスで終了する場合、またはCancelがos.ErrProcessDoneをラップするエラーを返す場合、
	// Waitや類似のメソッドは引き続きコマンドの通常の終了ステータスを返します。)
	//
	// Cancelがnilに設定されている場合、コマンドのContextが完了したときにはすぐには何も起こりませんが、
	// 非ゼロのWaitDelayは依然として効果を発揮します。これは、例えば、シャットダウンシグナルをサポートしていないが、
	// 常にすぐに終了することが期待されるコマンドのデッドロックを回避するために役立つかもしれません。
	//
	// Startが非nilのエラーを返す場合、Cancelは呼び出されません。
	Cancel func() error

	// WaitDelayが非ゼロの場合、Waitで予期しない遅延の2つの源に対する待機時間を制限します：
	// 関連するContextがキャンセルされた後も終了しない子プロセス、およびI/Oパイプを閉じずに終了する子プロセス。
	//
	// WaitDelayタイマーは、関連付けられたContextが完了したとき、または
	// Waitの呼び出しで子プロセスが終了したことが確認されたときのいずれか早い方から開始します。
	// 遅延が経過すると、コマンドは子プロセスと/またはそのI/Oパイプをシャットダウンします。
	//
	// 子プロセスが終了に失敗した場合 — たとえば、Cancel関数からのシャットダウンシグナルを無視したり、
	// 受信に失敗したりした場合、またはCancel関数が設定されていなかった場合 — それはos.Process.Killを使用して終了されます。
	//
	// その後、子プロセスと通信するI/Oパイプがまだ開いている場合、
	// それらのパイプは、現在ReadまたはWrite呼び出しでブロックされているgoroutineを解除するために閉じられます。
	//
	// WaitDelayによりパイプが閉じられ、Cancelの呼び出しが行われておらず、
	// コマンドがそれ以外の点で成功ステータスで終了した場合、Waitや類似のメソッドは
	// nilの代わりにErrWaitDelayを返します。
	//
	// WaitDelayがゼロ（デフォルト）の場合、I/OパイプはEOFまで読み取られます。
	// これは、コマンドの孤立したサブプロセスもパイプのディスクリプタを閉じるまで発生しないかもしれません。
	WaitDelay time.Duration

	// childIOFiles holds closers for any of the child process's
	// stdin, stdout, and/or stderr files that were opened by the Cmd itself
	// (not supplied by the caller). These should be closed as soon as they
	// are inherited by the child process.
	childIOFiles []io.Closer

	// parentIOPipes holds closers for the parent's end of any pipes
	// connected to the child's stdin, stdout, and/or stderr streams
	// that were opened by the Cmd itself (not supplied by the caller).
	// These should be closed after Wait sees the command and copying
	// goroutines exit, or after WaitDelay has expired.
	parentIOPipes []io.Closer

	// goroutine holds a set of closures to execute to copy data
	// to and/or from the command's I/O pipes.
	goroutine []func() error

	// If goroutineErr is non-nil, it receives the first error from a copying
	// goroutine once all such goroutines have completed.
	// goroutineErr is set to nil once its error has been received.
	goroutineErr <-chan error

	// If ctxResult is non-nil, it receives the result of watchCtx exactly once.
	ctxResult <-chan ctxResult

	// The stack saved when the Command was created, if GODEBUG contains
	// execwait=2. Used for debugging leaks.
	createdByStack []byte

	// For a security release long ago, we created x/sys/execabs,
	// which manipulated the unexported lookPathErr error field
	// in this struct. For Go 1.19 we exported the field as Err error,
	// above, but we have to keep lookPathErr around for use by
	// old programs building against new toolchains.
	// The String and Start methods look for an error in lookPathErr
	// in preference to Err, to preserve the errors that execabs sets.
	//
	// In general we don't guarantee misuse of reflect like this,
	// but the misuse of reflect was by us, the best of various bad
	// options to fix the security problem, and people depend on
	// those old copies of execabs continuing to work.
	// The result is that we have to leave this variable around for the
	// rest of time, a compatibility scar.
	//
	// See https://go.dev/blog/path-security
	// and https://go.dev/issue/43724 for more context.
	lookPathErr error

	// cachedLookExtensions caches the result of calling lookExtensions.
	// It is set when Command is called with an absolute path, letting it do
	// the work of resolving the extension, so Start doesn't need to do it again.
	// This is only used on Windows.
	cachedLookExtensions struct{ in, out string }
}

// Commandは、指定されたプログラムを
// 与えられた引数で実行するための [Cmd] 構造体を返します。
//
// それは返される構造体の中でPathとArgsだけを設定します。
//
// nameにパスセパレータが含まれていない場合、Commandは [LookPath] を使用して
// 可能な場合にはnameを完全なパスに解決します。それ以外の場合、nameを
// 直接Pathとして使用します。
//
// 返されるCmdのArgsフィールドは、コマンド名に続くargの要素から構築されます。
// したがって、argにはコマンド名自体を含めないでください。例えば、Command("echo", "hello")。
// Args[0]は常にnameで、解決されたPathではありません。
//
// Windowsでは、プロセスはコマンドライン全体を単一の文字列として受け取り、
// 自身でパースします。CommandはArgsを結合し、引用符で囲んで、
// CommandLineToArgvWを使用するアプリケーションと互換性のあるアルゴリズムで
// コマンドライン文字列にします（これが最も一般的な方法です）。注目すべき例外は、
// msiexec.exeとcmd.exe（したがって、すべてのバッチファイル）で、これらは異なる
// アンクォートアルゴリズムを持っています。これらまたは他の類似のケースでは、
// 自分で引用符を付けてSysProcAttr.CmdLineに完全なコマンドラインを提供し、
// Argsを空にすることができます。
func Command(name string, arg ...string) *Cmd

// CommandContextは [Command] と同様ですが、contextが含まれています。
//
// 提供されたcontextは、コマンドが自身で完了する前にcontextがdoneになった場合、
// プロセスを中断するために使用されます（cmd.Cancelまたは [os.Process.Kill] を呼び出す）。
//
// CommandContextは、コマンドのCancel関数をそのProcessのKillメソッドを呼び出すように設定し、
// WaitDelayは未設定のままにします。呼び出し元は、コマンドを開始する前にこれらのフィールドを
// 変更することでキャンセルの振る舞いを変更することができます。
func CommandContext(ctx context.Context, name string, arg ...string) *Cmd

// Stringは、cの人間が読める説明を返します。
// これはデバッグ専用です。
// 特に、シェルへの入力として使用するのには適していません。
// Stringの出力はGoのリリースによって異なる可能性があります。
func (c *Cmd) String() string

// Runは指定されたコマンドを開始し、その完了を待ちます。
//
// 返されるエラーは、コマンドが実行され、stdin、stdout、stderrのコピーに問題がなく、
// ゼロの終了ステータスで終了した場合にはnilです。
//
// コマンドが開始されるが正常に完了しない場合、エラーは
// [*ExitError] 型です。他の状況では他のエラータイプが返される可能性があります。
//
// 呼び出し元のgoroutineが [runtime.LockOSThread] でオペレーティングシステムのスレッドをロックし、
// 継承可能なOSレベルのスレッド状態（例えば、LinuxやPlan 9の名前空間）を変更した場合、
// 新しいプロセスは呼び出し元のスレッド状態を継承します。
func (c *Cmd) Run() error

// Startは指定されたコマンドを開始しますが、その完了を待ちません。
//
// Startが成功すると、c.Processフィールドが設定されます。
//
// Startの成功した呼び出しの後、関連するシステムリソースを解放するために
// [Cmd.Wait] メソッドを呼び出す必要があります。
func (c *Cmd) Start() error

// ExitErrorは、コマンドによる成功しない終了を報告します。
type ExitError struct {
	*os.ProcessState

	// Stderrは、標準エラーが他の方法で収集されていない場合、
	// Cmd.Outputメソッドからの標準エラー出力の一部を保持します。
	//
	// エラー出力が長い場合、Stderrは出力のプレフィックスと
	// サフィックスのみを含む可能性があり、中間部分は省略された
	// バイト数に関するテキストに置き換えられます。
	//
	// Stderrはデバッグ用に提供され、エラーメッセージに含めるためです。
	// 他のニーズを持つユーザーは、必要に応じてCmd.Stderrをリダイレクトしてください。
	Stderr []byte
}

func (e *ExitError) Error() string

// Waitは、コマンドが終了するのを待ち、stdinへのコピーまたは
// stdoutまたはstderrからのコピーが完了するのを待ちます。
//
// コマンドは [Cmd.Start] によって開始されていなければなりません。
//
// 返されるエラーは、コマンドが実行され、stdin、stdout、stderrのコピーに問題がなく、
// ゼロの終了ステータスで終了した場合にはnilです。
//
// コマンドが実行に失敗するか、正常に完了しない場合、
// エラーは [*ExitError] 型です。I/O問題に対しては他のエラータイプが
// 返される可能性があります。
//
// c.Stdin、c.Stdout、c.Stderrのいずれかが [*os.File] でない場合、
// Waitは、プロセスへのまたはプロセスからの対応するI/Oループのコピーが
// 完了するのを待ちます。
//
// Waitは、Cmdに関連付けられたリソースを解放します。
func (c *Cmd) Wait() error

// Outputはコマンドを実行し、その標準出力を返します。
// 返されるエラーは通常 [*ExitError] 型です。
// c.Stderrがnilで、返されたエラーが [*ExitError] 型の場合、Outputは返されたエラーのStderrフィールドを埋めます。
func (c *Cmd) Output() ([]byte, error)

// CombinedOutputはコマンドを実行し、その標準出力と標準エラーを結合したものを返します。
func (c *Cmd) CombinedOutput() ([]byte, error)

// StdinPipeは、コマンドが開始されたときにコマンドの標準入力に接続されるパイプを返します。
// パイプは、[Cmd.Wait] がコマンドの終了を確認した後、自動的に閉じられます。
// 呼び出し元は、パイプを早く閉じるためにCloseを呼び出すだけでよいです。
// 例えば、実行されるコマンドが標準入力が閉じるまで終了しない場合、呼び出し元はパイプを閉じる必要があります。
func (c *Cmd) StdinPipe() (io.WriteCloser, error)

// StdoutPipeは、コマンドが開始されたときにコマンドの標準出力に接続されるパイプを返します。
//
// [Cmd.Wait] は、コマンドの終了を確認した後にパイプを閉じるため、
// ほとんどの呼び出し元は自分でパイプを閉じる必要はありません。
// したがって、パイプからのすべての読み取りが完了する前にWaitを呼び出すことは誤りです。
// 同様の理由で、StdoutPipeを使用しているときに [Cmd.Run] を呼び出すことも誤りです。
// 一般的な使用法については、例を参照してください。
func (c *Cmd) StdoutPipe() (io.ReadCloser, error)

// StderrPipeは、コマンドが開始されたときにコマンドの標準エラーに接続されるパイプを返します。
//
// [Cmd.Wait] は、コマンドの終了を確認した後にパイプを閉じるため、
// ほとんどの呼び出し元は自分でパイプを閉じる必要はありません。
// したがって、パイプからのすべての読み取りが完了する前にWaitを呼び出すことは誤りです。
// 同様の理由で、StderrPipeを使用しているときに [Cmd.Run] を呼び出すことも誤りです。
// 一般的な使用法については、例を参照してください。
func (c *Cmd) StderrPipe() (io.ReadCloser, error)

// Environは、現在設定されている状態でコマンドが実行される環境のコピーを返します。
func (c *Cmd) Environ() []string

// ErrDotは、パスの検索が「.」がパスに含まれているために、
// 現在のディレクトリ内の実行可能ファイルに解決したことを示します。
// これは暗黙的または明示的に行われます。詳細はパッケージのドキュメンテーションを参照してください。
//
// このパッケージの関数はErrDotを直接返さないことに注意してください。
// コードはerr == ErrDotではなく、errors.Is(err, ErrDot)を使用して、
// 返されたエラーerrがこの条件によるものかどうかをテストする必要があります。
var ErrDot = errors.New("cannot run executable found relative to current directory")
