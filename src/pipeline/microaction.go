// Code generated, DO NOT EDIT.
package pipeline

import (
		"bit_utils"
		"isa"
		"math/bits"
		"memory"
)

const (
		XLEN = 64
)

type MicroAction struct {
	Name string
	Inst isa.Instruction
	IsBranch bool
	IsIndirectJump bool
	MemoryAccessFunction func(m *memory.Memory64,addr uint64, input uint64) (output uint64)
	ALUFunction func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64)
	ValCFunction func(pc uint64, Imm int32) (valC uint64)
	dstE, dstM uint8
	PositiveOptionPC func(pc uint64, Imm int32) uint64
	NegativeOptionPC func(pc uint64, Imm int32) uint64
	MStagePeriod uint
	EStagePeriod uint
	InstToString func() string
}

var MicroActionConfigTable = []struct{
	Name string
	Pred func (code uint32) bool
	Parse func (code uint32) *MicroAction
}{{
			Name: "add",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x0,0x00, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "add",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1+t2
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "add " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "mul",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x0,0x01, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "mul",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						_,lo:=bit_utils.MulI64I64(int64(t1),int64(t2)); output=uint64(lo)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "mul " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "sub",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x0,0x20, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "sub",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1-t2
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "sub " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "sll",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x1,0x00, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "sll",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1<<t2
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "sll " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "mulh",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x1,0x01, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "mulh",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						hi,_:=bit_utils.MulI64I64(int64(t1),int64(t2)); output=uint64(hi)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "mulh " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "slt",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x2,0x00, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "slt",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						if t1<t2{output=1}else{output=0}
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "slt " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "xor",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x4,0x00, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "xor",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1^t2
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "xor " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "div",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x4,0x01, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "div",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						 if t2==0 {output=bit_utils.SignExtU64(-1)} else if int64(t1)==-(1<<(XLEN-1)) && int64(t2)==-1 {output=uint64(1)<<(XLEN-1)} else {output=uint64(int64(t1) / int64(t2))}
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: 15,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "div " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "srl",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x5,0x00, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "srl",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1>>t2
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "srl " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "sra",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x5,0x20, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "sra",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=uint64(int64(t1)>>t2)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "sra " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "or",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x6,0x00, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "or",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1|t2
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "or " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "rem",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x6,0x01, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "rem",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						if t2==0 {output=t1} else if int64(t1)==-(1<<(XLEN-1)) && int64(t2)==-1 {output=0} else {output=uint64(int64(t1)%int64(t2))}
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: 15,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "rem " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "and",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0x33,0x7,0x00, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "and",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1&t2
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "and " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "lb",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x03,0x0, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "lb",
					Inst: isa.NewIInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						output=uint64(int64(int8(m.LoadU8(addr))))
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1 + bit_utils.SignExtU64(Imm)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Read_data_memory_period,
					EStagePeriod: Pipeline_step_period,
					dstE: 0,
					dstM: isa.NewIInstruction(code).Rd,
					InstToString:  func() string{
						return "lb " + isa.ToDISFormat(isa.NewIInstruction(code))
					},
				}
			},
		},
	{
			Name: "lh",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x03,0x1, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "lh",
					Inst: isa.NewIInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						output=uint64(int64(int16(m.LoadU16(addr))))
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1 + bit_utils.SignExtU64(Imm)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Read_data_memory_period,
					EStagePeriod: Pipeline_step_period,
					dstE: 0,
					dstM: isa.NewIInstruction(code).Rd,
					InstToString:  func() string{
						return "lh " + isa.ToDISFormat(isa.NewIInstruction(code))
					},
				}
			},
		},
	{
			Name: "lw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x03,0x2, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "lw",
					Inst: isa.NewIInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						output=uint64(int64(int32(m.LoadU32(addr))))
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1 + bit_utils.SignExtU64(Imm)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Read_data_memory_period,
					EStagePeriod: Pipeline_step_period,
					dstE: 0,
					dstM: isa.NewIInstruction(code).Rd,
					InstToString:  func() string{
						return "lw " + isa.ToDISFormat(isa.NewIInstruction(code))
					},
				}
			},
		},
	{
			Name: "ld",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x03,0x3, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "ld",
					Inst: isa.NewIInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						output=uint64(int64(m.LoadU64(addr)))
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1 + bit_utils.SignExtU64(Imm)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Read_data_memory_period,
					EStagePeriod: Pipeline_step_period,
					dstE: 0,
					dstM: isa.NewIInstruction(code).Rd,
					InstToString:  func() string{
						return "ld " + isa.ToDISFormat(isa.NewIInstruction(code))
					},
				}
			},
		},
	{
			Name: "lwu",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x03,0b110, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "lwu",
					Inst: isa.NewIInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						output=uint64(m.LoadU32(addr))
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1 + bit_utils.UnSignExtU64(Imm)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: 0,
					dstM: isa.NewIInstruction(code).Rd,
					InstToString:  func() string{
						return "lwu " + isa.ToDISFormat(isa.NewIInstruction(code))
					},
				}
			},
		},
	{
			Name: "addi",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x13,0x0, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "addi",
					Inst: isa.NewIInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1 + bit_utils.SignExtU64(Imm)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewIInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "addi " + isa.ToDSIFormat(isa.NewIInstruction(code))
					},
				}
			},
		},
	{
			Name: "slli",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x13,0x1, inst,[3]uint32{26, 31, 0})
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "slli",
					Inst: isa.NewIInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1 << bit_utils.U32Bits(uint32(Imm), 0, 5)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewIInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "slli " + isa.ToDSIFormat(isa.NewIInstruction(code))
					},
				}
			},
		},
	{
			Name: "slti",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x13,0x2, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "slti",
					Inst: isa.NewIInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						if int64(t1)<int64(Imm){output=1}else{output=0}
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewIInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "slti " + isa.ToDSIFormat(isa.NewIInstruction(code))
					},
				}
			},
		},
	{
			Name: "xori",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x13,0x4, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "xori",
					Inst: isa.NewIInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1 ^ bit_utils.SignExtU64(Imm)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewIInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "xori " + isa.ToDSIFormat(isa.NewIInstruction(code))
					},
				}
			},
		},
	{
			Name: "srli",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x13,0x5, inst,[3]uint32{26, 31, 0})
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "srli",
					Inst: isa.NewIInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1 >> bit_utils.U32Bits(uint32(Imm), 0, 5)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewIInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "srli " + isa.ToDSIFormat(isa.NewIInstruction(code))
					},
				}
			},
		},
	{
			Name: "srai",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x13,0x5, inst,[3]uint32{26, 31, 0x10})
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "srai",
					Inst: isa.NewIInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=uint64(int64(t1) >>bit_utils.U32Bits(uint32(Imm), 0, 5))
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewIInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "srai " + isa.ToDSIFormat(isa.NewIInstruction(code))
					},
				}
			},
		},
	{
			Name: "ori",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x13,0x6, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "ori",
					Inst: isa.NewIInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1 | bit_utils.SignExtU64(Imm)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewIInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "ori " + isa.ToDSIFormat(isa.NewIInstruction(code))
					},
				}
			},
		},
	{
			Name: "andi",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x13,0x7, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "andi",
					Inst: isa.NewIInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1 & bit_utils.SignExtU64(Imm)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewIInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "andi " + isa.ToDSIFormat(isa.NewIInstruction(code))
					},
				}
			},
		},
	{
			Name: "addiw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x1B,0x0, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "addiw",
					Inst: isa.NewIInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=bit_utils.SignExtU64(int32(t1) + Imm)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewIInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "addiw " + isa.ToDSIFormat(isa.NewIInstruction(code))
					},
				}
			},
		},
	{
			Name: "slliw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0b0011011,0b001, inst,[3]uint32{25, 31, 0x0})
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "slliw",
					Inst: isa.NewIInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=bit_utils.SignExtU64(int32(uint32(t1) << bit_utils.U32Bits(uint32(Imm), 0, 4)))
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewIInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "slliw " + isa.ToDSIFormat(isa.NewIInstruction(code))
					},
				}
			},
		},
	{
			Name: "srliw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0b0011011,0b101, inst,[3]uint32{25, 31, 0x0})
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "srliw",
					Inst: isa.NewIInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=bit_utils.SignExtU64(int32(uint32(t1) >> bit_utils.U32Bits(uint32(Imm), 0, 4)))
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewIInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "srliw " + isa.ToDSIFormat(isa.NewIInstruction(code))
					},
				}
			},
		},
	{
			Name: "sraiw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0b0011011,0b101, inst,[3]uint32{25, 31, 0b0100000})
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "sraiw",
					Inst: isa.NewIInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output= bit_utils.SignExtU64(int32(t1) >> bit_utils.U32Bits(uint32(Imm), 0, 4))
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewIInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "sraiw " + isa.ToDSIFormat(isa.NewIInstruction(code))
					},
				}
			},
		},
	{
			Name: "addw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0111011,0b000,0b0000000, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "addw",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=bit_utils.SignExtU64(int32(t1)+int32(t2))
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "addw " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "subw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0111011,0b000,0b0100000, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "subw",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=bit_utils.SignExtU64(int32(t1)-int32(t2))
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "subw " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "sllw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0111011,0b001,0b0000000, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "sllw",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						shamt:=bit_utils.U32Bits(uint32(t2), 0, 4); output=bit_utils.SignExtU64(int32(uint32(t1)<<shamt))
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "sllw " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "srlw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0111011,0b101,0b0000000, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "srlw",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						shamt:=bit_utils.U32Bits(uint32(t2), 0, 4); output=bit_utils.SignExtU64(int32(uint32(t1)>>shamt))
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "srlw " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "sraw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0111011,0b101,0b0100000, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "sraw",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						shamt:=bit_utils.U32Bits(uint32(t2), 0, 4); output=bit_utils.SignExtU64(int32(t1)>>shamt)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "sraw " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "jalr",
			Pred: func (inst uint32) bool {
				return isa.IsMatchIFormat(0x67,0x0, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "jalr",
					Inst: isa.NewIInstruction(code),
					IsBranch: false,
					IsIndirectJump: true,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=bit_utils.SetBit(t1 + bit_utils.SignExtU64(Imm), 0, 0)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewIInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "jalr " + isa.ToDSIFormat(isa.NewIInstruction(code))
					},
				}
			},
		},
	{
			Name: "sb",
			Pred: func (inst uint32) bool {
				return isa.IsMatchSFormat(0x23,0x0, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "sb",
					Inst: isa.NewSInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						m.StoreU8(addr, uint8(input))
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1+bit_utils.SignExtU64(Imm)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Store_data_memory_period,
					EStagePeriod: Pipeline_step_period,
					dstE: 0,
					dstM: 0,
					InstToString:  func() string{
						return "sb " + isa.ToSISFormat(isa.NewSInstruction(code))
					},
				}
			},
		},
	{
			Name: "sh",
			Pred: func (inst uint32) bool {
				return isa.IsMatchSFormat(0x23,0x1, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "sh",
					Inst: isa.NewSInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						m.StoreU16(addr, uint16(input))
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1+bit_utils.SignExtU64(Imm)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Store_data_memory_period,
					EStagePeriod: Pipeline_step_period,
					dstE: 0,
					dstM: 0,
					InstToString:  func() string{
						return "sh " + isa.ToSISFormat(isa.NewSInstruction(code))
					},
				}
			},
		},
	{
			Name: "sw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchSFormat(0x23,0x2, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "sw",
					Inst: isa.NewSInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						m.StoreU32(addr, uint32(input))
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1+bit_utils.SignExtU64(Imm)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Store_data_memory_period,
					EStagePeriod: Pipeline_step_period,
					dstE: 0,
					dstM: 0,
					InstToString:  func() string{
						return "sw " + isa.ToSISFormat(isa.NewSInstruction(code))
					},
				}
			},
		},
	{
			Name: "sd",
			Pred: func (inst uint32) bool {
				return isa.IsMatchSFormat(0x23,0x3, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "sd",
					Inst: isa.NewSInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						m.StoreU64(addr, input)
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1+bit_utils.SignExtU64(Imm)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Store_data_memory_period,
					EStagePeriod: Pipeline_step_period,
					dstE: 0,
					dstM: 0,
					InstToString:  func() string{
						return "sd " + isa.ToSISFormat(isa.NewSInstruction(code))
					},
				}
			},
		},
	{
			Name: "beq",
			Pred: func (inst uint32) bool {
				return isa.IsMatchBFormat(0x63,0x0, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "beq",
					Inst: isa.NewBInstruction(code),
					IsBranch: true,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						isPredictError=!(t1==t2)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+bit_utils.SignExtU64(Imm)
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: 0,
					dstM: 0,
					InstToString:  func() string{
						return "beq " + isa.ToSSIFormat(isa.NewBInstruction(code))
					},
				}
			},
		},
	{
			Name: "bne",
			Pred: func (inst uint32) bool {
				return isa.IsMatchBFormat(0x63,0x1, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "bne",
					Inst: isa.NewBInstruction(code),
					IsBranch: true,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						isPredictError=!(t1!=t2)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+bit_utils.SignExtU64(Imm)
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: 0,
					dstM: 0,
					InstToString:  func() string{
						return "bne " + isa.ToSSIFormat(isa.NewBInstruction(code))
					},
				}
			},
		},
	{
			Name: "blt",
			Pred: func (inst uint32) bool {
				return isa.IsMatchBFormat(0x63,0x4, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "blt",
					Inst: isa.NewBInstruction(code),
					IsBranch: true,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						isPredictError=!(int64(t1)<int64(t2))
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+bit_utils.SignExtU64(Imm)
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: 0,
					dstM: 0,
					InstToString:  func() string{
						return "blt " + isa.ToSSIFormat(isa.NewBInstruction(code))
					},
				}
			},
		},
	{
			Name: "bge",
			Pred: func (inst uint32) bool {
				return isa.IsMatchBFormat(0x63,0x5, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "bge",
					Inst: isa.NewBInstruction(code),
					IsBranch: true,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						isPredictError=!(int64(t1)>=int64(t2))
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+bit_utils.SignExtU64(Imm)
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: 0,
					dstM: 0,
					InstToString:  func() string{
						return "bge " + isa.ToSSIFormat(isa.NewBInstruction(code))
					},
				}
			},
		},
	{
			Name: "auipc",
			Pred: func (inst uint32) bool {
				return isa.IsMatchUFormat(0x17, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "auipc",
					Inst: isa.NewUInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=ValC
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						valC=pc + bit_utils.SignExtU64(Imm)
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewUInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "auipc " + isa.ToDIFormat(isa.NewUInstruction(code))
					},
				}
			},
		},
	{
			Name: "lui",
			Pred: func (inst uint32) bool {
				return isa.IsMatchUFormat(0x37, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "lui",
					Inst: isa.NewUInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=ValC
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						valC=bit_utils.SignExtU64(Imm)
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewUInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "lui " + isa.ToDIFormat(isa.NewUInstruction(code))
					},
				}
			},
		},
	{
			Name: "jal",
			Pred: func (inst uint32) bool {
				return isa.IsMatchJFormat(0x6f, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "jal",
					Inst: isa.NewJInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=ValC
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						valC=pc+4
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+bit_utils.SignExtU64(Imm)
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewJInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "jal " + isa.ToDIFormat(isa.NewJInstruction(code))
					},
				}
			},
		},
	{
			Name: "mulhsu",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0110011,0b010,0b0000001, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "mulhsu",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						hi,_:=bit_utils.MulI64U64(int64(t1),t2); output=uint64(hi)
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "mulhsu " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "mulhu",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0110011,0b011,0b0000001, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "mulhu",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						hi,_:=bits.Mul64(t1,t2); output=hi
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "mulhu " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "divu",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0110011,0b100,0b0000001, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "divu",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1/t2
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: 15,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "divu " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "remu",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0110011,0b111,0b0000001, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "remu",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=t1%t2
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: 15,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "remu " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "mulw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0111011,0b000,0b0000001, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "mulw",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=bit_utils.SignExtU64(int32(t1) * int32(t2))
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: Pipeline_step_period,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "mulw " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "divw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0111011,0b100,0b0000001, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "divw",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=bit_utils.SignExtU64(int32(t1) / int32(t2))
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: 10,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "divw " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "divuw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0111011,0b101,0b0000001, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "divuw",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=bit_utils.SignExtU64(int32(uint32(t1) / uint32(t2)))
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: 10,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "divuw " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "remw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0111011,0b110,0b0000001, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "remw",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=bit_utils.SignExtU64(int32(t1) / int32(t2))
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: 10,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "remw " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	{
			Name: "remuw",
			Pred: func (inst uint32) bool {
				return isa.IsMatchRFormat(0b0111011,0b111,0b0000001, inst)
			},
			Parse: func (code uint32) *MicroAction{
				return &MicroAction{
					Name: "remuw",
					Inst: isa.NewRInstruction(code),
					IsBranch: false,
					IsIndirectJump: false,
					MemoryAccessFunction: func(m *memory.Memory64,addr uint64, input uint64) (output uint64){
						_, _, _ = m, addr, input // Ensure any variable is used once
						output = 0
						
						return
					},
					ALUFunction: func(t1 uint64, t2 uint64, Imm int32, ValC uint64) (isPredictError bool,output uint64){
						_, _, _, _ = t1, t2, Imm, ValC // Ensure any variable is used once
						isPredictError = false
						output = 0
						output=bit_utils.SignExtU64(int32(uint32(t1) % uint32(t2)))
						return
					},
					ValCFunction: func(pc uint64, Imm int32) (valC uint64){
						_, _ = pc, Imm // Ensure any variable is used once
						valC = 0
						
						return
					},
					PositiveOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return pc+4
					},
					NegativeOptionPC: func(pc uint64, Imm int32) uint64{
						_, _ = pc, Imm // Ensure any variable is used once
						return 0
					},
					MStagePeriod: Pipeline_step_period,
					EStagePeriod: 10,
					dstE: isa.NewRInstruction(code).Rd,
					dstM: 0,
					InstToString:  func() string{
						return "remuw " + isa.ToDSSFormat(isa.NewRInstruction(code))
					},
				}
			},
		},
	
}
