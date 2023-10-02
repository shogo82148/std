// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package zip は、ZIP アーカイブの読み書きをサポートします。

詳細については、[ZIP specification] を参照してください。

このパッケージはディスクスパニングをサポートしていません。

ZIP64 についての注意点:

後方互換性を保つために、FileHeader には 32 ビットと 64 ビットの両方の Size フィールドがあります。
64 ビットフィールドには常に正しい値が含まれ、通常のアーカイブでは両方のフィールドが同じ値になります。
ZIP64 形式が必要なファイルの場合、32 ビットフィールドは 0xffffffff になり、代わりに 64 ビットフィールドを使用する必要があります。

[ZIP specification]: https://support.pkware.com/pkzip/appnote
*/
package zip

import (
	"github.com/shogo82148/std/io/fs"
	"github.com/shogo82148/std/time"
)

// 圧縮方式
const (
	Store   uint16 = 0
	Deflate uint16 = 8
)

// FileHeader は、ZIP ファイル内のファイルを説明します。
// 詳細については、[ZIP specification] を参照してください。
//
// [ZIP specification]: https://support.pkware.com/pkzip/appnote
type FileHeader struct {
	Name string

	Comment string

	NonUTF8 bool

	CreatorVersion uint16
	ReaderVersion  uint16
	Flags          uint16

	Method uint16

	Modified time.Time

	ModifiedTime uint16

	ModifiedDate uint16

	CRC32 uint32

	CompressedSize uint32

	UncompressedSize uint32

	CompressedSize64 uint64

	UncompressedSize64 uint64

	Extra         []byte
	ExternalAttrs uint32
}

// FileInfo は、FileHeader の fs.FileInfo を返します。
func (h *FileHeader) FileInfo() fs.FileInfo

// headerFileInfo implements fs.FileInfo.

// FileInfoHeader は、fs.FileInfo から部分的に設定された FileHeader を作成します。
// fs.FileInfo の Name メソッドは、説明するファイルの基本名のみを返すため、
// 返されたヘッダーの Name フィールドを変更して、ファイルの完全なパス名を提供する必要がある場合があります。
// 圧縮が必要な場合、呼び出し元は FileHeader.Method フィールドを設定する必要があります。
// デフォルトでは、Method は未設定です。
func FileInfoHeader(fi fs.FileInfo) (*FileHeader, error)

// ModTime は、旧来の ModifiedDate および ModifiedTime フィールドを使用して、UTC での変更時刻を返します。
//
// Deprecated: 代わりに Modified を使用してください。
func (h *FileHeader) ModTime() time.Time

// SetModTime は、与えられた時刻を UTC で指定して、Modified、ModifiedTime、および ModifiedDate フィールドを設定します。
//
// Deprecated: 代わりに Modified を使用してください。
func (h *FileHeader) SetModTime(t time.Time)

// Mode は、FileHeader のパーミッションとモードビットを返します。
func (h *FileHeader) Mode() (mode fs.FileMode)

// SetMode は、FileHeader のパーミッションとモードビットを変更します。
func (h *FileHeader) SetMode(mode fs.FileMode)
