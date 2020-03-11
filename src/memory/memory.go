package memory

import (
	"encoding/binary"
	"log"
)

type memSegment64 struct{
	begin, size uint64
	buf []byte
}

func (ms memSegment64) end() uint64{
	return ms.begin + ms.size
}

func (ms memSegment64) isContain(addr uint64) bool{
	return ms.begin <= addr && addr < ms.end()
}

func (ms memSegment64) isOverlay(other *memSegment64) bool{
	return !(other.end() <= ms.begin || ms.end() <= other.begin)
}

func (ms memSegment64) slice(addr uint64) []byte{
	return ms.buf[addr - ms.begin:]
}

type Memory64 struct {
	segs []memSegment64;
}

func (m *Memory64) findSegment(addr uint64) (*memSegment64, bool) {
	for _, ms := range m.segs{
		if ms.isContain(addr){
			return &ms, true
		}
	}
	log.Panicf("Segmentation Fault. addr = 0x%016x", addr)
	return nil, false
}

func (m *Memory64) Alloc(begin uint64, size uint64) bool{
	seg := memSegment64{begin, size, make([]byte, size)}
	for _, ms := range m.segs{
		if ms.isOverlay(&seg) {
			log.Panicf("Memory allocation failed. begin = 0x%016x, end = 0x%016x", begin, size)
			return false
		}
	}
	m.segs = append(m.segs, seg)
	return true
}

func (m *Memory64) LoadU8(addr uint64) uint8{
	if ms, ok := m.findSegment(addr); ok{
		return ms.slice(addr)[0];
	}
	return 0
}

func (m *Memory64) LoadU16(addr uint64) uint16{
	if ms, ok := m.findSegment(addr); ok{
		return binary.LittleEndian.Uint16(ms.slice(addr))
	}
	return 0
}

func (m *Memory64) LoadU32(addr uint64) uint32{
	if ms, ok := m.findSegment(addr); ok{
		return binary.LittleEndian.Uint32(ms.slice(addr))
	}
	return 0
}

func (m *Memory64) LoadU64(addr uint64) uint64{
	if ms, ok := m.findSegment(addr); ok{
		return binary.LittleEndian.Uint64(ms.slice(addr))
	}
	return 0
}

func (m *Memory64) StoreU8(addr uint64, value uint8){
	if ms, ok := m.findSegment(addr); ok{
		ms.slice(addr)[0] = value
	}
}

func (m *Memory64) StoreU16(addr uint64, value uint16){
	if ms, ok := m.findSegment(addr); ok{
		binary.LittleEndian.PutUint16(ms.slice(addr), value)
	}
}

func (m *Memory64) StoreU32(addr uint64, value uint32){
	if ms, ok := m.findSegment(addr); ok{
		binary.LittleEndian.PutUint32(ms.slice(addr), value)
	}
}

func (m *Memory64) StoreU64(addr uint64, value uint64){
	if ms, ok := m.findSegment(addr); ok{
		binary.LittleEndian.PutUint64(ms.slice(addr), value)
	}
}

func (m *Memory64) storeBytes(addr uint64, size uint64, bytes []byte){
	if ms, ok := m.findSegment(addr); ok{
		copy(ms.slice(addr), bytes[:size])
	}
}