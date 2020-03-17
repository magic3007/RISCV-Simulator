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
		{"add", "/../test_data/add.out", 0xfe010113},
		{"double-float", "/../test_data/double-float.out", 0xfe010113},
		{"mul-div", "/../test_data/mul-div.out", 0xfe010113},
		{"n!", "/../test_data/n!.out", 0xff010113},
		{"qsort", "/../test_data/qsort.out", 0xfe010113},
		{"simple-function", "/../test_data/simple-function.out", 0xff010113},
	}
	log.SetFlags(log.Llongfile | log.LstdFlags)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filepath := utils.GetCurrentPath() + test.relativePath
			sim := NewSimulator(filepath)
			sim.LoadMemory()
			assert.Equal(t, test.expected, sim.M.LoadU32(sim.PC))
		})
	}
}

func TestSingleStep(t *testing.T) {
	tests := []struct {
		name string
		relativePath string
		expected []string
	}{
		{"add", "/../test_data/add.out",
			[]string{"addi", "sd", "addi", "sw", "sw"},
			},
	}
	log.SetFlags(log.Llongfile | log.LstdFlags)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filepath := utils.GetCurrentPath() + test.relativePath
			sim := NewSimulator(filepath)
			sim.LoadMemory()
			for _, expectInst := range test.expected{
				item := sim.SingleStep()
				assert.Equal(t, expectInst, item.Name)
			}
		})
	}
}