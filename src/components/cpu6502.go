package components

import "reflect"

type Flag uint8

// FLAGS
const (
	carryBit          Flag = 1 << 0
	zero              Flag = 1 << 1
	disableInterrupts Flag = 1 << 2
	decimalMode       Flag = 1 << 3
	break_            Flag = 1 << 4
	unused            Flag = 1 << 5
	overflow          Flag = 1 << 6
	negative          Flag = 1 << 7
)

type instruction struct {
	name                        string
	operation                   func() uint8
	addressingMode              func() uint8
	requiredAmountOfClockCycles uint8
}

type CPU6502 struct {
	bus                 *Bus
	accumulatorReg      uint8
	xReg                uint8
	yReg                uint8
	stackPointerReg     uint8
	programCounterReg   uint16
	statusReg           Flag
	fetchedData         uint8
	absoluteAddress     uint16
	relativeAddress     uint16
	relativeAddr        uint16
	opCode              uint8
	amountOfClockCycles uint8
	lookup              []instruction
}

func NewCPU6502() *CPU6502 {
	newCPU6502 := &CPU6502{}

	newCPU6502.lookup = []instruction{
		{"BRK", newCPU6502.BRK, newCPU6502.IMM, 7}, {"ORA", newCPU6502.ORA, newCPU6502.IZX, 6}, {"???", newCPU6502.ERR, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 8}, {"???", newCPU6502.NOP, newCPU6502.IMP, 3}, {"ORA", newCPU6502.ORA, newCPU6502.ZP0, 3}, {"ASL", newCPU6502.ASL, newCPU6502.ZP0, 5}, {"???", newCPU6502.ERR, newCPU6502.IMP, 5}, {"PHP", newCPU6502.PHP, newCPU6502.IMP, 3}, {"ORA", newCPU6502.ORA, newCPU6502.IMM, 2}, {"ASL", newCPU6502.ASL, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 2}, {"???", newCPU6502.NOP, newCPU6502.IMP, 4}, {"ORA", newCPU6502.ORA, newCPU6502.ABS, 4}, {"ASL", newCPU6502.ASL, newCPU6502.ABS, 6}, {"???", newCPU6502.ERR, newCPU6502.IMP, 6},
		{"BPL", newCPU6502.BPL, newCPU6502.REL, 2}, {"ORA", newCPU6502.ORA, newCPU6502.IZY, 5}, {"???", newCPU6502.ERR, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 8}, {"???", newCPU6502.NOP, newCPU6502.IMP, 4}, {"ORA", newCPU6502.ORA, newCPU6502.ZPX, 4}, {"ASL", newCPU6502.ASL, newCPU6502.ZPX, 6}, {"???", newCPU6502.ERR, newCPU6502.IMP, 6}, {"CLC", newCPU6502.CLC, newCPU6502.IMP, 2}, {"ORA", newCPU6502.ORA, newCPU6502.ABY, 4}, {"???", newCPU6502.NOP, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 7}, {"???", newCPU6502.NOP, newCPU6502.IMP, 4}, {"ORA", newCPU6502.ORA, newCPU6502.ABX, 4}, {"ASL", newCPU6502.ASL, newCPU6502.ABX, 7}, {"???", newCPU6502.ERR, newCPU6502.IMP, 7},
		{"JSR", newCPU6502.JSR, newCPU6502.ABS, 6}, {"AND", newCPU6502.AND, newCPU6502.IZX, 6}, {"???", newCPU6502.ERR, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 8}, {"BIT", newCPU6502.BIT, newCPU6502.ZP0, 3}, {"AND", newCPU6502.AND, newCPU6502.ZP0, 3}, {"ROL", newCPU6502.ROL, newCPU6502.ZP0, 5}, {"???", newCPU6502.ERR, newCPU6502.IMP, 5}, {"PLP", newCPU6502.PLP, newCPU6502.IMP, 4}, {"AND", newCPU6502.AND, newCPU6502.IMM, 2}, {"ROL", newCPU6502.ROL, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 2}, {"BIT", newCPU6502.BIT, newCPU6502.ABS, 4}, {"AND", newCPU6502.AND, newCPU6502.ABS, 4}, {"ROL", newCPU6502.ROL, newCPU6502.ABS, 6}, {"???", newCPU6502.ERR, newCPU6502.IMP, 6},
		{"BMI", newCPU6502.BMI, newCPU6502.REL, 2}, {"AND", newCPU6502.AND, newCPU6502.IZY, 5}, {"???", newCPU6502.ERR, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 8}, {"???", newCPU6502.NOP, newCPU6502.IMP, 4}, {"AND", newCPU6502.AND, newCPU6502.ZPX, 4}, {"ROL", newCPU6502.ROL, newCPU6502.ZPX, 6}, {"???", newCPU6502.ERR, newCPU6502.IMP, 6}, {"SEC", newCPU6502.SEC, newCPU6502.IMP, 2}, {"AND", newCPU6502.AND, newCPU6502.ABY, 4}, {"???", newCPU6502.NOP, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 7}, {"???", newCPU6502.NOP, newCPU6502.IMP, 4}, {"AND", newCPU6502.AND, newCPU6502.ABX, 4}, {"ROL", newCPU6502.ROL, newCPU6502.ABX, 7}, {"???", newCPU6502.ERR, newCPU6502.IMP, 7},
		{"RTI", newCPU6502.RTI, newCPU6502.IMP, 6}, {"EOR", newCPU6502.EOR, newCPU6502.IZX, 6}, {"???", newCPU6502.ERR, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 8}, {"???", newCPU6502.NOP, newCPU6502.IMP, 3}, {"EOR", newCPU6502.EOR, newCPU6502.ZP0, 3}, {"LSR", newCPU6502.LSR, newCPU6502.ZP0, 5}, {"???", newCPU6502.ERR, newCPU6502.IMP, 5}, {"PHA", newCPU6502.PHA, newCPU6502.IMP, 3}, {"EOR", newCPU6502.EOR, newCPU6502.IMM, 2}, {"LSR", newCPU6502.LSR, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 2}, {"JMP", newCPU6502.JMP, newCPU6502.ABS, 3}, {"EOR", newCPU6502.EOR, newCPU6502.ABS, 4}, {"LSR", newCPU6502.LSR, newCPU6502.ABS, 6}, {"???", newCPU6502.ERR, newCPU6502.IMP, 6},
		{"BVC", newCPU6502.BVC, newCPU6502.REL, 2}, {"EOR", newCPU6502.EOR, newCPU6502.IZY, 5}, {"???", newCPU6502.ERR, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 8}, {"???", newCPU6502.NOP, newCPU6502.IMP, 4}, {"EOR", newCPU6502.EOR, newCPU6502.ZPX, 4}, {"LSR", newCPU6502.LSR, newCPU6502.ZPX, 6}, {"???", newCPU6502.ERR, newCPU6502.IMP, 6}, {"CLI", newCPU6502.CLI, newCPU6502.IMP, 2}, {"EOR", newCPU6502.EOR, newCPU6502.ABY, 4}, {"???", newCPU6502.NOP, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 7}, {"???", newCPU6502.NOP, newCPU6502.IMP, 4}, {"EOR", newCPU6502.EOR, newCPU6502.ABX, 4}, {"LSR", newCPU6502.LSR, newCPU6502.ABX, 7}, {"???", newCPU6502.ERR, newCPU6502.IMP, 7},
		{"RTS", newCPU6502.RTS, newCPU6502.IMP, 6}, {"ADC", newCPU6502.ADC, newCPU6502.IZX, 6}, {"???", newCPU6502.ERR, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 8}, {"???", newCPU6502.NOP, newCPU6502.IMP, 3}, {"ADC", newCPU6502.ADC, newCPU6502.ZP0, 3}, {"ROR", newCPU6502.ROR, newCPU6502.ZP0, 5}, {"???", newCPU6502.ERR, newCPU6502.IMP, 5}, {"PLA", newCPU6502.PLA, newCPU6502.IMP, 4}, {"ADC", newCPU6502.ADC, newCPU6502.IMM, 2}, {"ROR", newCPU6502.ROR, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 2}, {"JMP", newCPU6502.JMP, newCPU6502.IND, 5}, {"ADC", newCPU6502.ADC, newCPU6502.ABS, 4}, {"ROR", newCPU6502.ROR, newCPU6502.ABS, 6}, {"???", newCPU6502.ERR, newCPU6502.IMP, 6},
		{"BVS", newCPU6502.BVS, newCPU6502.REL, 2}, {"ADC", newCPU6502.ADC, newCPU6502.IZY, 5}, {"???", newCPU6502.ERR, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 8}, {"???", newCPU6502.NOP, newCPU6502.IMP, 4}, {"ADC", newCPU6502.ADC, newCPU6502.ZPX, 4}, {"ROR", newCPU6502.ROR, newCPU6502.ZPX, 6}, {"???", newCPU6502.ERR, newCPU6502.IMP, 6}, {"SEI", newCPU6502.SEI, newCPU6502.IMP, 2}, {"ADC", newCPU6502.ADC, newCPU6502.ABY, 4}, {"???", newCPU6502.NOP, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 7}, {"???", newCPU6502.NOP, newCPU6502.IMP, 4}, {"ADC", newCPU6502.ADC, newCPU6502.ABX, 4}, {"ROR", newCPU6502.ROR, newCPU6502.ABX, 7}, {"???", newCPU6502.ERR, newCPU6502.IMP, 7},
		{"???", newCPU6502.NOP, newCPU6502.IMP, 2}, {"STA", newCPU6502.STA, newCPU6502.IZX, 6}, {"???", newCPU6502.NOP, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 6}, {"STY", newCPU6502.STY, newCPU6502.ZP0, 3}, {"STA", newCPU6502.STA, newCPU6502.ZP0, 3}, {"STX", newCPU6502.STX, newCPU6502.ZP0, 3}, {"???", newCPU6502.ERR, newCPU6502.IMP, 3}, {"DEY", newCPU6502.DEY, newCPU6502.IMP, 2}, {"???", newCPU6502.NOP, newCPU6502.IMP, 2}, {"TXA", newCPU6502.TXA, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 2}, {"STY", newCPU6502.STY, newCPU6502.ABS, 4}, {"STA", newCPU6502.STA, newCPU6502.ABS, 4}, {"STX", newCPU6502.STX, newCPU6502.ABS, 4}, {"???", newCPU6502.ERR, newCPU6502.IMP, 4},
		{"BCC", newCPU6502.BCC, newCPU6502.REL, 2}, {"STA", newCPU6502.STA, newCPU6502.IZY, 6}, {"???", newCPU6502.ERR, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 6}, {"STY", newCPU6502.STY, newCPU6502.ZPX, 4}, {"STA", newCPU6502.STA, newCPU6502.ZPX, 4}, {"STX", newCPU6502.STX, newCPU6502.ZPY, 4}, {"???", newCPU6502.ERR, newCPU6502.IMP, 4}, {"TYA", newCPU6502.TYA, newCPU6502.IMP, 2}, {"STA", newCPU6502.STA, newCPU6502.ABY, 5}, {"TXS", newCPU6502.TXS, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 5}, {"???", newCPU6502.NOP, newCPU6502.IMP, 5}, {"STA", newCPU6502.STA, newCPU6502.ABX, 5}, {"???", newCPU6502.ERR, newCPU6502.IMP, 5}, {"???", newCPU6502.ERR, newCPU6502.IMP, 5},
		{"LDY", newCPU6502.LDY, newCPU6502.IMM, 2}, {"LDA", newCPU6502.LDA, newCPU6502.IZX, 6}, {"LDX", newCPU6502.LDX, newCPU6502.IMM, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 6}, {"LDY", newCPU6502.LDY, newCPU6502.ZP0, 3}, {"LDA", newCPU6502.LDA, newCPU6502.ZP0, 3}, {"LDX", newCPU6502.LDX, newCPU6502.ZP0, 3}, {"???", newCPU6502.ERR, newCPU6502.IMP, 3}, {"TAY", newCPU6502.TAY, newCPU6502.IMP, 2}, {"LDA", newCPU6502.LDA, newCPU6502.IMM, 2}, {"TAX", newCPU6502.TAX, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 2}, {"LDY", newCPU6502.LDY, newCPU6502.ABS, 4}, {"LDA", newCPU6502.LDA, newCPU6502.ABS, 4}, {"LDX", newCPU6502.LDX, newCPU6502.ABS, 4}, {"???", newCPU6502.ERR, newCPU6502.IMP, 4},
		{"BCS", newCPU6502.BCS, newCPU6502.REL, 2}, {"LDA", newCPU6502.LDA, newCPU6502.IZY, 5}, {"???", newCPU6502.ERR, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 5}, {"LDY", newCPU6502.LDY, newCPU6502.ZPX, 4}, {"LDA", newCPU6502.LDA, newCPU6502.ZPX, 4}, {"LDX", newCPU6502.LDX, newCPU6502.ZPY, 4}, {"???", newCPU6502.ERR, newCPU6502.IMP, 4}, {"CLV", newCPU6502.CLV, newCPU6502.IMP, 2}, {"LDA", newCPU6502.LDA, newCPU6502.ABY, 4}, {"TSX", newCPU6502.TSX, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 4}, {"LDY", newCPU6502.LDY, newCPU6502.ABX, 4}, {"LDA", newCPU6502.LDA, newCPU6502.ABX, 4}, {"LDX", newCPU6502.LDX, newCPU6502.ABY, 4}, {"???", newCPU6502.ERR, newCPU6502.IMP, 4},
		{"CPY", newCPU6502.CPY, newCPU6502.IMM, 2}, {"CMP", newCPU6502.CMP, newCPU6502.IZX, 6}, {"???", newCPU6502.NOP, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 8}, {"CPY", newCPU6502.CPY, newCPU6502.ZP0, 3}, {"CMP", newCPU6502.CMP, newCPU6502.ZP0, 3}, {"DEC", newCPU6502.DEC, newCPU6502.ZP0, 5}, {"???", newCPU6502.ERR, newCPU6502.IMP, 5}, {"INY", newCPU6502.INY, newCPU6502.IMP, 2}, {"CMP", newCPU6502.CMP, newCPU6502.IMM, 2}, {"DEX", newCPU6502.DEX, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 2}, {"CPY", newCPU6502.CPY, newCPU6502.ABS, 4}, {"CMP", newCPU6502.CMP, newCPU6502.ABS, 4}, {"DEC", newCPU6502.DEC, newCPU6502.ABS, 6}, {"???", newCPU6502.ERR, newCPU6502.IMP, 6},
		{"BNE", newCPU6502.BNE, newCPU6502.REL, 2}, {"CMP", newCPU6502.CMP, newCPU6502.IZY, 5}, {"???", newCPU6502.ERR, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 8}, {"???", newCPU6502.NOP, newCPU6502.IMP, 4}, {"CMP", newCPU6502.CMP, newCPU6502.ZPX, 4}, {"DEC", newCPU6502.DEC, newCPU6502.ZPX, 6}, {"???", newCPU6502.ERR, newCPU6502.IMP, 6}, {"CLD", newCPU6502.CLD, newCPU6502.IMP, 2}, {"CMP", newCPU6502.CMP, newCPU6502.ABY, 4}, {"NOP", newCPU6502.NOP, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 7}, {"???", newCPU6502.NOP, newCPU6502.IMP, 4}, {"CMP", newCPU6502.CMP, newCPU6502.ABX, 4}, {"DEC", newCPU6502.DEC, newCPU6502.ABX, 7}, {"???", newCPU6502.ERR, newCPU6502.IMP, 7},
		{"CPX", newCPU6502.CPX, newCPU6502.IMM, 2}, {"SBC", newCPU6502.SBC, newCPU6502.IZX, 6}, {"???", newCPU6502.NOP, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 8}, {"CPX", newCPU6502.CPX, newCPU6502.ZP0, 3}, {"SBC", newCPU6502.SBC, newCPU6502.ZP0, 3}, {"INC", newCPU6502.INC, newCPU6502.ZP0, 5}, {"???", newCPU6502.ERR, newCPU6502.IMP, 5}, {"INX", newCPU6502.INX, newCPU6502.IMP, 2}, {"SBC", newCPU6502.SBC, newCPU6502.IMM, 2}, {"NOP", newCPU6502.NOP, newCPU6502.IMP, 2}, {"???", newCPU6502.SBC, newCPU6502.IMP, 2}, {"CPX", newCPU6502.CPX, newCPU6502.ABS, 4}, {"SBC", newCPU6502.SBC, newCPU6502.ABS, 4}, {"INC", newCPU6502.INC, newCPU6502.ABS, 6}, {"???", newCPU6502.ERR, newCPU6502.IMP, 6},
		{"BEQ", newCPU6502.BEQ, newCPU6502.REL, 2}, {"SBC", newCPU6502.SBC, newCPU6502.IZY, 5}, {"???", newCPU6502.ERR, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 8}, {"???", newCPU6502.NOP, newCPU6502.IMP, 4}, {"SBC", newCPU6502.SBC, newCPU6502.ZPX, 4}, {"INC", newCPU6502.INC, newCPU6502.ZPX, 6}, {"???", newCPU6502.ERR, newCPU6502.IMP, 6}, {"SED", newCPU6502.SED, newCPU6502.IMP, 2}, {"SBC", newCPU6502.SBC, newCPU6502.ABY, 4}, {"NOP", newCPU6502.NOP, newCPU6502.IMP, 2}, {"???", newCPU6502.ERR, newCPU6502.IMP, 7}, {"???", newCPU6502.NOP, newCPU6502.IMP, 4}, {"SBC", newCPU6502.SBC, newCPU6502.ABX, 4}, {"INC", newCPU6502.INC, newCPU6502.ABX, 7}, {"???", newCPU6502.ERR, newCPU6502.IMP, 7},
	}

	return newCPU6502
}

