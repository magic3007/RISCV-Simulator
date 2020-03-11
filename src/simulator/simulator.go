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
}

func (sim *Simulator) LoadElfFile(file *elf.File) {
	for _, prog := range file.Progs{
		if prog.Type == elf.PT_LOAD{
			vaddr, memsz, filesz := prog.Vaddr, prog.Memsz, prog.Filesz
			buf := make([]byte, filesz)
			_, err :=  prog.ReadAt(buf, 0)
			utils.PANIC_CHECK(err)
			sim.m.Alloc(vaddr, memsz)
			sim.m.StoreBytes(vaddr, filesz, buf)
		}
	}
	sim.pc = file.Entry
}