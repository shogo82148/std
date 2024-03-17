// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// パッケージsyncは相互排他ロックなどの基本的な同期プリミティブを提供します。OnceとWaitGroup以外の型は、低レベルのライブラリルーチンでの使用を意図しています。より高レベルな同期はチャネルと通信を介して行う方が良いです。
=======
// Package sync provides basic synchronization primitives such as mutual
// exclusion locks. Other than the [Once] and [WaitGroup] types, most are intended
// for use by low-level library routines. Higher-level synchronization is
// better done via channels and communication.
>>>>>>> upstream/master
//
// このパッケージで定義された型を含む値は、コピーしないでください。
package sync

// Mutexは相互排他ロックです。
// Mutexのゼロ値はロックされていないMutexです。
//
// Mutexは最初の使用後にコピーしてはいけません。
//
<<<<<<< HEAD
// Goのメモリモデルの用語では、Unlockのn回目の呼び出しはm回目のLockの前に同期されます（n < m）。
// TryLockの成功した呼び出しはLockの呼び出しと同等です。
// TryLockの失敗した呼び出しはどのような「同期前の」関係も確立しません。
=======
// In the terminology of the Go memory model,
// the n'th call to [Mutex.Unlock] “synchronizes before” the m'th call to [Mutex.Lock]
// for any n < m.
// A successful call to [Mutex.TryLock] is equivalent to a call to Lock.
// A failed call to TryLock does not establish any “synchronizes before”
// relation at all.
>>>>>>> upstream/master
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
<<<<<<< HEAD
// ロックされたMutexは特定のゴルーチンに関連付けられていません。
// あるゴルーチンがMutexをロックし、別のゴルーチンがそれをアンロックするようにすることも許可されています。
=======
// A locked [Mutex] is not associated with a particular goroutine.
// It is allowed for one goroutine to lock a Mutex and then
// arrange for another goroutine to unlock it.
>>>>>>> upstream/master
func (m *Mutex) Unlock()
