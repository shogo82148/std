// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !go1.19
// +build !go1.19

//go:buildは、Goのビルドシステムで条件付きコンパイルを制御するためのディレクティブです。
// +buildは、ビルド制御行として認識され、その下に続くビルドタグが評価されます。

package main
