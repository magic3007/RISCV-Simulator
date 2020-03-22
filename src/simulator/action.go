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
	ToString func (isa.Instruction) string
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
			ToString: func (inst isa.Instruction) string {
				return "add " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "mul " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "sub " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "sll " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "mulh " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "slt " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "xor " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "div " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "srl " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "sra " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "or " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "rem " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "and " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "lb " + isa.ToDISFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.IInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "lh " + isa.ToDISFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.IInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "lw " + isa.ToDISFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.IInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "ld " + isa.ToDISFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.IInstruction); ok{
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
			Name: "lwu",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x03,0b110, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewIInstruction(inst)
			},
			ToString: func (inst isa.Instruction) string {
				return "lwu " + isa.ToDISFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.IInstruction); ok{
					rs1, rd, imm := t.Rs1, t.Rd, t.Imm
					_, _, _ = rs1, rd, imm
					t1:=r.Load(rs1); v:= uint64(int64(int32(m.LoadU32(t1 + bit_utils.UnSignExtU64(imm))))); r.Store(rd, v)
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
			ToString: func (inst isa.Instruction) string {
				return "addi " + isa.ToDSIFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.IInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "slli " + isa.ToDSIFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.IInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "slti " + isa.ToDSIFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.IInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "xori " + isa.ToDSIFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.IInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "srli " + isa.ToDSIFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.IInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "srai " + isa.ToDSIFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.IInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "ori " + isa.ToDSIFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.IInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "andi " + isa.ToDSIFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.IInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "addiw " + isa.ToDSIFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.IInstruction); ok{
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
			Name: "slliw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0b0011011,0b001, inst,[3]uint32{25, 31, 0x0})
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewIInstruction(inst)
			},
			ToString: func (inst isa.Instruction) string {
				return "slliw " + isa.ToDSIFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.IInstruction); ok{
					rs1, rd, imm := t.Rs1, t.Rd, t.Imm
					_, _, _ = rs1, rd, imm
					t1:=uint32(r.Load(rs1)); r.Store(rd, bit_utils.SignExtU64(int32(t1 << bit_utils.U32Bits(uint32(imm), 0, 4))))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "srliw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0b0011011,0b101, inst,[3]uint32{25, 31, 0x0})
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewIInstruction(inst)
			},
			ToString: func (inst isa.Instruction) string {
				return "srliw " + isa.ToDSIFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.IInstruction); ok{
					rs1, rd, imm := t.Rs1, t.Rd, t.Imm
					_, _, _ = rs1, rd, imm
					t1:=uint32(r.Load(rs1)); r.Store(rd, bit_utils.SignExtU64(int32(t1 >> bit_utils.U32Bits(uint32(imm), 0, 4))))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "sraiw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0b0011011,0b101, inst,[3]uint32{25, 31, 0b0100000})
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewIInstruction(inst)
			},
			ToString: func (inst isa.Instruction) string {
				return "sraiw " + isa.ToDSIFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.IInstruction); ok{
					rs1, rd, imm := t.Rs1, t.Rd, t.Imm
					_, _, _ = rs1, rd, imm
					t1:=int32(r.Load(rs1)); r.Store(rd, bit_utils.SignExtU64(int32(t1 >> bit_utils.U32Bits(uint32(imm), 0, 4))))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "addw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0111011,0b000,0b0000000, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			ToString: func (inst isa.Instruction) string {
				return "addw " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=r.Load(rs1), r.Load(rs2); r.Store(rd, bit_utils.SignExtU64(int32(t1)+int32(t2)))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "subw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0111011,0b000,0b0100000, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			ToString: func (inst isa.Instruction) string {
				return "subw " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=r.Load(rs1), r.Load(rs2); r.Store(rd, bit_utils.SignExtU64(int32(t1)-int32(t2)))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "sllw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0111011,0b001,0b0000000, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			ToString: func (inst isa.Instruction) string {
				return "sllw " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=uint32(r.Load(rs1)), r.Load(rs2); shamt:=bit_utils.U32Bits(uint32(t2), 0, 4); r.Store(rd, bit_utils.SignExtU64(int32(t1<<shamt)))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "srlw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0111011,0b101,0b0000000, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			ToString: func (inst isa.Instruction) string {
				return "srlw " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=uint32(r.Load(rs1)), r.Load(rs2); shamt:=bit_utils.U32Bits(uint32(t2), 0, 4); r.Store(rd, bit_utils.SignExtU64(int32(t1>>shamt)))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "sraw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0111011,0b101,0b0100000, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			ToString: func (inst isa.Instruction) string {
				return "sraw " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=int32(r.Load(rs1)), r.Load(rs2); shamt:=bit_utils.U32Bits(uint32(t2), 0, 4); r.Store(rd, bit_utils.SignExtU64(int32(t1>>shamt)))
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
			ToString: func (inst isa.Instruction) string {
				return "jalr " + isa.ToDSIFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.IInstruction); ok{
					rs1, rd, imm := t.Rs1, t.Rd, t.Imm
					_, _, _ = rs1, rd, imm
					t1:=r.Load(rs1); r.Store(rd, *pc + 4)
					*pc = bit_utils.SetBit(t1 + bit_utils.SignExtU64(imm), 0, 0)
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
			ToString: func (inst isa.Instruction) string {
				return "sb " + isa.ToSISFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.SInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "sh " + isa.ToSISFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.SInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "sw " + isa.ToSISFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.SInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "sd " + isa.ToSISFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.SInstruction); ok{
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
			ToString: func (inst isa.Instruction) string {
				return "beq " + isa.ToSSIFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.BInstruction); ok{
					rs1, rs2, imm := t.Rs1, t.Rs2, t.Imm
					_, _, _ = rs1, rs2, imm
					t1, t2:=r.Load(rs1), r.Load(rs2)
					if t1==t2 {*pc+= bit_utils.SignExtU64(imm)} else {*pc += 4}
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
			ToString: func (inst isa.Instruction) string {
				return "bne " + isa.ToSSIFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.BInstruction); ok{
					rs1, rs2, imm := t.Rs1, t.Rs2, t.Imm
					_, _, _ = rs1, rs2, imm
					t1, t2:=r.Load(rs1), r.Load(rs2)
					if t1!=t2 {*pc+= bit_utils.SignExtU64(imm)} else {*pc += 4}
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
			ToString: func (inst isa.Instruction) string {
				return "blt " + isa.ToSSIFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.BInstruction); ok{
					rs1, rs2, imm := t.Rs1, t.Rs2, t.Imm
					_, _, _ = rs1, rs2, imm
					t1, t2:=int64(r.Load(rs1)), int64(r.Load(rs2))
					if t1<t2 {*pc+= bit_utils.SignExtU64(imm)} else {*pc += 4}
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
			ToString: func (inst isa.Instruction) string {
				return "bge " + isa.ToSSIFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.BInstruction); ok{
					rs1, rs2, imm := t.Rs1, t.Rs2, t.Imm
					_, _, _ = rs1, rs2, imm
					t1, t2:=int64(r.Load(rs1)), int64(r.Load(rs2))
					if t1>=t2 {*pc+= bit_utils.SignExtU64(imm)} else {*pc += 4}
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
			ToString: func (inst isa.Instruction) string {
				return "auipc " + isa.ToDIFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.UInstruction); ok{
					rd, imm := t.Rd, t.Imm
					_, _ = rd, imm
					r.Store(rd, *pc + (bit_utils.SignExtU64(imm)))
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
			ToString: func (inst isa.Instruction) string {
				return "lui " + isa.ToDIFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.UInstruction); ok{
					rd, imm := t.Rd, t.Imm
					_, _ = rd, imm
					r.Store(rd, bit_utils.SignExtU64(imm))
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
			ToString: func (inst isa.Instruction) string {
				return "jal " + isa.ToDIFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.JInstruction); ok{
					rd, imm := t.Rd, t.Imm
					_, _ = rd, imm
					r.Store(rd, *pc + 4)
					*pc += bit_utils.SignExtU64(imm)
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "mulhsu",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0110011,0b010,0b0000001, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			ToString: func (inst isa.Instruction) string {
				return "mulhsu " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
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
			Name: "mulhu",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0110011,0b011,0b0000001, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			ToString: func (inst isa.Instruction) string {
				return "mulhu " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
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
			Name: "dviu",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0110011,0b100,0b0000001, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			ToString: func (inst isa.Instruction) string {
				return "dviu " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=r.Load(rs1), r.Load(rs2); r.Store(rd, t1 / t2)
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "remu",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0110011,0b111,0b0000001, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			ToString: func (inst isa.Instruction) string {
				return "remu " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=r.Load(rs1), r.Load(rs2); r.Store(rd, t1 % t2)
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "mulw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0111011,0b000,0b0000001, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			ToString: func (inst isa.Instruction) string {
				return "mulw " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=int32(r.Load(rs1)), int32(r.Load(rs2)); r.Store(rd, bit_utils.SignExtU64(int32(t1 * t2)))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "divw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0111011,0b100,0b0000001, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			ToString: func (inst isa.Instruction) string {
				return "divw " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=int32(r.Load(rs1)), int32(r.Load(rs2)); r.Store(rd, bit_utils.SignExtU64(int32(t1 / t2)))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "divuw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0111011,0b101,0b0000001, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			ToString: func (inst isa.Instruction) string {
				return "divuw " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=uint32(r.Load(rs1)), uint32(r.Load(rs2)); r.Store(rd, bit_utils.SignExtU64(int32(t1 / t2)))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "remw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0111011,0b110,0b0000001, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			ToString: func (inst isa.Instruction) string {
				return "remw " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=int32(r.Load(rs1)), int32(r.Load(rs2)); r.Store(rd, bit_utils.SignExtU64(int32(t1 % t2)))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},{
			Name: "remuw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0111011,0b111,0b0000001, inst)
			},
			Inst: func (inst uint32) isa.Instruction {
				return isa.NewRInstruction(inst)
			},
			ToString: func (inst isa.Instruction) string {
				return "remuw " + isa.ToDSSFormat(inst)
			},
			Exec: func (sim *Simulator, inst isa.Instruction){
				m, r, pc := &sim.M, &sim.R, &sim.PC
				_, _, _ = m, r, pc // Ensure any variable is used once
				if t, ok := inst.(isa.RInstruction); ok{
					rs1, rs2, rd := t.Rs1, t.Rs2, t.Rd
					_, _, _ = rs1, rs2, rd // Ensure any variable is used once
					t1, t2:=uint32(r.Load(rs1)), uint32(r.Load(rs2)); r.Store(rd, bit_utils.SignExtU64(int32(t1 % t2)))
					*pc += 4
				}else{
					log.Panicln("Instruction Type mismatch!")
				}
				return
				},
		},
}
