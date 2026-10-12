# 翻訳メモリー（glossary）

Go 標準ライブラリーの godoc を和訳するときの用語集・定型文集です。
`origin/upstream/release-branch.go1.27`（英語原文）と `release-branch.go1.27`（和訳）のドキュメントコメントを宣言単位で対応付け（約 6,800 組）、よく使われている訳語を集計してまとめました。

- 「推奨訳」は原則として既存の和訳で最も多く使われている訳です（カタカナ語の長音は後述の規則を優先します）。新しく訳すときはこちらに揃えてください。
- 括弧内の数字は、既存の和訳での出現回数の目安です。
- 「避ける訳」は既存の和訳に混在しているが、揃えたい表記です。既存の訳を見つけても、無理にまとめて直す必要はありません。

翻訳ルール全般（見出し・doc link・`Deprecated:` の扱いなど）は [AGENTS.md](AGENTS.md) を参照してください。

## 表記の基本方針

### カタカナ語の長音

もとの英語のつづりが **-er, -or, -ar, -y** で終わる語は、語尾に「ー」を付けます。
既存の和訳では長音なしの表記も多く混在していますが、新しく訳すときはこの規則に従ってください。

| 英語 | 推奨表記 | 避ける表記 |
| --- | --- | --- |
| server | サーバー | サーバ |
| header | ヘッダー | ヘッダ |
| parameter | パラメーター | パラメータ |
| security | セキュリティー | セキュリティ |
| user | ユーザー | ユーザ |
| reader / writer | リーダー／ライター | リーダ／ライタ |
| parser | パーサー | パーサ |
| receiver | レシーバー | レシーバ |
| handler | ハンドラー | ハンドラ |
| buffer | バッファー | バッファ |
| iterator | イテレーター | イテレータ |
| pointer | ポインター | ポインタ |
| compiler | コンパイラー | コンパイラ |
| driver | ドライバー | ドライバ |
| listener | リスナー | リスナ |
| resolver | リゾルバー | リゾルバ |
| collector | コレクター | コレクタ |
| counter | カウンター | カウンタ |
| descriptor | ディスクリプター | ディスクリプタ |
| encoder / decoder | エンコーダー／デコーダー | エンコーダ／デコーダ |
| scanner | スキャナー | スキャナ |
| helper | ヘルパー | ヘルパ |
| directory | ディレクトリー | ディレクトリ |
| entry | エントリー | エントリ |
| memory | メモリー | メモリ |
| library | ライブラリー | ライブラリ |
| binary | バイナリー | バイナリ |
| query | クエリー | クエリ |
| proxy | プロキシー | プロキシ |
| category | カテゴリー | カテゴリ |
| body | ボディー | ボディ |
| key | キー | |

規則の対象外の語は、既存の和訳での慣用に従います。

| 英語 | 推奨表記 | 避ける表記 |
| --- | --- | --- |
| interface | インターフェース (531) | インタフェース (22) |
| channel | チャネル (102) | チャンネル |
| goroutine | ゴルーチン (236) | goroutine、ゴールーチン |

### 英語のまま残すもの

- 識別子、型名、パッケージ名、引数名（`nil`, `true`, `false`, `error`, `io.Reader` など）
- doc link（`[Ident]`, `[pkg.Ident]`）。前後に空白か句読点を入れる
- `Deprecated:`、`BUG(...)` などのマーカー
- RFC 番号、プロトコル名、アルゴリズム名（`UTF-8`, `HTTP/2`, `SHA-256`, `TLS` など）

## 定型文

