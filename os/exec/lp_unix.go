// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

package exec

import (
	"github.com/shogo82148/std/errors"
)

// ErrNotFoundは、パスの検索が実行可能なファイルを見つけられなかった場合のエラーです。
var ErrNotFound = errors.New("executable file not found in $PATH")

// LookPathは、PATH環境変数で指定されたディレクトリ内の実行可能なファイルを検索します。
// もしファイルがスラッシュを含んでいる場合、直接試みられ、
// PATHは参照されません。
// 成功すると、結果は絶対パスになります。
//
// Goの古いバージョンでは、LookPathは現在のディレクトリに対する相対パスを返すことができました。
// Go 1.19以降では、LookPathはそのパスとともに、errors.Is(err, ErrDot)を満たすエラーを返します。
// 詳細はパッケージのドキュメンテーションを参照してください。
func LookPath(file string) (string, error)
