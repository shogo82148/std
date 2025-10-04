// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// contextパッケージは、期限、キャンセルシグナル、および他のAPI境界やプロセス間を超えたリクエストスコープの値を伝達するContext型を定義します。
//
// サーバーへの受信リクエストは [Context] を作成し、サーバーへの送信
// 呼び出しはContextを受け入れる必要があります。それらの間の関数
// 呼び出しのチェーンはContextを伝播する必要があり、オプションで
// [WithCancel]、[WithDeadline]、[WithTimeout]、または [WithValue] を使用して
// 作成された派生Contextで置き換えることができます。
//
// Contextは、その代わりに実行される作業を停止すべきことを示すためにキャンセルされる場合があります。
// 期限のあるContextは、期限が過ぎた後にキャンセルされます。
// Contextがキャンセルされると、そこから派生したすべてのContextもキャンセルされます。
//
// [WithCancel]、[WithDeadline]、および [WithTimeout] 関数は
// Context（親）を受け取り、派生Context（子）と
// [CancelFunc] を返します。CancelFuncを直接呼び出すと、子とその
// 子たちがキャンセルされ、親の子への参照が削除され、
// 関連するタイマーが停止されます。CancelFuncの呼び出しに失敗すると、
// 親がキャンセルされるまで子とその子たちがリークします。go vetツールは
// CancelFuncがすべての制御フローパスで使用されているかをチェックします。
//
// [WithCancelCause]、[WithDeadlineCause]、および [WithTimeoutCause] 関数は
// [CancelCauseFunc] を返します。これはエラーを受け取り、それを
// キャンセルの原因として記録します。キャンセルされたコンテキスト
// またはその子のいずれかで [Cause] を呼び出すと、原因が取得されます。原因が指定されていない場合、
// Cause(ctx)はctx.Err()と同じ値を返します。
//
// Contextを使用するプログラムは、これらのルールに従う必要があります。
// これにより、パッケージ間でインターフェースを一貫させ、静的解析ツールがコンテキストの伝播をチェックできるようになります。
//
// Contextを構造体型の内部に格納しないでください。代わりに、必要とする
// 各関数にContextを明示的に渡してください。これについては
// https://go.dev/blog/context-and-structs でさらに詳しく説明されています。Contextは最初の
// パラメータであるべきで、通常はctxという名前を付けます：
//
//	func DoSomething(ctx context.Context, arg Arg) error {
//		// ... use ctx ...
//	}
//
// 関数がnilの [Context] を許可していても、nilの [Context] を渡さないでください。
// どの [Context] を使用するかわからない場合は、 [context.TODO] を渡してください。
//
// コンテキストの値は、オプションのパラメータを関数に渡すためではなく、プロセスやAPIを超えるリクエストスコープのデータにのみ使用してください。
//
// 同じContextは、異なるゴルーチンで実行される関数に渡すことができます。Contextは、複数のゴルーチンによる同時使用に対して安全です。
//
// Contextを使用するサーバーのサンプルコードについては https://go.dev/blog/context を参照してください。
package context

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/time"
)

// Context は、期限、キャンセルシグナル、および他の値をAPI境界を超えて伝達します。
//
// Contextのメソッドは、複数のゴルーチンから同時に呼び出すことができます。
type Context interface {
	Deadline() (deadline time.Time, ok bool)

	Done() <-chan struct{}

	Err() error

	Value(key any) any
}

// Canceledは、期限の経過以外の理由でコンテキストがキャンセルされた場合に
// [Context.Err] によって返されるエラーです。
var Canceled = errors.New("context canceled")

// DeadlineExceededは、コンテキストの期限が切れた場合に [Context.Err] によって返されるエラーです。
var DeadlineExceeded error = deadlineExceededError{}

// Backgroundは、非nilで空の [Context] を返します。キャンセルされることはなく、値も期限もありません。
// 通常、main関数、初期化、テスト、および着信リクエストのトップレベルContextとして使用されます。
func Background() Context

// TODOは、非nilで空の [Context] を返します。
// コードがどの [Context] を使用するか不明である場合や、まだ [Context] パラメータを受け入れるように拡張されていない
// （周囲の関数がまだ [Context] を受け入れるように拡張されていない）場合に、コードは [context.TODO] を使用する必要があります。
func TODO() Context

// CancelFunc 操作がその作業を中止するように指示します。
// CancelFuncは、作業が停止するのを待ちません。
// CancelFuncは、複数のゴルーチンから同時に呼び出すことができます。
// 最初の呼び出しの後、CancelFuncへの後続の呼び出しは何もしません。
type CancelFunc func()

// WithCancelは親のコンテキストを指す派生コンテキストを返しますが、
// 新しいDoneチャネルを持ちます。返されたコンテキストのDoneチャネルは、
// 返されたキャンセル関数が呼び出されたとき、または親のコンテキストの
// Doneチャネルが閉じられたときのいずれか最初に発生したときに閉じられます。
//
// このコンテキストをキャンセルすると、それに関連するリソースが解放されるため、コードはこの [Context] で実行される操作が完了したらすぐにcancelを呼び出す必要があります。
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

