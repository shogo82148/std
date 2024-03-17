// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 一般的な環境変数。

package os

<<<<<<< HEAD
// Expandはマッピング関数に基づいて文字列内の${var}または$varを置き換えます。
// 例えば、os.ExpandEnv(s)はos.Expand(s, os.Getenv)と同等です。
=======
// Expand replaces ${var} or $var in the string based on the mapping function.
// For example, [os.ExpandEnv](s) is equivalent to [os.Expand](s, [os.Getenv]).
>>>>>>> upstream/master
func Expand(s string, mapping func(string) string) string

// ExpandEnvは、文字列内の${var}または$varを現在の環境変数の値に応じて置換します。未定義の変数への参照は空文字列に置換されます。
func ExpandEnv(s string) string

<<<<<<< HEAD
// Getenvはキーで指定された環境変数の値を取得します。
// もし変数が存在しない場合、空の値が返されます。
// 空の値と未設定の値を区別するためには、LookupEnvを使用してください。
=======
// Getenv retrieves the value of the environment variable named by the key.
// It returns the value, which will be empty if the variable is not present.
// To distinguish between an empty value and an unset value, use [LookupEnv].
>>>>>>> upstream/master
func Getenv(key string) string

// LookupEnvは、キーで指定された環境変数の値を取得します。
// もし環境変数が存在する場合、値（空である可能性があります）と真偽値trueが返されます。
// そうでなければ、返される値は空で、真偽値はfalseです。
func LookupEnv(key string) (string, bool)

// Setenvは、キーで指定された環境変数の値を設定します。
// エラーがある場合は、エラーを返します。
func Setenv(key, value string) error

// Unsetenvは1つの環境変数を削除します。
func Unsetenv(key string) error

// Clearenvはすべての環境変数を削除します。
func Clearenv()

// Environは環境を表す文字列のコピーを返します。
// 形式は「key=value」です。
func Environ() []string
