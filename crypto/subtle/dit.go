// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package subtle

// WithDataIndependentTimingは、fを実行する前に、特定の命令のタイミングが
// 入力に依存しないことを保証するアーキテクチャ固有の機能を有効にします。
// fが戻るとこれらの機能を無効にします。
//
// fによって生成されたゴルーチンも、その生存期間中、およびその子孫のゴルーチンも、
// データ非依存タイミングが有効になります。
//
// f内、またはfによって生成されたゴルーチンからcgo経由で呼び出されたCコードも、
// 呼び出しの間はデータ非依存タイミングが有効になります。Cコードがデータ非依存タイミングを
// 無効にした場合、Goに戻る際に再度有効になります。
//
// fまたはその他の場所からcgo経由で呼び出されたCコードがデータ非依存タイミングを
// 有効または無効にした場合、Goを呼び出してもその呼び出しの間はその状態が維持されます。
//
// WithDataIndependentTimingは、fが定数時間操作を利用するように書かれている場合にのみ
// 使用するべきです。WithDataIndependentTimingは可変時間のコードを定数時間にするものではありません。
//
// WithDataIndependentTimingの呼び出しはネストできます。
//
// FEAT_DITを備えたArm64プロセッサーでは、WithDataIndependentTimingはPSTATE.DITを有効にします。
// https://developer.arm.com/documentation/ka005181/1-0/?lang=en を参照してください。
//
// 現在、その他のすべてのアーキテクチャでは、WithDataIndependentTimingは
// 他の副作用なしにfを即座に実行します。
//
//go:noinline
func WithDataIndependentTiming(f func())
