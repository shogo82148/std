// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// DistpackはGoの配布用のtgzファイルとzipファイルを作成します。
// これはGOROOT/pkg/distpackに書き込みます：
//
//   - 現在のGOOSとGOARCH向けのバイナリ配布（tgzまたはzip）
//   - GOOS/GOARCHに依存しないソース配布
//   - ゴーグルコマンドで使用されるようにダウンロードするためのモジュールmod、info、zipファイル
//
// Distpackは通常、make.bashの-distpackフラグによって呼び出されます。
// goos/goarch向けのクロスコンパイル配布は次のようにしてビルドできます：
//
//   GOOS=goos GOARCH=goarch ./make.bash -distpack
//
// モジュールのダウンロードがgoコマンドで使用可能であるかをテストするには：
//
//   ./make.bash -distpack
//   mkdir -p /tmp/goproxy/golang.org/toolchain/
//   ln -sf $(pwd)/../pkg/distpack /tmp/goproxy/golang.org/toolchain/@v
//   GOPROXY=file:///tmp/goproxy GOTOOLCHAIN=$(sed 1q ../VERSION) gotip version
//
// gotipは、リリースされた古いGoのバージョンで置き換えることができます。
// make.bashがビルドしたバージョンであるため、それをスキップします。
package main
