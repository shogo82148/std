// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// Gosched はプロセッサを譲り、他のゴルーチンが実行されるようにします。現在のゴルーチンは一時停止されませんが、実行は自動的に再開されます。
//
//go:nosplit
func Gosched()

// ブレークポイントはブレークポイントトラップを実行します。
func Breakpoint()

// LockOSThreadは呼び出し側のゴルーチンを現在のオペレーティングシステムスレッドに接続します。
// 呼び出し側のゴルーチンは常にそのスレッドで実行され、他のゴルーチンは実行されません。
// それまでのLockOSThreadへの呼び出し回数と同じ数だけ、UnlockOSThreadへの呼び出しを行うまで、呼び出し側のゴルーチン以外は実行されません。
// 呼び出し側のゴルーチンがスレッドのロックを解除せずに終了すると、スレッドは終了します。
//
// すべてのinit関数は起動時のスレッド上で実行されます。init関数からLockOSThreadを呼び出すと、main関数がそのスレッド上で呼び出されます。
//
// ゴルーチンは、スレッドごとの状態に依存するOSサービスや非Goライブラリ関数を呼び出す前に、LockOSThreadを呼び出す必要があります。
//
//go:nosplit
func LockOSThread()

// UnlockOSThreadは、以前のLockOSThread呼び出しを取り消します。
// 呼び出し元のゴルーチンのアクティブなLockOSThread呼び出し数がゼロになると、
// 呼び出し元のゴルーチンは固定されたオペレーティングシステムスレッドからの接続が解除されます。
// アクティブなLockOSThread呼び出しがない場合、これは無効操作です。
//
// UnlockOSThreadを呼び出す前に、呼び出し元は他のゴルーチンを実行するためにOSスレッドが適していることを確認する必要があります。
// 呼び出し元が他のゴルーチンに影響を与えるスレッドの状態に対して恒久的な変更を行った場合、
// この関数を呼び出さずにゴルーチンをOSスレッドにロックしたままにしておくべきです。
//
//go:nosplit
func UnlockOSThread()
