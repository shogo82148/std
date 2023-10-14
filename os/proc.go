// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 処理など

package os

// Argsはコマンドラインの引数を保持し、プログラム名から開始します。
var Args []string

// Getuidは呼び出し元のユーザーの数値IDを返します。
//
// Windowsでは、-1を返します。
func Getuid() int

// Geteuidは呼び出し元の数値効果的ユーザーIDを返します。
//
// Windowsでは-1が返されます。
func Geteuid() int

// Getgidは呼び出し元のグループIDの数値を返します。
//
// Windowsでは、-1を返します。
func Getgid() int

// Getegidは呼び出し元の数値形式の有効グループIDを返します。
//
// Windowsでは、-1を返します。
func Getegid() int

// Getgroupsは、呼び出し元が所属しているグループの数値IDの一覧を返します。
//
// Windowsでは、syscall.EWINDOWSが返されます。代替手段については、os/userパッケージを参照してください。
func Getgroups() ([]int, error)

// Exitは指定されたステータスコードで現在のプログラムを終了させます。
// 慣習的に、コード0は成功を示し、非ゼロはエラーを示します。
// プログラムは直ちに終了します。延期された関数は実行されません。
//
// 移植性のために、ステータスコードは[0, 125]の範囲内にあるべきです。
func Exit(code int)
