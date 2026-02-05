package background

import (
	"talos/hardware/display"
	mem "talos/hardware/memory"
)

const (
	TILE_4BIT uint8 = 32
	TILE_8BIT uint8 = 64

	CHAR_BASE_BLOCK_SIZE uint32 = 0x4000
)

// An 8x8 tile with four bits-per-pixel. Each pair of pixels is is assigned
// right-to-left, meaning the lower 4 bits of a uint8 represent the right pixel,
// while the upper 4 bits represent the left pixel.
type Tile4 [TILE_4BIT]uint8

// An 8x8 tile with eight bits-per-pixel
type Tile8 [TILE_8BIT]uint8

func getCharBaseAddr(base display.CHAR_BASE) uint32 {
	return mem.VRAM_START.Offset(uint32(base) * CHAR_BASE_BLOCK_SIZE)
}

// LoadTile4 loads a single Tile4 into memory at the given offset, "at," from
// character base-block VRAM.
func LoadTile4(base display.CHAR_BASE, at uint16, data Tile4) {
	addr := getCharBaseAddr(base) + uint32(at)*uint32(TILE_4BIT)
	for i := uint32(0); i < uint32(TILE_4BIT)/2; i++ {
		// Combine two bytes into uint16
		val := uint16(data[i*2]) | (uint16(data[i*2+1]) << 8)
		mem.Write16(addr+i*2, val)
	}
}

// LoadTile8 loads a single Tile8 into memory at the given offset, "at," from
// character base-block VRAM.
func LoadTile8(base display.CHAR_BASE, at uint16, data Tile8) {
	addr := getCharBaseAddr(base) + uint32(at)*uint32(TILE_8BIT)
	for i := uint32(0); i < uint32(TILE_8BIT); i++ {
		mem.Write8(addr+i, data[i])
	}
}

// LoadTiles4 loads any number of Tile4's into memory at the given offset, "at,"
// from character base-block VRAM.
func LoadTiles4(base display.CHAR_BASE, start uint16, tiles []Tile4) {
	addr := getCharBaseAddr(base) + uint32(start)*uint32(TILE_4BIT)
	for _, tile := range tiles {
		for i := uint32(0); i < uint32(TILE_4BIT); i++ {
			mem.Write8(addr, tile[i])
			addr++
		}
	}
}

// LoadTiles8 loads any number of Tile8's into memory at the given offset, "at,"
// from character base-block VRAM.
func LoadTiles8(base display.CHAR_BASE, start uint16, tiles []Tile8) {
	addr := getCharBaseAddr(base) + uint32(start)*uint32(TILE_8BIT)
	for _, tile := range tiles {
		for i := uint32(0); i < uint32(TILE_8BIT); i++ {
			mem.Write8(addr, tile[i])
			addr++
		}
	}
}
