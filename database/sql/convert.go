// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Type conversions for Scan.

package sql

// ccChecker wraps the driver.ColumnConverter and allows it to be used
// as if it were a NamedValueChecker. If the driver ColumnConverter
// is not present then the NamedValueChecker will return driver.ErrSkip.
