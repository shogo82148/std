// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package http は HTTP クライアントとサーバーの実装を提供します。

<<<<<<< HEAD
Get、Head、Post、PostForm は HTTP (または HTTPS) リクエストを行います:
=======
[Get], [Head], [Post], and [PostForm] make HTTP (or HTTPS) requests:
>>>>>>> upstream/release-branch.go1.22

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

<<<<<<< HEAD
HTTP クライアントヘッダー、リダイレクトポリシー、その他の設定を制御するには、Client を作成してください。
=======
For control over HTTP client headers, redirect policy, and other
settings, create a [Client]:
>>>>>>> upstream/release-branch.go1.22

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

<<<<<<< HEAD
プロキシ、TLS 設定、Keep-Alive、圧縮、その他の設定を制御するには、Transport を作成してください。
=======
For control over proxies, TLS configuration, keep-alives,
compression, and other settings, create a [Transport]:
>>>>>>> upstream/release-branch.go1.22

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://example.com")

クライアントとトランスポートは、複数のゴルーチンによる同時使用に対して安全であり、効率的に使用するためには、1度だけ作成して再利用する必要があります。

# Servers

<<<<<<< HEAD
ListenAndServe は、指定されたアドレスとハンドラーで HTTP サーバーを開始します。
ハンドラーは通常 nil で、DefaultServeMux を使用することを意味します。
Handle と HandleFunc は、DefaultServeMux にハンドラーを追加します。
=======
ListenAndServe starts an HTTP server with a given address and handler.
The handler is usually nil, which means to use [DefaultServeMux].
[Handle] and [HandleFunc] add handlers to [DefaultServeMux]:
>>>>>>> upstream/release-branch.go1.22

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

<<<<<<< HEAD
Go 1.6 以降、HTTPS を使用する場合、http パッケージは HTTP/2 プロトコルの透過的なサポートを提供します。HTTP/2 を無効にする必要があるプログラムは、Transport.TLSNextProto (クライアント用) または Server.TLSNextProto (サーバー用) を nil でない空のマップに設定することで行えます。また、次の GODEBUG 設定が現在サポートされています。
=======
Starting with Go 1.6, the http package has transparent support for the
HTTP/2 protocol when using HTTPS. Programs that must disable HTTP/2
can do so by setting [Transport.TLSNextProto] (for clients) or
[Server.TLSNextProto] (for servers) to a non-nil, empty
map. Alternatively, the following GODEBUG settings are
currently supported:
>>>>>>> upstream/release-branch.go1.22

	GODEBUG=http2client=0  # HTTP/2 クライアントサポートを無効にする
	GODEBUG=http2server=0  # HTTP/2 サーバーサポートを無効にする
	GODEBUG=http2debug=1   # 詳細な HTTP/2 デバッグログを有効にする
	GODEBUG=http2debug=2   # ... フレームダンプを含めて、より詳細なログを有効にする

HTTP/2 サポートを無効にする前に、問題がある場合は報告してください: https://golang.org/s/http2bug

<<<<<<< HEAD
http パッケージの Transport と Server は、単純な構成に対して自動的に HTTP/2 サポートを有効にします。より複雑な構成で HTTP/2 を有効にする、より低レベルの HTTP/2 機能を使用する、またはより新しいバージョンの Go の http2 パッケージを使用するには、直接 "golang.org/x/net/http2" をインポートし、その ConfigureTransport および/または ConfigureServer 関数を使用します。golang.org/x/net/http2 パッケージを使用して HTTP/2 を手動で設定する場合、net/http パッケージの組み込みの HTTP/2 サポートよりも優先されます。
=======
The http package's [Transport] and [Server] both automatically enable
HTTP/2 support for simple configurations. To enable HTTP/2 for more
complex configurations, to use lower-level HTTP/2 features, or to use
a newer version of Go's http2 package, import "golang.org/x/net/http2"
directly and use its ConfigureTransport and/or ConfigureServer
functions. Manually configuring HTTP/2 via the golang.org/x/net/http2
package takes precedence over the net/http package's built-in HTTP/2
support.
>>>>>>> upstream/release-branch.go1.22
*/
package http