| 英語 | 推奨訳 | 備考 |
| --- | --- | --- |
| Package X implements ... | X パッケージは、...を実装します。 | 「パッケージ X は」より「X パッケージは」が多数派 (155 : 8) |
| Package X provides ... | X パッケージは、...を提供します。 | |
| X returns ... | X は ... を返します。 | |
| X returns a new Y. | X は新しい Y を返します。 | |
| X returns the number of ... | X は ... の数を返します。 | |
| X reports whether ... | X は、...かどうかを報告します。 | (296)。「かどうかを返します」(9) は避ける |
| X sets ... to ... | X は ... を ... に設定します。 | |
| A X represents ... | X は ... を表します。 | |
| A X describes ... | X は ... を記述します。／説明します。 | |
| X is the interface that wraps the Y method. | X は、Y メソッドをラップするインターフェースです。 | |
| X is the interface that groups the basic Y and Z methods. | X は、基本的な Y と Z メソッドをグループ化するインターフェースです。 | |
| X implements the [Y] interface. | X は [Y] インターフェースを実装します。 | |
| X is equivalent to Y. | X は Y と同等です。 | (65)。「等価です」(17) より多い |
| X is like Y but ... | X は Y と同様ですが、... | |
| X is a convenience wrapper around Y. | X は Y の便利なラッパーです。 | |
| See [X] for details. | 詳細は [X] を参照してください。 | 「をご覧ください」は使わない |
| See the X documentation for details. | 詳細については、X のドキュメントを参照してください。 | |
| For example, ... | 例えば、... | 括弧内の (for example, "...") は（例: "..."） |
| Note that ... | ...ことに注意してください。 | |
| It panics if ... | ...場合、パニックを引き起こします。 | 「パニックを起こします」「パニックします」「パニックを発生させます」「パニックが発生します」は使わない。ただし仕様の "a run-time panic occurs" は「ランタイムパニックが発生します」 |
| It returns an error if ... | ...場合はエラーを返します。 | |
| If X is nil, ... | X が nil の場合、... | |
| ..., or nil if ... | ...を返します。...の場合は nil を返します。 | |
| X is called when ... | X は ... ときに呼び出されます。 | |
| X blocks until ... | X は ... までブロックします。 | |
| The zero value for X is ... ready to use. | X のゼロ値は、すぐに使用できる...です。 | |
| ... is safe for concurrent use by multiple goroutines. | ...は複数のゴルーチンによる同時使用に対して安全です。 | 訳がばらついているので、この形に揃えたい |
| Multiple goroutines may invoke methods on a X simultaneously. | 複数のゴルーチンが X のメソッドを同時に呼び出すことができます。 | |
| It must not be modified. | 変更してはいけません。 | |
| ... should not be used. | ...は使用すべきではありません。 | |
| It is the caller's responsibility to ... | ...のは呼び出し元の責任です。 | |
| ... is ignored. | ...は無視されます。 | |
| ... has no effect. | ...は何もしません。／何の効果もありません。 | |
| ... is not supported. | ...はサポートされていません。 | |
| ... is kept for compatibility. | ...は互換性のために保持されています。 | |
| ... is reserved for future use. | ...は将来の使用のために予約されています。 | |
| at most N | 最大 N／N 以下 | |
| at least N | 少なくとも N／N 以上 | |
| Deprecated: Use X instead. | Deprecated: 代わりに X を使用してください。 | 「かわりに」は使わない。「非推奨：」を重ねない |
| Deprecated: No longer used. | Deprecated: もはや使用されていません。 | |

## 用語集

### Go 言語・型システム

| 英語 | 推奨訳 | 備考 |
| --- | --- | --- |
| type | 型 | 「タイプ」は Mach-O の cpu type などファイル形式上の用語や、`Type 'x'` のような原文の表記に限る |
| value | 値 | |
| zero value | ゼロ値 | |
| function | 関数 | |
| method | メソッド | |
| method set | メソッドセット | |
| receiver | レシーバー | |
| argument | 引数 | |
| parameter | パラメーター | 型パラメーター (type parameter) も同様 |
| type argument | 型引数 | |
| return value | 戻り値 | 「返り値」(34) より「戻り値」(97) |
| result | 結果 | |
| field | フィールド | |
| struct | 構造体 | |
| embedded field | 埋め込みフィールド | |
| promoted | 昇格された | |
| slice | スライス | |
| byte slice | バイトスライス | |
| array | 配列 | |
| map | マップ | |
| channel | チャネル | |
| pointer | ポインター | |
| string | 文字列 | |
| integer | 整数 | |
| unsigned | 符号なし | |
| signed | 符号付き | 暗号の signed は「署名された」 |
| boolean | ブール値 | 「真偽値」も可 |
| floating-point | 浮動小数点数 | |
| constant | 定数 | |
| untyped constant | 型なし定数 | 既存訳では用例が少ない |
| variable | 変数 | |
| alias | エイリアス | |
| generic | ジェネリック | |
| instantiate | インスタンス化する | |
| instance | インスタンス | |
| constraint | 制約 | |
| comparable | 比較可能な | |
| assignable | 代入可能な | |
| underlying | 基礎となる | (118)。「基になる」(37) も見られるが、新規は「基礎となる」に揃える。underlying type は「基底型」でもよい |
| exported | エクスポートされた | |
| unexported | エクスポートされていない | |
| identifier | 識別子 | |
| keyword | キーワード | |
| declaration | 宣言 | |
| statement | 文／ステートメント | go/ast など構文の文脈では「文」。database/sql の prepared statement は「プリペアドステートメント」 |
| expression | 式 | regular expression は「正規表現」 |
| operator | 演算子 | |
| operand | オペランド | |
| assignment | 代入 | |
| scope | スコープ | |
| literal | リテラル | |
| syntax | 構文 | |
| directive | ディレクティブ | |
| package | パッケージ | |
| module | モジュール | |
| interface | インターフェース | |
| implementation | 実装 | |
| implements | 実装します | |
| nil / non-nil | nil／非 nil | |

