package components

import (
	"os"
	"unsafe"
)

type Cartridge struct {
	fileName  string
	PRGMemory []uint8
	CHRMemory []uint8
	mapperID  uint8
	PRGBanks  uint8
	CHRBanks  uint8
}

func NewCartridge() *Cartridge {

	return &Cartridge{
		fileName: "",
		mapperID: 0,
		PRGBanks: 0,
		CHRBanks: 0,
	}
}

func (cart *Cartridge) openFile(fileName string) {
	type FormatHeader struct {
		name           [4]byte
		PRG_ROM_chunks uint8
		CHR_ROM_chunks uint8
		mapper1        uint8
		mapper2        uint8
		PRG_RAM_size   uint8
		TVsystem1      uint8
		TVsystem2      uint8
		unused         [5]byte
	}

	f, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	headerSize := make([]byte, int(unsafe.Sizeof(&FormatHeader{})))

	header, err := f.Read(headerSize)

	if err != nil {
		panic(err)
	}

}

func (cart *Cartridge) CPUWrite(addr uint16, data *uint8) {

}

func (cart *Cartridge) CPURead(addr uint16, readOnly bool) uint8 {
	return 0
}

func (cart *Cartridge) PPUWrite(addr uint16, data *uint8) {

}

func (cart *Cartridge) PPURead(addr uint16, readOnly bool) uint8 {
	return 0
}
