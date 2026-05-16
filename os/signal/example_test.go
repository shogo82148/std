// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package signal_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/os/signal"
)

func ExampleNotify() {
	// シグナル通知を送るためのチャネルを用意します。
	// バッファ付きチャネルを使わないと、シグナル送信時に
	// 受信の準備ができていない場合に見逃すおそれがあります。
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// シグナルを受け取るまでブロックします。
	s := <-c
	fmt.Println("Got signal:", s)
}

func ExampleNotify_allSignals() {
	// シグナル通知を送るためのチャネルを用意します。
	// バッファ付きチャネルを使わないと、シグナル送信時に
	// 受信の準備ができていない場合に見逃すおそれがあります。
	c := make(chan os.Signal, 1)

	// Notify にシグナルを渡さない場合は、
	// すべてのシグナルがチャネルへ送られます。
	signal.Notify(c)

	// いずれかのシグナルを受け取るまでブロックします。
	s := <-c
	fmt.Println("Got signal:", s)
}
