# AGENTS.md

このリポジトリは、Go 標準ライブラリの godoc（ドキュメントコメント）を日本語に翻訳したものです。
モジュールパスは `github.com/shogo82148/std` です。

## リポジトリの構成

- 各パッケージのディレクトリ（`bytes/`, `net/http/`, `cmd/` など）は golang/go の `src/` と同じ構成になっています。
- ファイルには **公開 API の宣言とドキュメントコメントだけ** が含まれます。関数の本体、非公開の宣言、テストは `_tools/importgodoc` によって取り除かれています。
- 標準ライブラリの import は `github.com/shogo82148/std/...` に書き換えられています。
- `_tools/` は独立した Go モジュール（`_tools/go.mod`）で、生成用のツールが入っています。
  - `importgodoc`: golang/go のソースから godoc だけを抜き出して、このリポジトリに書き出します。
  - `rm-all.sh`: 同期の前に、`_` で始まるもの以外のディレクトリをすべて削除します。

## ブランチ運用

| ブランチ | 内容 |
| --- | --- |
| `upstream/<branch>` | sync ワークフロー（`.github/workflows/sync.yml`）が毎日自動で生成する英語の原文です。**手で編集しないでください。** |
| `<branch>`（`master`, `release-branch.go1.XX`） | 和訳済みのブランチです。作業は基本的にこちらに対して行います。 |

### upstream の変更を取り込む

1. `<branch>` から `merge-to-<branch>-YYYY-MM-DD` というブランチを作ります。
2. `origin/upstream/<branch>` をマージします。
3. コンフリクトは、upstream の変更内容を反映した和訳にして解消します（英語の原文を残したままにしない）。
4. `<branch>` に向けて PR を作ります。PR の本文には、コンフリクトを解消したファイルと変更内容の要約を書きます。

### 翻訳作業

`translate-<内容>-go1.XX` のようなブランチを切り、対象の `<branch>` に向けて PR を作ります。

## 翻訳ルール

- 翻訳するのは **コメントだけ** です。コード（宣言、型、import など）は変更しないでください。
- 先頭の `// Copyright ...` のライセンスヘッダーは翻訳しません。
- `#` で始まる見出し行は翻訳せず、upstream の英語のまま残します（pkg.go.dev のアンカーになるため）。本文だけを訳します。
- doc link（`[Ident]`, `[pkg.Ident]`, `[Type.Method]`）は、前後を空白か句読点で区切ります。日本語と直接くっつけると `go/doc/comment` がリンクとして認識しません。
  - 良い例: `これのインスタンスを [FileInfoHeader] に渡すと、`
  - 悪い例: `これのインスタンスを[FileInfoHeader]に渡すと、`
- `Deprecated:` や `BUG(...)` などのマーカーは英語のまま残し、その後ろの説明文だけを訳します。
- インデントされたコードブロック、エスケープシーケンス、識別子名、URL は原文どおりに残します。
- 識別子名は実在するものと一致しているか確認します（フィールド名のタイプミスやコピペミスに注意）。
- `//go:` ディレクティブや `// Output:` などの特別なコメントには手を加えません。

## 検証

CI（`.github/workflows/godoc.yml`）と同じく、すべてのパッケージの godoc がエラーなく出力できることを確認します。

```bash
set -o pipefail
go list ./... | xargs -n1 go doc -all > /dev/null
```

CI は `gotip` で実行しています。新しい言語機能を使ったファイルがあり、手元の Go で失敗する場合は `gotip` を使ってください。

`_tools` を変更した場合は、そのディレクトリでテストを実行します。

```bash
cd _tools && go test ./...
```

## コミット・PR

- コミットメッセージは英語で書きます（例: `Translate runtime/metrics package documentation into Japanese`）。
- PR の説明は日本語で構いません。