### ランタイム・並行処理

| 英語 | 推奨訳 | 備考 |
| --- | --- | --- |
| goroutine | ゴルーチン | |
| concurrently | 同時に／並行して | |
| parallel | 並列 | testing の t.Parallel など |
| atomically | アトミックに | |
| compare-and-swap | 比較交換操作 | |
| lock / unlock | ロック／アンロック | |
| mutex | ミューテックス | |
| block | ブロックする | |
| panic (動詞) | パニックを引き起こす | 名詞は「パニック」 |
| runtime | ランタイム | |
| heap | ヒープ | |
| stack | スタック | |
| stack trace | スタックトレース | |
| allocate / allocation | 割り当てる／割り当て | 「アロケーション」より「割り当て」 |
| garbage collector | ガベージコレクター | 「GC」とそのまま書いてもよい |
| thread | スレッド | |
| signal | シグナル | |
| process | プロセス | |
| program counter | プログラムカウンター | |
| profile / profiling | プロファイル／プロファイリング | |
| trace | トレース | |
| metric | メトリクス | |
| context | コンテキスト | |
| cancel / canceled | キャンセル／キャンセルされた | |
| deadline | 期限 | 「デッドライン」より多数派 |
| timeout | タイムアウト | |
| callback | コールバック | |
| iterator | イテレーター | |
| iterate | 反復する | |

### エラー・動作

| 英語 | 推奨訳 | 備考 |
| --- | --- | --- |
| error | エラー | |
| wrap / wrapped | ラップする／ラップされた | |
| unwrap | アンラップする | |
| caller | 呼び出し元 | 「呼び出し側」(25) より「呼び出し元」(200) |
| call / invoke | 呼び出す | |
| behavior | 動作 | 「振る舞い」「挙動」より「動作」 |
| default | デフォルト | |
| valid / invalid | 有効な／無効な | |
| ignored | 無視される | |
| equivalent | 同等 | |
| corresponding | 対応する | |
| associated | 関連付けられた | |
| deprecated | 非推奨 | マーカーの `Deprecated:` は英語のまま |
| compatibility | 互換性 | backward compatibility は「後方互換性」 |
| deterministic | 決定的 | |
| canonical | 正規 | canonical form は「正規形式」 |
| immutable | 不変 | |
| equal | 等しい | |
| documentation | ドキュメント | 「ドキュメンテーション」より「ドキュメント」 |
| doc comment | ドキュメントコメント | |

### I/O・ファイル

| 英語 | 推奨訳 | 備考 |
| --- | --- | --- |
| read / reads | 読み取る／読み込む | |
| write / writes | 書き込む | |
| append | 追加する | |
| flush | フラッシュする | |
| buffer / buffered | バッファー／バッファリングされた | |
| close | 閉じる | |
| seek | シークする | |
| offset | オフセット | |
| truncate | 切り詰める | |
| stream | ストリーム | |
| input / output | 入力／出力 | |
| standard output / standard error | 標準出力／標準エラー | |
| file | ファイル | |
| file system | ファイルシステム | |
| directory | ディレクトリー | |
| working directory | 作業ディレクトリー | |
| path | パス | |
| absolute path / relative path | 絶対パス／相対パス | |
| symbolic link | シンボリックリンク | |
| file descriptor | ファイルディスクリプター | |
| permission bits | パーミッションビット | |
| archive | アーカイブ | |
| entry | エントリー | |
| environment variable | 環境変数 | |
| exit / exit status | 終了／終了ステータス | |

