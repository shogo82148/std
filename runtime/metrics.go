// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// statDep is a dependency on a group of statistics
// that a metric might have.

// statDepSet represents a set of statDeps.
//
// Under the hood, it's a bitmap.

// heapStatsAggregate represents memory stats obtained from the
// runtime. This set of stats is grouped together because they
// depend on each other in some way to make sense of the runtime's
// current heap memory use. They're also sharded across Ps, so it
// makes sense to grab them all at once.

// sysStatsAggregate represents system memory stats obtained
// from the runtime. This set of stats is grouped together because
// they're all relatively cheap to acquire and generally independent
// of one another and other runtime memory stats. The fact that they
// may be acquired at different times, especially with respect to
// heapStatsAggregate, means there could be some skew, but because of
// these stats are independent, there's no real consistency issue here.

// cpuStatsAggregate represents CPU stats obtained from the runtime
// acquired together to avoid skew and inconsistencies.

// statAggregate is the main driver of the metrics implementation.
//
// It contains multiple aggregates of runtime statistics, as well
// as a set of these aggregates that it has populated. The aggergates
// are populated lazily by its ensure method.

// metricKind is a runtime copy of runtime/metrics.ValueKind and
// must be kept structurally identical to that type.

// metricSample is a runtime copy of runtime/metrics.Sample and
// must be kept structurally identical to that type.

// metricValue is a runtime copy of runtime/metrics.Sample and
// must be kept structurally identical to that type.

// metricFloat64Histogram is a runtime copy of runtime/metrics.Float64Histogram
// and must be kept structurally identical to that type.

// agg is used by readMetrics, and is protected by metricsSema.
//
// Managed as a global variable because its pointer will be
// an argument to a dynamically-defined function, and we'd
// like to avoid it escaping to the heap.
