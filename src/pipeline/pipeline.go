package pipeline

import (
	"errors"
	"isa"
	"log"
	"reflect"
	"simulator"
)

type Pipe struct {
	*simulator.Simulator

	// stage registers
	WReg WriteBackReg
	MReg AccessMemoryReg
	EReg ExecuteReg
	DReg DecodeReg

	IsWaitingIndirectJump bool

	// combinatorial logic circuit
	M_IsPredictError bool
	M_UnselectedPC   uint64
	M_IsBranch       bool
	M_IsIndirectJump bool
	M_valE           uint64
	m_valM           uint64
	M_dstE, M_dstM   uint8
	E_dstE, E_dstM   uint8
	e_valE           uint64
}

func (p *Pipe) WriteBack() {
	if !p.WReg.IsBubble {
		p.R.Store(p.WReg.DstE, p.WReg.ValE)
		p.R.Store(p.WReg.DstM, p.WReg.ValM)
	}
	p.WReg.Status = &NormalStatus{}
}

func (p *Pipe) AccessMemory() {
	if p.MReg.IsBubble {
		p.M_IsBranch = false
		p.M_IsIndirectJump = false
		p.M_IsPredictError = false
		p.M_dstE = 0
		p.M_dstM = 0
	} else {
		p.M_IsBranch = p.MReg.AccumulatedPeriod == 0 && p.MReg.MicroInst.IsBranch
		p.M_IsIndirectJump = p.MReg.AccumulatedPeriod == 0 && p.MReg.MicroInst.IsIndirectJump
		p.M_IsPredictError = p.MReg.IsPredictError
		p.M_dstE = p.MReg.DstE
		p.M_dstM = p.MReg.DstM
		p.M_UnselectedPC = p.MReg.UnselectedPC
		p.M_valE = p.MReg.ValE
	}

	if p.M_IsPredictError {
		p.EReg.IsBubble = true
		p.DReg.IsBubble = true
	}

	if p.MReg.IsBubble {
		p.WReg.IsBubble = true
		p.MReg.Status = &NormalStatus{}
		return
	}
	p.MReg.AccumulatedPeriod += Pipeline_step_period
	if p.MReg.AccumulatedPeriod < p.MReg.MicroInst.MStagePeriod {
		p.WReg.IsBubble = true
		p.MReg.Status = &StallStatus{}
		return
	}
	memAccessFunc := p.MReg.MicroInst.MemoryAccessFunction
	p.m_valM = memAccessFunc(&p.M, p.MReg.ValE, p.MReg.ValRs2)
	p.WReg = WriteBackReg{
		PipeReg{&NormalStatus{}, false, 0},
		p.MReg.MicroInst,
		p.MReg.ValE, p.m_valM,
		p.MReg.DstE, p.MReg.DstM,
	}
	p.MReg.Status = &NormalStatus{}
}

func (p *Pipe) Execute() {
	if p.EReg.IsBubble {
		p.E_dstE = 0
		p.E_dstM = 0
	} else {
		p.E_dstE = p.EReg.DstE
		p.E_dstM = p.EReg.DstM
	}

	if p.MReg.Status != nil && p.MReg.Status.ToString() == "STALL" {
		p.EReg.Status = &StallStatus{}
		return
	}
	if p.EReg.IsBubble {
		p.MReg.IsBubble = true
		p.EReg.Status = &NormalStatus{}
		return
	}
	p.EReg.AccumulatedPeriod += Pipeline_step_period
	if p.EReg.AccumulatedPeriod < p.EReg.MicroInst.EStagePeriod {
		p.MReg.IsBubble = true
		p.EReg.Status = &StallStatus{}
		return
	}
	ValA, ValB, Imm, ValC := p.EReg.ValRs1, p.EReg.ValRs2, p.EReg.Imm, p.EReg.ValC
	ALUFunc := p.EReg.MicroInst.ALUFunction
	isPredictError, ValE := ALUFunc(ValA, ValB, Imm, ValC)
	p.MReg = AccessMemoryReg{
		PipeReg{&NormalStatus{}, false, 0},
		p.EReg.MicroInst, isPredictError, p.EReg.MicroInst.SelectM_valE(ValE, ValC),
		p.EReg.ValRs2, p.EReg.DstE, p.EReg.DstM, p.EReg.UnselectedPC,
	}
	p.e_valE = ValE
	p.EReg.Status = &NormalStatus{}
}

