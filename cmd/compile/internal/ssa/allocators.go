// Code generated from _gen/allocators.go using 'go generate'; DO NOT EDIT.

package ssa

func (c *Cache) AllocValueSlice(n int) []*Value

func (c *Cache) FreeValueSlice(s []*Value)

func (c *Cache) AllocLimitSlice(n int) []Limit

func (c *Cache) FreeLimitSlice(s []Limit)

func (c *Cache) AllocSparseSet(n int) *SparseSet

func (c *Cache) FreeSparseSet(s *SparseSet)

func (c *Cache) AllocSparseMap(n int) *sparseMap

func (c *Cache) FreeSparseMap(s *sparseMap)

func (c *Cache) AllocSparseMapPos(n int) *SparseMapPos

func (c *Cache) FreeSparseMapPos(s *SparseMapPos)

func (c *Cache) AllocBlockSlice(n int) []*Block

func (c *Cache) FreeBlockSlice(s []*Block)

func (c *Cache) AllocInt64(n int) []int64

func (c *Cache) FreeInt64(s []int64)

func (c *Cache) AllocIntSlice(n int) []int

func (c *Cache) FreeIntSlice(s []int)

func (c *Cache) AllocInt32Slice(n int) []int32

func (c *Cache) FreeInt32Slice(s []int32)

func (c *Cache) AllocInt8Slice(n int) []int8

func (c *Cache) FreeInt8Slice(s []int8)

func (c *Cache) AllocBoolSlice(n int) []bool

func (c *Cache) FreeBoolSlice(s []bool)

func (c *Cache) AllocIDSlice(n int) []ID

func (c *Cache) FreeIDSlice(s []ID)

func (c *Cache) AllocUintSlice(n int) []uint

func (c *Cache) FreeUintSlice(s []uint)

func (c *Cache) AllocKnownBitsEntriesSlice(n int) []knownBitsEntry

func (c *Cache) FreeKnownBitsEntriesSlice(s []knownBitsEntry)
