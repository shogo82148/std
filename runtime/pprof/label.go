// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pprof

import (
	"github.com/shogo82148/std/context"
)

// LabelSetはラベルのセットです。
type LabelSet struct {
	list []label
}

<<<<<<< HEAD
// WithLabelsは指定されたラベルが追加された新しいcontext.Contextを返します。
// ラベルは同じキーを持つ以前のラベルを上書きします。
func WithLabels(ctx context.Context, labels LabelSet) context.Context

// Labelsは、キーと値のペアを表す文字列の偶数個を受け取り、それらを含むLabelSetを作成します。
// ラベルは、同じキーを持つ以前のラベルを上書きします。
// 現在、CPUプロファイルとゴルーチンプロファイルのみがラベル情報を利用しています。
// 詳細は、https://golang.org/issue/23458を参照してください。
=======
// WithLabels returns a new [context.Context] with the given labels added.
// A label overwrites a prior label with the same key.
func WithLabels(ctx context.Context, labels LabelSet) context.Context

// Labels takes an even number of strings representing key-value pairs
// and makes a [LabelSet] containing them.
// A label overwrites a prior label with the same key.
// Currently only the CPU and goroutine profiles utilize any labels
// information.
// See https://golang.org/issue/23458 for details.
>>>>>>> upstream/master
func Labels(args ...string) LabelSet

// Labelは与えられたキーに対応するラベルの値と、そのラベルが存在するかを示すブール値をctxから返します。
func Label(ctx context.Context, key string) (string, bool)

// ForLabelsはコンテキストに設定された各ラベルを持ってfを呼び出します。
// 関数fは繰り返しを続けるためにtrueを返すか、繰り返しを早期に停止するためにfalseを返す必要があります。
func ForLabels(ctx context.Context, f func(key, value string) bool)