func (p *Pipe) Decode() {
	if p.EReg.Status != nil && p.EReg.Status.ToString() == "STALL" {
		p.DReg.Status = &StallStatus{}
		return
	}
	if p.DReg.IsBubble {
		p.EReg.IsBubble = true
		p.EReg.Status = &NormalStatus{}
		return
	}
	rs1, rs2 := p.DReg.SrcRs1, p.DReg.SrcRs2
	if (p.E_dstM != 0) && (rs1 == p.E_dstM || rs2 == p.E_dstM) {
		p.EReg.IsBubble = true
		p.DReg.Status = &StallStatus{}
		return
	}
	SelectRegVal := func(index uint8) uint64 {
		if index == 0 {
			return 0
		}
		val := p.R.Load(index)
		if p.M_dstM == index {
			val = p.m_valM
		}
		if p.M_dstE == index {
			val = p.M_valE
		}
		if p.E_dstE == index {
			val = p.e_valE
		}
		return val
	}
	valRs1, valRs2 := SelectRegVal(rs1), SelectRegVal(rs2)
	p.EReg = ExecuteReg{
		PipeReg{&NormalStatus{}, false, 0},
		p.DReg.MicroInst, valRs1, valRs2, p.DReg.ValC, p.DReg.Imm,
		p.DReg.DstE, p.DReg.DstM, p.DReg.UnselectedPC,
	}
	p.DReg.Status = &NormalStatus{}
}

func (p *Pipe) HasFinished() bool {
	return p.PC == 0 &&
		p.DReg.IsBubble &&
		p.EReg.IsBubble &&
		p.MReg.IsBubble &&
		p.WReg.IsBubble
}

func Match(code uint32) (item *MicroAction, err error) {
	for _, item := range MicroActionConfigTable {
		if item.Pred(code) == true {
			return item.Parse(code), nil
		}
	}
	return nil, errors.New("unknown instruction format")
}

func (p *Pipe) Fetch() (*MicroAction, error) {
	if p.DReg.Status != nil && p.DReg.Status.ToString() == "STALL" {
		return nil, nil
	}
	if p.IsWaitingIndirectJump && !p.M_IsIndirectJump {
		p.DReg.IsBubble = true
		// Insert bubble for indirect jump
		return nil, nil
	}
	p.IsWaitingIndirectJump = false

	SelectPC := func() uint64 {
		rv := p.PC
		if p.M_IsIndirectJump {
			rv = p.e_valE
		}
		if p.M_IsBranch && p.M_IsPredictError {
			rv = p.M_UnselectedPC
		}
		return rv
	}

	p.PC = SelectPC()
	pc := p.PC

	if pc == 0 {
		p.DReg.IsBubble = true
		// Insert bubble as the program is almost finished
		return nil, nil
	}
	m := p.M
	code := m.LoadU32(pc)
	if isa.InstructionLength(code) != 32 {
		log.Panicln("Only support 32-bit instruction")
	}
	microInst, err := Match(code)
	if err != nil {
		log.Fatalf(" 0x%08x: %v", code, err)
	}

	// the register indexes and Immediate number are assigned to zero if no field was found.
	var rs1, rs2 uint8 = 0, 0
	var Imm int32 = 0
	if val := reflect.ValueOf(microInst.Inst).FieldByName("Rs1"); val.IsValid() {
		rs1 = uint8(val.Uint())
	}
	if val := reflect.ValueOf(microInst.Inst).FieldByName("Rs2"); val.IsValid() {
		rs2 = uint8(val.Uint())
	}
	if val := reflect.ValueOf(microInst.Inst).FieldByName("Imm"); val.IsValid() {
		Imm = int32(val.Int())
	}

	p.DReg = DecodeReg{
		PipeReg{&NormalStatus{}, false, 0},
		*microInst, rs1, rs2, Imm,
		microInst.dstE, microInst.dstM, microInst.ValCFunction(pc, Imm),
		microInst.NegativeOptionPC(pc, Imm),
	}

	if microInst.IsIndirectJump {
		p.IsWaitingIndirectJump = true
	} else {
		p.PC = microInst.PositiveOptionPC(pc, Imm)
	}
	return microInst, nil
}

func (p *Pipe) Initialize() {
	p.WReg.IsBubble = true
	p.MReg.IsBubble = true
	p.EReg.IsBubble = true
	p.DReg.IsBubble = true
}
func (p *Pipe) SingleStep() (*MicroAction, error) {
	if p.HasFinished() {
		return nil, errors.New("the program has finished")
	}
	p.WriteBack()
	p.AccessMemory()
	p.Execute()
	p.Decode()
	return p.Fetch()
}
