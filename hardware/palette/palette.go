package palette

import mem "talos/hardware/memory"

type Color uint16

const (
	intensity_max = 0x1F
	green_offset  = 5
	blue_offset   = 10
)

const (
	COLOR_BLACK      Color = 0x0000
	COLOR_WHITE      Color = 0x7FFF
	COLOR_RED        Color = 0x001F
	COLOR_GREEN      Color = 0x03E0
	COLOR_BLUE       Color = 0x7C00
	COLOR_YELLOW     Color = 0x03FF
	COLOR_CYAN       Color = 0x7FE0
	COLOR_MAGENTA    Color = 0x7C1F
	COLOR_DARK_RED   Color = 0x000F
	COLOR_DARK_GREEN Color = 0x01E0
	COLOR_DARK_BLUE  Color = 0x3C00
	COLOR_GRAY       Color = 0x39E7
)

func RGB(r, g, b uint8) Color {
	return Color(
		(uint16(r) & intensity_max) |
			((uint16(g) & intensity_max) << green_offset) |
			((uint16(b) & intensity_max) << blue_offset))
}

func SetBGColor(idx uint8, color Color) {
	mem.Write16(uint32(mem.PALETTE_BG_START)+uint32(idx)*2, uint16(color))
}

func SetOBJColor(idx uint8, color Color) {
	mem.Write16(uint32(mem.PALETTE_OBJ_START)+uint32(idx)*2, uint16(color))
}

func SetBGPalette16(pidx uint8, colors [16]Color) {
	if pidx < 16 {
		idx := uint8(pidx) * 16
		for i, color := range colors {
			SetBGColor(idx+uint8(i), color)
		}
	}
}

func SetOBJPalette16(pidx uint8, colors [16]Color) {
	if pidx < 16 {
		idx := uint8(pidx) * 16
		for i, color := range colors {
			SetOBJColor(idx+uint8(i), color)
		}
	}
}

func SetBackdropColor(color Color) {
	SetBGColor(0, color)
}
