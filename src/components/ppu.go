package components

type PPU struct {
	cartridge         *Cartridge
	vram_nameTable    [2][1024]uint8
	vram_paletteTable [32]uint8
	vram_patternTable [2][4096]uint8 // for future implementation
}

func (this *PPU) CPUWrite(addr uint16, data *uint8) {
	switch addr {
	case 0x0000:
		break
	case 0x0001:
		break
	case 0x0002:
		break
	case 0x0003:
		break
	case 0x0004:
		break
	case 0x0005:
		break
	case 0x0006:
		break
	case 0x0007:
		break
	}
}

func (this *PPU) CPURead(addr uint16, readOnly bool) uint8 {
	var data uint8 = 0x00

	switch addr {
	case 0x0000:
		break
	case 0x0001:
		break
	case 0x0002:
		break
	case 0x0003:
		break
	case 0x0004:
		break
	case 0x0005:
		break
	case 0x0006:
		break
	case 0x0007:
		break
	}

	return data
}

func (this *PPU) PPUWrite(addr uint16, data *uint8) {
	addr &= 0x3FFF
}

func (this *PPU) PPURead(addr uint16, readOnly bool) uint8 {
	var data uint8 = 0x00
	addr &= 0x3FFF
	return data
}

func (this *PPU) ConnectCartridge(cartridge *Cartridge) {
	this.cartridge = cartridge
}

func (this *PPU) clock() {

}