func (this *CPU6502) ConnectBus(bus *Bus) {
	this.bus = bus
}

func (this *CPU6502) setFlag(flag Flag, value bool) {
	if value {
		this.statusReg |= flag
	} else {
		this.statusReg &= ^flag
	}
}

func (this *CPU6502) getFlag(flag Flag) uint8 {
	if (this.statusReg & flag) > 0 {
		return 1
	}
	return 0
}

func (this *CPU6502) ClockSignal() {
	if this.amountOfClockCycles == 0 {
		this.opCode = this.Read(this.programCounterReg, false)
		this.programCounterReg++
		this.amountOfClockCycles = this.lookup[this.opCode].requiredAmountOfClockCycles

		additionalCycle1 := (this.lookup[this.opCode].addressingMode)()
		additionalCycle2 := (this.lookup[this.opCode].operation)()

		this.amountOfClockCycles += (additionalCycle1 & additionalCycle2)
	}
	this.amountOfClockCycles--
}

func (this *CPU6502) ResetSignal() {

}

func (this *CPU6502) InterruptRequestSignal() {

}

func (this *CPU6502) NonMaskableInterruptRequestSignal() {

}

func (this *CPU6502) FetchData() uint8 {
	addressingModeIsNotImplied := !(reflect.ValueOf(this.lookup[this.opCode].addressingMode).Pointer() == reflect.ValueOf(this.IMP).Pointer())

	if addressingModeIsNotImplied {
		this.fetchedData = this.Read(this.absoluteAddress, false)
	}

	return this.fetchedData
}

