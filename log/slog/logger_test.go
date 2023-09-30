// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

// textTimeRE is a regexp to match log timestamps for Text handler.
// This is RFC3339Nano with the fixed 3 digit sub-second precision.

// jsonTimeRE is a regexp to match log timestamps for Text handler.
// This is RFC3339Nano with an arbitrary sub-second precision.

// panicTextAndJsonMarshaler is a type that panics in MarshalText and MarshalJSON.
