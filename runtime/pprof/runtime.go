// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pprof

import (
	"github.com/shogo82148/std/context"
)

// SetGoroutineLabelsは現在のゴルーチンのラベルをctxと一致させます。
// 新しいゴルーチンは、その作成元のゴルーチンのラベルを継承します。
// これは、可能な場合は代わりに使用するべきDoよりも低レベルのAPIです。
func SetGoroutineLabels(ctx context.Context)

// 親のコンテキストのコピーを使用して f を呼び出します。
// 親のラベルマップに指定されたラベルが追加されます。
// f を実行する間に生成されたゴルーチンは、拡張されたラベルセットを継承します。
// labels の各キー/値ペアは、提供された順序でラベルマップに挿入され、同じキーの以前の値を上書きします。
// 拡張されたラベルマップは、f の呼び出しの間、設定され、f の戻り値時に復元されます。
func Do(ctx context.Context, labels LabelSet, f func(context.Context))
