// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package context_test

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
)

// This example demonstrates the use of a cancelable context to prevent a
// goroutine leak. By the end of the example function, the goroutine started
// by gen will return without leaking.
func ExampleWithCancel() {
	// genは、別のゴルーチンで整数を生成し、それらを返されたチャネルに送信します。
	// genの呼び出し元は、生成された整数を消費し終わったらコンテキストをキャンセルする必要があります。
	// そうしないと、genによって開始された内部ゴルーチンがリークすることになります。
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 整数を消費し終わったらキャンセル

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

var shortDuration time.Duration = 1
var neverReady <-chan time.Time

// This example passes a context with an arbitrary deadline to tell a blocking
// function that it should abandon its work as soon as it gets to it.
func ExampleWithDeadline() {
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// ctxが期限切れになっている場合でも、そのキャンセル関数を呼び出すことは良い習慣です。
	// そうしないと、コンテキストとその親が必要以上に長く生き残る可能性があります。
	defer cancel()

	select {
	case <-neverReady:
		fmt.Println("ready")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

	// Output:
	// context deadline exceeded
}

// This example passes a context with a timeout to tell a blocking function that
// it should abandon its work after the timeout elapses.
func ExampleWithTimeout() {
	// タイムアウト付きのコンテキストを渡すことで、ブロッキング関数に、タイムアウトが経過した後に作業を中止するように指示します。
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-neverReady:
		fmt.Println("ready")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // "context deadline exceeded" を出力します
	}

	// Output:
	// context deadline exceeded
}

// This example demonstrates how a value can be passed to the context
// and also how to retrieve it if it exists.
func ExampleWithValue() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))

	// Output:
	// found value: Go
	// key not found: color
}

// This example uses AfterFunc to define a function which waits on a sync.Cond,
// stopping the wait when a context is canceled.
func ExampleAfterFunc_cond() {
	waitOnCond := func(ctx context.Context, cond *sync.Cond) error {
		stopf := context.AfterFunc(ctx, cond.Broadcast)
		defer stopf()
		cond.Wait()
		return ctx.Err()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	mu.Lock()
	err := waitOnCond(ctx, cond)
	fmt.Println(err)

	// Output:
	// context deadline exceeded
}

// This example uses AfterFunc to define a function which reads from a net.Conn,
// stopping the read when a context is canceled.
func ExampleAfterFunc_connection() {
	readFromConn := func(ctx context.Context, conn net.Conn, b []byte) (n int, err error) {
		stopc := make(chan struct{})
		stop := context.AfterFunc(ctx, func() {
			conn.SetReadDeadline(time.Now())
			close(stopc)
		})
		n, err = conn.Read(b)
		if !stop() {
			// AfterFuncが開始されました。
			// 完了するまで待ち、Connの期限をリセットします。
			<-stopc
			conn.SetReadDeadline(time.Time{})
			return n, ctx.Err()
		}
		return n, err
	}

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	conn, err := net.Dial(listener.Addr().Network(), listener.Addr().String())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	b := make([]byte, 1024)
	_, err = readFromConn(ctx, conn, b)
	fmt.Println(err)

	// Output:
	// context deadline exceeded
}

// This example uses AfterFunc to define a function which combines
// the cancellation signals of two Contexts.
func ExampleAfterFunc_merge() {
	// mergeCancelは、ctxの値を含み、ctxまたはcancelCtxのいずれかがキャンセルされたときにキャンセルされるコンテキストを返します。
	mergeCancel := func(ctx, cancelCtx context.Context) (context.Context, context.CancelFunc) {
		ctx, cancel := context.WithCancelCause(ctx)
		stop := context.AfterFunc(cancelCtx, func() {
			cancel(context.Cause(cancelCtx))
		})
		return ctx, func() {
			stop()
			cancel(context.Canceled)
		}
	}

	ctx1, cancel1 := context.WithCancelCause(context.Background())
	defer cancel1(errors.New("ctx1 canceled"))

	ctx2, cancel2 := context.WithCancelCause(context.Background())

	mergedCtx, mergedCancel := mergeCancel(ctx1, ctx2)
	defer mergedCancel()

	cancel2(errors.New("ctx2 canceled"))
	<-mergedCtx.Done()
	fmt.Println(context.Cause(mergedCtx))

	// Output:
	// ctx2 canceled
}