func (this *CPU6502) Write(addr uint16, data uint8) {
	this.bus.Write(addr, data)
}

func (this *CPU6502) Read(addr uint16, readOnly bool) uint8 {
	return this.bus.Read(addr, false)
}

// Addressing Modes
func (this *CPU6502) IMP() uint8 {
	this.fetchedData = this.accumulatorReg
	return 0
}

func (this *CPU6502) IMM() uint8 {
	this.programCounterReg++
	this.absoluteAddress = this.programCounterReg
	return 0
}

func (this *CPU6502) ZP0() uint8 {
	this.absoluteAddress = uint16(this.Read(this.programCounterReg, false))
	this.programCounterReg++
	this.absoluteAddress &= 0x00FF
	return 0
}

func (this *CPU6502) ZPX() uint8 {
	this.absoluteAddress = uint16(this.Read(this.programCounterReg, false) + this.xReg)
	this.programCounterReg++
	this.absoluteAddress &= 0x00FF
	return 0
}

func (this *CPU6502) ZPY() uint8 {
	this.absoluteAddress = uint16(this.Read(this.programCounterReg, false) + this.yReg)
	this.programCounterReg++
	this.absoluteAddress &= 0x00FF
	return 0
}

func (this *CPU6502) REL() uint8 {
	this.relativeAddress = uint16(this.Read(this.programCounterReg, false))
	this.programCounterReg++

	if this.relativeAddress&0x80 != 0 {
		this.relativeAddress |= 0xFF00
	}

	return 0
}

