// Code generated, DO NOT EDIT.
package simulator

import (
		"bit_utils"
		"isa"
		"log"
		"math/bits"
)

const (
		XLEN = 64
)

type Action struct{
	Name string
	Pred func (uint32) bool
	Inst func (uint32) isa.Instruction
	Exec func (*Simulator, isa.Instruction)
}

var ActionSet = []Action{{
			Name: "add",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x0,0x00, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=r.Load(rs1), r.Load(rs2); r.Store(rd, t1 + t2)
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "mul",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x0,0x01, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=r.Load(rs1), r.Load(rs2); _,lo:=bits.Mul64(t1,t2); r.Store(rd, lo)
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "sub",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x0,0x20, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=r.Load(rs1), r.Load(rs2); r.Store(rd, t1 - t2)
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "sll",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x1,0x00, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=r.Load(rs1), r.Load(rs2); r.Store(rd, t1 << t2)
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "mulh",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x1,0x01, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=r.Load(rs1), r.Load(rs2); hi,_:=bits.Mul64(t1,t2); r.Store(rd, hi)
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "slt",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x2,0x00, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=int64(r.Load(rs1)), int64(r.Load(rs2)); if t1<t2 {r.Store(rd, 1)} else{r.Store(rd, 0)}
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "xor",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x4,0x00, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=r.Load(rs1), r.Load(rs2); r.Store(rd, t1 ^ t2)
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "div",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x4,0x01, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=int64(r.Load(rs1)), int64(r.Load(rs2)); if t2==0 {r.Store(rd, bit_utils.SignExtU64(-1))} else if t1==-(1<<(XLEN-1)) && t2==-1 {r.Store(rd, uint64(1)<<(XLEN-1))} else {r.Store(rd, uint64(t1 / t2))}
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "srl",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x5,0x00, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=r.Load(rs1), r.Load(rs2); r.Store(rd, t1 >> t2)
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "sra",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x5,0x20, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=r.Load(rs1), r.Load(rs2); v:=int64(t1)>>t2;r.Store(rd, uint64(v))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "or",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x6,0x00, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=r.Load(rs1), r.Load(rs2); r.Store(rd, t1 | t2)
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "rem",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x6,0x01, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=int64(r.Load(rs1)), int64(r.Load(rs2)); if t2==0 {r.Store(rd, uint64(t1))} else if t1==-(1<<(XLEN-1)) && t2==-1 {r.Store(rd, 0)} else {r.Store(rd, uint64(t1 % t2))}
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "and",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x7,0x00, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=r.Load(rs1), r.Load(rs2); r.Store(rd, t1 & t2)
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "lb",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x03,0x0, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewIInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.IInstruction); ok{
					rs1, rd, imm := t.Rs1, t.Rd, t.Imm
					_, _, _ = rs1, rd, imm
					t1:=r.Load(rs1); v:= uint64(int64(int8(m.LoadU8(t1 + bit_utils.SignExtU64(imm))))); r.Store(rd, v)
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "lh",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x03,0x1, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewIInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.IInstruction); ok{
					rs1, rd, imm := t.Rs1, t.Rd, t.Imm
					_, _, _ = rs1, rd, imm
					t1:=r.Load(rs1); v:= uint64(int64(int16(m.LoadU16(t1 + bit_utils.SignExtU64(imm))))); r.Store(rd, v)
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "lw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x03,0x2, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewIInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.IInstruction); ok{
					rs1, rd, imm := t.Rs1, t.Rd, t.Imm
					_, _, _ = rs1, rd, imm
					t1:=r.Load(rs1); v:= uint64(int64(int32(m.LoadU32(t1 + bit_utils.SignExtU64(imm))))); r.Store(rd, v)
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "ld",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x03,0x3, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewIInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.IInstruction); ok{
					rs1, rd, imm := t.Rs1, t.Rd, t.Imm
					_, _, _ = rs1, rd, imm
					t1:=r.Load(rs1); v:= uint64(int64(m.LoadU64(t1 + bit_utils.SignExtU64(imm)))); r.Store(rd, v)
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "addi",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x13,0x0, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewIInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.IInstruction); ok{
					rs1, rd, imm := t.Rs1, t.Rd, t.Imm
					_, _, _ = rs1, rd, imm
					t1:=r.Load(rs1); r.Store(rd, t1 + bit_utils.SignExtU64(imm))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "slli",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x13,0x1, inst,[3]uint32{26, 31, 0})
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewIInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.IInstruction); ok{
					rs1, rd, imm := t.Rs1, t.Rd, t.Imm
					_, _, _ = rs1, rd, imm
					t1:=r.Load(rs1); r.Store(rd, t1 << bit_utils.U32Bits(uint32(imm), 0, 5))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "slti",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x13,0x2, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewIInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.IInstruction); ok{
					rs1, rd, imm := t.Rs1, t.Rd, t.Imm
					_, _, _ = rs1, rd, imm
					t1:=int64(r.Load(rs1)); if t1 < int64(imm) { r.Store(rd, 1)} else {r.Store(rd,0)} 
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "xori",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x13,0x4, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewIInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.IInstruction); ok{
					rs1, rd, imm := t.Rs1, t.Rd, t.Imm
					_, _, _ = rs1, rd, imm
					t1:=r.Load(rs1); r.Store(rd, t1 ^ bit_utils.SignExtU64(imm))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "srli",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x13,0x5, inst,[3]uint32{26, 31, 0})
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewIInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.IInstruction); ok{
					rs1, rd, imm := t.Rs1, t.Rd, t.Imm
					_, _, _ = rs1, rd, imm
					t1:=r.Load(rs1); r.Store(rd, t1 >> bit_utils.U32Bits(uint32(imm), 0, 5))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "srai",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x13,0x5, inst,[3]uint32{26, 31, 0x10})
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewIInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.IInstruction); ok{
					rs1, rd, imm := t.Rs1, t.Rd, t.Imm
					_, _, _ = rs1, rd, imm
					t1:=int64(r.Load(rs1)); r.Store(rd, uint64(t1 >>bit_utils.U32Bits(uint32(imm), 0, 5)))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "ori",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x13,0x6, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewIInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.IInstruction); ok{
					rs1, rd, imm := t.Rs1, t.Rd, t.Imm
					_, _, _ = rs1, rd, imm
					t1:=r.Load(rs1); r.Store(rd, t1 | bit_utils.SignExtU64(imm))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "andi",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x13,0x7, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewIInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.IInstruction); ok{
					rs1, rd, imm := t.Rs1, t.Rd, t.Imm
					_, _, _ = rs1, rd, imm
					t1:=r.Load(rs1); r.Store(rd, t1 & bit_utils.SignExtU64(imm))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "addiw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x1B,0x0, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewIInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.IInstruction); ok{
					rs1, rd, imm := t.Rs1, t.Rd, t.Imm
					_, _, _ = rs1, rd, imm
					t1:=r.Load(rs1); r.Store(rd, bit_utils.SignExtU64(int32(t1) + imm))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "jalr",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x67,0x0, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewIInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.IInstruction); ok{
					rs1, rd, imm := t.Rs1, t.Rd, t.Imm
					_, _, _ = rs1, rd, imm
					r.Store(rd, *pc + 4)
					*pc = bit_utils.SetBit(r.Load(rs1) + bit_utils.SignExtU64(imm), 0, 0)
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "sb",
			Pred: func (inst uint32) bool {
				return isa.IsMatchSFormat(0x23,0x0, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewSInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.SInstruction); ok{
					rs1, rs2, imm := t.Rs1, t.Rs2, t.Imm
					_, _, _ = rs1, rs2, imm
					t1,t2:=r.Load(rs1),r.Load(rs2);addr:=t1+bit_utils.SignExtU64(imm); m.StoreU8(addr, uint8(t2))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "sh",
			Pred: func (inst uint32) bool {
				return isa.IsMatchSFormat(0x23,0x1, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewSInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.SInstruction); ok{
					rs1, rs2, imm := t.Rs1, t.Rs2, t.Imm
					_, _, _ = rs1, rs2, imm
					t1,t2:=r.Load(rs1),r.Load(rs2);addr:=t1+bit_utils.SignExtU64(imm); m.StoreU16(addr, uint16(t2))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "sw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchSFormat(0x23,0x2, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewSInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.SInstruction); ok{
					rs1, rs2, imm := t.Rs1, t.Rs2, t.Imm
					_, _, _ = rs1, rs2, imm
					t1,t2:=r.Load(rs1),r.Load(rs2);addr:=t1+bit_utils.SignExtU64(imm); m.StoreU32(addr, uint32(t2))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "sd",
			Pred: func (inst uint32) bool {
				return isa.IsMatchSFormat(0x23,0x3, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewSInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.SInstruction); ok{
					rs1, rs2, imm := t.Rs1, t.Rs2, t.Imm
					_, _, _ = rs1, rs2, imm
					t1,t2:=r.Load(rs1),r.Load(rs2);addr:=t1+bit_utils.SignExtU64(imm); m.StoreU64(addr, t2)
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "beq",
			Pred: func (inst uint32) bool {
				return isa.IsMatchBFormat(0x63,0x0, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewBInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.BInstruction); ok{
					rs1, rs2, imm := t.Rs1, t.Rs2, t.Imm
					_, _, _ = rs1, rs2, imm
					t1, t2:=r.Load(rs1), r.Load(rs2)
					if t1==t2 {*pc+= bit_utils.SignExtU64(imm<<1)} else {*pc += 4}
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "bne",
			Pred: func (inst uint32) bool {
				return isa.IsMatchBFormat(0x63,0x1, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewBInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.BInstruction); ok{
					rs1, rs2, imm := t.Rs1, t.Rs2, t.Imm
					_, _, _ = rs1, rs2, imm
					t1, t2:=r.Load(rs1), r.Load(rs2)
					if t1!=t2 {*pc+= bit_utils.SignExtU64(imm<<1)} else {*pc += 4}
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "blt",
			Pred: func (inst uint32) bool {
				return isa.IsMatchBFormat(0x63,0x4, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewBInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.BInstruction); ok{
					rs1, rs2, imm := t.Rs1, t.Rs2, t.Imm
					_, _, _ = rs1, rs2, imm
					t1, t2:=int64(r.Load(rs1)), int64(r.Load(rs2))
					if t1<t2 {*pc+= bit_utils.SignExtU64(imm<<1)} else {*pc += 4}
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "bge",
			Pred: func (inst uint32) bool {
				return isa.IsMatchBFormat(0x63,0x5, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewBInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.BInstruction); ok{
					rs1, rs2, imm := t.Rs1, t.Rs2, t.Imm
					_, _, _ = rs1, rs2, imm
					t1, t2:=int64(r.Load(rs1)), int64(r.Load(rs2))
					if t1>=t2 {*pc+= bit_utils.SignExtU64(imm<<1)} else {*pc += 4}
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "auipc",
			Pred: func (inst uint32) bool {
				return isa.IsMatchUFormat(0x17, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewUInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.UInstruction); ok{
					rd, imm := t.Rd, t.Imm
					_, _ = rd, imm
					r.Store(rd, *pc + (bit_utils.SignExtU64(imm)<<12))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "lui",
			Pred: func (inst uint32) bool {
				return isa.IsMatchUFormat(0x37, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewUInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.UInstruction); ok{
					rd, imm := t.Rd, t.Imm
					_, _ = rd, imm
					r.Store(rd, bit_utils.SignExtU64(imm)<<12)
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "jal",
			Pred: func (inst uint32) bool {
				return isa.IsMatchJFormat(0x6f, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewJInstruction(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(*isa.JInstruction); ok{
					rd, imm := t.Rd, t.Imm
					_, _ = rd, imm
					r.Store(rd, *pc + 4)
					*pc += bit_utils.SignExtU64(imm) << 1
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},
}
