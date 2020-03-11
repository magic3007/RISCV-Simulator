package simulator

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"utils"
)

func TestLoadElf(t *testing.T) {
	tests := []struct {
		name string
		relativePath string
		expected uint32
	}{
		{"add", "/../test_data/add.out", 0x00002197},
		{"double-float", "/../test_data/double-float.out", 0x00002197},
		{"mul-div", "/../test_data/mul-div.out", 0x00002197},
		{"qsort", "/../test_data/qsort.out", 0x00002197},
		{"simple-function", "/../test_data/simple-function.out", 0x00002197},
	}
	log.SetFlags(log.Llongfile | log.LstdFlags)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filepath := utils.GetCurrentPath() + test.relativePath
			sim := NewSimulator(filepath)
			sim.LoadMemory()
			assert.Equal(t, test.expected, sim.m.LoadU32(sim.pc))
		})
	}
}