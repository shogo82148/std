// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// contextパッケージは、期限、キャンセルシグナル、および他のAPI境界やプロセス間を超えたリクエストスコープの値を伝達するContext型を定義します。
//
<<<<<<< HEAD
// サーバーへの入力リクエストは [Context] を作成し、サーバーへの出力呼び出しは [Context] を受け入れる必要があります。
// それらの間の関数呼び出しのチェーンは、Contextを伝播させ、[WithCancel]、[WithDeadline]、[WithTimeout]、または [WithValue] を使用して作成された派生Contextで置き換えることができます。
// Contextがキャンセルされると、それから派生したすべてのContextもキャンセルされます。
//
// [WithCancel]、[WithDeadline]、および [WithTimeout] 関数は、Context（親）を取得し、派生Context（子）と [CancelFunc] を返します。
// CancelFuncを呼び出すと、子とその子がキャンセルされ、親の子への参照が削除され、関連するタイマーが停止します。
// CancelFuncを呼び出さないと、子とその子は親がキャンセルされるか、タイマーが発火するまでリークします。
// go vetツールは、CancelFuncがすべての制御フローパスで使用されていることを確認します。
//
// [WithCancelCause] 関数は [CancelCauseFunc] を返し、エラーを受け取り、キャンセルの原因として記録します。
// キャンセルされたコンテキストまたはその子のいずれかで [Cause] を呼び出すと、原因が取得されます。
// 原因が指定されていない場合、 Cause(ctx) は ctx.Err() と同じ値を返します。
=======
// Incoming requests to a server should create a [Context], and outgoing
// calls to servers should accept a Context. The chain of function
// calls between them must propagate the Context, optionally replacing
// it with a derived Context created using [WithCancel], [WithDeadline],
// [WithTimeout], or [WithValue].
//
// A Context may be canceled to indicate that work done on its behalf should stop.
// A Context with a deadline is canceled after the deadline passes.
// When a Context is canceled, all Contexts derived from it are also canceled.
//
// The [WithCancel], [WithDeadline], and [WithTimeout] functions take a
// Context (the parent) and return a derived Context (the child) and a
// [CancelFunc]. Calling the CancelFunc directly cancels the child and its
// children, removes the parent's reference to the child, and stops
// any associated timers. Failing to call the CancelFunc leaks the
// child and its children until the parent is canceled. The go vet tool
// checks that CancelFuncs are used on all control-flow paths.
//
// The [WithCancelCause], [WithDeadlineCause], and [WithTimeoutCause] functions
// return a [CancelCauseFunc], which takes an error and records it as
// the cancellation cause. Calling [Cause] on the canceled context
// or any of its children retrieves the cause. If no cause is specified,
// Cause(ctx) returns the same value as ctx.Err().
>>>>>>> upstream/release-branch.go1.25
//
// Contextを使用するプログラムは、これらのルールに従う必要があります。
// これにより、パッケージ間でインターフェースを一貫させ、静的解析ツールがコンテキストの伝播をチェックできるようになります。
//
<<<<<<< HEAD
// 構造体型の内部にContextを格納しないでください。
// 代わりに、それが必要な各関数に明示的にContextを渡してください。
// 通常、最初のパラメーターにctxという名前を付けます。
=======
// Do not store Contexts inside a struct type; instead, pass a Context
// explicitly to each function that needs it. This is discussed further in
// https://go.dev/blog/context-and-structs. The Context should be the first
// parameter, typically named ctx:
>>>>>>> upstream/release-branch.go1.25
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
<<<<<<< HEAD
// サーバーでContextを使用する例のコードについては、https://blog.golang.org/context を参照してください。
=======
// See https://go.dev/blog/context for example code for a server that uses
// Contexts.
>>>>>>> upstream/release-branch.go1.25
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

<<<<<<< HEAD
// Canceled コンテキストがキャンセルされた場合に [Context.Err] が返すエラーです。
var Canceled = errors.New("context canceled")

type deadlineExceededError struct{}

func (deadlineExceededError) Error() string { return "context deadline exceeded" }

// DeadlineExceeded コンテキストの期限が切れた場合に [Context.Err] が返すエラーです。
=======
// Canceled is the error returned by [Context.Err] when the context is canceled
// for some reason other than its deadline passing.
var Canceled = errors.New("context canceled")

// DeadlineExceeded is the error returned by [Context.Err] when the context is canceled
// due to its deadline passing.
>>>>>>> upstream/release-branch.go1.25
var DeadlineExceeded error = deadlineExceededError{}

// Backgroundは、非nilで空の [Context] を返します。キャンセルされることはなく、値も期限もありません。
// 通常、main関数、初期化、テスト、および着信リクエストのトップレベルContextとして使用されます。
func Background() Context

// TODO 非nilで空の [Context] を返します。
// コードがどの [Context] を使用するか不明である場合や、まだ [Context] パラメータを受け入れるように拡張されていない
// （周囲の関数がまだ [Context] を受け入れるように拡張されていない）場合に、コードは [context.TODO] を使用する必要があります。
func TODO() Context

// CancelFunc 操作がその作業を中止するように指示します。
// CancelFuncは、作業が停止するのを待ちません。
// CancelFuncは、複数のゴルーチンから同時に呼び出すことができます。
// 最初の呼び出しの後、CancelFuncへの後続の呼び出しは何もしません。
type CancelFunc func()

