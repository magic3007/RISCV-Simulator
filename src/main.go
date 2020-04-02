package main

import (
	"debug/elf"
	"debugger"
	"flag"
	"fmt"
	"log"
	"pipedebug"
	"pipeline"
	"simulator"
	"utils"
)

var (
	verbose  = flag.Bool("v", true, "Display verbose info of ELF file. To disable displaying these info, type \"-v=0\"")
	filepath = flag.String("f", "", "Filepath of the ELF file")
	mode = flag.String("m", "debug", "Simulation mode. Valid modes are: debug, pipeline")
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
	log.SetFlags(log.LstdFlags | log.Llongfile)
	file, err := elf.Open(*filepath)
	utils.PANIC_CHECK(err)
	defer file.Close()
	if *verbose == true {
		statElf(file)
	}
	sim := simulator.NewSimulator(*filepath)
	sim.LoadMemory()
	switch *mode{
	case "debug":
		var d debugger.Debugger
		d.RunPrompt(sim)
	case "pipeline":
		var d pipedebug.PipeDebugger
		pipe := pipeline.Pipe{Simulator: sim}
		d.RunPrompt(&pipe)
	default:
		err := fmt.Errorf("unvalid mode")
		fmt.Print(err.Error())
	}
}
