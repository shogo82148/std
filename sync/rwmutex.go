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
// もしロックが既に1つ以上のリーダーによって保持されている場合、
// いずれかのゴルーチンがLockを呼び出すと、RLockへの並行呼び出しがライターがロックを取得（および解放）するまでブロックされます。
// これにより、ロックが最終的にライターに利用可能になることが保証されます。
// なお、これにより再帰的な読み込みロックが禁止されます。
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
