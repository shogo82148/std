// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This Go implementation is derived in part from the reference
// ANSI C implementation, which carries the following notice:
//
//	rijndael-alg-fst.c
//
//	@version 3.0 (December 2000)
//
//	Optimised ANSI C code for the Rijndael cipher (now AES)
//
//	@author Vincent Rijmen <vincent.rijmen@esat.kuleuven.ac.be>
//	@author Antoon Bosselaers <antoon.bosselaers@esat.kuleuven.ac.be>
//	@author Paulo Barreto <paulo.barreto@terra.com.br>
//
//	This code is hereby placed in the public domain.
//
//	THIS SOFTWARE IS PROVIDED BY THE AUTHORS ''AS IS'' AND ANY EXPRESS
//	OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
//	WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
//	ARE DISCLAIMED.  IN NO EVENT SHALL THE AUTHORS OR CONTRIBUTORS BE
//	LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
//	CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
//	SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR
//	BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
//	WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE
//	OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE,
//	EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// See FIPS 197 for specification, and see Daemen and Rijmen's Rijndael submission
// for implementation details.
//	https://csrc.nist.gov/csrc/media/publications/fips/197/final/documents/fips-197.pdf
//	https://csrc.nist.gov/archive/aes/rijndael/Rijndael-ammended.pdf

package aes
