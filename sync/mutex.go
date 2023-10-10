// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージsyncは相互排他ロックなどの基本的な同期プリミティブを提供します。OnceとWaitGroup以外の型は、低レベルのライブラリルーチンでの使用を意図しています。より高レベルな同期はチャネルと通信を介して行う方が良いです。
//
// このパッケージで定義された型を含む値は、コピーしないでください。
package sync

// Mutexは相互排他ロックです。
// Mutexのゼロ値はロックされていないMutexです。
//
// Mutexは最初の使用後にコピーしてはいけません。
//
// Goのメモリモデルの用語では、Unlockのn回目の呼び出しはm回目のLockの前に同期されます（n < m）。
// TryLockの成功した呼び出しはLockの呼び出しと同等です。
// TryLockの失敗した呼び出しはどのような「同期前の」関係も確立しません。
type Mutex struct {
	state int32
	sema  uint32
}

// Lockerはロックおよびアンロックできるオブジェクトを表します。
type Locker interface {
	Lock()
	Unlock()
}

// Lock は m をロックします。
// もし既にロックが使用中である場合、呼び出し元のゴルーチンは
// ミューテックスが利用可能になるまでブロックされます。
func (m *Mutex) Lock()

// TryLock はmをロックしようとし、成功したかどうかを報告します。
//
// TryLockの正しい使用方法が存在することに注意してくださいが、それらはまれであり、
// TryLockの使用はしばしばミューテックスの特定の使用において深刻な問題の兆候です。
func (m *Mutex) TryLock() bool

// Unlockはmをアンロックします。
// mがUnlockされる前にロックされていない場合、ランタイムエラーになります。
//
// ロックされたMutexは特定のゴルーチンに関連付けられていません。
// あるゴルーチンがMutexをロックし、別のゴルーチンがそれをアンロックするようにすることも許可されています。
func (m *Mutex) Unlock()