<<<<<<< HEAD
// WithCancel 新しいDoneチャネルを持つ親のコピーを返します。
// 返されたコンテキストのDoneチャネルは、返されたキャンセル関数が呼び出されるか、
// または親のコンテキストのDoneチャネルが閉じられたとき、より早く閉じられます。
=======
// WithCancel returns a derived context that points to the parent context
// but has a new Done channel. The returned context's Done channel is closed
// when the returned cancel function is called or when the parent context's
// Done channel is closed, whichever happens first.
>>>>>>> upstream/release-branch.go1.25
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

<<<<<<< HEAD
// AfterFunc ctxが完了（キャンセルまたはタイムアウト）した後、fを独自のゴルーチンで呼び出すように設定します。
// もしctxが既に完了している場合、AfterFuncは独自のゴルーチンで直ちにfを呼び出します。
=======
// AfterFunc arranges to call f in its own goroutine after ctx is canceled.
// If ctx is already canceled, AfterFunc calls f immediately in its own goroutine.
>>>>>>> upstream/release-branch.go1.25
//
// ContextでのAfterFuncの複数回の呼び出しは独立して動作し、1つが他を置き換えることはありません。
//
<<<<<<< HEAD
// 返されたstop関数を呼び出すと、ctxとfの関連付けが停止します。
// 呼び出しがfの実行を停止した場合、trueを返します。
// stopがfalseを返す場合、
// コンテキストが完了し、fが独自のゴルーチンで開始されたか、
// またはfが既に停止されています。
// stop関数は、fが完了するのを待ってから戻りません。
// 呼び出し元がfが完了したかどうかを知る必要がある場合、
// 明示的にfと調整する必要があります。
=======
// Calling the returned stop function stops the association of ctx with f.
// It returns true if the call stopped f from being run.
// If stop returns false,
// either the context is canceled and f has been started in its own goroutine;
// or f was already stopped.
// The stop function does not wait for f to complete before returning.
// If the caller needs to know whether f is completed,
// it must coordinate with f explicitly.
>>>>>>> upstream/release-branch.go1.25
//
// ctxに「AfterFunc(func()) func() bool」メソッドがある場合、
// AfterFuncはそれを使用して呼び出しをスケジュールします。
func AfterFunc(ctx Context, f func()) (stop func() bool)

<<<<<<< HEAD
// WithoutCancelは、親がキャンセルされたときにキャンセルされない親のコピーを返します。
// 返されたコンテキストは、DeadlineやErrを返さず、Doneチャネルはnilです。
// 返されたコンテキストで [Cause] を呼び出すとnilが返されます。
func WithoutCancel(parent Context) Context

// WithDeadlineは、親の期限をdよりも遅くならないように調整した親のコピーを返します。
// 親の期限がすでにdよりも早い場合、 WithDeadline(parent, d) は親と意味的に等価です。
// 返された [Context.Done] チャネルは、期限が切れたとき、返されたキャンセル関数が呼び出されたとき、または親のコンテキストのDoneチャネルが閉じられたときのいずれかが最初に発生したときに閉じられます。
=======
// WithoutCancel returns a derived context that points to the parent context
// and is not canceled when parent is canceled.
// The returned context returns no Deadline or Err, and its Done channel is nil.
// Calling [Cause] on the returned context returns nil.
func WithoutCancel(parent Context) Context

// WithDeadline returns a derived context that points to the parent context
// but has the deadline adjusted to be no later than d. If the parent's
// deadline is already earlier than d, WithDeadline(parent, d) is semantically
// equivalent to parent. The returned [Context.Done] channel is closed when
// the deadline expires, when the returned cancel function is called,
// or when the parent context's Done channel is closed, whichever happens first.
>>>>>>> upstream/release-branch.go1.25
//
// このコンテキストをキャンセルすると、それに関連するリソースが解放されるため、コードはこの [Context] で実行される操作が完了したらすぐにcancelを呼び出す必要があります。
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)

// WithDeadlineCause [WithDeadline] と同様に動作しますが、期限が切れたときに返された [Context] の原因も設定します。
// 返された [CancelFunc] は原因を設定しません。
func WithDeadlineCause(parent Context, d time.Time, cause error) (Context, CancelFunc)

// WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
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

<<<<<<< HEAD
// WithValueは、キーに関連付けられた値がvalである親のコピーを返します。
=======
// WithValue returns a derived context that points to the parent Context.
// In the derived context, the value associated with key is val.
>>>>>>> upstream/release-branch.go1.25
//
// コンテキストの値は、プロセスやAPIを超えて転送されるリクエストスコープのデータにのみ使用し、関数にオプションのパラメータを渡すために使用しないでください。
//
// 提供されたキーは比較可能である必要があり、衝突を避けるためにstringまたは他の組み込み型であってはなりません。
// WithValueを使用するユーザーは、キーのために独自の型を定義する必要があります。
// interface{} に代入するときのアロケーションを避けるために、コンテキストキーは通常、具体的な型 struct{} を持ちます。
// 代替案として、エクスポートされたコンテキストキー変数の静的型はポインタまたはインターフェースである必要があります。
func WithValue(parent Context, key, val any) Context
