// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Native Client SRPC message passing.
// This code is needed to invoke SecureRandom, the NaCl equivalent of /dev/random.

package syscall

// An srpcClient represents the client side of an SRPC connection.

// An srpcService is a single method that the server offers.

// An srpc represents a single srpc issued by a client.

// The current protocol number.
// Kind of useless, since there have been backwards-incompatible changes
// to the wire protocol that did not update the protocol number.
// At this point it's really just a sanity check.

// An srpcErrno is an SRPC status code.

// A msgHdr is the data argument to the imc_recvmsg
// and imc_sendmsg system calls.

// A single region for I/O.

// A msgReceiver receives messages from a file descriptor.

// A msgSender sends messages on a file descriptor.

// A msg is the Go representation of an SRPC message.

// At startup, connect to the name service.
