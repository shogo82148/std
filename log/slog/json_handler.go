// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/io"
)

// JSONHandler is a Handler that writes Records to an io.Writer as
// line-delimited JSON objects.
type JSONHandler struct {
	*commonHandler
}

// NewJSONHandler creates a JSONHandler that writes to w,
// using the given options.
// If opts is nil, the default options are used.
func NewJSONHandler(w io.Writer, opts *HandlerOptions) *JSONHandler

// Enabled reports whether the handler handles records at the given level.
// The handler ignores records whose level is lower.
func (h *JSONHandler) Enabled(_ context.Context, level Level) bool

// WithAttrs returns a new JSONHandler whose attributes consists
// of h's attributes followed by attrs.
func (h *JSONHandler) WithAttrs(attrs []Attr) Handler

func (h *JSONHandler) WithGroup(name string) Handler

// Handle formats its argument Record as a JSON object on a single line.
//
// If the Record's time is zero, the time is omitted.
// Otherwise, the key is "time"
// and the value is output as with json.Marshal.
//
// If the Record's level is zero, the level is omitted.
// Otherwise, the key is "level"
// and the value of [Level.String] is output.
//
// If the AddSource option is set and source information is available,
// the key is "source", and the value is a record of type [Source].
//
// The message's key is "msg".
//
// To modify these or other attributes, or remove them from the output, use
// [HandlerOptions.ReplaceAttr].
//
// Values are formatted as with an [encoding/json.Encoder] with SetEscapeHTML(false),
// with two exceptions.
//
// First, an Attr whose Value is of type error is formatted as a string, by
// calling its Error method. Only errors in Attrs receive this special treatment,
// not errors embedded in structs, slices, maps or other data structures that
// are processed by the encoding/json package.
//
// Second, an encoding failure does not cause Handle to return an error.
// Instead, the error message is formatted as a string.
//
// Each call to Handle results in a single serialized call to io.Writer.Write.
func (h *JSONHandler) Handle(_ context.Context, r Record) error

// Copied from encoding/json/tables.go.
//
// safeSet holds the value true if the ASCII character with the given array
// position can be represented inside a JSON string without any further
// escaping.
//
// All values are true except for the ASCII control characters (0-31), the
// double quote ("), and the backslash character ("\").
