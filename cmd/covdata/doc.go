// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Covdataは、アプリケーションや統合テストを実行することによって生成される、2世代のカバレッジテストの出力ファイルを操作し、レポートを生成するためのプログラムです。例えば、

	$ mkdir ./profiledir
	$ go build -cover -o myapp.exe .
	$ GOCOVERDIR=./profiledir ./myapp.exe <arguments>
	$ ls ./profiledir
	covcounters.cce1b350af34b6d0fb59cc1725f0ee27.821598.1663006712821344241
	covmeta.cce1b350af34b6d0fb59cc1725f0ee27
	$

次のようにしてcovdataを実行します。"go tool covdata <モード>"と入力し、'モード'は、特定のレポート、マージ、またはデータ操作操作を選択するサブコマンドです。各モードの説明（指定されたモードの使用法の詳細については、"go tool cover <モード> -help"を実行してください）：

1. プロファイルされたパッケージごとのステートメントのカバー率を報告する：

	$ go tool covdata percent -i=profiledir
	cov-example/p	カバレッジ：ステートメントの41.1%
	main	カバレッジ：ステートメントの87.5%
	$

2. プロファイルされたパッケージのインポートパスを報告する：

	$ go tool covdata pkglist -i=profiledir
	cov-example/p
	main
	$

3. 関数ごとのステートメントのカバー率を報告する：

	$ go tool covdata func -i=profiledir
	cov-example/p/p.go:12:		emptyFn			0.0%
	cov-example/p/p.go:32:		Small			100.0%
	cov-example/p/p.go:47:		Medium			90.9%
	...
	$

4. カバレッジデータを古いテキスト形式に変換する：

	$ go tool covdata textfmt -i=profiledir -o=cov.txt
	$ head cov.txt
	mode: set
	cov-example/p/p.go:12.22,13.2 0 0
	cov-example/p/p.go:15.31,16.2 1 0
	cov-example/p/p.go:16.3,18.3 0 0
	cov-example/p/p.go:19.3,21.3 0 0
	...
	$ go tool cover -html=cov.txt
	$

5. プロファイルをマージする：

	$ go tool covdata merge -i=indir1,indir2 -o=outdir -modpaths=github.com/go-delve/delve
	$

6. プロファイルから別のプロファイルを差し引く：

	$ go tool covdata subtract -i=indir1,indir2 -o=outdir
	$

7. プロファイルの交差点を求める：

	$ go tool covdata intersect -i=indir1,indir2 -o=outdir
	$

8. デバッグ目的でプロファイルをダンプする：

	$ go tool covdata debugdump -i=indir
	<人間に読みやすい出力>
	$
*/package main
