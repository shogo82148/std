// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
FIPS-140 Verification Support

See ../../../internal/obj/fips.go for a basic overview.
This file is concerned with computing the hash of the FIPS code+data.
Package obj has taken care of marking the FIPS symbols with the
special types STEXTFIPS, SRODATAFIPS, SNOPTRDATAFIPS, and SDATAFIPS.

# FIPS Symbol Layout

The first order of business is collecting the FIPS symbols into
contiguous sections of the final binary and identifying the start and
end of those sections. The linker already tracks the start and end of
the text section as runtime.text and runtime.etext, and similarly for
other sections, but the implementation of those symbols is tricky and
platform-specific. The problem is that they are zero-length
pseudo-symbols that share addresses with other symbols, which makes
everything harder. For the FIPS sections, we avoid that subtlety by
defining actual non-zero-length symbols bracketing each section and
use those symbols as the boundaries.

Specifically, we define a 1-byte symbol go:textfipsstart of type
STEXTFIPSSTART and a 1-byte symbol go:textfipsend of type STEXTFIPSEND,
and we place those two symbols immediately before and after the
STEXTFIPS symbols. We do the same for SRODATAFIPS, SNOPTRDATAFIPS,
and SDATAFIPS. Because the symbols are real (but otherwise unused) data,
they can be treated as normal symbols for symbol table purposes and
don't need the same kind of special handling that runtime.text and
friends do.

Note that treating the FIPS text as starting at &go:textfipsstart and
ending at &go:textfipsend means that go:textfipsstart is included in
the verified data while go:textfipsend is not. That's fine: they are
only framing and neither strictly needs to be in the hash.

The new special symbols are created by [loadfips].

# FIPS Info Layout

Having collated the FIPS symbols, we need to compute the hash
and then leave both the expected hash and the FIPS address ranges
for the run-time check in crypto/internal/fips140/check.
We do that by creating a special symbol named go:fipsinfo of the form

	struct {
		sum   [32]byte
		self  uintptr // points to start of struct
		sects [4]struct{
			start uintptr
			end   uintptr
		}
	}

The crypto/internal/fips140/check uses linkname to access this symbol,
which is of course not included in the hash.

# FIPS Info Calculation

When using internal linking, [asmbfips] runs after writing the output
binary but before code-signing it. It reads the relevant sections
back from the output file, hashes them, and then writes the go:fipsinfo
content into the output file.

When using external linking, especially with -buildmode=pie, we cannot
predict the specific PLT index references that the linker will insert
into the FIPS code sections, so we must read the final linked executable
after external linking, compute the hash, and then write it back to the
executable in the go:fipsinfo sum field. [hostlinkfips] does this.
It finds go:fipsinfo easily because that symbol is given its own section
(.go.fipsinfo on ELF, __go_fipsinfo on Mach-O), and then it can use the
sections field to find the relevant parts of the executable, hash them,
and fill in sum.

Both [asmbfips] and [hostlinkfips] need the same hash calculation code.
The [fipsObj] type provides that calculation.

# Debugging

It is of course impossible to debug a mismatched hash directly:
two random 32-byte strings differ. For debugging, the linker flag
-fipso can be set to the name of a file (such as /tmp/fips.o)
where the linker will write the “FIPS object” that is being hashed.

There is also commented-out code in crypto/internal/fips140/check that
will write /tmp/fipscheck.o during the run-time verification.

When the hashes differ, the first step is to uncomment the
/tmp/fipscheck.o-writing code and then rebuild with
-ldflags=-fipso=/tmp/fips.o. Then when the hash check fails,
compare /tmp/fips.o and /tmp/fipscheck.o to find the differences.
*/

package ld
