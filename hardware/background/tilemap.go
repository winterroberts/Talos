package background

import (
	"talos/hardware/display"
	mem "talos/hardware/memory"
)

const (
	h_bit_offset        uint16 = 10
	v_bit_offset        uint16 = 11
	palette_bits_offset uint16 = 12

	tile_num_bits uint16 = 0x3FF
	palette_bits  uint16 = 0xF
)

type TileMapEntry uint16

func NewTileMapEntry(tile uint16, hflip, vflip bool, pallete uint8) TileMapEntry {
	e := TileMapEntry((tile & tile_num_bits) |
		((uint16(pallete) & palette_bits) << palette_bits_offset))

	if hflip {
		e |= 1 << h_bit_offset
	}
	if vflip {
		e |= 1 << v_bit_offset
	}

	return e
}

const (
	SCREEN_BASE_BLOCK_SIZE = 0x800
)

// Get screen base address
func getScreenBaseAddr(screenBase display.SCREEN_BASE) uint32 {
	return uint32(mem.VRAM_START) + uint32(screenBase)*SCREEN_BASE_BLOCK_SIZE
}

// Set a single tilemap entry
func SetTileMapEntry(screenBase display.SCREEN_BASE, x, y uint8, entry TileMapEntry) {
	addr := getScreenBaseAddr(screenBase) + uint32(y)*32*2 + uint32(x)*2
	mem.Write16(addr, uint16(entry))
}

// Fill entire tilemap with one tile
func FillTileMap(screenBase display.SCREEN_BASE, entry TileMapEntry) {
	addr := getScreenBaseAddr(screenBase)
	for i := uint32(0); i < 32*32; i++ {
		mem.Write16(addr+i*2, uint16(entry))
	}
}

// Load a tilemap (32x32 entries = 2KB)
func LoadTileMap(screenBase display.SCREEN_BASE, entries [32][32]TileMapEntry) {
	addr := getScreenBaseAddr(screenBase)
	for y := uint8(0); y < 32; y++ {
		for x := uint8(0); x < 32; x++ {
			mem.Write16(addr, uint16(entries[y][x]))
			addr += 2
		}
	}
}
