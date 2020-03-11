package simulator

import (
	"debug/elf"
	"memory"
	"register"
	"utils"
)

type Simulator struct {
	m  memory.Memory64
	rh register.Heap64
	pc uint64
	elfFile *elf.File
}

func NewSimulator(elfFilePath string) *Simulator {
	file, err := elf.Open(elfFilePath)
	utils.PANIC_CHECK(err)
	return &Simulator{elfFile: file}
}

func (sim *Simulator) LoadMemory() {
	for _, prog := range sim.elfFile.Progs{
		if prog.Type == elf.PT_LOAD{
			vaddr, memsz, filesz := prog.Vaddr, prog.Memsz, prog.Filesz
			buf := make([]byte, filesz)
			_, err :=  prog.ReadAt(buf, 0)
			utils.PANIC_CHECK(err)
			sim.m.Alloc(vaddr, memsz)
			sim.m.StoreBytes(vaddr, filesz, buf)
		}
	}
	sim.pc = sim.elfFile.Entry
}