func (this *CPU6502) ABS() uint8 {
	lowByte := uint16(this.Read(this.programCounterReg, false))
	this.programCounterReg++
	highByte := uint16(this.Read(this.programCounterReg, false))
	this.programCounterReg++

	this.absoluteAddress = (highByte << 8) | lowByte

	return 0
}

func (this *CPU6502) ABX() uint8 {
	lowByte := uint16(this.Read(this.programCounterReg, false))
	this.programCounterReg++
	highByte := uint16(this.Read(this.programCounterReg, false))
	this.programCounterReg++

	this.absoluteAddress = (highByte << 8) | lowByte
	this.absoluteAddress += uint16(this.xReg)

	highByteHasChangedDueToOverflow := (this.absoluteAddress & 0xFF00) != (highByte << 8)

	if highByteHasChangedDueToOverflow {
		return 1
	} else {
		return 0
	}
}

func (this *CPU6502) ABY() uint8 {
	lowByte := uint16(this.Read(this.programCounterReg, false))
	this.programCounterReg++
	highByte := uint16(this.Read(this.programCounterReg, false))
	this.programCounterReg++

	this.absoluteAddress = (highByte << 8) | lowByte
	this.absoluteAddress += uint16(this.yReg)

	highByteHasChangedDueToOverflow := (this.absoluteAddress & 0xFF00) != (highByte << 8)

	if highByteHasChangedDueToOverflow {
		return 1
	} else {
		return 0
	}
}

