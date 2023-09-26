// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package context は、期限、キャンセルシグナル、および他のAPI境界やプロセス間を超えたリクエストスコープの値を伝達するContext型を定義します。
//
// サーバーへの入力リクエストは [Context] を作成し、サーバーへの出力呼び出しは[Context]を受け入れる必要があります。
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
//
// Contextを使用するプログラムは、これらのルールに従う必要があります。
// これにより、パッケージ間でインターフェースを一貫させ、静的解析ツールがコンテキストの伝播をチェックできるようになります。
//
// 構造体型の内部にContextを格納しないでください。
// 代わりに、それが必要な各関数に明示的にContextを渡してください。
// 通常、最初のパラメーターにctxという名前を付けます。
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
// サーバーでContextを使用する例のコードについては、https://blog.golang.org/contextを参照してください。
package context

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/time"
)

// A Context carries a deadline, a cancellation signal, and other values across
// API boundaries.
//
// Context's methods may be called by multiple goroutines simultaneously.
type Context interface {
	Deadline() (deadline time.Time, ok bool)

	Done() <-chan struct{}

	Err() error

	Value(key any) any
}

// Canceled is the error returned by [Context.Err] when the context is canceled.
var Canceled = errors.New("context canceled")

// DeadlineExceeded is the error returned by [Context.Err] when the context's
// deadline passes.
var DeadlineExceeded error = deadlineExceededError{}

// An emptyCtx is never canceled, has no values, and has no deadline.
// It is the common base of backgroundCtx and todoCtx.

// Background returns a non-nil, empty [Context]. It is never canceled, has no
// values, and has no deadline. It is typically used by the main function,
// initialization, and tests, and as the top-level Context for incoming
// requests.
func Background() Context

// TODO returns a non-nil, empty [Context]. Code should use context.TODO when
// it's unclear which Context to use or it is not yet available (because the
// surrounding function has not yet been extended to accept a Context
// parameter).
func TODO() Context

// A CancelFunc tells an operation to abandon its work.
// A CancelFunc does not wait for the work to stop.
// A CancelFunc may be called by multiple goroutines simultaneously.
// After the first call, subsequent calls to a CancelFunc do nothing.
type CancelFunc func()

// WithCancel returns a copy of parent with a new Done channel. The returned
// context's Done channel is closed when the returned cancel function is called
// or when the parent context's Done channel is closed, whichever happens first.
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this Context complete.
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

// A CancelCauseFunc behaves like a [CancelFunc] but additionally sets the cancellation cause.
// This cause can be retrieved by calling [Cause] on the canceled Context or on
// any of its derived Contexts.
//
// If the context has already been canceled, CancelCauseFunc does not set the cause.
// For example, if childContext is derived from parentContext:
//   - if parentContext is canceled with cause1 before childContext is canceled with cause2,
//     then Cause(parentContext) == Cause(childContext) == cause1
//   - if childContext is canceled with cause2 before parentContext is canceled with cause1,
//     then Cause(parentContext) == cause1 and Cause(childContext) == cause2
type CancelCauseFunc func(cause error)

// WithCancelCause behaves like [WithCancel] but returns a [CancelCauseFunc] instead of a [CancelFunc].
// Calling cancel with a non-nil error (the "cause") records that error in ctx;
// it can then be retrieved using Cause(ctx).
// Calling cancel with nil sets the cause to Canceled.
//
// Example use:
//
//	ctx, cancel := context.WithCancelCause(parent)
//	cancel(myError)
//	ctx.Err() // returns context.Canceled
//	context.Cause(ctx) // returns myError
func WithCancelCause(parent Context) (ctx Context, cancel CancelCauseFunc)

// Cause returns a non-nil error explaining why c was canceled.
// The first cancellation of c or one of its parents sets the cause.
// If that cancellation happened via a call to CancelCauseFunc(err),
// then [Cause] returns err.
// Otherwise Cause(c) returns the same value as c.Err().
// Cause returns nil if c has not been canceled yet.
func Cause(c Context) error

// AfterFunc arranges to call f in its own goroutine after ctx is done
// (cancelled or timed out).
// If ctx is already done, AfterFunc calls f immediately in its own goroutine.
//
// Multiple calls to AfterFunc on a context operate independently;
// one does not replace another.
//
// Calling the returned stop function stops the association of ctx with f.
// It returns true if the call stopped f from being run.
// If stop returns false,
// either the context is done and f has been started in its own goroutine;
// or f was already stopped.
// The stop function does not wait for f to complete before returning.
// If the caller needs to know whether f is completed,
// it must coordinate with f explicitly.
//
// If ctx has a "AfterFunc(func()) func() bool" method,
// AfterFunc will use it to schedule the call.
func AfterFunc(ctx Context, f func()) (stop func() bool)

// A stopCtx is used as the parent context of a cancelCtx when
// an AfterFunc has been registered with the parent.
// It holds the stop function used to unregister the AfterFunc.

// goroutines counts the number of goroutines ever created; for testing.

// &cancelCtxKey is the key that a cancelCtx returns itself for.

// A canceler is a context type that can be canceled directly. The
// implementations are *cancelCtx and *timerCtx.

// closedchan is a reusable closed channel.

// A cancelCtx can be canceled. When canceled, it also cancels any children
// that implement canceler.

// WithoutCancel returns a copy of parent that is not canceled when parent is canceled.
// The returned context returns no Deadline or Err, and its Done channel is nil.
// Calling [Cause] on the returned context returns nil.
func WithoutCancel(parent Context) Context

// WithDeadline returns a copy of the parent context with the deadline adjusted
// to be no later than d. If the parent's deadline is already earlier than d,
// WithDeadline(parent, d) is semantically equivalent to parent. The returned
// [Context.Done] channel is closed when the deadline expires, when the returned
// cancel function is called, or when the parent context's Done channel is
// closed, whichever happens first.
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this [Context] complete.
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)

// WithDeadlineCause behaves like [WithDeadline] but also sets the cause of the
// returned Context when the deadline is exceeded. The returned [CancelFunc] does
// not set the cause.
func WithDeadlineCause(parent Context, d time.Time, cause error) (Context, CancelFunc)

// A timerCtx carries a timer and a deadline. It embeds a cancelCtx to
// implement Done and Err. It implements cancel by stopping its timer then
// delegating to cancelCtx.cancel.

// WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this [Context] complete:
//
//	func slowOperationWithTimeout(ctx context.Context) (Result, error) {
//		ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
//		defer cancel()  // releases resources if slowOperation completes before timeout elapses
//		return slowOperation(ctx)
//	}
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

// WithTimeoutCause behaves like [WithTimeout] but also sets the cause of the
// returned Context when the timeout expires. The returned [CancelFunc] does
// not set the cause.
func WithTimeoutCause(parent Context, timeout time.Duration, cause error) (Context, CancelFunc)

// WithValue returns a copy of parent in which the value associated with key is
// val.
//
// Use context Values only for request-scoped data that transits processes and
// APIs, not for passing optional parameters to functions.
//
// The provided key must be comparable and should not be of type
// string or any other built-in type to avoid collisions between
// packages using context. Users of WithValue should define their own
// types for keys. To avoid allocating when assigning to an
// interface{}, context keys often have concrete type
// struct{}. Alternatively, exported context key variables' static
// type should be a pointer or interface.
func WithValue(parent Context, key, val any) Context

// A valueCtx carries a key-value pair. It implements Value for that key and
// delegates all other calls to the embedded Context.
