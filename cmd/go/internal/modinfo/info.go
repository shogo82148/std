// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modinfo

import (
	"github.com/shogo82148/std/cmd/go/internal/modfetch/codehost"
	"github.com/shogo82148/std/time"
)

type ModulePublic struct {
	Path       string           `json:",omitempty"`
	Version    string           `json:",omitempty"`
	Query      string           `json:",omitempty"`
	Versions   []string         `json:",omitempty"`
	Replace    *ModulePublic    `json:",omitempty"`
	Time       *time.Time       `json:",omitempty"`
	Update     *ModulePublic    `json:",omitempty"`
	Main       bool             `json:",omitempty"`
	Indirect   bool             `json:",omitempty"`
	Dir        string           `json:",omitempty"`
	GoMod      string           `json:",omitempty"`
	GoVersion  string           `json:",omitempty"`
	Retracted  []string         `json:",omitempty"`
	Deprecated string           `json:",omitempty"`
	Error      *ModuleError     `json:",omitempty"`
	Sum        string           `json:",omitempty"`
	GoModSum   string           `json:",omitempty"`
	Origin     *codehost.Origin `json:",omitempty"`
	Reuse      bool             `json:",omitempty"`
}

type ModuleError struct {
	Err string
}

// UnmarshalJSON accepts both {"Err":"text"} and "text",
// so that the output of go mod download -json can still
// be unmarshaled into a ModulePublic during -reuse processing.
func (e *ModuleError) UnmarshalJSON(data []byte) error

func (m *ModulePublic) String() string
