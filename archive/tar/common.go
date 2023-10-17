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
	// Typeflag is the type of header entry.
	// The zero value is automatically promoted to either TypeReg or TypeDir
	// depending on the presence of a trailing slash in Name.
	Typeflag byte

	Name     string
	Linkname string

	Size  int64
	Mode  int64
	Uid   int
	Gid   int
	Uname string
	Gname string

	// If the Format is unspecified, then Writer.WriteHeader rounds ModTime
	// to the nearest second and ignores the AccessTime and ChangeTime fields.
	//
	// To use AccessTime or ChangeTime, specify the Format as PAX or GNU.
	// To use sub-second resolution, specify the Format as PAX.
	ModTime    time.Time
	AccessTime time.Time
	ChangeTime time.Time

	Devmajor int64
	Devminor int64

	// Xattrs stores extended attributes as PAX records under the
	// "SCHILY.xattr." namespace.
	//
	// The following are semantically equivalent:
	//  h.Xattrs[key] = value
	//  h.PAXRecords["SCHILY.xattr."+key] = value
	//
	// When Writer.WriteHeader is called, the contents of Xattrs will take
	// precedence over those in PAXRecords.
	//
	// Deprecated: Use PAXRecords instead.
	Xattrs map[string]string

	// PAXRecords is a map of PAX extended header records.
	//
	// User-defined records should have keys of the following form:
	//	VENDOR.keyword
	// Where VENDOR is some namespace in all uppercase, and keyword may
	// not contain the '=' character (e.g., "GOLANG.pkg.version").
	// The key and value should be non-empty UTF-8 strings.
	//
	// When Writer.WriteHeader is called, PAX records derived from the
	// other fields in Header take precedence over PAXRecords.
	PAXRecords map[string]string

	// Format specifies the format of the tar header.
	//
	// This is set by Reader.Next as a best-effort guess at the format.
	// Since the Reader liberally reads some non-compliant files,
	// it is possible for this to be FormatUnknown.
	//
	// If the format is unspecified when Writer.WriteHeader is called,
	// then it uses the first format (in the order of USTAR, PAX, GNU)
	// capable of encoding this Header (see Format).
	Format Format
}

// fs.FileInfoのNameメソッドは、説明するファイルのベース名のみを返すため、
// ファイルの完全なパス名を提供するためにHeader.Nameを変更する必要がある場合があります。
//
// fiが[FileInfoNames]を実装している場合、ヘッダーのGnameとUnameは、
// インターフェースのメソッドによって提供されます。
// FileInfoは、Headerのfs.FileInfoを返します。
func (h *Header) FileInfo() fs.FileInfo

// FileInfoHeaderは、fiから部分的に設定された [Header] を作成します。
// fiがシンボリックリンクを記述している場合、FileInfoHeaderはlinkをリンクターゲットとして記録します。
// fiがディレクトリを記述している場合、名前にスラッシュが追加されます。
//
// fs.FileInfoのNameメソッドは、
// 記述するファイルのベース名のみを返すため、
// ファイルの完全なパス名を提供するためにHeader.Nameを変更する必要がある場合があります。
func FileInfoHeader(fi fs.FileInfo, link string) (*Header, error)

// FileInfoNamesは、UID/GIDを名前に変換するために[FileInfo]を拡張します。
// これを[FileInfoHeader]に渡すことで、UID/GIDの解決をコントロールできます。
type FileInfoNames interface {
	fs.FileInfo
	// Uname should translate a UID into a user name.
	Uname(uid int) (string, error)
	// Gname should translate a GID into a group name.
	Gname(gid int) (string, error)
}
