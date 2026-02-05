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

type Tile4 [TILE_4BIT]uint8

type Tile8 [TILE_8BIT]uint8

func getCharBaseAddr(base display.CHAR_BASE) uint32 {
	return mem.VRAM_START.Offset(uint32(base) * CHAR_BASE_BLOCK_SIZE)
}

// Write tile using Write16 instead
func LoadTile4(base display.CHAR_BASE, at uint16, data Tile4) {
	addr := getCharBaseAddr(base) + uint32(at)*uint32(TILE_4BIT)
	for i := uint32(0); i < uint32(TILE_4BIT)/2; i++ {
		// Combine two bytes into uint16
		val := uint16(data[i*2]) | (uint16(data[i*2+1]) << 8)
		mem.Write16(addr+i*2, val)
	}
}

func LoadTile8(base display.CHAR_BASE, at uint16, data Tile8) {
	addr := getCharBaseAddr(base) + uint32(at)*uint32(TILE_8BIT)
	for i := uint32(0); i < uint32(TILE_8BIT); i++ {
		mem.Write8(addr+i, data[i])
	}
}

func LoadTiles4(base display.CHAR_BASE, start uint16, tiles []Tile4) {
	addr := getCharBaseAddr(base) + uint32(start)*uint32(TILE_4BIT)
	for _, tile := range tiles {
		for i := uint32(0); i < uint32(TILE_4BIT); i++ {
			mem.Write8(addr, tile[i])
			addr++
		}
	}
}

func LoadTiles8(base display.CHAR_BASE, start uint16, tiles []Tile8) {
	addr := getCharBaseAddr(base) + uint32(start)*uint32(TILE_8BIT)
	for _, tile := range tiles {
		for i := uint32(0); i < uint32(TILE_8BIT); i++ {
			mem.Write8(addr, tile[i])
			addr++
		}
	}
}
