package main

import (
	"debug/elf"
	"flag"
	"fmt"
	"simulator"
	"utils"
)

var (
	verbose  = flag.Bool("v", true, "Display verbose info of ELF file")
	filepath = flag.String("f", "", "filepath of the ELF file")
)

func statElf(file *elf.File){
	utils.RecursiveStatField(*file, 0)
	for index, prog := range file.Progs{
		fmt.Println("===========================================")
		fmt.Printf("Porgram %d:\n", index)
		utils.RecursiveStatField(*prog, 2)
	}
}

func main() {
	utils.FlagInit([]string{"f"})
	file, err := elf.Open(*filepath)
	utils.PANIC_CHECK(err)
	defer file.Close()
	if *verbose == true {
		statElf(file)
	}
	sim := simulator.Simulator{}
	sim.LoadElfFile(file)
}
