// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

var _ WriterTo = (*multiReader)(nil)

// MultiReaderは、提供された入力リーダーの論理的な連結であるReaderを返します。
// これらは順次読み取られます。すべての入力がEOFを返したら、ReadはEOFを返します。
// リーダーのいずれかがnilでない、EOFでないエラーを返す場合、Readはそのエラーを返します。
func MultiReader(readers ...Reader) Reader

var _ StringWriter = (*multiWriter)(nil)

// MultiWriterは、Unixのtee(1)コマンドに似た、書き込みを提供されたすべてのライターに複製するライターを作成します。
//
// 各書き込みは、1つずつリストされたすべてのライターに書き込まれます。
// リストされたライターのいずれかがエラーを返すと、その全体の書き込み操作が停止し、エラーが返されます。
// リストは、書き込み操作が完了するまで変更されるべきではありません。
func MultiWriter(writers ...Writer) Writer
