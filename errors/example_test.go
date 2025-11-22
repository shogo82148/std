// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors_test

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/io/fs"
	"github.com/shogo82148/std/os"
)

func Example() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
	// Output: 1989-03-15 22:30:00 +0000 UTC: the file system has gone away
}

func ExampleNew() {
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}
	// Output: emit macho dwarf: elf header corrupted
}

var ErrSentinel = errors.New("an error")

// Each call to [errors.New] returns an unique instance of the error,
// even if the arguments are the same. To match against errors
// created by [errors.New], declare a sentinel error and reuse it.
func ExampleNew_unique() {
	err1 := OopsNew()
	err2 := OopsNew()
	fmt.Println("Errors using distinct errors.New calls:")
	fmt.Printf("Is(%q, %q) = %v\n", err1, err2, errors.Is(err1, err2))

	err3 := OopsSentinel()
	err4 := OopsSentinel()
	fmt.Println()
	fmt.Println("Errors using a sentinel error:")
	fmt.Printf("Is(%q, %q) = %v\n", err3, err4, errors.Is(err3, err4))

	// Output:
	// Errors using distinct errors.New calls:
	// Is("an error", "an error") = false
	//
	// Errors using a sentinel error:
	// Is("an error", "an error") = true
}

// The fmt package's Errorf function lets us use the package's formatting
// features to create descriptive error messages.
func ExampleNew_errorf() {
	const name, id = "bimmler", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	if err != nil {
		fmt.Print(err)
	}
	// Output: user "bimmler" (id 17) not found
}

func ExampleJoin() {
	err1 := errors.New("err1")
	err2 := errors.New("err2")
	err := errors.Join(err1, err2)
	fmt.Println(err)
	if errors.Is(err, err1) {
		fmt.Println("err is err1")
	}
	if errors.Is(err, err2) {
		fmt.Println("err is err2")
	}
	fmt.Println(err.(interface{ Unwrap() []error }).Unwrap())
	// Output:
	// err1
	// err2
	// err is err1
	// err is err2
	// [err1 err2]
}

func ExampleIs() {
	if _, err := os.Open("non-existing"); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("file does not exist")
		} else {
			fmt.Println(err)
		}
	}

	// Output:
	// file does not exist
}

// Custom errors can implement a method "Is(error) bool" to match other error values,
// overriding the default matching of [errors.Is].
func ExampleIs_custom_match() {
	var err error = MyIsError{"an error"}
	fmt.Println("Error equals fs.ErrPermission:", err == fs.ErrPermission)
	fmt.Println("Error is fs.ErrPermission:", errors.Is(err, fs.ErrPermission))

	// Output:
	// Error equals fs.ErrPermission: false
	// Error is fs.ErrPermission: true
}

func ExampleAs() {
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}

	// Output:
	// Failed at path: non-existing
}

func ExampleAsType() {
	if _, err := os.Open("non-existing"); err != nil {
		if pathError, ok := errors.AsType[*fs.PathError](err); ok {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
	// Output:
	// Failed at path: non-existing
}

// Custom errors can implement a method "As(any) bool" to match against other error types,
// overriding the default matching of [errors.As].
func ExampleAs_custom_match() {
	var err error = MyAsError{"an error"}
	fmt.Println("Error:", err)
	fmt.Printf("TypeOf err: %T\n", err)

	var pathError *fs.PathError
	ok := errors.As(err, &pathError)
	fmt.Println("Error as fs.PathError:", ok)
	fmt.Println("fs.PathError:", pathError)

	// Output:
	// Error: an error
	// TypeOf err: errors_test.MyAsError
	// Error as fs.PathError: true
	// fs.PathError: custom /: an error
}

// Custom errors can implement a method "As(any) bool" to match against other error types,
// overriding the default matching of [errors.AsType].
func ExampleAsType_custom_match() {
	var err error = MyAsError{"an error"}
	fmt.Println("Error:", err)
	fmt.Printf("TypeOf err: %T\n", err)

	pathError, ok := errors.AsType[*fs.PathError](err)
	fmt.Println("Error as fs.PathError:", ok)
	fmt.Println("fs.PathError:", pathError)

	// Output:
	// Error: an error
	// TypeOf err: errors_test.MyAsError
	// Error as fs.PathError: true
	// fs.PathError: custom /: an error
}

func ExampleUnwrap() {
	err1 := errors.New("error1")
	err2 := fmt.Errorf("error2: [%w]", err1)
	fmt.Println(err2)
	fmt.Println(errors.Unwrap(err2))
	// Output:
	// error2: [error1]
	// error1
}
