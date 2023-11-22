// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package traceviewer

import (
	"github.com/shogo82148/std/net/http"
)

func MainHandler(views []View) http.Handler

const CommonStyle = `
/* See https://github.com/golang/pkgsite/blob/master/static/shared/typography/typography.css */
body {
  font-family:	-apple-system, BlinkMacSystemFont, 'Segoe UI', Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji';
  font-size:	1rem;
  line-height:	normal;
  max-width:	9in;
  margin:	1em;
}
h1 { font-size: 1.5rem; }
h2 { font-size: 1.375rem; }
h1,h2 {
  font-weight: 600;
  line-height: 1.25em;
  word-break: break-word;
}
p  { color: grey85; font-size:85%; }
code,
pre,
textarea.code {
  font-family: SFMono-Regular, Consolas, 'Liberation Mono', Menlo, monospace;
  font-size: 0.875rem;
  line-height: 1.5em;
}

pre,
textarea.code {
  background-color: var(--color-background-accented);
  border: var(--border);
  border-radius: var(--border-radius);
  color: var(--color-text);
  overflow-x: auto;
  padding: 0.625rem;
  tab-size: 4;
  white-space: pre;
}
`

type View struct {
	Type   ViewType
	Ranges []Range
}

type ViewType string

const (
	ViewProc   ViewType = "proc"
	ViewThread ViewType = "thread"
)

func (v View) URL(rangeIdx int) string

type Range struct {
	Name      string
	Start     int
	End       int
	StartTime int64
	EndTime   int64
}

func (r Range) URL(viewType ViewType) string

func TraceHandler() http.Handler

func StaticHandler() http.Handler
