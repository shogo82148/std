// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Doc (通常は go doc として実行される) は 0 個、1 個、または 2 個の引数を受け付けます。
//
// 0 個の引数:
//
//	go doc
//
// 現在のディレクトリに含まれるパッケージのドキュメントを表示します。
//
// 1 個の引数:
//
//	go doc <pkg>
//	go doc <sym>[.<methodOrField>]
//	go doc [<pkg>.]<sym>[.<methodOrField>]
//	go doc [<pkg>.][<sym>.]<methodOrField>
//
// 成功する最初の項目のドキュメントが表示されます。シンボルが指定されているがパッケージが指定されていない場合、現在のディレクトリのパッケージが選択されます。ただし、引数が大文字で始まる場合は常に現在のディレクトリのシンボルと見なされます。
//
// 2 個の引数:
//
//	go doc <pkg> <sym>[.<methodOrField>]
//
// パッケージ、シンボル、およびメソッドまたはフィールドのドキュメントを表示します。最初の引数は完全なパッケージパスである必要があります。これは godoc コマンドのコマンドライン使用法と似ています。
//
// コマンドの場合、-cmd フラグが存在しない限り、"go doc コマンド" はパッケージレベルのドキュメントのみ表示します。
//
// -src フラグを指定すると、doc は構造体、関数、またはメソッドの本体などのシンボルの全ソースコードを表示します。
//
// -all フラグを指定すると、doc はパッケージとその可視なシンボルのすべてのドキュメントを表示します。引数はパッケージを識別する必要があります。
//
// 完全なドキュメントについては、「go help doc」を実行してください。
package main
