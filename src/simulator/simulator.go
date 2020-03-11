package simulator

import (
	"debug/elf"
	"memory"
	"register"
)

type Simulator struct {
	m  memory.Memory64
	rh register.Heap64
	pc uint64
}

func (sim *Simulator) LoadElfFile(file *elf.File) {

}