// CancelCauseFunc [CancelFunc]と同様に動作しますが、キャンセルの原因を設定します。
// この原因は、キャンセルされたContextまたはその派生Contextのいずれかで [Cause] を呼び出すことで取得できます。
//
// コンテキストが既にキャンセルされている場合、 [CancelCauseFunc] は原因を設定しません。
// たとえば、childContextがparentContextから派生している場合：
//   - childContextがcause2でキャンセルされる前に、parentContextがcause1でキャンセルされた場合、
//     その後、Cause(parentContext) == Cause(childContext) == cause1
//   - parentContextがcause1でキャンセルされる前に、childContextがcause2でキャンセルされた場合、
//     その後、Cause(parentContext) == cause1 および Cause(childContext) == cause2
type CancelCauseFunc func(cause error)

// WithCancelCause [WithCancel] と同様に動作しますが、 [CancelFunc] の代わりに [CancelCauseFunc] を返します。
// エラー（「原因」と呼ばれる）を非nilで渡すと、そのエラーがctxに記録されます。
// その後、[Cause] を使用して取得できます。
// nilでキャンセルすると、原因は [Canceled] に設定されます。
//
// 使用例：
//
//	ctx、cancel := context.WithCancelCause(parent)
//	cancel(myError)
//	ctx.Err() // context.Canceledを返します
//	context.Cause(ctx) // myErrorを返します
func WithCancelCause(parent Context) (ctx Context, cancel CancelCauseFunc)

// Causeは、cがキャンセルされた理由を説明する非nilのエラーを返します。
// cまたはその親の最初のキャンセルは原因を設定します。
// そのキャンセルがCancelCauseFunc(err)の呼び出しによって行われた場合、 [Cause] はerrを返します。
// そうでない場合、Cause(c)はc.Err()と同じ値を返します。
// cがまだキャンセルされていない場合、Causeはnilを返します。
func Cause(c Context) error

// AfterFuncは、ctxがキャンセルされた後に独自のゴルーチンでfを呼び出すように手配します。
// ctxが既にキャンセルされている場合、AfterFuncは独自のゴルーチンで即座にfを呼び出します。
//
// ContextでのAfterFuncの複数回の呼び出しは独立して動作し、1つが他を置き換えることはありません。
//
// 返されたstop関数を呼び出すと、ctxとfの関連付けが停止されます。
// fの実行が停止された場合、trueを返します。
// stopがfalseを返す場合、
// コンテキストがキャンセルされてfが独自のゴルーチンで開始されているか、
// またはfが既に停止されています。
// stop関数は、fの完了を待たずに戻ります。
// 呼び出し元がfが完了したかどうかを知る必要がある場合、
// fと明示的に協調する必要があります。
//
// ctxに「AfterFunc(func()) func() bool」メソッドがある場合、
// AfterFuncはそれを使用して呼び出しをスケジュールします。
func AfterFunc(ctx Context, f func()) (stop func() bool)

// WithoutCancelは親のコンテキストを指す派生コンテキストを返し、
// 親がキャンセルされたときにキャンセルされません。
// 返されたコンテキストはDeadlineやErrを返さず、そのDoneチャネルはnilです。
// 返されたコンテキストで[Cause]を呼び出すとnilを返します。
func WithoutCancel(parent Context) Context

// WithDeadlineは親のコンテキストを指す派生コンテキストを返しますが、
// 期限はd以降にならないように調整されます。親の
// 期限が既にdより早い場合、WithDeadline(parent, d)は意味的に
// 親と同等です。返された[Context.Done]チャネルは、
// 期限が切れたとき、返されたキャンセル関数が呼び出されたとき、
// または親のコンテキストのDoneチャネルが閉じられたときのいずれか最初に発生したときに閉じられます。
//
// このコンテキストをキャンセルすると、それに関連するリソースが解放されるため、コードはこの [Context] で実行される操作が完了したらすぐにcancelを呼び出す必要があります。
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)

// WithDeadlineCause [WithDeadline] と同様に動作しますが、期限が切れたときに返された [Context] の原因も設定します。
// 返された [CancelFunc] は原因を設定しません。
func WithDeadlineCause(parent Context, d time.Time, cause error) (Context, CancelFunc)

// WithTimeoutはWithDeadline(parent, time.Now().Add(timeout))を返します。
//
// このコンテキストをキャンセルすると、それに関連するリソースが解放されるため、
// コードはこの [Context] で実行される操作が完了したらすぐにcancelを呼び出す必要があります。
//
//	func slowOperationWithTimeout(ctx context.Context) (Result, error) {
//		ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
//		defer cancel()  // slowOperationがタイムアウトが経過する前に完了した場合、リソースが解放されます
//		return slowOperation(ctx)
//	}
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

// WithTimeoutCause [WithTimeout]と同様に動作しますが、タイムアウトが切れたときに返された [Context] の原因も設定します。
// 返された [CancelFunc] は原因を設定しません。
func WithTimeoutCause(parent Context, timeout time.Duration, cause error) (Context, CancelFunc)

// WithValueは親のContextを指す派生コンテキストを返します。
// 派生コンテキストでは、keyに関連付けられた値はvalです。
//
// コンテキストの値は、プロセスやAPIを超えて転送されるリクエストスコープのデータにのみ使用し、関数にオプションのパラメータを渡すために使用しないでください。
//
// 提供されたキーは比較可能である必要があり、衝突を避けるためにstringまたは他の組み込み型であってはなりません。
// WithValueを使用するユーザーは、キーのために独自の型を定義する必要があります。
// interface{} に代入するときのアロケーションを避けるために、コンテキストキーは通常、具体的な型 struct{} を持ちます。
// 代替案として、エクスポートされたコンテキストキー変数の静的型はポインタまたはインターフェースである必要があります。
func WithValue(parent Context, key, val any) Context
