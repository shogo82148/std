// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// CreateTempはディレクトリdirに新しい一時ファイルを作成し、
// ファイルを読み書きするために開き、結果のファイルを返します。
// ファイル名はpatternを取り、末尾にランダムな文字列を追加して生成されます。
// もしpatternに"*"が含まれている場合、最後の"*"はランダムな文字列に置き換えられます。
// dirが空の文字列の場合、CreateTempは一時ファイル用のデフォルトディレクトリ(TempDirが返す)を使用します。
// 同時にCreateTempを呼び出す複数のプログラムやゴルーチンは同じファイルを選びません。
// 呼び出し元は、ファイルのNameメソッドを使用してファイルのパス名を取得できます。
// ファイルが不要になったら、呼び出し元の責任でファイルを削除する必要があります。
func CreateTemp(dir, pattern string) (*File, error)

// MkdirTempはディレクトリdir内に新しい一時ディレクトリを作成し、
// 新しいディレクトリのパス名を返します。
// 新しいディレクトリの名前は、patternの末尾にランダムな文字列を追加することで生成されます。
// patternに"*"が含まれている場合、ランダムな文字列は最後の"*"に置換されます。
// dirが空の文字列の場合、MkdirTempは一時ファイルのデフォルトディレクトリ（TempDirによって返される）を使用します。
// 同時に複数のプログラムやゴルーチンがMkdirTempを呼び出しても、同じディレクトリを選択しません。
// ディレクトリは不要になった時に削除するのは呼び出し元の責任です。
func MkdirTemp(dir, pattern string) (string, error)
