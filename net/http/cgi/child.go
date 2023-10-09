// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルは子プロセスの視点からCGIを実装しています。

package cgi

import (
	"github.com/shogo82148/std/net/http"
)

// Request は現在の環境で表される HTTP リクエストを返します。
// この関数は、現在のプログラムが CGI 環境で実行されていることを前提としています。
// 適用される場合、返された Request の Body は取得されます。
func Request() (*http.Request, error)

// RequestFromMapはCGI変数からhttp.Requestを作成します。
// 返されたRequestのBodyフィールドは入力されません。
func RequestFromMap(params map[string]string) (*http.Request, error)

// Serveは現在アクティブなCGIリクエストに提供されたHandlerを実行します。もし現在のCGI環境がない場合、エラーが返されます。提供されたハンドラーがnilの場合、http.DefaultServeMuxが使用されます。
func Serve(handler http.Handler) error
