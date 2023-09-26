// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package slogtest implements support for testing implementations of log/slog.Handler.
package slogtest

import (
	"github.com/shogo82148/std/log/slog"
	"github.com/shogo82148/std/testing"
)

// TestHandler tests a [slog.Handler].
// If TestHandler finds any misbehaviors, it returns an error for each,
// combined into a single error with errors.Join.
//
// TestHandler installs the given Handler in a [slog.Logger] and
// makes several calls to the Logger's output methods.
// The Handler should be enabled for levels Info and above.
//
// The results function is invoked after all such calls.
// It should return a slice of map[string]any, one for each call to a Logger output method.
// The keys and values of the map should correspond to the keys and values of the Handler's
// output. Each group in the output should be represented as its own nested map[string]any.
// The standard keys slog.TimeKey, slog.LevelKey and slog.MessageKey should be used.
//
// If the Handler outputs JSON, then calling [encoding/json.Unmarshal] with a `map[string]any`
// will create the right data structure.
//
// If a Handler intentionally drops an attribute that is checked by a test,
// then the results function should check for its absence and add it to the map it returns.
func TestHandler(h slog.Handler, results func() []map[string]any) error

// Run exercises a [slog.Handler] on the same test cases as [TestHandler], but
// runs each case in a subtest. For each test case, it first calls newHandler to
// get an instance of the handler under test, then runs the test case, then
// calls result to get the result. If the test case fails, it calls t.Error.
func Run(t *testing.T, newHandler func(*testing.T) slog.Handler, result func(*testing.T) map[string]any)
