// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package modget implements the module-aware “go get” command.
package modget

import (
	"github.com/shogo82148/std/cmd/go/internal/base"
)

var CmdGet = &base.Command{

	UsageLine: "go get [-t] [-u] [-v] [build flags] [packages]",
	Short:     "add dependencies to current module and install them",
	Long: `
Get resolves its command-line arguments to packages at specific module versions,
updates go.mod to require those versions, and downloads source code into the
module cache.

To add a dependency for a package or upgrade it to its latest version:

	go get example.com/pkg

To upgrade or downgrade a package to a specific version:

	go get example.com/pkg@v1.2.3

To remove a dependency on a module and downgrade modules that require it:

	go get example.com/mod@none

To upgrade the minimum required Go version to the latest released Go version:

	go get go@latest

To upgrade the Go toolchain to the latest patch release of the current Go toolchain:

	go get toolchain@patch

See https://golang.org/ref/mod#go-get for details.

In earlier versions of Go, 'go get' was used to build and install packages.
Now, 'go get' is dedicated to adjusting dependencies in go.mod. 'go install'
may be used to build and install commands instead. When a version is specified,
'go install' runs in module-aware mode and ignores the go.mod file in the
current directory. For example:

	go install example.com/pkg@v1.2.3
	go install example.com/pkg@latest

See 'go help install' or https://golang.org/ref/mod#go-install for details.

'go get' accepts the following flags.

The -t flag instructs get to consider modules needed to build tests of
packages specified on the command line.

The -u flag instructs get to update modules providing dependencies
of packages named on the command line to use newer minor or patch
releases when available.

The -u=patch flag (not -u patch) also instructs get to update dependencies,
but changes the default to select patch releases.

When the -t and -u flags are used together, get will update
test dependencies as well.

The -x flag prints commands as they are executed. This is useful for
debugging version control commands when a module is downloaded directly
from a repository.

For more about build flags, see 'go help build'.

For more about modules, see https://golang.org/ref/mod.

For more about using 'go get' to update the minimum Go version and
suggested Go toolchain, see https://go.dev/doc/toolchain.

For more about specifying packages, see 'go help packages'.

This text describes the behavior of get using modules to manage source
code and dependencies. If instead the go command is running in GOPATH
mode, the details of get's flags and effects change, as does 'go help get'.
See 'go help gopath-get'.

See also: go build, go install, go clean, go mod.
	`,
}

var HelpVCS = &base.Command{
	UsageLine: "vcs",
	Short:     "controlling version control with GOVCS",
	Long: `
The 'go get' command can run version control commands like git
to download imported code. This functionality is critical to the decentralized
Go package ecosystem, in which code can be imported from any server,
but it is also a potential security problem, if a malicious server finds a
way to cause the invoked version control command to run unintended code.

To balance the functionality and security concerns, the 'go get' command
by default will only use git and hg to download code from public servers.
But it will use any known version control system (bzr, fossil, git, hg, svn)
to download code from private servers, defined as those hosting packages
matching the GOPRIVATE variable (see 'go help private'). The rationale behind
allowing only Git and Mercurial is that these two systems have had the most
attention to issues of being run as clients of untrusted servers. In contrast,
Bazaar, Fossil, and Subversion have primarily been used in trusted,
authenticated environments and are not as well scrutinized as attack surfaces.

The version control command restrictions only apply when using direct version
control access to download code. When downloading modules from a proxy,
'go get' uses the proxy protocol instead, which is always permitted.
By default, the 'go get' command uses the Go module mirror (proxy.golang.org)
for public packages and only falls back to version control for private
packages or when the mirror refuses to serve a public package (typically for
legal reasons). Therefore, clients can still access public code served from
Bazaar, Fossil, or Subversion repositories by default, because those downloads
use the Go module mirror, which takes on the security risk of running the
version control commands using a custom sandbox.

The GOVCS variable can be used to change the allowed version control systems
for specific packages (identified by a module or import path).
The GOVCS variable applies when building package in both module-aware mode
and GOPATH mode. When using modules, the patterns match against the module path.
When using GOPATH, the patterns match against the import path corresponding to
the root of the version control repository.

The general form of the GOVCS setting is a comma-separated list of
pattern:vcslist rules. The pattern is a glob pattern that must match
one or more leading elements of the module or import path. The vcslist
is a pipe-separated list of allowed version control commands, or "all"
to allow use of any known command, or "off" to disallow all commands.
Note that if a module matches a pattern with vcslist "off", it may still be
downloaded if the origin server uses the "mod" scheme, which instructs the
go command to download the module using the GOPROXY protocol.
The earliest matching pattern in the list applies, even if later patterns
might also match.

For example, consider:

	GOVCS=github.com:git,evil.com:off,*:git|hg

With this setting, code with a module or import path beginning with
github.com/ can only use git; paths on evil.com cannot use any version
control command, and all other paths (* matches everything) can use
only git or hg.

The special patterns "public" and "private" match public and private
module or import paths. A path is private if it matches the GOPRIVATE
variable; otherwise it is public.

If no rules in the GOVCS variable match a particular module or import path,
the 'go get' command applies its default rule, which can now be summarized
in GOVCS notation as 'public:git|hg,private:all'.

To allow unfettered use of any version control system for any package, use:

	GOVCS=*:all

To disable all use of version control, use:

	GOVCS=*:off

The 'go env -w' command (see 'go help env') can be used to set the GOVCS
variable for future go command invocations.
`,
}
