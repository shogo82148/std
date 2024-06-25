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
// もし任意のゴルーチンが [RWMutex.Lock] を呼び出し、そのロックがすでに
// 1つ以上のリーダーによって保持されている場合、並行する [RWMutex.RLock] への呼び出しは
// ライターがロックを取得（そして解放）するまでブロックされます。これにより、
// ロックが最終的にライターに利用可能になることを保証します。
// これは再帰的な読み取りロックを禁止することに注意してください。
//
<<<<<<< HEAD
// Goのメモリモデルの用語では、
// n番目の [RWMutex.Unlock] への呼び出しは、任意のn < mに対するm番目のLockへの呼び出しよりも
// 「先に同期します」。これは [Mutex] と同様です。
// RLockへの任意の呼び出しに対して、nが存在し、
// n番目のUnlockへの呼び出しはそのRLockへの呼び出しよりも「先に同期します」、
// そして対応する [RWMutex.RUnlock] への呼び出しも同様に「先に同期します」。
=======
// In the terminology of [the Go memory model],
// the n'th call to [RWMutex.Unlock] “synchronizes before” the m'th call to Lock
// for any n < m, just as for [Mutex].
// For any call to RLock, there exists an n such that
// the n'th call to Unlock “synchronizes before” that call to RLock,
// and the corresponding call to [RWMutex.RUnlock] “synchronizes before”
// the n+1'th call to Lock.
//
// [the Go memory model]: https://go.dev/ref/mem
>>>>>>> 41b4a7d0008e48dd077e189fd86911de2b36d90d
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
// 新しい読み取り者がロックを取得することを排除します。[RWMutex] 型のドキュメントを参照してください。
func (rw *RWMutex) RLock()

// TryRLockはrwを読み取りロックしようとし、成功したかどうかを報告します。
//
// TryRLockの正しい使用方法は存在しますが、それらは稀であり、
// TryRLockの使用はしばしばミューテックスの特定の使用法におけるより深刻な問題の兆候です。
func (rw *RWMutex) TryRLock() bool

// RUnlockは1回の [RWMutex.RLock] 呼び出しを元に戻します。
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
// Mutexと同様に、ロックされた [RWMutex] は特定のゴルーチンに関連付けられていません。あるゴルーチンがRWMutexを [RWMutex.RLock]（[RWMutex.Lock]）し、別のゴルーチンが [RWMutex.RUnlock]（[RWMutex.Unlock]）するようにすることができます。
func (rw *RWMutex) Unlock()

// RLockerは [Locker] インターフェースを返します。このインターフェースは、rw.RLockとrw.RUnlockを呼び出して [RWMutex.Lock] と [RWMutex.Unlock] メソッドを実装します。
func (rw *RWMutex) RLocker() Locker
