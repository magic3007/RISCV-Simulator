package cache

type Memory struct{
	stats_ 		StorageStats
	latency_	StorageLatency
}

func NewMemory(latency StorageLatency) *Memory {
	c := new(Memory)
	c.latency_ = latency
	return c
}

func (m *Memory) HandleRequest(addr uint64, bytes uint, request StorageRequestType) uint{
	rv := m.latency_.HitLatency
	m.stats_.AccessCounter += 1
	m.stats_.AccessTime += rv
	return rv
}