func (this *CPU6502) IND() uint8 {
	ptrlowByte := uint16(this.Read(this.programCounterReg, false))
	this.programCounterReg++
	ptrhighByte := uint16(this.Read(this.programCounterReg, false))
	this.programCounterReg++

	ptr := (ptrhighByte << 8) | ptrlowByte

	if ptrlowByte == 0x00FF {
		this.absoluteAddress = uint16((this.Read(ptr&0xFF00, false) << 8) | this.Read(ptr+0, false))
	} else {
		this.absoluteAddress = uint16((this.Read(ptr+1, false) << 8) | this.Read(ptr+0, false))
	}

	return 0
}

func (this *CPU6502) IZX() uint8 {
	offsetIntoTheZerothPage := this.Read(this.programCounterReg, false)
	this.programCounterReg++

	lowByte := uint16(this.Read(uint16((offsetIntoTheZerothPage+this.xReg))&0x00FF, false))
	highByte := uint16(this.Read(uint16((offsetIntoTheZerothPage+this.xReg+1))&0x00FF, false))

	this.absoluteAddress = (highByte << 8) | lowByte

	return 0
}

func (this *CPU6502) IZY() uint8 {
	offsetIntoTheZerothPage := this.Read(this.programCounterReg, false)
	this.programCounterReg++

	lowByte := uint16(this.Read(uint16(offsetIntoTheZerothPage)&0x00FF, false))
	highByte := uint16(this.Read(uint16(offsetIntoTheZerothPage+1)&0x00FF, false))

	this.absoluteAddress = (highByte << 8) | lowByte
	this.absoluteAddress += uint16(this.yReg)

	highByteHasChangedDueToOverflow := (this.absoluteAddress & 0xFF00) != (highByte << 8)

	if highByteHasChangedDueToOverflow {
		return 1
	} else {
		return 0
	}
}

