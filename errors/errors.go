// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// errorsパッケージは、エラーを操作するための関数を実装しています。
//
// [New]関数は、テキストメッセージのみが含まれるエラーを作成します。
//
// エラーeの型にメソッドが1つある場合、eは別のエラーをラップします。
//
//	Unwrap() error
//	Unwrap() []error
//
// もしe.Unwrap()がnilでないエラーwまたはwを含むスライスを返した場合、
// eはwをラップしていると言います。e.Unwrap()がnilを返した場合、
// eは他のエラーをラップしていないことを示します。Unwrapメソッドが、
// nilエラー値を含む[]errorを返すことは無効です。
//
// ラップされたエラーを作成する簡単な方法は、 [fmt.Errorf] を呼び出し、
// エラー引数に%w動詞を適用することです。
//
//	wrapsErr := fmt.Errorf("... %w ...", ..., err, ...)
//
// エラーの連続的なアンラップにより、ツリーが作成されます。
// [Is] および [As] 関数は、エラーのツリーを調べるために、
// 最初にエラー自体を調べ、次にその子のツリーを順番に調べます
// （前順、深さ優先のトラバース）。
//
<<<<<<< HEAD
// Isは、第1引数のエラーツリーを調べ、第2引数に一致するエラーを探します。
// 一致するエラーが見つかった場合、trueを返します。
// 単純な等価性チェックよりも使用することをお勧めします。
=======
// [Is] examines the tree of its first argument looking for an error that
// matches the second. It reports whether it finds a match. It should be
// used in preference to simple equality checks:
>>>>>>> upstream/master
//
//	if errors.Is(err, fs.ErrExist)
//
// これは、次のようなエラーの比較よりも好ましいです。
//
//	if err == fs.ErrExist
//
// なぜなら、前者はエラーが [io/fs.ErrExist] をラップしている場合に成功するからです。
//
<<<<<<< HEAD
// Asは、第1引数のエラーツリーを調べ、第2引数に割り当て可能なエラーを探します。
// 第2引数はポインタである必要があります。成功した場合、Asは割り当てを実行し、trueを返します。
// それ以外の場合、falseを返します。フォームは次のようになります。
=======
// [As] examines the tree of its first argument looking for an error that can be
// assigned to its second argument, which must be a pointer. If it succeeds, it
// performs the assignment and returns true. Otherwise, it returns false. The form
>>>>>>> upstream/master
//
//	var perr *fs.PathError
//	if errors.As(err, &perr) {
//		fmt.Println(perr.Path)
//	}
//
// これは、次のようなエラーの比較よりも好ましいです。
//
//	if perr, ok := err.(*fs.PathError); ok {
//		fmt.Println(perr.Path)
//	}
//
// なぜなら、前者はエラーが[*io/fs.PathError]をラップしている場合に成功するからです。
package errors

// Newは、指定されたテキストをフォーマットするエラーを返します。
// テキストが同じであっても、Newの各呼び出しは異なるエラー値を返します。
func New(text string) error

// ErrUnsupportedは、サポートされていないため、要求された操作を実行できないことを示します。
// たとえば、ハードリンクをサポートしていないファイルシステムを使用して [os.Link] を呼び出す場合。
//
// 関数やメソッドは、このエラーを返すべきではありません。
// 代わりに、適切な文脈を含むエラーを返すべきです。
//
//	errors.Is(err, errors.ErrUnsupported)
//
<<<<<<< HEAD
// ErrUnsupportedを直接ラップするか、Isメソッドを実装することによって、
// サポートされていないことを示すエラーを返すことができます。
=======
// either by directly wrapping ErrUnsupported or by implementing an [Is] method.
>>>>>>> upstream/master
//
// 関数やメソッドは、このエラーをラップして返す場合があることを文書化する必要があります。
var ErrUnsupported = New("unsupported operation")
