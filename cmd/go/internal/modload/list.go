// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modload

import (
	"github.com/shogo82148/std/context"

	"github.com/shogo82148/std/cmd/go/internal/modinfo"
)

type ListMode int

const (
	ListU ListMode = 1 << iota
	ListRetracted
	ListDeprecated
	ListVersions
	ListRetractedVersions
)

// ListModules returns a description of the modules matching args, if known,
// along with any error preventing additional matches from being identified.
//
// The returned slice can be nonempty even if the error is non-nil.
func ListModules(ctx context.Context, args []string, mode ListMode, reuseFile string) ([]*modinfo.ModulePublic, error)

// ParsePathVersionは、argがpath@versionの形式であることを想定して解析します。
// argに'@'がない場合、foundはfalse、versは""、pathはargになります。これは
// strings.Cutの典型的な使い方に対応しています。ParsePathVersionは、モジュール
// バージョンの解析におけるstrings.Cutの汎用的な代替として使われることを意図しています。
// バージョンが不正な場合はエラーを返します。バージョンが'-'または'/'で始まる場合は
// 不正とみなされます。これらは、そのバージョンを使って実行するコマンドを組み立てる際に
// セキュリティ上の問題を引き起こす可能性があるためです。
func ParsePathVersion(arg string) (path, vers string, found bool, err error)
