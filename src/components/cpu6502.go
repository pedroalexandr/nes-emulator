package components

// FLAGS
const (
	carryBit          = 1 << 0
	zero              = 1 << 1
	disableInterrupts = 1 << 2
	decimalMode       = 1 << 3
	break_            = 1 << 4
	unused            = 1 << 5
	overflow          = 1 << 6
	negative          = 1 << 7
)

type CPU6502 struct {
	bus               *Bus
	accumulatorReg    uint8
	xReg              uint8
	yReg              uint8
	stackPointerReg   uint8
	programCounterReg uint8
	statusReg         uint8
}

func (cpu *CPU6502) ConnectBus(bus *Bus) {
	cpu.bus = bus
}

func (cpu *CPU6502) IMP() {

}

func (cpu *CPU6502) write(addr uint16, data uint8) {
	cpu.bus.write(addr, data)
}

func (cpu *CPU6502) read(addr uint16, readOnly bool) uint8 {
	return cpu.bus.read(addr, false)
}
