package pipeline

import (
	"fmt"
	"register"
	"strconv"
	"strings"
)

type PipeStatus interface {
	ToString() string
}

type NormalStatus struct {
	PipeStatus
}

func (s *NormalStatus) ToString() string {
	return "NORMAL"
}

type StallStatus struct {
	PipeStatus
}

func (s *StallStatus) ToString() string {
	return "STALL"
}

func mapForEach(arr []string, fn func(it string) string) []string {
	var newArray = []string{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

type PipeReg struct {
	Status            PipeStatus
	IsBubble          bool
	AccumulatedPeriod uint
}

type WriteBackReg struct {
	PipeReg
	MicroInst  MicroAction
	ValE, ValM uint64
	DstE, DstM uint8
}

func (r *WriteBackReg) ToString() string {
	rv := "[   WriteBack  ] "
	if r.IsBubble {
		return rv + "(bubble)"
	}
	rv += strings.Join(mapForEach(
		[]string{
			"Inst: " + r.MicroInst.InstToString(),
			"STATUS: " + r.Status.ToString(),
			"ValE: " + "0x" + strconv.FormatUint(r.ValE, 16),
			"ValM: " + "0x" + strconv.FormatUint(r.ValM, 16),
			"DstE: " + register.NamefromIndex(r.DstE),
			"DstM: " + register.NamefromIndex(r.DstM),
		}, func(it string) string { return fmt.Sprintf("%-20s", it) }),
		" ")
	return rv
}

type AccessMemoryReg struct {
	PipeReg
	MicroInst      MicroAction
	IsPredictError bool
	ValE, ValRs2   uint64
	DstE, DstM     uint8
	UnselectedPC   uint64
}

func (r *AccessMemoryReg) ToString() string {
	rv := "[ AccessMemory ] "
	if r.IsBubble {
		return rv + "(bubble)"
	}
	rv += strings.Join(mapForEach(
		[]string{
			"Inst: " + r.MicroInst.InstToString(),
			"STATUS: " + r.Status.ToString() + "(" + strconv.Itoa(int(r.AccumulatedPeriod)) + "/" +
				strconv.Itoa(int(r.MicroInst.MStagePeriod)) + ")",
			"ValE: " + "0x" + strconv.FormatUint(r.ValE, 16),
			"ValRs2: " + "0x" + strconv.FormatUint(r.ValRs2, 16),
			"DstE: " + register.NamefromIndex(r.DstE),
			"DstM: " + register.NamefromIndex(r.DstM),
			"UnselectedPC: " + "0x" + strconv.FormatUint(r.UnselectedPC, 16),
		}, func(it string) string { return fmt.Sprintf("%-20s", it) }),
		" ")
	return rv
}

type ExecuteReg struct {
	PipeReg
	MicroInst      MicroAction
	ValRs1, ValRs2 uint64
	ValC           uint64
	Imm            int32
	DstE, DstM     uint8
	UnselectedPC   uint64
}

func (r *ExecuteReg) ToString() string {
	rv := "[   Execute    ] "
	if r.IsBubble {
		return rv + "(bubble)"
	}
	rv += strings.Join(mapForEach(
		[]string{
			"Inst: " + r.MicroInst.InstToString(),
			"STATUS: " + r.Status.ToString() + "(" + strconv.Itoa(int(r.AccumulatedPeriod)) + "/" +
				strconv.Itoa(int(r.MicroInst.EStagePeriod)) + ")",
			"ValRs1: " + "0x" + strconv.FormatUint(r.ValRs1, 16),
			"ValRs2: " + "0x" + strconv.FormatUint(r.ValRs2, 16),
			"ValC: " + "0x" + strconv.FormatUint(r.ValC, 16),
			"Imm: " + strconv.FormatInt(int64(r.Imm), 10),
			"DstE: " + register.NamefromIndex(r.DstE),
			"DstM: " + register.NamefromIndex(r.DstM),
			"UnselectedPC: " + "0x" + strconv.FormatUint(r.UnselectedPC, 16),
		}, func(it string) string { return fmt.Sprintf("%-20s", it) }),
		" ")
	return rv
}

type DecodeReg struct {
	PipeReg
	MicroInst      MicroAction
	SrcRs1, SrcRs2 uint8
	Imm            int32
	DstE, DstM     uint8
	ValC           uint64
	UnselectedPC   uint64
}

func (r *DecodeReg) ToString() string {
	rv := "[   Decode     ] "
	if r.IsBubble {
		return rv + "(bubble)"
	}
	rv += strings.Join(mapForEach(
		[]string{
			"Inst: " + r.MicroInst.InstToString(),
			"STATUS: " + r.Status.ToString(),
			"SrcRs1: " + register.NamefromIndex(r.SrcRs1),
			"SrcRs2: " + register.NamefromIndex(r.SrcRs2),
			"ValC: " + "0x" + strconv.FormatUint(r.ValC, 16),
			"Imm: " + strconv.FormatInt(int64(r.Imm), 10),
			"DstE: " + register.NamefromIndex(r.DstE),
			"ValM: " + register.NamefromIndex(r.DstM),
			"UnselectedPC: " + "0x" + strconv.FormatUint(r.UnselectedPC, 16),
		}, func(it string) string { return fmt.Sprintf("%-20s", it) }),
		" ")
	return rv
}
