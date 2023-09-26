// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

// certCache implements an intern table for reference counted x509.Certificates,
// implemented in a similar fashion to BoringSSL's CRYPTO_BUFFER_POOL. This
// allows for a single x509.Certificate to be kept in memory and referenced from
// multiple Conns. Returned references should not be mutated by callers. Certificates
// are still safe to use after they are removed from the cache.
//
// Certificates are returned wrapped in a activeCert struct that should be held by
// the caller. When references to the activeCert are freed, the number of references
// to the certificate in the cache is decremented. Once the number of references
// reaches zero, the entry is evicted from the cache.
//
// The main difference between this implementation and CRYPTO_BUFFER_POOL is that
// CRYPTO_BUFFER_POOL is a more  generic structure which supports blobs of data,
// rather than specific structures. Since we only care about x509.Certificates,
// certCache is implemented as a specific cache, rather than a generic one.
//
// See https://boringssl.googlesource.com/boringssl/+/master/include/openssl/pool.h
// and https://boringssl.googlesource.com/boringssl/+/master/crypto/pool/pool.c
// for the BoringSSL reference.

// activeCert is a handle to a certificate held in the cache. Once there are
// no alive activeCerts for a given certificate, the certificate is removed
// from the cache by a finalizer.
