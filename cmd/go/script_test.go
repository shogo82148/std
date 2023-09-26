// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Script-driven tests.
// See testdata/script/README for an overview.

package main_test

// A testScript holds execution state for a single test script.

// scriptCmds are the script command implementations.
// Keep list and the implementations below sorted by name.
//
// NOTE: If you make changes here, update testdata/script/README too!
//

// When expanding shell variables for these commands, we apply regexp quoting to
// expanded strings within the first argument.

// A condition guards execution of a command.

// A command is a complete command parsed from a script.
