// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xml

import (
	"github.com/shogo82148/std/reflect"
)

// typeInfo holds details for the xml representation of a type.

// fieldInfo holds details for the xml representation of a single field.

// A TagPathError represents an error in the unmarshalling process
// caused by the use of field tags with conflicting paths.
type TagPathError struct {
	Struct       reflect.Type
	Field1, Tag1 string
	Field2, Tag2 string
}

func (e *TagPathError) Error() string
