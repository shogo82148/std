// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// A vcsCmd describes how to use a version control system
// like Mercurial, Git, or Subversion.

// A tagCmd describes a command to list available tags
// that can be passed to tagSyncCmd.

// vcsList lists the known version control systems

// vcsHg describes how to use Mercurial.

// vcsGit describes how to use Git.

// vcsBzr describes how to use Bazaar.

// vcsSvn describes how to use Subversion.

// A vcsPath describes how to convert an import path into a
// version control system and repository name.

// repoRoot represents a version control system, a repo, and a root of
// where to put it on disk.

// metaImport represents the parsed <meta name="go-import"
// content="prefix vcs reporoot" /> tags from HTML files.

// errNoMatch is returned from matchGoImport when there's no applicable match.

// vcsPaths lists the known vcs paths.
