// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package context_test

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/time"
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
	waitOnCond := func(ctx context.Context, cond *sync.Cond, conditionMet func() bool) error {
		stopf := context.AfterFunc(ctx, func() {
			// We need to acquire cond.L here to be sure that the Broadcast
			// below won't occur before the call to Wait, which would result
			// in a missed signal (and deadlock).
			cond.L.Lock()
			defer cond.L.Unlock()

			// If multiple goroutines are waiting on cond simultaneously,
			// we need to make sure we wake up exactly this one.
			// That means that we need to Broadcast to all of the goroutines,
			// which will wake them all up.
			//
			// If there are N concurrent calls to waitOnCond, each of the goroutines
			// will spuriously wake up O(N) other goroutines that aren't ready yet,
			// so this will cause the overall CPU cost to be O(N²).
			cond.Broadcast()
		})
		defer stopf()

		// Since the wakeups are using Broadcast instead of Signal, this call to
		// Wait may unblock due to some other goroutine's context being canceled,
		// so to be sure that ctx is actually canceled we need to check it in a loop.
		for !conditionMet() {
			cond.Wait()
			if ctx.Err() != nil {
				return ctx.Err()
			}
		}

		return nil
	}

	cond := sync.NewCond(new(sync.Mutex))

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
			defer cancel()

			cond.L.Lock()
			defer cond.L.Unlock()

			err := waitOnCond(ctx, cond, func() bool { return false })
			fmt.Println(err)
		}()
	}
	wg.Wait()

	// Output:
	// context deadline exceeded
	// context deadline exceeded
	// context deadline exceeded
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

	listener, err := net.Listen("tcp", "localhost:0")
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
