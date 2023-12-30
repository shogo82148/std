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
	// Nameはファイルの名前です。
	//
	// それは相対パスでなければならず、ドライブレター（"C:"など）で始まってはならず、
	// バックスラッシュの代わりにフォワードスラッシュを使用しなければなりません。末尾のスラッシュは
	// このファイルがディレクトリであり、データを持つべきではないことを示します。
	Name string

	// Commentは64KiB未満の任意のユーザー定義文字列です。
	Comment string

	// NonUTF8は、NameとCommentがUTF-8でエンコードされていないことを示します。
	//
	// 仕様によれば、許可される他のエンコーディングはCP-437のみですが、
	// 歴史的に多くのZIPリーダーはNameとCommentをシステムのローカル文字エンコーディングとして解釈します。
	//
	// このフラグは、ユーザーが特定のローカライズされた地域の非ポータブルなZIPファイルをエンコードするつもりである場合にのみ設定するべきです。
	// それ以外の場合、Writerは有効なUTF-8文字列のZIP形式のUTF-8フラグを自動的に設定します。
	NonUTF8 bool

	CreatorVersion uint16
	ReaderVersion  uint16
	Flags          uint16

	// Methodは圧縮方法です。ゼロの場合、Storeが使用されます。
	Method uint16

	// Modifiedはファイルの変更時間です。
	//
	// 読み取り時には、レガシーなMS-DOSの日付フィールドよりも拡張タイムスタンプが優先され、
	// 時間のオフセットがタイムゾーンとして使用されます。
	// MS-DOSの日付のみが存在する場合、タイムゾーンはUTCとみなされます。
	//
	// 書き込み時には、タイムゾーンに依存しない拡張タイムスタンプが常に出力されます。
	// レガシーなMS-DOSの日付フィールドは、Modified時間の位置に従ってエンコードされます。
	Modified time.Time

	// ModifiedTimeはMS-DOSでエンコードされた時間です。
	//
	// Deprecated: 代わりにModifiedを使用してください。
	ModifiedTime uint16

	// ModifiedDateはMS-DOSでエンコードされた日付です。
	//
	// Deprecated: 代わりにModifiedを使用してください。
	ModifiedDate uint16

	// CRC32は、ファイル内容のCRC32チェックサムです。
	CRC32 uint32

	// CompressedSizeは、ファイルの圧縮サイズ（バイト単位）です。
	// ファイルの非圧縮または圧縮サイズが32ビットに収まらない場合、
	// CompressedSizeは^uint32(0)に設定されます。
	//
	// Deprecated: 代わりにCompressedSize64を使用してください。
	CompressedSize uint32

	// UncompressedSizeは、ファイルの非圧縮サイズ（バイト単位）です。
	// ファイルの非圧縮または圧縮サイズが32ビットに収まらない場合、
	// CompressedSizeは^uint32(0)に設定されます。
	//
	// Deprecated: 代わりにUncompressedSize64を使用してください。
	UncompressedSize uint32

	// CompressedSize64は、ファイルの圧縮サイズ（バイト単位）です。
	CompressedSize64 uint64

	// UncompressedSize64は、ファイルの非圧縮サイズ（バイト単位）です。
	UncompressedSize64 uint64

	Extra         []byte
	ExternalAttrs uint32
}

// FileInfo は、[FileHeader] の fs.FileInfo を返します。
func (h *FileHeader) FileInfo() fs.FileInfo

// FileInfoHeaderは、fs.FileInfoから部分的に設定された [FileHeader] を作成します。
// fs.FileInfoのNameメソッドは、記述するファイルのベース名のみを返すため、
// ファイルの完全なパス名を提供するために、返されたヘッダーのNameフィールドを変更する必要がある場合があります。
// 圧縮が必要な場合は、呼び出し元はFileHeader.Methodフィールドを設定する必要があります。デフォルトでは設定されていません。
func FileInfoHeader(fi fs.FileInfo) (*FileHeader, error)

// ModTime は、旧来の ModifiedDate および [ModifiedTime] フィールドを使用して、UTC での変更時刻を返します。
//
// Deprecated: 代わりに [Modified] を使用してください。
func (h *FileHeader) ModTime() time.Time

// SetModTime は、与えられた時刻を UTC で指定して、 [Modified] 、 [ModifiedTime] 、および [ModifiedDate] フィールドを設定します。
//
// Deprecated: 代わりに [Modified] を使用してください。
func (h *FileHeader) SetModTime(t time.Time)

// Mode は、 [FileHeader] のパーミッションとモードビットを返します。
func (h *FileHeader) Mode() (mode fs.FileMode)

// SetMode は、 [FileHeader] のパーミッションとモードビットを変更します。
func (h *FileHeader) SetMode(mode fs.FileMode)
