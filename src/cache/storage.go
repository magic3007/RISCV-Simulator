package cache

type StorageStats struct {
	AccessCounter 	uint
	MissNum      	uint
	AccessTime   	uint // in nanoseconds
	ReplaceNum   	uint // evict old lines
	FetchNum     	uint // fetch lower layer
	PrefetchNum  	uint // prefetch
}

type StorageLatency struct {
	HitLatency 		uint // in nanoseconds
	BusLatency 		uint // added to each request
}

type StorageRequestType uint

const(
	Read StorageRequestType = iota
	Write
)

type Storage interface {
	HandleRequest(addr uint64, bytes uint, request StorageRequestType) uint
}