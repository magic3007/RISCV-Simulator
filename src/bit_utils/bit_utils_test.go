package bit_utils

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestSetBit(t *testing.T) {
	tests := []struct {
		name string
		value uint64
		index uint32
		targetBit uint64
		expected uint64
	}{
		{"test1", 0b10, 0, 1, uint64(0b11)},
		{"test2", 0b11, 1, 0, uint64(0b1)},
	}
	log.SetFlags(log.Llongfile | log.LstdFlags)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, SetBit(test.value, test.index, test.targetBit))
		})
	}
}

func TestConcatenateSignExtBits(t *testing.T) {
	tests := []struct {
		name string
		value uint32
		intervals [][2]uint32
		expected int32
	}{
		{"test1", 0b1101_1011_1100_1001, [][2]uint32{{4,7}, {0,1}}, 0b01_1100},
		{"test2", 0b1101_1011_1100_1001, [][2]uint32{{4,7}, {0,3}}, -100},
	}
	log.SetFlags(log.Llongfile | log.LstdFlags)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, ConcatenateSignExtBits(test.value, test.intervals...))
		})
	}
}

func TestMulI64I64(t *testing.T) {
	tests := []struct {
		name string
		x int64
		y int64
		expected_hi int64
		expected_lo int64
	}{
		{"test1", -1, 1, -1, -1},
	}
	log.SetFlags(log.Llongfile | log.LstdFlags)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			hi, lo := MulI64I64(test.x, test.y)
			assert.Equal(t, test.expected_hi, hi)
			assert.Equal(t, test.expected_lo, lo)
		})
	}
}

func TestMulI64U64(t *testing.T) {
	tests := []struct {
		name string
		x int64
		y uint64
		expected_hi int64
		expected_lo int64
	}{
		{"test1", -1<<63, 1<<3, (^0)-3, 0},
	}
	log.SetFlags(log.Llongfile | log.LstdFlags)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			hi, lo := MulI64U64(test.x, test.y)
			assert.Equal(t, test.expected_hi, hi)
			assert.Equal(t, test.expected_lo, lo)
		})
	}
}