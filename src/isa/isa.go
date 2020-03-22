package isa

import (
	"bit_utils"
	"fmt"
	"log"
	"reflect"
	"register"
	"strings"
)

const (
	R_FORMAT_MASK   uint32 = 0b1111111_00000_00000_111_00000_1111111
	ISB_FORMAT_MASK uint32 = 0b00000000000000000_111_00000_1111111
	UJ_FORMAT_MASK  uint32 = 0b0000000000000000000000000_1111111
)

func IsMatchRFormat(opcode uint32, funct3 uint32, funct7 uint32, inst uint32) bool {
	var v = funct7<<25 | funct3<<12 | opcode
	return (inst & R_FORMAT_MASK) == v
}

func IsMatchIFormat(opcode uint32, funct3 uint32, inst uint32, bitConstraints ...[3]uint32) bool {
	var v = funct3<<12 | opcode
	if (inst & ISB_FORMAT_MASK) != v {
		return false
	}
	for _, con := range bitConstraints {
		from, to, val := con[0], con[1], con[2]
		if bit_utils.U32Bits(inst, from, to) != val {
			return false
		}
	}
	return true
}

func IsMatchSFormat(opcode uint32, funct3 uint32, inst uint32) bool {
	var v = funct3<<12 | opcode
	return (inst & ISB_FORMAT_MASK) == v
}

func IsMatchBFormat(opcode uint32, funct3 uint32, inst uint32) bool {
	var v = funct3<<12 | opcode
	return (inst & ISB_FORMAT_MASK) == v
}

func IsMatchJFormat(opcode uint32, inst uint32) bool {
	return (inst & UJ_FORMAT_MASK) == opcode
}

func IsMatchUFormat(opcode uint32, inst uint32) bool {
	return (inst & UJ_FORMAT_MASK) == opcode
}

func InstructionLength(inst uint32) uint32 {
	if bit_utils.U32Bits(inst, 0, 1) != 0b11 {
		return 16
	}
	if bit_utils.U32Bits(inst, 0, 1) == 0b11 && bit_utils.U32Bits(inst, 2, 4) != 0b111 {
		return 32
	}
	log.Panicln("Don't support instruction with more than 32-bit")
	return 0
}

type Instruction interface {
}

type RInstruction struct {
	Rs1, Rs2, Rd uint8
}

type IInstruction struct {
	Rs1, Rd uint8
	Imm     int32
}

type SInstruction struct {
	Rs1, Rs2 uint8
	Imm      int32
}

type BInstruction struct {
	Rs1, Rs2 uint8
	Imm      int32
}

type UInstruction struct {
	Rd  uint8
	Imm int32
}

type JInstruction struct {
	Rd  uint8
	Imm int32
}

func NewRInstruction(inst uint32) RInstruction {
	return RInstruction{
		uint8(bit_utils.U32Bits(inst, 15, 19)),
		uint8(bit_utils.U32Bits(inst, 20, 24)),
		uint8(bit_utils.U32Bits(inst, 7, 11)),
	}
}

func NewIInstruction(inst uint32) IInstruction {
	return IInstruction{
		uint8(bit_utils.U32Bits(inst, 15, 19)),
		uint8(bit_utils.U32Bits(inst, 7, 11)),
		bit_utils.ConcatenateSignExtBits(inst, [2]uint32{20, 31}),
	}
}

func NewSInstruction(inst uint32) SInstruction {
	return SInstruction{
		Rs1: uint8(bit_utils.U32Bits(inst, 15, 19)),
		Rs2: uint8(bit_utils.U32Bits(inst, 20, 24)),
		Imm: bit_utils.ConcatenateSignExtBits(inst, [2]uint32{7, 11}, [2]uint32{25, 31}),
	}
}

func NewBInstruction(inst uint32) BInstruction {
	return BInstruction{
		Rs1: uint8(bit_utils.U32Bits(inst, 15, 19)),
		Rs2: uint8(bit_utils.U32Bits(inst, 20, 24)),
		Imm: bit_utils.ConcatenateSignExtBits(inst, [2]uint32{8, 11}, [2]uint32{25, 30}, [2]uint32{7, 7}, [2]uint32{31, 31}) << 1,
	}
}

func NewUInstruction(inst uint32) UInstruction {
	return UInstruction{
		Rd:  uint8(bit_utils.U32Bits(inst, 7, 11)),
		Imm: bit_utils.ConcatenateSignExtBits(inst, [2]uint32{12, 31}) << 12,
	}
}

func NewJInstruction(inst uint32) JInstruction {
	return JInstruction{
		Rd:  uint8(bit_utils.U32Bits(inst, 7, 11)),
		Imm: bit_utils.ConcatenateSignExtBits(inst, [2]uint32{21, 30}, [2]uint32{20, 20}, [2]uint32{12, 19}, [2]uint32{31, 31}) << 1,
	}
}

//////////////////////////////////////////////////////////////// display format

/*
* DSS format
*  eg. add rd, rs1, rs2
 */
func ToDSSFormat(inst Instruction) string {
	values := reflect.ValueOf(inst)
	return strings.Join([]string{
		register.NamefromIndex(uint8(values.FieldByName("Rd").Uint())),
		register.NamefromIndex(uint8(values.FieldByName("Rs1").Uint())),
		register.NamefromIndex(uint8(values.FieldByName("Rs2").Uint())),
	}, ", ")
}

/*
* DIS format
*  eg. lb rd, offset(rs1)
 */
func ToDISFormat(inst Instruction) string {
	values := reflect.ValueOf(inst)
	return strings.Join([]string{
		register.NamefromIndex(uint8(values.FieldByName("Rd").Uint())),
		fmt.Sprintf("%d", values.FieldByName("Imm").Int()) +
			"(" + register.NamefromIndex(uint8(values.FieldByName("Rs1").Uint())) + ")",
	}, ", ")
}

/*
* DSI format
*  eg. addi rd, rs1, imm
 */
func ToDSIFormat(inst Instruction) string {
	values := reflect.ValueOf(inst)
	return strings.Join([]string{
		register.NamefromIndex(uint8(values.FieldByName("Rd").Uint())),
		register.NamefromIndex(uint8(values.FieldByName("Rs1").Uint())),
		fmt.Sprintf("%d", values.FieldByName("Imm").Int()),
	}, ", ")
}

/*
* SIS format
*  eg. sb rs2, offset(rs1)
 */
func ToSISFormat(inst Instruction) string {
	values := reflect.ValueOf(inst)
	return strings.Join([]string{
		register.NamefromIndex(uint8(values.FieldByName("Rs2").Uint())),
		fmt.Sprintf("%d", values.FieldByName("Imm").Int()) +
			"(" + register.NamefromIndex(uint8(values.FieldByName("Rs1").Uint())) + ")",
	}, ", ")
}

/*
* SSI format
*  eg. beq rs1, rs2, imm
 */
func ToSSIFormat(inst Instruction) string {
	values := reflect.ValueOf(inst)
	return strings.Join([]string{
		register.NamefromIndex(uint8(values.FieldByName("Rs1").Uint())),
		register.NamefromIndex(uint8(values.FieldByName("Rs2").Uint())),
		fmt.Sprintf("%d", values.FieldByName("Imm").Int()),
	}, ", ")
}

/*
* DI format
*  eg. auipc rd, offset
 */
func ToDIFormat(inst Instruction) string {
	values := reflect.ValueOf(inst)
	return strings.Join([]string{
		register.NamefromIndex(uint8(values.FieldByName("Rd").Uint())),
		fmt.Sprintf("%d", values.FieldByName("Imm").Int()),
	}, ", ")
}

