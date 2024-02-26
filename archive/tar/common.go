// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// tarパッケージは、tarアーカイブへのアクセスを実装します。
//
// テープアーカイブ（tar）は、ストリーミング方式で読み書きできるファイル形式で、
// 一連のファイルを格納するために使用されます。
// このパッケージは、GNUおよびBSD tarツールによって生成されたものを含め、
// このフォーマットのほとんどのバリエーションをカバーすることを目的としています。
package tar

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io/fs"
	"github.com/shogo82148/std/time"
)

var (
	ErrHeader          = errors.New("archive/tar: invalid tar header")
	ErrWriteTooLong    = errors.New("archive/tar: write too long")
	ErrFieldTooLong    = errors.New("archive/tar: header field too long")
	ErrWriteAfterClose = errors.New("archive/tar: write after close")
	ErrInsecurePath    = errors.New("archive/tar: insecure file path")
)

// Type flags for Header.Typeflag.
const (
	// Type '0' は通常のファイルを示します。
	TypeReg = '0'

	// Deprecated: 非推奨：かわりにTypeRegを使用してください。
	TypeRegA = '\x00'

	// Type '1'から'6'は、ヘッダーのみのフラグであり、データ本体を持たない場合があります。
	TypeLink    = '1'
	TypeSymlink = '2'
	TypeChar    = '3'
	TypeBlock   = '4'
	TypeDir     = '5'
	TypeFifo    = '6'

	// Type '7' は予約されています。
	TypeCont = '7'

	// Type 'x' は、PAXフォーマットで、次のファイルにのみ関連するキー-値レコードを格納するために使用されます。
	// このパッケージは、これらのタイプを透過的に処理します。
	TypeXHeader = 'x'

	// 'g' 型は、すべての後続ファイルに関連するキーと値のレコードを格納するために PAX 形式で使用されます。
	// このパッケージは、このようなヘッダーの解析と構成のみをサポートしていますが、現在はファイル間でグローバル状態を永続化することはできません。
	TypeXGlobalHeader = 'g'

	// 'S' 型は、GNU 形式でスパースファイルを示します。
	TypeGNUSparse = 'S'

	// 'L' 型と 'K' 型は、GNU 形式でメタファイルに使用されます。
	// このメタファイルは、次のファイルのパスまたはリンク名を格納するために使用されます。
	// このパッケージは、これらのタイプを透過的に処理します。
	TypeGNULongName = 'L'
	TypeGNULongLink = 'K'
)

// Header は、tar アーカイブ内の単一のヘッダーを表します。
// 一部のフィールドは、値が設定されていない場合があります。
//
// 将来の互換性のために、Reader.Next から Header を取得し、
// いくつかの方法で変更し、Writer.WriteHeader に戻すユーザーは、
// 新しい Header を作成し、保存する必要があるフィールドをコピーすることで行う必要があります。
type Header struct {
	// Typeflagはヘッダーエントリのタイプです。
	// ゼロ値は、Nameの末尾のスラッシュの有無に応じて、自動的にTypeRegまたはTypeDirに昇格します。
	Typeflag byte

	Name     string
	Linkname string

	Size  int64
	Mode  int64
	Uid   int
	Gid   int
	Uname string
	Gname string

	// Formatが指定されていない場合、Writer.WriteHeaderはModTimeを最も近い秒に丸め、
	// AccessTimeとChangeTimeフィールドを無視します。
	//
	// AccessTimeまたはChangeTimeを使用するには、FormatをPAXまたはGNUとして指定します。
	// サブセカンドの解像度を使用するには、FormatをPAXとして指定します。
	ModTime    time.Time
	AccessTime time.Time
	ChangeTime time.Time

	Devmajor int64
	Devminor int64

	// Xattrsは、"SCHILY.xattr."名前空間の下のPAXレコードとして拡張属性を保存します。
	//
	// 以下は意味的に等価です：
	//  h.Xattrs[key] = value
	//  h.PAXRecords["SCHILY.xattr."+key] = value
	//
	// Writer.WriteHeaderが呼び出されると、Xattrsの内容がPAXRecordsのものよりも優先されます。
	//
	// Deprecated: 代わりにPAXRecordsを使用してください。
	Xattrs map[string]string

	// PAXRecordsは、PAX拡張ヘッダーレコードのマップです。
	//
	// ユーザー定義のレコードは、次の形式のキーを持つべきです：
	//	VENDOR.keyword
	// ここで、VENDORはすべて大文字の何らかの名前空間であり、keywordは
	// '='文字を含んではなりません（例："GOLANG.pkg.version"）。
	// キーと値は非空のUTF-8文字列でなければなりません。
	//
	// Writer.WriteHeaderが呼び出されると、Headerの他のフィールドから派生した
	// PAXレコードは、PAXRecordsよりも優先されます。
	PAXRecords map[string]string

	// Formatはtarヘッダーの形式を指定します。
	//
	// これは、Reader.Nextによって形式の最善の推測として設定されます。
	// Readerは一部の非準拠ファイルを寛大に読み取るため、
	// これがFormatUnknownである可能性があります。
	//
	// Writer.WriteHeaderが呼び出されたときに形式が指定されていない場合、
	// このHeaderをエンコードできる最初の形式（USTAR、PAX、GNUの順）を使用します（Formatを参照）。
	Format Format
}

// fs.FileInfoのNameメソッドは、説明するファイルのベース名のみを返すため、
// ファイルの完全なパス名を提供するためにHeader.Nameを変更する必要がある場合があります。
//
// fiが [FileInfoNames] を実装している場合、ヘッダーのGnameとUnameは、
// インターフェースのメソッドによって提供されます。
// FileInfoは、Headerのfs.FileInfoを返します。
func (h *Header) FileInfo() fs.FileInfo

// FileInfoHeaderは、fiから部分的に設定された [Header] を作成します。
// fiがシンボリックリンクを記述している場合、FileInfoHeaderはlinkをリンクターゲットとして記録します。
// fiがディレクトリを記述している場合、名前にスラッシュが追加されます。
//
// fs.FileInfoのNameメソッドは、それが記述するファイルのベース名のみを返すため、
// ファイルの完全なパス名を提供するためにHeader.Nameを変更する必要があるかもしれません。
func FileInfoHeader(fi fs.FileInfo, link string) (*Header, error)
