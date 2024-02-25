// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fcgi

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/http"
)

// ErrRequestAborted は、ウェブサーバーによって中止されたリクエストのボディを読み込もうとするハンドラがReadを呼び出した場合に返されます。
var ErrRequestAborted = errors.New("fcgi: request aborted by web server")

// ErrConnClosedは、接続がウェブサーバーとの間で閉じられた後に、ハンドラがリクエストのボディを読み取ろうとした場合に、Readで返されます。
var ErrConnClosed = errors.New("fcgi: connection to web server closed")

// Serveはリスナーlで受け入れた入力FastCGI接続を処理し、それぞれのために新しいゴルーチンを作成します。ゴルーチンはリクエストを読み取り、その後ハンドラを呼び出して応答します。
// lがnilの場合、Serveはos.Stdinからの接続を受け入れます。
// handlerがnilの場合、[http.DefaultServeMux] が使用されます。
func Serve(l net.Listener, handler http.Handler) error

// ProcessEnvは、リクエストrに関連するFastCGI環境変数を返します。
// リクエスト自体に含まれるための努力がなされなかったデータは、リクエストのコンテキストに隠されています。
// たとえば、リクエストに対してREMOTE_USERが設定されている場合、r内のどこにも見つけることはできませんが、ProcessEnvの応答（rのコンテキストを介して）に含まれます。
func ProcessEnv(r *http.Request) map[string]string
