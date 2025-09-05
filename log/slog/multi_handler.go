// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

import (
	"github.com/shogo82148/std/context"
)

// NewMultiHandler creates a [MultiHandler] with the given Handlers.
func NewMultiHandler(handlers ...Handler) *MultiHandler

// MultiHandler is a [Handler] that invokes all the given Handlers.
// Its Enable method reports whether any of the handlers' Enabled methods return true.
// Its Handle, WithAttr and WithGroup methods call the corresponding method on each of the enabled handlers.
type MultiHandler struct {
	multi []Handler
}

func (h *MultiHandler) Enabled(ctx context.Context, l Level) bool

func (h *MultiHandler) Handle(ctx context.Context, r Record) error

func (h *MultiHandler) WithAttrs(attrs []Attr) Handler

func (h *MultiHandler) WithGroup(name string) Handler
