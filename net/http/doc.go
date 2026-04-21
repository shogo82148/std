// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
httpパッケージは HTTP クライアントとサーバーの実装を提供します。

[Get]、[Head]、[Post]、[PostForm] は HTTP (または HTTPS) リクエストを行います:

	resp, err := http.Get("http://example.com/")
	...
	resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
	...
	resp, err := http.PostForm("http://example.com/form",
		url.Values{"key": {"Value"}, "id": {"123"}})

関数を呼び出した後、レスポンスボディを閉じる必要があります。

	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	// ...

# Clients and Transports

HTTP クライアントヘッダー、リダイレクトポリシー、その他の設定を制御するには、[Client] を作成してください。

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	resp, err := client.Get("http://example.com")
	// ...

	req, err := http.NewRequest("GET", "http://example.com", nil)
	// ...
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	// ...

プロキシ、TLS 設定、Keep-Alive、圧縮、その他の設定を制御するには、[Transport] を作成してください。

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://example.com")

クライアントとトランスポートは、複数のゴルーチンによる同時使用に対して安全であり、効率的に使用するためには、1度だけ作成して再利用する必要があります。

# Servers

ListenAndServe は、指定されたアドレスとハンドラーで HTTP サーバーを開始します。
ハンドラーは通常 nil で、[DefaultServeMux] を使用することを意味します。
[Handle] と [HandleFunc] は、[DefaultServeMux] にハンドラーを追加します。

	http.Handle("/foo", fooHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

サーバーの動作に関するより詳細な制御は、カスタムサーバーを作成することで利用できます。

	s := &http.Server{
		Addr:           ":8080",
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

# HTTP/2

httpパッケージは HTTP/2 プロトコルに対する透過的なサポートを提供します。

[Server] と [DefaultTransport] は HTTPS を使用する際に自動的に HTTP/2 サポートを有効にします。
[Transport] はデフォルトで HTTP/2 を有効にしません。

HTTP/1、HTTP/2、および/または暗号化されていない HTTP/2 のサポートを有効または無効にするには、
[Server.Protocols] と [Transport.Protocols] の設定フィールドを参照してください。

高度な HTTP/2 機能を設定するには、[Server.HTTP2] と [Transport.HTTP2] の設定フィールドを参照してください。

代替として、現在以下の GODEBUG 設定がサポートされています：

	GODEBUG=http2client=0  # HTTP/2 クライアントサポートを無効にする
	GODEBUG=http2server=0  # HTTP/2 サーバーサポートを無効にする
	GODEBUG=http2debug=1   # 詳細な HTTP/2 デバッグログを有効にする
	GODEBUG=http2debug=2   # ... フレームダンプを含めて、より詳細なログを有効にする

"omithttp2" ビルドタグを使用して、http パッケージに含まれる HTTP/2 実装を無効にすることができます。
*/
package http
