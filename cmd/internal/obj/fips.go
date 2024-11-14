// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
FIPS-140 Verification Support

# Overview

For FIPS-140 crypto certification, one of the requirements is that the
“cryptographic module” perform a power-on self-test that includes
verification of its code+data at startup, ostensibly to guard against
corruption. (Like most of FIPS, the actual value here is as questionable
as it is non-negotiable.) Specifically, at startup we need to compute
an HMAC-SHA256 of the cryptographic code+data and compare it against a
build-time HMAC-SHA256 that has been stored in the binary as well.
This obviously guards against accidental corruption only, not attacks.

We could compute an HMAC-SHA256 of the entire binary, but that's more
startup latency than we'd like. (At 500 MB/s, a large 50MB binary
would incur a 100ms hit.) Also, as we'll see, there are some
limitations imposed on the code+data being hashed, and it's nice to
restrict those to the actual cryptographic packages.

# FIPS Symbol Types

Since we're not hashing the whole binary, we need to record the parts
of the binary that contain FIPS code, specifically the part of the
binary corresponding to the crypto/internal/fips package subtree.
To do that, we create special symbol types STEXTFIPS, SRODATAFIPS,
SNOPTRDATAFIPS, and SDATAFIPS, which those packages use instead of
STEXT, SRODATA, SNOPTRDATA, and SDATA. The linker groups symbols by
their type, so that naturally makes the FIPS parts contiguous within a
given type. The linker then writes out in a special symbol the start
and end of each of these FIPS-specific sections, alongside the
expected HMAC-SHA256 of them. At startup, the crypto/internal/fips/check
package has an init function that recomputes the hash and checks it
against the recorded expectation.

The first important functionality in this file, then, is converting
from the standard symbol types to the FIPS symbol types, in the code
that needs them. Every time an LSym.Type is set, code must call
[LSym.setFIPSType] to update the Type to a FIPS type if appropriate.

# Relocation Restrictions

Of course, for the hashes to match, the FIPS code+data written by the
linker has to match the FIPS code+data in memory at init time.
This means that there cannot be an load-time relocations that modify
the FIPS code+data. In a standard -buildmode=exe build, that's vacuously
true, since those binaries have no load-time relocations at all.
For a -buildmode=pie build, there's more to be done.
Specifically, we have to make sure that all the relocations needed are
position-independent, so that they can be applied a link time with no
load-time component. For the code segment (the STEXTFIPS symbols),
that means only using PC-relative relocations. For the data segment,
that means basically having no relocations at all. In particular,
there cannot be R_ADDR relocations.

For example, consider the compilation of code like the global variables:

	var array = [...]int{10, 20, 30}
	var slice = array[:]

The standard implementation of these globals is to fill out the array
values in an SDATA symbol at link time, and then also to fill out the
slice header at link time as {nil, 3, 3}, along with a relocation to
fill in the first word of the slice header with the pointer &array at
load time, once the address of array is known.

A similar issue happens with:

	var slice = []int{10, 20, 30}

The compiler invents an anonymous array and then treats the code as in
the first example. In both cases, a load-time relocation applied
before the crypto/internal/fips/check init function would invalidate
the hash. Instead, we disable the “link time initialization” optimizations
in the compiler (package staticinit) for the fips packages.
That way, the slice initialization is deferred to its own init function.
As long as the package in question imports crypto/internal/fips/check,
the hash check will happen before the package's own init function
runs, and so the hash check will see the slice header written by the
linker, with a slice base pointer predictably nil instead of the
unpredictable &array address.

The details of disabling the static initialization appropriately are
left to the compiler (see ../../compile/internal/staticinit).
This file is only concerned with making sure that no hash-invalidating
relocations sneak into the object files. [LSym.checkFIPSReloc] is called
for every new relocation in a symbol in a FIPS package (as reported by
[Link.IsFIPS]) and rejects invalid relocations.

# FIPS and Non-FIPS Symbols

The cryptographic code+data must be included in the hash-verified
data. In general we accomplish that by putting all symbols from
crypto/internal/fips/... packages into the hash-verified data.
But not all.

Note that wrapper code that layers a Go API atop the cryptographic
core is unverified. For example, crypto/internal/fips/sha256 is part of
the FIPS module and verified but the crypto/sha256 package that wraps
it is outside the module and unverified. Also, runtime support like
the implementation of malloc and garbage collection is outside the
FIPS module. Again, only the core cryptographic code and data is in
scope for the verification.

By analogy with these cases, we treat function wrappers like foo·f
(the function pointer form of func foo) and runtime support data like
runtime type descriptors, generic dictionaries, stack maps, and
function argument data as being outside the FIPS module. That's
important because some of them need to be contiguous with other
non-FIPS data, and all of them include data relocations that would be
incompatible with the hash verification.

# Debugging

Bugs in the handling of FIPS symbols can be mysterious. It is very
helpful to narrow the bug down to a specific symbol that causes a
problem when treated as a FIPS symbol. Rather than work that out
manually, if “go test strings” is failing, then you can use

	go install golang.org/x/tools/cmd/bisect@latest
	bisect -compile=fips go test strings

to automatically bisect which symbol triggers the bug.

# Link-Time Hashing

The link-time hash preparation is out of scope for this file;
see ../../link/internal/ld/fips.go for those details.
*/

package obj

// IsFIPS reports whether we are compiling one of the crypto/internal/fips/... packages.
func (ctxt *Link) IsFIPS() bool

// SetFIPSDebugHash sets the bisect pattern for debugging FIPS changes.
// The compiler calls this with the pattern set by -d=fipshash=pattern,
// so that if FIPS symbol type conversions are causing problems,
// you can use 'bisect -compile fips go test strings' to identify exactly
// which symbol is not being handled correctly.
func SetFIPSDebugHash(pattern string)

// EnableFIPS reports whether FIPS should be enabled at all
// on the current buildcfg GOOS and GOARCH.
func EnableFIPS() bool
