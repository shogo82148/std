// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Go 1.20以降では、ブートストラップツールチェーンとしてGo 1.17が必要です。
// もしcmd/distが以前のGoバージョンを使用してビルドされている場合、このファイルはビルドに含まれ、以下のようなエラーを引き起こします：
//
// % GOROOT_BOOTSTRAP=$HOME/sdk/go1.16 ./make.bash
// Building Go cmd/dist using /Users/rsc/sdk/go1.16. (go1.16 darwin/amd64)
// found packages main (build.go) and building_Go_requires_Go_1_17_13_or_later (notgo117.go) in /Users/rsc/go/src/cmd/dist
// %
//
// 状況の中では、これが最善の対応策です。
//
// GoがブートストラップにおいてGo 1.4から移行した背景については、go.dev/issue/44505を参照してください。

//go:build !go1.17
// +build !go1.17

//go:buildでは、特定のGoバージョンでのビルド制約を指定することができます。

package building_Go_requires_Go_1_17_13_or_later
