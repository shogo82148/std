// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package rpcは、ネットワークやその他のI/O接続を介してオブジェクトのエクスポートされたメソッドへアクセスする機能を提供します。

net/rpcパッケージは凍結されており、新しい機能は受け付けていません。

サーバーはオブジェクトを登録し、その型名をサービス名として公開します。
登録後、オブジェクトのエクスポートされたメソッドはリモートからアクセス可能になります。
サーバーは異なる型の複数のオブジェクト（サービス）を登録できますが、同じ型のオブジェクトを複数登録することはエラーとなります。

以下の条件を満たすメソッドのみがリモートアクセス可能になります。それ以外のメソッドは無視されます：

  - メソッドの型がエクスポートされていること。
  - メソッドがエクスポートされていること。
  - メソッドが2つの引数を持ち、両方の引数がエクスポートされている（または組み込み）型であること。
  - メソッドの2番目の引数がポインターであること。
  - メソッドが戻り値としてerror型を持つこと。

要するに、メソッドは次のようなスキーマである必要があります。

	func (t *T) MethodName(argType T1, replyType *T2) error

ここで、T1とT2はencoding/gobでマーシャリングできる型です。
これらの要件は、異なるコーデックが使用されている場合でも適用されます。
（将来的には、カスタムコーデックに対してこれらの要件は緩和されるかもしれません。）

メソッドの最初の引数は呼び出し元から提供される引数を表し、
2番目の引数は呼び出し元に返される結果パラメーターを表します。
メソッドの戻り値がnilでない場合、それはクライアントが [errors.New] によって作成されたかのようにクライアントが確認する文字列として送り返されます。
エラーが返された場合、応答パラメーターはクライアントに送り返されません。

サーバーは、[ServeConn] を呼び出すことによって単一の接続上のリクエストを処理することができます。また、通常はネットワークリスナーを作成し、[Accept] を呼び出すか、HTTPリスナーの場合は [HandleHTTP] と [http.Serve] を呼び出します。

サービスを使用するためには、クライアントは接続を確立し、その後、接続上で [NewClient] を呼び出します。[Dial]（[DialHTTP]）という便利な関数は、生のネットワーク接続（HTTP接続）に対して両方の手順を実行します。結果として得られる [Client] オブジェクトには、サービスとメソッドを指定するための2つのメソッド、[Call] とGoがあり、引数を含むポインターと結果パラメーターを受け取るポインターを指定します。

Callメソッドは、リモート呼び出しが完了するまで待機し、
Goメソッドは非同期に呼び出しを開始し、Call構造体のDoneチャネルを使用して完了をシグナルします。

明示的なコーデックが設定されていない場合、データの転送には [encoding/gob] パッケージが使用されます。

以下にシンプルな例を示します。サーバーはArithタイプのオブジェクトをエクスポートしたい場合です。

	package server

	import "errors"

	type Args struct {
		A, B int
	}

	type Quotient struct {
		Quo, Rem int
	}

	type Arith int

	func (t *Arith) Multiply(args *Args, reply *int) error {
		*reply = args.A * args.B
		return nil
	}

	func (t *Arith) Divide(args *Args, quo *Quotient) error {
		if args.B == 0 {
			return errors.New("divide by zero")
		}
		quo.Quo = args.A / args.B
		quo.Rem = args.A % args.B
		return nil
	}

サーバーの呼び出し（HTTPサービスの場合）：

	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	go http.Serve(l, nil)

この時点で、クライアントは"Arith"というサービスとそのメソッド"Arith.Multiply"、"Arith.Divide"を見ることができます。呼び出すためには、クライアントはまずサーバーにダイヤルします。

	client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

そして、リモート呼び出しを行うことができます。

	// 同期呼び出し
	args := &server.Args{7,8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)

または

	// 非同期呼び出し
	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done	// divCallと等しい
	// エラーをチェックし、出力などを行います。

サーバーの実装では、クライアントのためのシンプルで型セーフなラッパーを提供することがよくあります。

# Security Considerations

このパッケージは悪意のあるピアに対して堅牢化されておらず、
[Go Security Policy] の対象外です。

[Go Security Policy]: https://go.dev/security/policylient.
*/
package rpc
