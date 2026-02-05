package memory

import (
	"unsafe"
)

type Addr uint32

const (
	KB = 0x400

	MB = 0x100000

	CharOffset = 16 * (KB / 2)

	ScreenOffset = 16 * (KB / 2)

	PALETTEOffset = 16

	TileOffset = 16
)

const (
	// general rom/ram
	BIOS_START Addr = 0x0
	BIOS_END   Addr = 16*KB + BIOS_START - 1

	WSLOW_START Addr = 0x2000000
	WSLOW_END   Addr = 256*KB + WSLOW_START - 1

	WFAST_START Addr = 0x3000000
	WFAST_END        = 32*KB + WFAST_START - 1

	// registers
	IO_START Addr = 0x4000000
	IO_END   Addr = KB + IO_START - 2 - 1

	// video
	PALETTE_BG_START  Addr = 0x5000000
	PALETTE_BG_END    Addr = (KB / 2) + PALETTE_BG_START - 1
	PALETTE_OBJ_START Addr = 0x5000200
	PALETTE_OBJ_END   Addr = (KB / 2) + PALETTE_OBJ_START - 1

	VRAM_START Addr = 0x6000000
	VRAM_END   Addr = 96*KB - 1

	OAM_START Addr = 0x7000000
	OAM_END   Addr = KB + OAM_START - 1

	//gamepak
	FLASH_0_START Addr = 0x8000000
	FLASH_0_END   Addr = 32*MB + FLASH_0_START - 1

	FLASH_1_START Addr = 0xA000000
	FLASH_1_END   Addr = 32*MB + FLASH_1_START - 1

	FLASH_2_START Addr = 0xC000000
	FLASH_2_END   Addr = 32*MB + FLASH_2_START - 1

	SRAM_START Addr = 0xE000000
	SRAM_END   Addr = 64*KB + SRAM_START - 1
)

// Offset gets an offset from a base address
func (a Addr) Offset(o uint32) uint32 {
	return uint32(a) + o
}

// Read16 reads a uint16 from a memory address
func Read16(addr uint32) uint16 {
	return *(*uint16)(unsafe.Pointer(uintptr(addr)))
}

// Write16b writes a uint16 to a memory address
func Write16(addr uint32, value uint16) {
	*(*uint16)(unsafe.Pointer(uintptr(addr))) = value
}

// Read32 reads a uint32 from a memory address
func Read32(addr uint32) uint32 {
	return *(*uint32)(unsafe.Pointer(uintptr(addr)))
}

// Write32 writes a uint32 to a memory address
func Write32(addr uint32, value uint32) {
	*(*uint32)(unsafe.Pointer(uintptr(addr))) = value
}

// Write8 writes a uint8 to a memory address
func Write8(addr uint32, value uint8) {
	*(*uint8)(unsafe.Pointer(uintptr(addr))) = value
}
