package simulator

import (
	"debug/elf"
	"errors"
	"isa"
	"log"
	"memory"
	"register"
	"utils"
)

var STACK_TOP uint64 = 0x7fffff_fff000
var STACK_SIZE uint64 = 4096

type Simulator struct {
	M       memory.Memory64
	R       register.Heap64
	PC      uint64
	elfFile *elf.File
}

func NewSimulator(elfFilePath string) *Simulator {
	file, err := elf.Open(elfFilePath)
	utils.PANIC_CHECK(err)
	return &Simulator{elfFile: file}
}

func (sim *Simulator) LoadMemory() {
	file := sim.elfFile
	for _, prog := range file.Progs{
		if prog.Type == elf.PT_LOAD{
			vaddr, memsz, filesz := prog.Vaddr, prog.Memsz, prog.Filesz
			buf := make([]byte, filesz)
			_, err :=  prog.ReadAt(buf, 0)
			utils.PANIC_CHECK(err)
			sim.M.Alloc(vaddr, memsz)
			sim.M.StoreBytes(vaddr, filesz, buf)
		}
	}

	// Allocate stack memory and set stack pointer
	sim.M.Alloc(STACK_TOP - STACK_SIZE, STACK_SIZE)
	sim.R.StoreByName("sp", STACK_TOP)

	//! Magic: start from |main| function instead of the entry in elf file
	symbols, _  := file.Symbols()
	found := false
	for _, symbol := range symbols {
		if symbol.Name == "main"{
			sim.PC = symbol.Value
			found = true
			break
		}
	}
	if found == false{
		log.Fatalln("Can not find the symbol named \"main\"");
	}
}

func Match(code uint32) (item *Action, err error) {
	for _, item := range ActionSet {
		if item.Pred(code) == true {
			return &item, nil
		}
	}
	return nil, errors.New("unknown instruction format")
}

func (sim *Simulator) SingleStep() *Action {
	m := sim.M
	code := m.LoadU32(sim.PC)
	if isa.InstructionLength(code) != 32{
		log.Panicln("Only support 32-bit instruction")
	}
	item, err := Match(code)
	if err != nil {
		log.Fatalf(" 0x%08x: %v", code, err)
	}
	inst := item.Inst(code)
	item.Exec(sim, inst)
	return item
}