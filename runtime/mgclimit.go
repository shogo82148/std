// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// gcCPULimiter is a mechanism to limit GC CPU utilization in situations
// where it might become excessive and inhibit application progress (e.g.
// a death spiral).
//
// The core of the limiter is a leaky bucket mechanism that fills with GC
// CPU time and drains with mutator time. Because the bucket fills and
// drains with time directly (i.e. without any weighting), this effectively
// sets a very conservative limit of 50%. This limit could be enforced directly,
// however, but the purpose of the bucket is to accommodate spikes in GC CPU
// utilization without hurting throughput.
//
// Note that the bucket in the leaky bucket mechanism can never go negative,
// so the GC never gets credit for a lot of CPU time spent without the GC
// running. This is intentional, as an application that stays idle for, say,
// an entire day, could build up enough credit to fail to prevent a death
// spiral the following day. The bucket's capacity is the GC's only leeway.
//
// The capacity thus also sets the window the limiter considers. For example,
// if the capacity of the bucket is 1 cpu-second, then the limiter will not
// kick in until at least 1 full cpu-second in the last 2 cpu-second window
// is spent on GC CPU time.

// gcCPULimiterUpdatePeriod dictates the maximum amount of wall-clock time
// we can go before updating the limiter.

// capacityPerProc is the limiter's bucket capacity for each P in GOMAXPROCS.

// limiterEventType indicates the type of an event occuring on some P.
//
// These events represent the full set of events that the GC CPU limiter tracks
// to execute its function.
//
// This type may use no more than limiterEventBits bits of information.

// limiterEventTypeMask is a mask for the bits in p.limiterEventStart that represent
// the event type. The rest of the bits of that field represent a timestamp.

// limiterEventStamp is a nanotime timestamp packed with a limiterEventType.

// limiterEvent represents tracking state for an event tracked by the GC CPU limiter.
