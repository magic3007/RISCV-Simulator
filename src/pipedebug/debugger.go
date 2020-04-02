package pipedebug

import (
	"bufio"
	"debugger"
	"fmt"
	"os"
	"pipeline"
	"regexp"
	"strconv"
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

func DisplayStatus(p *pipeline.Pipe) {
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
		if !p.HasFinished() {
			fmt.Print("PC -> ")
			d.ParseMemoryCommand(sim, []string{"x/1i", strconv.FormatUint(sim.PC, 10)})
		} else {
			fmt.Println("The program has ended.")
		}
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
			goto done
		case commands[0] == "status":
			DisplayStatus(p)
		case commands[0] == "c":
			var err error = nil
			count := 0
			for !sim.HasFinished() && err == nil {
				d.ParseMemoryCommand(sim, []string{"x/1i", strconv.FormatUint(sim.PC, 10)})
				_, err = p.SingleStep()
				count += 1
			}
		case commands[0] == "si":
			if _, err := p.SingleStep(); err != nil {
				fmt.Println(err)
			}
			DisplayStatus(p)
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
