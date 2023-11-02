// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package buildinfoは、Goバイナリに埋め込まれた情報にアクセスするための機能を提供します。
// これには、Goツールチェーンのバージョン、および使用されたモジュールのセット（モジュールモードでビルドされたバイナリの場合）が含まれます。
//
// ビルド情報は、現在実行中のバイナリでruntime/debug.ReadBuildInfoを使用して利用できます。
package buildinfo

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/runtime/debug"
)

// ビルド情報のための型エイリアスです。
// ここに型を移動することはできません。なぜなら、
// runtime/debugがこのパッケージをインポートする必要があるため、
// 依存関係が大きくなるためです。
type BuildInfo = debug.BuildInfo

// ReadFileは、指定されたパスのGoバイナリファイルに埋め込まれたビルド情報を返します。
// ほとんどの情報は、モジュールサポートでビルドされたバイナリでのみ利用可能です。
func ReadFile(name string) (info *BuildInfo, err error)

// Readは、指定されたReaderAtを介してアクセスされるGoバイナリファイルに埋め込まれたビルド情報を返します。
// ほとんどの情報は、モジュールサポートでビルドされたバイナリでのみ利用可能です。
func Read(r io.ReaderAt) (*BuildInfo, error)
