// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Go 1.22以降では、ブートストラップツールチェインとしてGo 1.20が必要です。
// もしcmd/distが以前のGoバージョンを使用してビルドされた場合、このファイルはビルドに含まれ、次のようなエラーが発生します：
//
// % GOROOT_BOOTSTRAP=$HOME/sdk/go1.16 ./make.bash
// /Users/rsc/sdk/go1.16を使用してGo cmd/distをビルドしています。(go1.16 darwin/amd64)
// /Users/rsc/go/src/cmd/dist内のmain (build.go)とbuilding_Go_requires_Go_1_20_6_or_later (notgo120.go)というパッケージが見つかりました。
// %
//
// これが、現状では最善の方法です。
//
// GoがブートストラップにGo 1.4から移行した背景については、go.dev/issue/44505を参照してください。

//go:build !go1.20
//go:build !go1.20
// +build !go1.20
// +build !go1.20

package building_Go_requires_Go_1_20_6_or_later
