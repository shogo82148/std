// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements type unification.

package types

// A unifier maintains the current type parameters for x and y
// and the respective types inferred for each type parameter.
// A unifier is created by calling newUnifier.

// A tparamsList describes a list of type parameters and the types inferred for them.
