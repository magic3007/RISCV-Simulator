package memory

import (
	"fmt"
	"log"
)

func ExampleT_memory64(){
	log.SetFlags(log.LstdFlags | log.Llongfile)
	m := Memory64{}
	m.Alloc(0, 1024)
	m.StoreU16(0, 0x1010)
	m.StoreU32(1016, 0x89d2defa)
	fmt.Printf("%02x\n", m.LoadU8(1))
	fmt.Printf("%04x\n", m.LoadU16(1020))
	fmt.Printf("%04x\n", m.LoadU16(1016))
	fmt.Printf("%04x\n", m.LoadU16(1018))
	// Output:
	// 10
	// 0000
	// defa
	// 89d2
}