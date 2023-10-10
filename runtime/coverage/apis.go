// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package coverage

import (
	"github.com/shogo82148/std/io"
)

// WriteMetaDirは現在実行中のプログラムのカバレッジのメタデータファイルを'dir'で指定されたディレクトリに書き込みます。操作が正常に完了できない場合にはエラーが返されます（たとえば、現在実行中のプログラムが"-cover"でビルドされていない場合、またはディレクトリが存在しない場合など）。
func WriteMetaDir(dir string) error

// WriteMetaは、現在実行中のプログラムのメタデータコンテンツ（通常はメタデータファイルに出力されるペイロード）をライター 'w'に書き込みます。操作が正常に完了できない場合（例えば、現在実行中のプログラムが "-cover" でビルドされていない場合や書き込みに失敗した場合など）、エラーが返されます。
func WriteMeta(w io.Writer) error

// WriteCountersDirは、現在実行中のプログラムのカバレッジカウンターデータファイルを'dir'で指定されたディレクトリに書き込みます。操作を正常に完了できない場合（たとえば、現在実行中のプログラムが'-cover'でビルドされていない場合や、ディレクトリが存在しない場合など）、エラーが返されます。書き込まれるカウンターデータは、呼び出し時のスナップショットとなります。
func WriteCountersDir(dir string) error

// WriteCountersは現在実行中のプログラムのカバレッジカウンターデータの内容をライター'w'に書き込みます。現在実行中のプログラムが"-cover"でビルドされていない場合や書き込みが失敗した場合など、操作が正常に完了できない場合はエラーが返されます。書き込まれるカウンターデータは、呼び出し時のスナップショットになります。
func WriteCounters(w io.Writer) error

// ClearCountersは現在実行中のプログラム内のカバレッジカウンタ変数をクリア/リセットします。
// "-cover"フラグでビルドされたプログラムではない場合、エラーが返されます。
// カウンタのクリアは、アトミックカウンタモードを使用しないプログラムに対してもサポートされていません
// (詳細なコメントについては、下記を参照してください）。
func ClearCounters() error
