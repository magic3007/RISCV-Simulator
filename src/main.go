package main

import (
	"debug/elf"
	"flag"
	"simulator"
	"utils"
)

var (
	verbose  = flag.Bool("v", true, "Display verbose info of ELF file")
	filepath = flag.String("f", "", "filepath of the ELF file")
)

func main() {
	utils.FlagInit([]string{"f"})
	file, err := elf.Open(*filepath)
	utils.PANIC_CHECK(err)
	defer file.Close()
	if *verbose == true {
		utils.RecursiveStatField(*file, 0)
	}
	sim := simulator.Simulator{}
	sim.LoadElfFile(file)
}
