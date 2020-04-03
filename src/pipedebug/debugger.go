package pipedebug

import (
	"bufio"
	"debugger"
	"fmt"
	"os"
	"pipeline"
	"regexp"
	"strings"
	"utils"
)

type PipeDebugger struct {
	debugger.Debugger
}

func (d *PipeDebugger) Usage() {
	fmt.Println(strings.Join([]string{
		"Usage:",
		"   help                          : Display this message",
		"   exit                          : Exit the debug prompt",
		"   c                             : Continue running until the program comes to the end",
		"   si                            : Run a single pipeline stage",
		"   reg                           : Display register information",
		"   status                        : Display the status of each pipeline register",
		"   info [symbol name]            : Display the address of this symbol defined in ELF file. if \"symbol name\" is not specified, display all the symbols.",
		"   x/<length><format> [address]  : Display the memory contents at a given address using the specified format. The address is PC by default",
		"									Valid format specifiers are:",
		"											i - instruction",
		"											b - byte",
		"											h - half word(16-bit value)",
		"											w - word(32-bit value)",
		"											g - giant word (64-bit value)",
	}, "\n"))
}


func (d *PipeDebugger) DisplayStatus(p *pipeline.Pipe) {
	fmt.Println(p.WReg.ToString())
	fmt.Println(p.MReg.ToString())
	fmt.Println(p.EReg.ToString())
	fmt.Println(p.DReg.ToString())
}

func (d *PipeDebugger) RunPrompt(p *pipeline.Pipe) {
	memCommandRe := regexp.MustCompile("^x/([1-9]\\d*)([ibhwg])$")
	lineReader := bufio.NewReader(os.Stdin)
	sim := p.Simulator
	fmt.Println("Start running...")
	p.Initialize()
	for true {
		if p.HasFinished(){
			fmt.Println("The program has ended.")
		}
		d.DisplayPCInstruction(sim)
		fmt.Print("(PipeDebug) ")
		line, _ := lineReader.ReadString('\n')
		line = strings.Trim(line, "\n")
		commands := strings.Split(line, " ")
		commands = utils.Filter(commands, func(str interface{}) bool { return str != "" }).([]string)
		switch {
		case len(commands) == 0:
			fmt.Print("Unknown command. Type \"help\" for more info.")
		case commands[0] == "help":
			d.Usage()
		case commands[0] == "exit":
			fmt.Println("[Statics]")
			fmt.Printf("%-40s: %d\n", "Cycle Per Step", pipeline.Pipeline_step_period)
			fmt.Printf("%-40s: %d\n", "# of Steps", p.CycleCounter.Count)
			fmt.Printf("%-40s: %d\n", "# of Cycles", p.CycleCounter.Count * pipeline.Pipeline_step_period)
			fmt.Printf("%-40s: %d\n", "# of Valid Instructions", p.CycleCounter.Count)
			fmt.Printf("%-40s: %.5f\n", "CPI",
				float32(p.CycleCounter.Count * pipeline.Pipeline_step_period)/float32(p.ValidInstCountCounter.Count))
			fmt.Printf("%-40s: %d\n", "# of Indirect Jump", p.IndirectJumpCounter.Count)
			fmt.Printf("%-40s: %.5f%%\n", "Jump Prediction Success Rate",
				 100.0 * (1- float32(p.JumpPredictionCounter.Numerator.Count) /
				float32(p.JumpPredictionCounter.Denominator.Count)))
			goto done
		case commands[0] == "status":
			d.DisplayStatus(p)
		case commands[0] == "c":
			var err error = nil
			for !p.HasFinished() && err == nil {
				fmt.Println()
				d.DisplayPCInstruction(sim)
				_, err = p.SingleStep()
				d.DisplayStatus(p)
			}
		case commands[0] == "si":
			if _, err := p.SingleStep(); err != nil {
				fmt.Println(err)
			}
			d.DisplayStatus(p)
		case commands[0] == "reg":
			d.DisplayRegisters(&sim.R)
			fmt.Printf("\n %4s: 0x%016x\n", "PC", sim.PC)
		case commands[0] == "info":
			if len(commands) == 1 {
				d.DisplayAllSymbolAddr(sim)
			} else if len(commands) == 2 {
				d.DisplaySymbolAddr(sim, commands[1])
			} else {
				fmt.Println("Command Format Error. Type \"help\" for more info.")
			}
		case memCommandRe.MatchString(commands[0]):
			d.ParseMemoryCommand(sim, commands)
		default:
			fmt.Print("Unknown command. Type \"help\" for more info.")
		}
		fmt.Println()
	}
done:
	fmt.Println("exit...")
}
