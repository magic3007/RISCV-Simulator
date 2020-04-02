package register

import (
	"log"
)

const RegNum = 32

type Heap64 struct {
	regs [RegNum]uint64
}

func (rh *Heap64) Load(index uint8) uint64 {
	return rh.regs[index]
}

func (rh *Heap64) Store(index uint8, value uint64) {
	if index != 0 {
		rh.regs[index] = value
	}
}

func (rh *Heap64) LoadByName(name string) uint64 {
	return rh.Load(fromString(name))
}

func (rh *Heap64) StoreByName(name string, value uint64) {
	rh.Store(fromString(name), value)
}

func fromString(name string) uint8 {
	for index, lst := range RegName {
		for _, str := range lst {
			if str == name {
				return uint8(index)
			}
		}
	}
	log.Panicf("No such register name: %s", name)
	return 0
}

/*
* @return use the last items in register name list by default
 */
func NamefromIndex(index uint8) string {
	lst := RegName[index]
	return lst[len(lst)-1]
}

/**
Credit to https://content.riscv.org/wp-content/uploads/2017/05/riscv-spec-v2.2.pdf.
*/
var RegName = [RegNum][]string{
	/* 00 */ []string{"x0", "zero"},
	/* 01 */ []string{"x1", /* return address */ "ra"},
	/* 02 */ []string{"x2", /* stack pointer */ "sp"},
	/* 03 */ []string{"x3", /* global pointer */ "gp"},
	/* 04 */ []string{"x4", /* thread pointer */ "tp"},
	/* 05 */ []string{"x5", /* alternate link register */ "t0"},
	/* 06 */ []string{"x6", "t1"},
	/* 07 */ []string{"x7", "t2"},
	/* 08 */ []string{"x8", /* save register */ "s0", /* frame pointer */ "fp"},
	/* 09 */ []string{"x9", "s1"},
	/* 10 */ []string{"x10", /* function arguments/return values */ "a0"},
	/* 11 */ []string{"x11", "a1"},
	/* 12 */ []string{"x12", "a2"},
	/* 13 */ []string{"x13", "a3"},
	/* 14 */ []string{"x14", "a4"},
	/* 15 */ []string{"x15", "a5"},
	/* 16 */ []string{"x16", "a6"},
	/* 17 */ []string{"x17", "a7"},
	/* 18 */ []string{"x18", "s3"},
	/* 19 */ []string{"x19", "s3"},
	/* 20 */ []string{"x20", "s4"},
	/* 21 */ []string{"x21", "s5"},
	/* 22 */ []string{"x22", "s6"},
	/* 23 */ []string{"x23", "s7"},
	/* 24 */ []string{"x24", "s8"},
	/* 25 */ []string{"x25", "s9"},
	/* 26 */ []string{"x26", "s10"},
	/* 27 */ []string{"x27", "s11"},
	/* 28 */ []string{"x28", "t3"},
	/* 29 */ []string{"x29", "t4"},
	/* 30 */ []string{"x30", "t5"},
	/* 31 */ []string{"x31", "t6"},
}
