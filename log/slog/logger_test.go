// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

import (
	loginternal "log/internal"
)

// panicTextAndJsonMarshaler is a type that panics in MarshalText and MarshalJSON.
