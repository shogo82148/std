// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルは、CGIのホスト側（ウェブサーバーの親プロセス）を実装しています。

// パッケージcgiは、RFC 3875で指定されているCGI（Common Gateway Interface）を実装しています。
//
// CGIを使用すると、各リクエストを処理するために新しいプロセスを起動することを意味しますが、
// これは通常、長時間実行されるサーバーを使用するよりも効率が低いです。
// このパッケージは、主に既存のシステムとの互換性を保つためのものです。
package cgi

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/net/http"
)

// HandlerはCGI環境でサブプロセス内で実行可能なプログラムです。
type Handler struct {
	Path string
	Root string

	// DirはCGI実行可能ファイルの作業ディレクトリを指定します。
	// Dirが空の場合、Pathのベースディレクトリが使用されます。
	// Pathにベースディレクトリがない場合、現在の作業ディレクトリが使用されます。
	Dir string

	Env        []string
	InheritEnv []string
	Logger     *log.Logger
	Args       []string
	Stderr     io.Writer

	// PathLocationHandlerは、CGIプロセスが"/"で始まるLocationヘッダーの値を返す場合に、内部リダイレクトを処理するためのルートHTTP Handlerを指定します。これは、RFC 3875 § 6.3.2で指定されていますが、おそらくhttp.DefaultServeMuxになるでしょう。
	// もしnilの場合、ローカルURIパスを持つCGIレスポンスがクライアントに送信され、内部リダイレクトは行われません。
	PathLocationHandler http.Handler
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request)
