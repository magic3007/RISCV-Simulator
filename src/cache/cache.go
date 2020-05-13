package cache

import (
	"math"
)

type WriteUpPolicyType uint
type WriteBackPolicyType uint
type EvictPolicyType uint

const (
	WriteThrough WriteUpPolicyType = iota
	WriteBack
)

const (
	WriteAllocate WriteBackPolicyType = iota
	NoWriteAllocate
)

const (
	LRU 	EvictPolicyType = iota
	LFU
)

type CacheConfig struct {
	LineSize		uint
	Associativity	uint
	CacheSize		uint
	WriteUpPolicy	WriteUpPolicyType
	WriteBackPolicy WriteBackPolicyType
	EvictPolicy		EvictPolicyType
}

func (c *CacheConfig) SetNum() uint{
	return c.CacheSize / c.LineSize / c.Associativity
}

func (c *CacheConfig) OffsetBits() uint{
	return uint(math.Log2(float64(c.LineSize)))
}

func (c *CacheConfig) IndexBits() uint{
	return uint(math.Log2(float64(c.SetNum())))
}

func (c *CacheConfig) Index(addr uint64) uint64{
	return (addr >> c.OffsetBits()) & ((uint64(1) << c.IndexBits()) - 1)
}

func (c *CacheConfig) Tag(addr uint64) uint64{
	return addr >> (c.IndexBits() + c.OffsetBits())
}

type CacheLine struct{
	Addr			uint64
	Tag				uint64
	IsDirty 		bool
	IsValid 		bool
	TimeStamp		uint
	Counter			uint
}

func (c *CacheLine) Update(addr uint64, ExtractTagFunc func(uint64)uint64){
	c.Addr = addr
	c.Tag = ExtractTagFunc(addr)
	c.IsDirty = false
	c.IsValid = true
	c.Counter = 0
}

type Cache struct{
	Stats     StorageStats
	latency_  StorageLatency
	config_   CacheConfig
	lower_    Storage
	lines_    [][]CacheLine
	TimeStamp uint
}

func NewCache(latency StorageLatency, config CacheConfig, lower Storage) *Cache {
	c := new(Cache)
	c.latency_ = latency
	c.config_ = config
	c.lower_ = lower
	c.lines_ = make([][]CacheLine, c.config_.SetNum())
	for i := range c.lines_ {
		c.lines_[i] = make([]CacheLine, c.config_.Associativity)
	}
	return c
}

func (c *Cache) FindCacheLine(addr uint64) (*CacheLine, bool) {
	index := c.config_.Index(addr)
	tag := c.config_.Tag(addr)
	for i := range c.lines_[index] {
		if c.lines_[index][i].Tag == tag && c.lines_[index][i].IsValid{
			return &c.lines_[index][i], true
		}
	}
	return nil, false
}


func (c *Cache) Insert(addr uint64) (evicteeAddr uint64, IsRequirePushDown bool) {
	index := c.config_.Index(addr)
	evictee := &c.lines_[index][0]
	if c.config_.EvictPolicy == LRU {
		for i := range c.lines_[index] {
			if c.lines_[index][i].TimeStamp < evictee.TimeStamp {
				evictee = &c.lines_[index][i]
			}
		}
	}else if c.config_.EvictPolicy == LFU {
		for i := range c.lines_[index] {
			if c.lines_[index][i].Counter < evictee.Counter {
				evictee = &c.lines_[index][i]
			}
		}
	}else{
		panic("unknown eviction policy")
	}
	evicteeAddr = evictee.Addr
	IsRequirePushDown = evictee.IsValid && evictee.IsDirty
	evictee.Update(addr, c.config_.Tag)
	return
}

func (c *Cache) Read(addr uint64) uint{
	c.TimeStamp += 1
	time := c.latency_.HitLatency
	line, isFound := c.FindCacheLine(addr)
	if !isFound{
		c.Stats.MissNum += 1
		if evicteeAddr, IsRequirePushDown := c.Insert(addr); IsRequirePushDown{
			time += c.lower_.HandleRequest(evicteeAddr, uint(1), Write)
		}
		time += c.lower_.HandleRequest(addr, uint(1), Read)
		line, isFound = c.FindCacheLine(addr)
	}
	line.Counter += 1
	line.TimeStamp = c.TimeStamp
	return time
}

func (c *Cache) Write(addr uint64) uint{
	c.TimeStamp += 1
	time := c.latency_.HitLatency
	line, isFound := c.FindCacheLine(addr)
	if isFound{
		if c.config_.WriteUpPolicy == WriteThrough {
			line.IsDirty = false
			time += c.lower_.HandleRequest(addr, uint(1), Write)
		}else{
			line.IsDirty = true
		}
	}else{
		c.Stats.MissNum += 1
		if c.config_.WriteBackPolicy == WriteAllocate{
			if evicteeAddr, IsRequirePushDown := c.Insert(addr); IsRequirePushDown{
				time += c.lower_.HandleRequest(evicteeAddr, uint(1), Write)
			}
			time += c.lower_.HandleRequest(addr, uint(1), Read)
			line, isFound = c.FindCacheLine(addr)
			if isFound{
				line.IsDirty = true
			}else{
				panic("Internal error")
			}
		}else{
			time += c.lower_.HandleRequest(addr, uint(1), Write)
			return time
		}
	}
	line.Counter += 1
	line.TimeStamp = c.TimeStamp
	return time
}

func (c *Cache) HandleRequest(addr uint64, bytes uint, request StorageRequestType) uint{
	rv := uint(0)
	for i := uint64(0); i < uint64(bytes); i++ {
		switch request{
		case Read: rv += c.Read(addr + i); break
		case Write: rv += c.Write(addr + i); break
		}
	}
	c.Stats.AccessCounter += bytes
	c.Stats.AccessTime += rv
	return rv
}