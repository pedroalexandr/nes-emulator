package components

const RAM_SIZE_KB = 2 * 1024

type Bus struct {
	cpu                CPU6502
	ppu                PPU
	cartridge          *Cartridge
	cpuRAM             [RAM_SIZE_KB]uint8
	systemClockCounter uint32
}

func NewBus() *Bus {
	newBus := &Bus{
		cpu:    CPU6502{},
		cpuRAM: [RAM_SIZE_KB]uint8{},
	}

	newBus.cpu.ConnectBus(newBus)

	return newBus
}

func (this *Bus) CPUWrite(addr uint16, data uint8) {
	isWithinCPUAddressRange := addr >= 0x0000 && addr <= 0x1FFF
	isWithinPPUAddressRange := addr >= 0x2000 && addr <= 0x3FFF

	if isWithinCPUAddressRange {
		this.cpuRAM[addr&0x7FF] = data
	} else if isWithinPPUAddressRange {
		this.ppu.CPUWrite(addr&0x0007, &data)
	}
}

func (this *Bus) CPURead(addr uint16, readOnly bool) uint8 {
	var data uint8 = 0x00
	isWithinCPUAddressRange := addr >= 0x0000 && addr <= 0x1FFF
	isWithinPPUAddressRange := addr >= 0x2000 && addr <= 0x3FFF

	if isWithinCPUAddressRange {
		data = this.cpuRAM[addr&0x7FF]
	} else if isWithinPPUAddressRange {
		data = this.ppu.CPURead(addr&0x0007, readOnly)
	}

	return data
}

func (this *Bus) InsertCartridge(cartridge *Cartridge) {
	this.cartridge = cartridge
	this.ppu.ConnectCartridge(cartridge)
}

func (this *Bus) Reset() {
	this.cpu.ResetSignal()
	this.systemClockCounter = 0
}

func (this *Bus) Clock() {

}
