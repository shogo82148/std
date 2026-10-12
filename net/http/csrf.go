// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/sync/atomic"
)

// CrossOriginProtectionは、安全でないクロスオリジンのブラウザリクエストを拒否することで、
// [Cross-Site Request Forgery (CSRF)]に対する保護を実装します。
//
// クロスオリジンリクエストは現在、2023年以降すべてのブラウザで利用可能な
// [Sec-Fetch-Site]ヘッダーによって、または [Origin] ヘッダーのホスト名と
// Hostヘッダーを比較することによって検出されます。
//
// GET、HEAD、OPTIONSメソッドは[safe methods]であり、常に許可されます。
// アプリケーションが安全なメソッドのリクエストによって状態を変更する
// アクションを実行しないことが重要です。
//
// Sec-Fetch-SiteヘッダーもOriginヘッダーもないリクエストは、現在、
// 同一オリジンまたはブラウザ以外からのリクエストとみなされ、許可されます。
//
// CrossOriginProtectionのゼロ値は有効であり、信頼されたオリジンや
// バイパスパターンを持ちません。
//
// [Sec-Fetch-Site]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-Fetch-Site
// [Origin]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin
// [Cross-Site Request Forgery (CSRF)]: https://developer.mozilla.org/en-US/docs/Web/Security/Attacks/CSRF
// [safe methods]: https://developer.mozilla.org/en-US/docs/Glossary/Safe/HTTP
type CrossOriginProtection struct {
	bypass    atomic.Pointer[ServeMux]
	trustedMu sync.RWMutex
	trusted   map[string]bool
	deny      atomic.Pointer[Handler]
}

// NewCrossOriginProtectionは新しい [CrossOriginProtection] の値を返します。
func NewCrossOriginProtection() *CrossOriginProtection

// AddTrustedOriginは、指定された値と完全に一致する [Origin] ヘッダーを持つ
// すべてのリクエストを許可します。
//
// Originヘッダーの値は"scheme://host[:port]"の形式です。
//
// AddTrustedOriginは他のメソッドやリクエスト処理と並行して呼び出すことができ、
// 以降のリクエストに適用されます。
//
// [Origin]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin
func (c *CrossOriginProtection) AddTrustedOrigin(origin string) error

// AddInsecureBypassPatternは、指定されたパターンに一致するすべてのリクエストを許可します。
//
// パターンの構文と優先順位のルールは [ServeMux] と同じです。パターンに直接一致する
// リクエストのみが許可されます。ServeMuxがパターンにリダイレクトするリクエスト
// （例えばパスのクリーンアップや末尾のスラッシュの追加の後）は許可されません。
//
// AddInsecureBypassPatternは他のメソッドやリクエスト処理と並行して呼び出すことができ、
// 以降のリクエストに適用されます。
func (c *CrossOriginProtection) AddInsecureBypassPattern(pattern string)

// SetDenyHandlerは、リクエストが拒否されたときに呼び出すハンドラーを設定します。
// デフォルトのエラーハンドラーは403 Forbiddenステータスで応答します。
//
// SetDenyHandlerは他のメソッドやリクエスト処理と並行して呼び出すことができ、
// 以降のリクエストに適用されます。
//
// Checkはエラーハンドラーを呼び出しません。
func (c *CrossOriginProtection) SetDenyHandler(h Handler)

// Checkはリクエストにクロスオリジンチェックを適用します。
// リクエストを拒否すべき場合はエラーを返します。
func (c *CrossOriginProtection) Check(req *Request) error

// Handlerは、ハンドラーhを呼び出す前にクロスオリジンチェックを適用する
// ハンドラーを返します。
//
// リクエストがクロスオリジンチェックに失敗した場合、そのリクエストは
// 403 Forbiddenステータスで拒否されるか、[CrossOriginProtection.SetDenyHandler] に
// 渡されたハンドラーで処理されます。
func (c *CrossOriginProtection) Handler(h Handler) Handler
