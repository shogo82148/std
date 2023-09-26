// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Support for test coverage with redesigned coverage implementation.

package testing

// cover2 variable stores the current coverage mode and a
// tear-down function to be called at the end of the testing run.
