// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

import (
	"github.com/shogo82148/std/sync/atomic"
)

// RWMutexは、読み込み/書き込み相互排他的なロックです。
// ロックは任意の数の読み込み者または単一の書き込み者によって保持することができます。
// RWMutexのゼロ値はロックされていないミューテックスです。
//
// RWMutexは、最初の使用後にコピーしてはいけません。
//
<<<<<<< HEAD
// もしロックが既に1つ以上のリーダーによって保持されている場合、
// いずれかのゴルーチンがLockを呼び出すと、RLockへの並行呼び出しがライターがロックを取得（および解放）するまでブロックされます。
// これにより、ロックが最終的にライターに利用可能になることが保証されます。
// なお、これにより再帰的な読み込みロックが禁止されます。
//
// Goのメモリモデルの用語では、Unlockのn回目の呼び出しは、
// 任意のn < mに対して、m回目のLockの呼び出しの前に同期化します（Mutexと同様）。
// また、RLockの呼び出しに対して、n回目のUnlockの呼び出しがあるために
// n+1回目のLockの呼び出しの前に同期化されます。
=======
// If any goroutine calls [RWMutex.Lock] while the lock is already held by
// one or more readers, concurrent calls to [RWMutex.RLock] will block until
// the writer has acquired (and released) the lock, to ensure that
// the lock eventually becomes available to the writer.
// Note that this prohibits recursive read-locking.
//
// In the terminology of the Go memory model,
// the n'th call to [RWMutex.Unlock] “synchronizes before” the m'th call to Lock
// for any n < m, just as for [Mutex].
// For any call to RLock, there exists an n such that
// the n'th call to Unlock “synchronizes before” that call to RLock,
// and the corresponding call to [RWMutex.RUnlock] “synchronizes before”
// the n+1'th call to Lock.
>>>>>>> upstream/master
type RWMutex struct {
	w           Mutex
	writerSem   uint32
	readerSem   uint32
	readerCount atomic.Int32
	readerWait  atomic.Int32
}

// RLockはrwの読み取りのためにロックします。
//
<<<<<<< HEAD
// 再帰的な読み取りのために使用すべきではありません。ブロックされたLock呼び出しは、
// 新しい読み取り者がロックを取得することを排除します。RWMutex型のドキュメントを参照してください。
=======
// It should not be used for recursive read locking; a blocked Lock
// call excludes new readers from acquiring the lock. See the
// documentation on the [RWMutex] type.
>>>>>>> upstream/master
func (rw *RWMutex) RLock()

// TryRLockはrwを読み取りロックしようとし、成功したかどうかを報告します。
//
// TryRLockの正しい使用方法は存在しますが、それらは稀であり、
// TryRLockの使用はしばしばミューテックスの特定の使用法におけるより深刻な問題の兆候です。
func (rw *RWMutex) TryRLock() bool

<<<<<<< HEAD
// RUnlockは1回のRLock呼び出しを元に戻します。
// 他の同時読み取りプロセスには影響しません。
// RUnlockが呼び出される時にrwが読み取りロックされていない場合、ランタイムエラーが発生します。
=======
// RUnlock undoes a single [RWMutex.RLock] call;
// it does not affect other simultaneous readers.
// It is a run-time error if rw is not locked for reading
// on entry to RUnlock.
>>>>>>> upstream/master
func (rw *RWMutex) RUnlock()

// Lockはrwを書き込み用にロックします。
// もし既に読み込みや書き込みのためにロックされている場合、
// Lockは利用可能になるまでブロックします。
func (rw *RWMutex) Lock()

// TryLockは、rwを書き込み用にロックしようとし、成功したかどうかを報告します。
//
// TryLockの正しい使用法は存在しますが、それらはまれであり、
// mutexの特定の使用法におけるより深刻な問題の兆候であることが多いため、
// TryLockの使用は避けるべきです。
func (rw *RWMutex) TryLock() bool

// Unlockは書き込みのためにrwをアンロックします。Unlockに入る前にrwが書き込み用にロックされていない場合、ランタイムエラーとなります。
//
<<<<<<< HEAD
// Mutexと同様に、ロックされたRWMutexは特定のゴルーチンに関連付けられていません。あるゴルーチンがRWMutexをRLock（Lock）し、別のゴルーチンがRUnlock（Unlock）するようにすることができます。
func (rw *RWMutex) Unlock()

// RLockerはLockerインターフェースを返します。このインターフェースは、rw.RLockとrw.RUnlockを呼び出してLockとUnlockメソッドを実装します。
=======
// As with Mutexes, a locked [RWMutex] is not associated with a particular
// goroutine. One goroutine may [RWMutex.RLock] ([RWMutex.Lock]) a RWMutex and then
// arrange for another goroutine to [RWMutex.RUnlock] ([RWMutex.Unlock]) it.
func (rw *RWMutex) Unlock()

// RLocker returns a [Locker] interface that implements
// the [RWMutex.Lock] and [RWMutex.Unlock] methods by calling rw.RLock and rw.RUnlock.
>>>>>>> upstream/master
func (rw *RWMutex) RLocker() Locker
