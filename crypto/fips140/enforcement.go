// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fips140

// WithoutEnforcementは、fの実行中に厳格なFIPS 140-3の強制を無効にします。
// 厳格な強制が有効になっていない状態（GODEBUG=fips140=onlyが設定されていないか、
// すでにWithoutEnforcementの呼び出し内にある場合）でWithoutEnforcementを
// 呼び出しても何もしません。
//
// WithoutEnforcementは、fの実行中に生成されたすべてのゴルーチンに継承されます。
//
// これは強制を無効にするため、厳密にスコープが限定された関数に
// 慎重に適用するべきです。
func WithoutEnforcement(f func())

// Enforcedは、厳格なFIPS 140-3の強制が有効かどうかを示します。
// 厳格な強制は、プログラムがGODEBUG=fips140=onlyで実行され、
// [WithoutEnforcement]の呼び出しによって強制が無効化されていない場合に有効になります。
func Enforced() bool