// Operations
func (this *CPU6502) ERR() uint8 { // Non-legitimate OpCode handling

}

func (this *CPU6502) ADC() uint8 {
	this.FetchData()
	temp := uint16(this.accumulatorReg) + uint16(this.fetchedData) + uint16(this.getFlag(carryBit))
	this.setFlag(carryBit, temp > 255)
	this.setFlag(zero, (temp&0x00FF) == 0)
	this.setFlag(negative, temp&0x80 != 0)

	hasOverflowed := (^(uint16(this.accumulatorReg) ^ uint16(this.fetchedData)) & (uint16(this.accumulatorReg) ^ uint16(temp)) & 0x0080) != 0
	this.setFlag(overflow, hasOverflowed)

	this.accumulatorReg = uint8(temp & 0x00FF)

	return 1
}

func (this *CPU6502) AND() uint8 {
	this.FetchData()
	this.accumulatorReg &= this.fetchedData
	this.setFlag(zero, this.accumulatorReg == 0x00)
	this.setFlag(negative, this.accumulatorReg&0x80 != 0)
	return 1
}

func (this *CPU6502) BCS() uint8 {
	if this.getFlag(carryBit) == 1 {
		this.amountOfClockCycles++
		this.absoluteAddress = this.programCounterReg + this.relativeAddress

		needsToCrossAPageBoundary := (this.absoluteAddress & 0xFF00) != (this.programCounterReg & 0xFF00)

		if needsToCrossAPageBoundary {
			this.amountOfClockCycles++
		}

		this.programCounterReg = this.absoluteAddress
	}
	return 0
}

