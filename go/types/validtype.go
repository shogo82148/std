// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// A tparamEnv provides the environment for looking up the type arguments
// with which type parameters for a given instance were instantiated.
// If we don't have an instance, the corresponding tparamEnv is nil.
