// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

import (
	"crypto/x509"
)

// recordingConn is a net.Conn that records the traffic that passes through it.
// WriteTo can be used to produce output that can be later be loaded with
// ParseTestData.

// localListener is set up by TestMain and used by localPipe to create Conn
// pairs like net.Pipe, but connected by an actual buffered TCP connection.

// zeroSource is an io.Reader that returns an unlimited number of zero bytes.

// testRSAPSSCertificate has signatureAlgorithm rsassaPss, but subjectPublicKeyInfo
// algorithm rsaEncryption, for use with the rsa_pss_rsae_* SignatureSchemes.
// See also TestRSAPSSKeyError. testRSAPSSCertificate is self-signed.

var _ = x509.ParsePKCS1PrivateKey(fromHex("3082025b02010002818100db467d932e12270648bc062821ab7ec4b6a25dfe1e5245887a3647a5080d92425bc281c0be97799840fb4f6d14fd2b138bc2a52e67d8d4099ed62238b74a0b74732bc234f1d193e596d9747bf3589f6c613cc0b041d4d92b2b2423775b1c3bbd755dce2054cfa163871d1e24c4f31d1a508baab61443ed97a77562f414c852d702030100010281800b07fbcf48b50f1388db34b016298b8217f2092a7c9a04f77db6775a3d1279b62ee9951f7e371e9de33f015aea80660760b3951dc589a9f925ed7de13e8f520e1ccbc7498ce78e7fab6d59582c2386cc07ed688212a576ff37833bd5943483b5554d15a0b9b4010ed9bf09f207e7e9805f649240ed6c1256ed75ab7cd56d9671024100fded810da442775f5923debae4ac758390a032a16598d62f059bb2e781a9c2f41bfa015c209f966513fe3bf5a58717cbdb385100de914f88d649b7d15309fa49024100dd10978c623463a1802c52f012cfa72ff5d901f25a2292446552c2568b1840e49a312e127217c2186615aae4fb6602a4f6ebf3f3d160f3b3ad04c592f65ae41f02400c69062ca781841a09de41ed7a6d9f54adc5d693a2c6847949d9e1358555c9ac6a8d9e71653ac77beb2d3abaf7bb1183aa14278956575dbebf525d0482fd72d90240560fe1900ba36dae3022115fd952f2399fb28e2975a1c3e3d0b679660bdcb356cc189d611cfdd6d87cd5aea45aa30a2082e8b51e94c2f3dd5d5c6036a8a615ed0240143993d80ece56f877cb80048335701eb0e608cc0c1ca8c2227b52edf8f1ac99c562f2541b5ce81f0515af1c5b4770dba53383964b4b725ff46fdec3d08907df"))

var _ = x509.ParseECPrivateKey(fromHex("3081dc0201010442019883e909ad0ac9ea3d33f9eae661f1785206970f8ca9a91672f1eedca7a8ef12bd6561bb246dda5df4b4d5e7e3a92649bc5d83a0bf92972e00e62067d0c7bd99d7a00706052b81040023a18189038186000400c4a1edbe98f90b4873367ec316561122f23d53c33b4d213dcd6b75e6f6b0dc9adf26c1bcb287f072327cb3642f1c90bcea6823107efee325c0483a69e0286dd33700ef0462dd0da09c706283d881d36431aa9e9731bd96b068c09b23de76643f1a5c7fe9120e5858b65f70dd9bd8ead5d7f5d5ccb9b69f30665b669a20e227e5bffe3b"))

var _ = x509.ParseECPrivateKey(fromHex("30770201010420012f3b52bc54c36ba3577ad45034e2e8efe1e6999851284cb848725cfe029991a00a06082a8648ce3d030107a14403420004c02c61c9b16283bbcc14956d886d79b358aa614596975f78cece787146abf74c2d5dc578c0992b4f3c631373479ebf3892efe53d21c4f4f1cc9a11c3536b7f75"))
