// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build race

package runtime

import (
	"github.com/shogo82148/std/unsafe"
)

func RaceRead(addr unsafe.Pointer)
func RaceWrite(addr unsafe.Pointer)
func RaceReadRange(addr unsafe.Pointer, len int)
func RaceWriteRange(addr unsafe.Pointer, len int)

func RaceErrors() int

// RaceAcquire/RaceRelease/RaceReleaseMergeを使用すると、
// ゴルーチン間のhappens-before関係を確立できます。
// これによって、競争検出器に、一部の理由で見えない実際の同期について情報を提供できます
// （例：RaceDisable/RaceEnableのコード節内の同期）。
// RaceAcquireは、addrの前のRaceReleaseMergeからaddrの最後のRaceReleaseまでの
// happens-before関係を確立します。
// Cのメモリモデル（C11 §5.1.2.4、§7.17.3）において、RaceAcquireは
// atomic_load(memory_order_acquire)に相当します。
//
//go:nosplit
func RaceAcquire(addr unsafe.Pointer)

// RaceReleaseは、addr上のリリース操作を実行します。
// これは、後のaddr上のRaceAcquireと同期することができます。
//
// Cのメモリモデルの観点からは、RaceReleaseはatomic_store(memory_order_release)に相当します。
//
//go:nosplit
func RaceRelease(addr unsafe.Pointer)

// RaceReleaseMergeはRaceReleaseと似ていますが、addrの前にあるRaceReleaseまたはRaceReleaseMergeとのhappens-before関係も確立します。
//
// Cのメモリモデルにおいて、RaceReleaseMergeはatomic_exchange(memory_order_release)と等価です。
//
//go:nosplit
func RaceReleaseMerge(addr unsafe.Pointer)

// RaceDisableは現在のゴルーチンでの競争同期イベントの処理を無効にします。
// RaceEnableで処理は再有効化されます。RaceDisable/RaceEnableは入れ子にすることができます。
// 非同期イベント（メモリアクセス、関数の入退場）は引き続き競争検出器に影響します。
//
//go:nosplit
func RaceDisable()

// RaceEnableは現在のgoroutineでのレースイベントの処理を再有効化します。
//
//go:nosplit
func RaceEnable()
