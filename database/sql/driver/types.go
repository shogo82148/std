// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package driver

// ValueConverter is the interface providing the ConvertValue method.
//
// Various implementations of ValueConverter are provided by the
// driver package to provide consistent implementations of conversions
// between drivers. The ValueConverters have several uses:
//
//   - converting from the [Value] types as provided by the sql package
//     into a database table's specific column type and making sure it
//     fits, such as making sure a particular int64 fits in a
//     table's uint16 column.
//
//   - converting a value as given from the database into one of the
//     driver [Value] types.
//
//   - by the [database/sql] package, for converting from a driver's [Value] type
//     to a user's type in a scan.
type ValueConverter interface {
	ConvertValue(v any) (Value, error)
}

// Valuer is the interface providing the Value method.
//
// Errors returned by the [Value] method are wrapped by the database/sql package.
// This allows callers to use [errors.Is] for precise error handling after operations
// like [database/sql.Query], [database/sql.Exec], or [database/sql.QueryRow].
//
// Types implementing Valuer interface are able to convert
// themselves to a driver [Value].
type Valuer interface {
	Value() (Value, error)
}

// Bool is a [ValueConverter] that converts input values to bool.
//
// The conversion rules are:
//   - booleans are returned unchanged
//   - for integer types,
//     1 is true
//     0 is false,
//     other integers are an error
//   - for strings and []byte, same rules as [strconv.ParseBool]
//   - all other types are an error
var Bool boolType

var _ ValueConverter = boolType{}

// Int32 is a [ValueConverter] that converts input values to int64,
// respecting the limits of an int32 value.
var Int32 int32Type

var _ ValueConverter = int32Type{}

// String is a [ValueConverter] that converts its input to a string.
// If the value is already a string or []byte, it's unchanged.
// If the value is of another type, conversion to string is done
// with fmt.Sprintf("%v", v).
var String stringType

// Null is a type that implements [ValueConverter] by allowing nil
// values but otherwise delegating to another [ValueConverter].
type Null struct {
	Converter ValueConverter
}

func (n Null) ConvertValue(v any) (Value, error)

// NotNull is a type that implements [ValueConverter] by disallowing nil
// values but otherwise delegating to another [ValueConverter].
type NotNull struct {
	Converter ValueConverter
}

func (n NotNull) ConvertValue(v any) (Value, error)

// IsValue reports whether v is a valid [Value] parameter type.
func IsValue(v any) bool

// IsScanValue is equivalent to [IsValue].
// It exists for compatibility.
func IsScanValue(v any) bool

// DefaultParameterConverter is the default implementation of
// [ValueConverter] that's used when a [Stmt] doesn't implement
// [ColumnConverter].
//
// DefaultParameterConverter returns its argument directly if
// IsValue(arg). Otherwise, if the argument implements [Valuer], its
// Value method is used to return a [Value]. As a fallback, the provided
// argument's underlying type is used to convert it to a [Value]:
// underlying integer types are converted to int64, floats to float64,
// bool, string, and []byte to themselves. If the argument is a nil
// pointer, defaultConverter.ConvertValue returns a nil [Value].
// If the argument is a non-nil pointer, it is dereferenced and
// defaultConverter.ConvertValue is called recursively. Other types
// are an error.
var DefaultParameterConverter defaultConverter

var _ ValueConverter = defaultConverter{}
