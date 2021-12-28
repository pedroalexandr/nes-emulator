package components

const RAM_SIZE_KB = 64 * 1024

type Bus struct {
	cpu CPU6502
	ram [RAM_SIZE_KB]uint8
}

func NewBus() *Bus {
	newBus := &Bus{
		cpu: CPU6502{},
		ram: [RAM_SIZE_KB]uint8{},
	}

	newBus.cpu.ConnectBus(newBus)

	return newBus
}

func (bus *Bus) Write(addr uint16, data uint8) {
	if addr >= 0x0000 && addr <= 0xFFFF {
		bus.ram[addr] = data
	}
}

func (bus *Bus) Read(addr uint16, readOnly bool) uint8 {
	if addr >= 0x0000 && addr <= 0xFFFF {
		return bus.ram[addr]
	}
	return 0x00
}