func (this *CPU6502) BCC() uint8 {
	if this.getFlag(carryBit) == 1 {
		this.amountOfClockCycles++
		this.absoluteAddress = this.programCounterReg + this.relativeAddress

		needsToCrossAPageBoundary := (this.absoluteAddress & 0xFF00) != (this.programCounterReg & 0xFF00)

		if needsToCrossAPageBoundary {
			this.amountOfClockCycles++
		}

		this.programCounterReg = this.absoluteAddress
	}
	return 0
}

func (this *CPU6502) BEQ() uint8 {
	if this.getFlag(carryBit) == 1 {
		this.amountOfClockCycles++
		this.absoluteAddress = this.programCounterReg + this.relativeAddress

		needsToCrossAPageBoundary := (this.absoluteAddress & 0xFF00) != (this.programCounterReg & 0xFF00)

		if needsToCrossAPageBoundary {
			this.amountOfClockCycles++
		}

		this.programCounterReg = this.absoluteAddress
	}
	return 0
}

func (this *CPU6502) BMI() uint8 {
	if this.getFlag(carryBit) == 1 {
		this.amountOfClockCycles++
		this.absoluteAddress = this.programCounterReg + this.relativeAddress

		needsToCrossAPageBoundary := (this.absoluteAddress & 0xFF00) != (this.programCounterReg & 0xFF00)

		if needsToCrossAPageBoundary {
			this.amountOfClockCycles++
		}

		this.programCounterReg = this.absoluteAddress
	}
	return 0
}

func (this *CPU6502) BNE() uint8 {
	if this.getFlag(carryBit) == 1 {
		this.amountOfClockCycles++
		this.absoluteAddress = this.programCounterReg + this.relativeAddress

		needsToCrossAPageBoundary := (this.absoluteAddress & 0xFF00) != (this.programCounterReg & 0xFF00)

		if needsToCrossAPageBoundary {
			this.amountOfClockCycles++
		}

		this.programCounterReg = this.absoluteAddress
	}
	return 0
}

func (this *CPU6502) BPL() uint8 {
	if this.getFlag(carryBit) == 1 {
		this.amountOfClockCycles++
		this.absoluteAddress = this.programCounterReg + this.relativeAddress

		needsToCrossAPageBoundary := (this.absoluteAddress & 0xFF00) != (this.programCounterReg & 0xFF00)

		if needsToCrossAPageBoundary {
			this.amountOfClockCycles++
		}

		this.programCounterReg = this.absoluteAddress
	}
	return 0
}

func (this *CPU6502) BVC() uint8 {
	if this.getFlag(carryBit) == 1 {
		this.amountOfClockCycles++
		this.absoluteAddress = this.programCounterReg + this.relativeAddress

		needsToCrossAPageBoundary := (this.absoluteAddress & 0xFF00) != (this.programCounterReg & 0xFF00)

		if needsToCrossAPageBoundary {
			this.amountOfClockCycles++
		}

		this.programCounterReg = this.absoluteAddress
	}
	return 0
}

func (this *CPU6502) BVS() uint8 {
	if this.getFlag(carryBit) == 1 {
		this.amountOfClockCycles++
		this.absoluteAddress = this.programCounterReg + this.relativeAddress

		needsToCrossAPageBoundary := (this.absoluteAddress & 0xFF00) != (this.programCounterReg & 0xFF00)

		if needsToCrossAPageBoundary {
			this.amountOfClockCycles++
		}

		this.programCounterReg = this.absoluteAddress
	}
	return 0
}

func (this *CPU6502) CLC() uint8 {
	this.setFlag(carryBit, false)
	return 0
}

func (this *CPU6502) CLD() uint8 {
	this.setFlag(decimalMode, false)
	return 0
}

func (this *CPU6502) CLI() uint8 {
	this.setFlag(disableInterrupts, false)
	return 0
}

func (this *CPU6502) CLV() uint8 {
	this.setFlag(overflow, false)
	return 0
}
