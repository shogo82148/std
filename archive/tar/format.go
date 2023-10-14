// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tar

// Formatはtarアーカイブのフォーマットを表します。
//
// オリジナルのtarフォーマットはUnix V7で導入されました。
// その後、V7フォーマットの制限を克服するために標準化または拡張を試みる複数の競合するフォーマットがありました。
// 最も一般的なフォーマットは、USTAR、PAX、GNUフォーマットで、それぞれ独自の利点と制限があります。
//
// 次の表は、各フォーマットの機能を示しています：
//
//	                  |  USTAR |       PAX |       GNU
//	------------------+--------+-----------+----------
//	Name              |   256B | unlimited | unlimited
//	Linkname          |   100B | unlimited | unlimited
//	Size              | uint33 | unlimited |    uint89
//	Mode              | uint21 |    uint21 |    uint57
//	Uid/Gid           | uint21 | unlimited |    uint57
//	Uname/Gname       |    32B | unlimited |       32B
//	ModTime           | uint33 | unlimited |     int89
//	AccessTime        |    n/a | unlimited |     int89
//	ChangeTime        |    n/a | unlimited |     int89
//	Devmajor/Devminor | uint21 |    uint21 |    uint57
//	------------------+--------+-----------+----------
//	string encoding   |  ASCII |     UTF-8 |    binary
//	sub-second times  |     no |       yes |        no
//	sparse files      |     no |       yes |       yes
//
<<<<<<< HEAD
// この表の上部は、ヘッダーフィールドを示しており、各フォーマットが各文字列フィールドに許可される最大バイト数と、
// 各数値フィールドを格納するために使用される整数型を報告します
// （タイムスタンプは、Unixエポックからの秒数として格納されます）。
=======
// The table's upper portion shows the [Header] fields, where each format reports
// the maximum number of bytes allowed for each string field and
// the integer type used to store each numeric field
// (where timestamps are stored as the number of seconds since the Unix epoch).
>>>>>>> upstream/master
//
// 表の下部は、各フォーマットの特殊な機能を示しています。
// たとえば、サポートされる文字列エンコーディング、サブセカンドタイムスタンプのサポート、スパースファイルのサポートなどがあります。
//
// Writerは現在、スパースファイルに対するサポートを提供していません。
type Format int

// Constants to identify various tar formats.
const (
	// Deliberately hide the meaning of constants from public API.
	_ Format = (1 << iota) / 4

	// FormatUnknownは、フォーマットが不明であることを示します。
	FormatUnknown

	// FormatUSTARは、POSIX.1-1988で定義されたUSTARヘッダーフォーマットを表します。
	//
	// このフォーマットは、ほとんどのtarリーダーと互換性がありますが、
	// このフォーマットにはいくつかの制限があるため、一部の用途には適していません。
	// 特に、スパースファイル、8GiBを超えるファイル、256文字を超えるファイル名、および非ASCIIファイル名をサポートできません。
	//
	// 参考：
	//	http://pubs.opengroup.org/onlinepubs/9699919799/utilities/pax.html#tag_20_92_13_06
	FormatUSTAR

	// FormatPAXは、POSIX.1-2001で定義されたPAXヘッダーフォーマットを表します。
	//
	// PAXは、Typeflag TypeXHeaderを持つ特別なファイルを書き込むことで、USTARを拡張します。
	// このファイルには、USTARの制限を克服するために使用される一連のキー-値レコードが含まれています。
	// さらに、タイムスタンプのサブセカンド解像度を提供する機能もあります。
	//
	// 一部の新しいフォーマットは、独自のキーを定義し、関連する値に特定の意味を割り当てることで、PAXに独自の拡張機能を追加します。
	// たとえば、PAXでのスパースファイルのサポートは、GNUマニュアルで定義されたキー（例：「GNU.sparse.map」）を使用して実装されています。
	//
	// 参考：
	//	http://pubs.opengroup.org/onlinepubs/009695399/utilities/pax.html
	FormatPAX

	// FormatGNUは、GNUヘッダーフォーマットを表します。
	//
	// GNUヘッダーフォーマットは、USTARおよびPAX規格よりも古く、
	// それらとの互換性はありません。
	// GNUフォーマットは、任意のファイルサイズ、任意のエンコーディングと長さのファイル名、スパースファイルなどをサポートします。
	//
	// GNUフォーマットのアーカイブの解析ができるアプリケーションしかない場合を除き、PAXをGNUよりも選択することが推奨されます。
	//
	// 参考：
	//	https://www.gnu.org/software/tar/manual/html_node/Standard.html
	FormatGNU
)

func (f Format) String() string
