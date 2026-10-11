// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

import (
	"github.com/shogo82148/std/context"
)

// NewMultiHandlerは、指定されたHandlerを持つ[MultiHandler]を作成します。
func NewMultiHandler(handlers ...Handler) *MultiHandler

// MultiHandlerは、指定されたすべてのHandlerを呼び出す[Handler]です。
// そのEnabledメソッドは、いずれかのハンドラのEnabledメソッドがtrueを返すかどうかを報告します。
// そのHandle、WithAttrs、WithGroupメソッドは、有効な各ハンドラの対応するメソッドを呼び出します。
type MultiHandler struct {
	multi []Handler
}

func (h *MultiHandler) Enabled(ctx context.Context, l Level) bool

func (h *MultiHandler) Handle(ctx context.Context, r Record) error

func (h *MultiHandler) WithAttrs(attrs []Attr) Handler

func (h *MultiHandler) WithGroup(name string) Handler
