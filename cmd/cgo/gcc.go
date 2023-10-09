// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// gccのデバッグ出力を解析して、Prog内のRefにCタイプのアノテーションを付ける。
// デバッグ出力をGoのタイプに変換する。

package main

<<<<<<< HEAD
// ProcessCgoDirectives processes the import C preamble:
//  1. discards all #cgo CFLAGS, LDFLAGS, nocallback and noescape directives,
//     so they don't make their way into _cgo_export.h.
//  2. parse the nocallback and noescape directives.
func (f *File) ProcessCgoDirectives()
=======
// DiscardCgoDirectivesはimport Cの前文を処理し、すべての＃cgo CFLAGSおよびLDFLAGSの指示を破棄します。これにより、_cgo_export.hに含まれないようになります。
func (f *File) DiscardCgoDirectives()
>>>>>>> release-branch.go1.21

// Translateは、インポートされたパッケージCへの参照を削除し、
// 対応するGoの型、関数、変数への参照に置き換えて、f.AST（元のGoの入力）を書き換えます。
func (p *Package) Translate(f *File)

// Stringは現在の型の表現を返します。フォーマット引数はこのメソッド内で組み立てられるため、可変な値の変更が考慮されます。
func (tr *TypeRepr) String() string

func (tr *TypeRepr) Empty() bool

// Setは型の表現を変更します。
// fargsが指定されている場合、reprはfmt.Sprintfのフォーマットとして使用されます。
// それ以外の場合、reprは処理されずに型の表現として使用されます。
func (tr *TypeRepr) Set(repr string, fargs ...interface{})
