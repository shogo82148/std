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
	// Name is the name of the file.
	//
	// It must be a relative path, not start with a drive letter (such as "C:"),
	// and must use forward slashes instead of back slashes. A trailing slash
	// indicates that this file is a directory and should have no data.
	Name string

	// Comment is any arbitrary user-defined string shorter than 64KiB.
	Comment string

	// NonUTF8 indicates that Name and Comment are not encoded in UTF-8.
	//
	// By specification, the only other encoding permitted should be CP-437,
	// but historically many ZIP readers interpret Name and Comment as whatever
	// the system's local character encoding happens to be.
	//
	// This flag should only be set if the user intends to encode a non-portable
	// ZIP file for a specific localized region. Otherwise, the Writer
	// automatically sets the ZIP format's UTF-8 flag for valid UTF-8 strings.
	NonUTF8 bool

	CreatorVersion uint16
	ReaderVersion  uint16
	Flags          uint16

	// Method is the compression method. If zero, Store is used.
	Method uint16

	// Modified is the modified time of the file.
	//
	// When reading, an extended timestamp is preferred over the legacy MS-DOS
	// date field, and the offset between the times is used as the timezone.
	// If only the MS-DOS date is present, the timezone is assumed to be UTC.
	//
	// When writing, an extended timestamp (which is timezone-agnostic) is
	// always emitted. The legacy MS-DOS date field is encoded according to the
	// location of the Modified time.
	Modified time.Time

	// ModifiedTime is an MS-DOS-encoded time.
	//
	// Deprecated: Use Modified instead.
	ModifiedTime uint16

	// ModifiedDate is an MS-DOS-encoded date.
	//
	// Deprecated: Use Modified instead.
	ModifiedDate uint16

	// CRC32 is the CRC32 checksum of the file content.
	CRC32 uint32

	// CompressedSize is the compressed size of the file in bytes.
	// If either the uncompressed or compressed size of the file
	// does not fit in 32 bits, CompressedSize is set to ^uint32(0).
	//
	// Deprecated: Use CompressedSize64 instead.
	CompressedSize uint32

	// UncompressedSize is the compressed size of the file in bytes.
	// If either the uncompressed or compressed size of the file
	// does not fit in 32 bits, CompressedSize is set to ^uint32(0).
	//
	// Deprecated: Use UncompressedSize64 instead.
	UncompressedSize uint32

	// CompressedSize64 is the compressed size of the file in bytes.
	CompressedSize64 uint64

	// UncompressedSize64 is the uncompressed size of the file in bytes.
	UncompressedSize64 uint64

	Extra         []byte
	ExternalAttrs uint32
}

// FileInfo は、FileHeader の fs.FileInfo を返します。
func (h *FileHeader) FileInfo() fs.FileInfo

// FileInfoHeaderは、fs.FileInfoから部分的に設定されたFileHeaderを作成します。
// fs.FileInfoのNameメソッドは、記述するファイルのベース名のみを返すため、
// ファイルの完全なパス名を提供するために、返されたヘッダーのNameフィールドを変更する必要がある場合があります。
// 圧縮が必要な場合は、呼び出し元はFileHeader.Methodフィールドを設定する必要があります。デフォルトでは設定されていません。
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
