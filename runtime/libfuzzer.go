// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build libfuzzer

package runtime

// Keep in sync with the definition of ret_sled in src/runtime/libfuzzer_amd64.s

//go:linkname __sanitizer_cov_trace_cmp1 __sanitizer_cov_trace_cmp1
//go:cgo_import_static __sanitizer_cov_trace_cmp1

//go:linkname __sanitizer_cov_trace_cmp2 __sanitizer_cov_trace_cmp2
//go:cgo_import_static __sanitizer_cov_trace_cmp2

//go:linkname __sanitizer_cov_trace_cmp4 __sanitizer_cov_trace_cmp4
//go:cgo_import_static __sanitizer_cov_trace_cmp4

//go:linkname __sanitizer_cov_trace_cmp8 __sanitizer_cov_trace_cmp8
//go:cgo_import_static __sanitizer_cov_trace_cmp8

//go:linkname __sanitizer_cov_trace_const_cmp1 __sanitizer_cov_trace_const_cmp1
//go:cgo_import_static __sanitizer_cov_trace_const_cmp1

//go:linkname __sanitizer_cov_trace_const_cmp2 __sanitizer_cov_trace_const_cmp2
//go:cgo_import_static __sanitizer_cov_trace_const_cmp2

//go:linkname __sanitizer_cov_trace_const_cmp4 __sanitizer_cov_trace_const_cmp4
//go:cgo_import_static __sanitizer_cov_trace_const_cmp4

//go:linkname __sanitizer_cov_trace_const_cmp8 __sanitizer_cov_trace_const_cmp8
//go:cgo_import_static __sanitizer_cov_trace_const_cmp8

//go:linkname __sanitizer_cov_8bit_counters_init __sanitizer_cov_8bit_counters_init
//go:cgo_import_static __sanitizer_cov_8bit_counters_init

//go:linkname __start___sancov_cntrs __start___sancov_cntrs
//go:cgo_import_static __start___sancov_cntrs

//go:linkname __stop___sancov_cntrs __stop___sancov_cntrs
//go:cgo_import_static __stop___sancov_cntrs

//go:linkname __sanitizer_cov_pcs_init __sanitizer_cov_pcs_init
//go:cgo_import_static __sanitizer_cov_pcs_init

//go:linkname __sanitizer_weak_hook_strcmp __sanitizer_weak_hook_strcmp
//go:cgo_import_static __sanitizer_weak_hook_strcmp
