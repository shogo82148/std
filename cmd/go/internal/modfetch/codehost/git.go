// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package codehost

import (
	"github.com/shogo82148/std/context"
)

// LocalGitRepo is like Repo but accepts both Git remote references
// and paths to repositories on the local file system.
func LocalGitRepo(ctx context.Context, remote string) (Repo, error)
