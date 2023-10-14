// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package time

// Locationは、時間の瞬間をその時点で使用されているタイムゾーンにマップします。
// 通常、Locationは、地理的な地域で使用される時間オフセットのコレクションを表します。
// 多くのLocationでは、時間オフセットは、その時点で夏時間が使用されているかどうかによって異なります。
//
// Locationは、印刷されたTime値でタイムゾーンを提供し、夏時間の境界をまたぐ可能性のある間隔に関する計算に使用されます。
type Location struct {
	name string
	zone []zone
	tx   []zoneTrans

	// tzdata情報の後に、zoneTransに記録されていないDSTの移行をどのように処理するかを説明する文字列が続く場合があります。
	// フォーマットは、コロンを含まないTZ環境変数です。詳細は、以下を参照してください。
	// https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap08.html。
	// 例として、America/Los_Angelesの場合は、PST8PDT,M3.2.0,M11.1.0となります。
	extend string

	// ほとんどの検索は現在の時刻で行われます。
	// txをバイナリサーチしないように、
	// Locationが作成された時点の正しいタイムゾーンを
	// 提供する静的な単一要素キャッシュを保持します。
	// もし cacheStart <= t < cacheEnd であれば、
	// キャッシュは cacheZone を返すことができます。
	// cacheStart と cacheEnd の単位は、
	// lookupへの引数と一致するように、
	// 1970年1月1日UTCからの秒数です。
	cacheStart int64
	cacheEnd   int64
	cacheZone  *zone
}

// UTCは協定世界時（UTC）を表します。
var UTC *Location = &utcLoc

// Local はシステムのローカルなタイムゾーンを表します。
// Unix システムでは、Local は TZ 環境変数を参照して使用するタイムゾーンを見つけます。TZ が指定されていない場合は、システムのデフォルトの /etc/localtime を使用します。
// TZ="" は UTC を使用します。
// TZ="foo" はシステムのタイムゾーンディレクトリ内のファイル foo を使用します。
var Local *Location = &localLoc

// Stringは、LoadLocationまたはFixedZoneのname引数に対応する、タイムゾーン情報の記述的な名前を返します。
func (l *Location) String() string

// FixedZoneは、常に指定されたタイムゾーン名とオフセット（UTCからの秒数）を使用するLocationを返します。
func FixedZone(name string, offset int) *Location

// LoadLocationは指定された名前を持つLocationを返します。
//
// 名前が""または"UTC"の場合、LoadLocationはUTCを返します。
// 名前が"Local"の場合、LoadLocationはLocalを返します。
//
// それ以外の場合、名前はファイルに対応する場所の名前であり、
// IANAタイムゾーンデータベースの中に存在します。例えば"America/New_York"です。
//
// LoadLocationは以下の順序でIANAタイムゾーンデータベースを探します:
//
//   - ZONEINFO環境変数によって指定されたディレクトリまたは解凍されたzipファイル
//   - Unixシステムでは、システムの標準インストール場所
//   - $GOROOT/lib/time/zoneinfo.zip
//   - インポートされていた場合、time/tzdataパッケージ
func LoadLocation(name string) (*Location, error)
