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
// If any goroutine calls Lock while the lock is already held by
// one or more readers, concurrent calls to RLock will block until
// the writer has acquired (and released) the lock, to ensure that
// the lock eventually becomes available to the writer.
// Note that this prohibits recursive read-locking.
=======
// ゴルーチンが読み込みのためにRWMutexを保持しており、他のゴルーチンが
// Lockを呼び出す可能性がある場合、初期の読み込みロックが解放されるまで、
// 他のゴルーチンは読み込みロックを獲得できることは期待できません。
// 特に、再帰的な読み込みロックは禁止されています。これは、ロックが最終的に利用可能になるようにするためです。
// ブロックされたLock呼び出しは、新しい読み込み者がロックを獲得するのを排除します。
>>>>>>> release-branch.go1.21
//
// Goのメモリモデルの用語では、Unlockのn回目の呼び出しは、
// 任意のn < mに対して、m回目のLockの呼び出しの前に同期化します（Mutexと同様）。
// また、RLockの呼び出しに対して、n回目のUnlockの呼び出しがあるために
// n+1回目のLockの呼び出しの前に同期化されます。
type RWMutex struct {
	w           Mutex
	writerSem   uint32
	readerSem   uint32
	readerCount atomic.Int32
	readerWait  atomic.Int32
}

// RLockはrwの読み取りのためにロックします。
//
// 再帰的な読み取りのために使用すべきではありません。ブロックされたLock呼び出しは、
// 新しい読み取り者がロックを取得することを排除します。RWMutex型のドキュメントを参照してください。
func (rw *RWMutex) RLock()

// TryRLockはrwを読み取りロックしようとし、成功したかどうかを報告します。
//
// TryRLockの正しい使用方法は存在しますが、それらは稀であり、
// TryRLockの使用はしばしばミューテックスの特定の使用法におけるより深刻な問題の兆候です。
func (rw *RWMutex) TryRLock() bool

// RUnlockは1回のRLock呼び出しを元に戻します。
// 他の同時読み取りプロセスには影響しません。
// RUnlockが呼び出される時にrwが読み取りロックされていない場合、ランタイムエラーが発生します。
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
// Mutexと同様に、ロックされたRWMutexは特定のゴルーチンに関連付けられていません。あるゴルーチンがRWMutexをRLock（Lock）し、別のゴルーチンがRUnlock（Unlock）するようにすることができます。
func (rw *RWMutex) Unlock()

// RLockerはLockerインターフェースを返します。このインターフェースは、rw.RLockとrw.RUnlockを呼び出してLockとUnlockメソッドを実装します。
func (rw *RWMutex) RLocker() Locker
