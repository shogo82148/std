// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package signal

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/os"
)

// Ignoreは、指定されたシグナルを無視するようにします。プログラムが
// それらを受信しても、何も起こりません。Ignoreは、指定された
// シグナルに対するそれ以前の [Notify] 呼び出しの効果を取り消します。
// シグナルが指定されなかった場合、すべての受信シグナルは無視されます。
func Ignore(sig ...os.Signal)

// Ignoredは、sig が現在無視されているかどうかを報告します。
func Ignored(sig os.Signal) bool

// Notifyは、受信したシグナルを c に中継するよう package signal に
// 指示します。
// シグナルが指定されなかった場合、すべての受信シグナルが c に
// 中継されます。そうでない場合は、指定されたシグナルのみが
// 中継されます。
//
// package signal は c への送信でブロックしません。呼び出し側は、
// 予想されるシグナル頻度に追随できるだけの十分なバッファを c が
// 持つようにしなければなりません。1つのシグナル値だけの通知に
// 使うチャネルなら、サイズ1のバッファで十分です。
//
// 同じチャネルに対して Notify を複数回呼び出すことができます。
// 各呼び出しで、そのチャネルに送られるシグナル集合が拡張されます。
// 集合からシグナルを取り除く唯一の方法は [Stop] を呼ぶことです。
//
// 異なるチャネルに対して、同じシグナルで Notify を複数回呼び出すことも
// できます。各チャネルは、受信したシグナルのコピーを独立に受け取ります。
func Notify(c chan<- os.Signal, sig ...os.Signal)

// Resetは、指定されたシグナルに対するそれ以前の [Notify] 呼び出しの
// 効果を取り消します。
// シグナルが指定されなかった場合、すべての signal handler が
// リセットされます。
func Reset(sig ...os.Signal)

// Stopは、受信したシグナルを c へ中継するのを package signal に
// やめさせます。
// これは、c を使ったそれ以前のすべての [Notify] 呼び出しの効果を
// 取り消します。
// Stop が戻った時点で、c がこれ以上シグナルを受け取らないことが
// 保証されます。
func Stop(c chan<- os.Signal)

// NotifyContextは、列挙されたシグナルのいずれかが到着したとき、
// 返される stop 関数が呼ばれたとき、または親コンテキストの Done
// チャネルが閉じられたときのうち、最初に起こった時点で done
// （Done チャネルが閉じられた状態）になる、親コンテキストのコピーを
// 返します。
//
// stop 関数はシグナル動作の登録を解除します。これは [signal.Reset] と
// 同様に、指定されたシグナルの既定動作を復元する場合があります。たとえば、
// [os.Interrupt] を受信したGoプログラムの既定動作は終了です。
// NotifyContext(parent, os.Interrupt) を呼ぶと、その動作は返された
// コンテキストをキャンセルする動作に変わります。返された stop 関数が
// 呼ばれるまでは、その後に受信した割り込みは既定の（終了）動作を
// 引き起こしません。
//
// シグナルによって返されたコンテキストがキャンセルされた場合、
// それに対して [context.Cause] を呼ぶと、そのシグナルを説明する
// エラーが返ります。
//
// stop 関数は関連するリソースを解放するため、この Context で実行している
// 操作が完了し、もはやシグナルをコンテキストへ振り向ける必要がなくなったら、
// できるだけ早く stop を呼ぶべきです。
func NotifyContext(parent context.Context, signals ...os.Signal) (ctx context.Context, stop context.CancelFunc)
