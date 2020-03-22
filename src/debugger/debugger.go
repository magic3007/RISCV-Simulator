package debugger

import (
	"bufio"
	"fmt"
	"memory"
	"os"
	"regexp"
	"register"
	"simulator"
	"strconv"
	"strings"
	"utils"
)

type Debugger struct {
}

func (d *Debugger) displayRegisters(r *register.Heap64) {
	for i := uint8(0); i < register.RegNum; i++ {
		fmt.Printf(" %4s: 0x%016x", register.NamefromIndex(i), r.Load(uint8(i)))
		if i%5 == 4 {
			fmt.Println()
		}
	}
}

func (d *Debugger) displaySymbolAddr(sim *simulator.Simulator, name string) {
	if symbol, found := sim.FindSymbolFromName(name); found {
		fmt.Printf("%-15s: 0x%016x\n", name, symbol.Value)
	} else {
		fmt.Printf("Couldn't find symbol \"%s\"\n", name)
	}
}

func (d *Debugger) displayAllSymbolAddr(sim *simulator.Simulator) {
	symbols, _  := sim.ElfFile.Symbols()
	for _, symbol := range symbols {
		fmt.Printf("%-40s: 0x%016x\n", symbol.Name, symbol.Value)
	}
}


func (d *Debugger) displayMemory(m *memory.Memory64, addr uint64, transformer func(*memory.Memory64, uint64) string) {
	format := transformer(m, addr)
	fmt.Printf("0x%016x:           "+format+"\n", addr)
}

func (d *Debugger) parserMemoryCommand(sim *simulator.Simulator, commands []string){
	memCommandRe := regexp.MustCompile("^x/([1-9]\\d*)([ibhwg])$")
	submatches := memCommandRe.FindStringSubmatch(commands[0])
	var addr, counter uint64
	var err error
	if len(commands) > 2 {
		goto error
	}
	if len(commands) == 1{
		addr = sim.PC
	}else if addr, err = strconv.ParseUint(commands[1], 0, 64); err != nil {
		goto error
	}
	if counter, err = strconv.ParseUint(submatches[1], 0, 64); err != nil {
		goto error
	}
	for i := uint64(0); i < counter; i++ {
		switch submatches[2] {
		case "b":
			d.displayMemory(&sim.M, addr, func(m *memory.Memory64, addr uint64) string {
				return fmt.Sprintf("0x%02x", m.LoadU8(addr))
			})
			addr += 1
		case "h":
			d.displayMemory(&sim.M, addr, func(m *memory.Memory64, addr uint64) string {
				return fmt.Sprintf("0x%04x", m.LoadU16(addr))
			})
			addr += 2
		case "w":
			d.displayMemory(&sim.M, addr, func(m *memory.Memory64, addr uint64) string {
				return fmt.Sprintf("0x%08x", m.LoadU32(addr))
			})
			addr += 4
		case "g":
			d.displayMemory(&sim.M, addr, func(m *memory.Memory64, addr uint64) string {
				return fmt.Sprintf("0x%016x", m.LoadU64(addr))
			})
			addr += 8
		case "i":
			d.displayMemory(&sim.M, addr, func(m *memory.Memory64, addr uint64) string {
				code := m.LoadU32(addr)
				item, err := simulator.Match(code)
				if err != nil {
					return "<unknown instruction>"
				}
				inst := item.Inst(code)
				return item.ToString(inst)
			})
			addr += 4
		}
	}
	goto done
error:
	fmt.Println("Command Format Error. Type \"help\" for more info.")
done:
}

func (d *Debugger) Usage() {
	fmt.Println(strings.Join([]string{
		"Usage:",
		"   help                          : Display this message",
		"   exit                          : Exit the debug prompt",
		"   c                             : Continue running until the program comes to the end",
		"   si                            : Run a single machine instruction",
		"   reg                           : Display register information",
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


func (d *Debugger) RunPrompt(sim *simulator.Simulator) {
	memCommandRe := regexp.MustCompile("^x/([1-9]\\d*)([ibhwg])$")
	lineReader := bufio.NewReader(os.Stdin)
	fmt.Println("Start running...")
	for true {
		if !sim.HasFinished() {
			fmt.Print("PC -> ")
			d.parserMemoryCommand(sim, []string{"x/1i", strconv.FormatUint(sim.PC, 10)})
		}else{
			fmt.Println("The program has ended.")
		}
		fmt.Print("(Debug) ")
		line, _ := lineReader.ReadString('\n')
		line = strings.Trim(line, "\n")
		commands := strings.Split(line, " ")
		commands = utils.Filter(commands, func(str interface{}) bool { return str != "" }).([]string)
		switch {
		case len(commands)==0:
			fmt.Print("Unknown command. Type \"help\" for more info.")
		case commands[0] == "help":
			d.Usage()
		case commands[0] == "exit":
			goto done
		case commands[0] == "c":
			var err error = nil
			for !sim.HasFinished() && err == nil {
				d.parserMemoryCommand(sim, []string{"x/1i", strconv.FormatUint(sim.PC, 10)})
				_, err = sim.SingleStep()
			}
		case commands[0] == "si":
			if _, err := sim.SingleStep(); err != nil {
				fmt.Print(err)
			}
		case commands[0] == "reg":
			d.displayRegisters(&sim.R)
			fmt.Printf("\n %4s: 0x%016x\n","PC", sim.PC)
		case commands[0] == "info":
			if len(commands) == 1{
				d.displayAllSymbolAddr(sim)
			} else if len(commands) == 2 {
				d.displaySymbolAddr(sim, commands[1])
			} else {
				fmt.Println("Command Format Error. Type \"help\" for more info.")
			}
		case memCommandRe.MatchString(commands[0]):
			d.parserMemoryCommand(sim, commands)
		default:
			fmt.Print("Unknown command. Type \"help\" for more info.")
		}
		fmt.Println()
	}
	done:
		fmt.Println("exit...")
}