### テキスト・エンコーディング

| 英語 | 推奨訳 | 備考 |
| --- | --- | --- |
| encode / encoding | エンコード／エンコーディング | 「符号化」は使わない |
| decode | デコード | |
| parse | 解析する | (284)。「パース」(84) も見られるが、新規は「解析」に揃える |
| marshal / unmarshal | マーシャル／アンマーシャル | |
| serialize | シリアライズ | |
| format | 形式／フォーマット | 「〜形式の文字列」のように名詞は「形式」、動詞の format は「フォーマットする」 |
| character | 文字 | |
| rune | ルーン | |
| code point | コードポイント | |
| byte | バイト | |
| sequence | シーケンス | |
| whitespace | 空白 | |
| space | スペース／空白 | |
| newline | 改行 | |
| escape | エスケープ | escape sequence は「エスケープシーケンス」 |
| prefix | プレフィックス | (96)。「接頭辞」(32) も可 |
| suffix | 接尾辞 | (38)。「サフィックス」(20) も可 |
| token | トークン | |
| template | テンプレート | |
| pattern | パターン | |
| match | 一致する | 正規表現の文脈では「マッチする」も可 |
| replacement character | 置換文字 | |
| comment | コメント | |
| text | テキスト | |

### ネットワーク・HTTP

| 英語 | 推奨訳 | 備考 |
| --- | --- | --- |
| network | ネットワーク | |
| connection | 接続 | |
| address | アドレス | |
| host | ホスト | |
| port | ポート | |
| dial | ダイヤルする | |
| listener | リスナー | |
| client / server | クライアント／サーバー | |
| request | リクエスト | (508)。「要求」より「リクエスト」 |
| response | レスポンス | (135)。「応答」(61) も見られる |
| header | ヘッダー | |
| body | ボディー | |
| handler | ハンドラー | |
| transport | トランスポート | |
| redirect | リダイレクト | |
| cookie | クッキー | |
| proxy | プロキシー | |
| scheme | スキーム | |
| resolver | リゾルバー | |
| lookup | 検索 | DNS lookup は「DNS 検索」または「ルックアップ」 |
| handshake | ハンドシェイク | |
| session | セッション | |
| status code | ステータスコード | |

### 暗号

| 英語 | 推奨訳 | 備考 |
| --- | --- | --- |
| key | 鍵 | 暗号の文脈では「鍵」。map のキーや一般的な key は「キー」 |
| private key | 秘密鍵 | 「プライベートキー」は使わない |
| public key | 公開鍵 | |
| certificate | 証明書 | |
| signature | 署名 | 関数シグネチャは「シグネチャ」 |
| sign | 署名する | 数値の sign は「符号」 |
| verify | 検証する | |
| encrypt / encryption | 暗号化する／暗号化 | |
| decrypt / decryption | 復号する／復号 | 既存訳では「復号化」が多いが、どちらも可 |
| plaintext | 平文 | |
| ciphertext | 暗号文 | |
| cipher | 暗号 | cipher suite は「暗号スイート」 |
| hash | ハッシュ | |
| checksum | チェックサム | |
| curve | 曲線 | |
| padding | パディング | |
| authentication | 認証 | |
| random / pseudo-random | ランダム／擬似乱数 | 「疑似乱数」ではなく「擬似乱数」 |
| seed | シード | |

### 数値

| 英語 | 推奨訳 | 備考 |
| --- | --- | --- |
| precision | 精度 | |
| rounding | 丸め | rounding mode は「丸めモード」 |
| exponent | 指数 | |
| mantissa | 仮数部 | |
| infinity | 無限大 | 楕円曲線の point at infinity は「無限遠点」 |
| overflow | オーバーフロー | |
| quotient / remainder | 商／余り | |
| modulus | モジュラス | |
| absolute value | 絶対値 | |
| decimal / hexadecimal | 10 進数／16 進数 | |

### テスト

| 英語 | 推奨訳 | 備考 |
| --- | --- | --- |
| test | テスト | |
| benchmark | ベンチマーク | |
| fuzz test | ファズテスト | |
| seed corpus | シードコーパス | |
| coverage | カバレッジ | |
| helper | ヘルパー | |
| cleanup | クリーンアップ | |
| skip | スキップする | |
| fail | 失敗する | |
| log | ログ／記録する | |
