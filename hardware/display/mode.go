package display

// Specifies rendering display mode
type Mode uint8

const (
	mode_bits = 0x7
)

const (
	MODE_0 Mode = 0x0
	MODE_1 Mode = 0x1
	MODE_2 Mode = 0x2
	MODE_3 Mode = 0x3
	MODE_4 Mode = 0x4
	MODE_5 Mode = 0x5
)
