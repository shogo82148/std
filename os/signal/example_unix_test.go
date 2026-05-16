// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

package signal_test

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/os/signal"
)

// この例では、シグナル付きのコンテキストを渡して、ブロックする関数に
// シグナル受信後は処理を中止するよう伝えます。
func ExampleNotifyContext() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		log.Fatal(err)
	}

	// Unix系システムでは、キーボードでCtrl+Cを押すと、
	// 実行中のプログラムのプロセスにSIGINTシグナルが送られます。
	//
	// この例では、自分自身にSIGINTシグナルを送ることでそれを再現します。
	if err := p.Signal(os.Interrupt); err != nil {
		log.Fatal(err)
	}

	select {
	case <-neverReady:
		fmt.Println("ready")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // "context canceled" を表示します。
		stop()                 // できるだけ早くシグナル通知の受信を停止します。
	}

	// Output:
	// context canceled
}
