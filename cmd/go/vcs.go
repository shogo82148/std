// Copyright 2012 The Go Authors. All rights reserved.
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

// scpSyntaxRe matches the SCP-like addresses used by Git to access
// repositories by SSH.

// vcsBzr describes how to use Bazaar.

// vcsSvn describes how to use Subversion.

// A vcsPath describes how to convert an import path into a
// version control system and repository name.

// repoRoot represents a version control system, a repo, and a root of
// where to put it on disk.

// securityMode specifies whether a function should make network
// calls using insecure transports (eg, plain text HTTP).
// The zero value is "secure".

// metaImport represents the parsed <meta name="go-import"
// content="prefix vcs reporoot" /> tags from HTML files.

// errNoMatch is returned from matchGoImport when there's no applicable match.

// vcsPaths defines the meaning of import paths referring to
// commonly-used VCS hosting sites (github.com/user/dir)
// and import paths referring to a fully-qualified importPath
// containing a VCS type (foo.com/repo.git/dir)

// vcsPathsAfterDynamic gives additional vcsPaths entries
// to try after the dynamic HTML check.
// This gives those sites a chance to introduce <meta> tags
// as part of a graceful transition away from the hard-coded logic.
