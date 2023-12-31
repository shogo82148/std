// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// slogtestパッケージは、log/slog.Handlerの実装をテストするためのサポートを提供します。
package slogtest

import (
	"github.com/shogo82148/std/log/slog"
	"github.com/shogo82148/std/testing"
)

// TestHandlerは、[slog.Handler] をテストします。
// もしTestHandlerが何かしらの不適切な動作を見つけた場合、それら全てを報告するエラーを返します。
// これらは [errors.Join] を使用して単一のエラーに結合されます。
//
// TestHandlerは、指定されたHandlerを [slog.Logger] にインストールし、
// Loggerの出力メソッドに対して複数回の呼び出しを行います。
//
// results関数は、そのような呼び出しの全ての後で呼び出されます。
// それはmap[string]anyのスライスを返すべきで、Loggerの出力メソッドへの各呼び出しに対して一つです。
// マップのキーと値は、Handlerの出力のキーと値に対応しているべきです。
// 出力内の各グループは、それ自体がネストしたmap[string]anyとして表現されるべきです。
// 標準のキーである [slog.TimeKey]、[slog.LevelKey]、[slog.MessageKey] を使用すべきです。
//
// もしHandlerがJSONを出力するなら、`map[string]any`を引数にして
// [encoding/json.Unmarshal] を呼び出すことで、適切なデータ構造が作成されます。
//
// もしHandlerが意図的にテストでチェックされる属性を削除するなら、
// results関数はその欠如をチェックし、それを返すマップに追加すべきです。
func TestHandler(h slog.Handler, results func() []map[string]any) error

// Runは、[TestHandler]　と同じテストケースで　[slog.Handler]　を実行しますが、
// 各ケースをサブテストで実行します。各テストケースでは、まずnewHandlerを呼び出して
// テスト対象のハンドラのインスタンスを取得し、次にテストケースを実行し、
// 次に結果を取得するためにresultを呼び出します。テストケースが失敗すると、t.Errorを呼び出します。
func Run(t *testing.T, newHandler func(*testing.T) slog.Handler, result func(*testing.T) map[string]any)
