// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Use an external test to avoid os/exec -> internal/testenv -> os/exec
// circular dependency.

package exec_test
