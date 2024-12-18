// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tls13 implements the TLS 1.3 Key Schedule as specified in RFC 8446,
// Section 7.1 and allowed by FIPS 140-3 IG 2.4.B Resolution 7.
package tls13

import (
	"github.com/shogo82148/std/crypto/internal/fips140"
)

// ExpandLabel implements HKDF-Expand-Label from RFC 8446, Section 7.1.
func ExpandLabel[H fips140.Hash](hash func() H, secret []byte, label string, context []byte, length int) []byte

type EarlySecret struct {
	secret []byte
	hash   func() fips140.Hash
}

func NewEarlySecret[H fips140.Hash](hash func() H, psk []byte) *EarlySecret

func (s *EarlySecret) ResumptionBinderKey() []byte

// ClientEarlyTrafficSecret derives the client_early_traffic_secret from the
// early secret and the transcript up to the ClientHello.
func (s *EarlySecret) ClientEarlyTrafficSecret(transcript fips140.Hash) []byte

type HandshakeSecret struct {
	secret []byte
	hash   func() fips140.Hash
}

func (s *EarlySecret) HandshakeSecret(sharedSecret []byte) *HandshakeSecret

// ClientHandshakeTrafficSecret derives the client_handshake_traffic_secret from
// the handshake secret and the transcript up to the ServerHello.
func (s *HandshakeSecret) ClientHandshakeTrafficSecret(transcript fips140.Hash) []byte

// ServerHandshakeTrafficSecret derives the server_handshake_traffic_secret from
// the handshake secret and the transcript up to the ServerHello.
func (s *HandshakeSecret) ServerHandshakeTrafficSecret(transcript fips140.Hash) []byte

type MasterSecret struct {
	secret []byte
	hash   func() fips140.Hash
}

func (s *HandshakeSecret) MasterSecret() *MasterSecret

// ClientApplicationTrafficSecret derives the client_application_traffic_secret_0
// from the master secret and the transcript up to the server Finished.
func (s *MasterSecret) ClientApplicationTrafficSecret(transcript fips140.Hash) []byte

// ServerApplicationTrafficSecret derives the server_application_traffic_secret_0
// from the master secret and the transcript up to the server Finished.
func (s *MasterSecret) ServerApplicationTrafficSecret(transcript fips140.Hash) []byte

// ResumptionMasterSecret derives the resumption_master_secret from the master secret
// and the transcript up to the client Finished.
func (s *MasterSecret) ResumptionMasterSecret(transcript fips140.Hash) []byte

type ExporterMasterSecret struct {
	secret []byte
	hash   func() fips140.Hash
}

// ExporterMasterSecret derives the exporter_master_secret from the master secret
// and the transcript up to the server Finished.
func (s *MasterSecret) ExporterMasterSecret(transcript fips140.Hash) *ExporterMasterSecret

// EarlyExporterMasterSecret derives the exporter_master_secret from the early secret
// and the transcript up to the ClientHello.
func (s *EarlySecret) EarlyExporterMasterSecret(transcript fips140.Hash) *ExporterMasterSecret

func (s *ExporterMasterSecret) Exporter(label string, context []byte, length int) []byte

func TestingOnlyExporterSecret(s *ExporterMasterSecret) []byte
