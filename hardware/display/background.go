package display

// bg control
const (
	char_base_offset   uint8 = 2
	screen_base_offset uint8 = 8
	screen_size_offset uint8 = 14

	bg_priority_bits uint16 = 0b11
	char_base_bits   uint16 = 0b11 << char_base_offset
	screen_base_bits uint16 = 0b11111 << screen_base_offset
	screen_size_bits uint16 = 0b11 << screen_size_offset
)

// Specifies priority of rendering a background layer
type BG_PRIORITY uint8

const (
	BG_PRIORITY_HIGHEST BG_PRIORITY = 0
	BG_PRIORITY_HIGH    BG_PRIORITY = 1
	BG_PRIORITY_LOW     BG_PRIORITY = 2
	BG_PRIORITY_LOWEST  BG_PRIORITY = 3
)

// Specifies a section of VRAM to be used for tiles
type CHAR_BASE uint8

const (
	CHAR_BASE_16K CHAR_BASE = 0
	CHAR_BASE_32K CHAR_BASE = 1
	CHAR_BASE_48K CHAR_BASE = 3
	CHAR_BASE_64K CHAR_BASE = 4
)

// Specifies a section of VRAM to be used for tilemaps
type SCREEN_BASE uint8

const (
	SCREEN_BASE_2K  SCREEN_BASE = 0
	SCREEN_BASE_4K  SCREEN_BASE = 1
	SCREEN_BASE_6K  SCREEN_BASE = 2
	SCREEN_BASE_8K  SCREEN_BASE = 3
	SCREEN_BASE_10K SCREEN_BASE = 4
	SCREEN_BASE_12K SCREEN_BASE = 5
	SCREEN_BASE_14K SCREEN_BASE = 6
	SCREEN_BASE_16K SCREEN_BASE = 7
	SCREEN_BASE_18K SCREEN_BASE = 8
	SCREEN_BASE_20K SCREEN_BASE = 9
	SCREEN_BASE_22K SCREEN_BASE = 10
	SCREEN_BASE_24K SCREEN_BASE = 11
	SCREEN_BASE_26K SCREEN_BASE = 12
	SCREEN_BASE_28K SCREEN_BASE = 13
	SCREEN_BASE_30K SCREEN_BASE = 14
	SCREEN_BASE_32K SCREEN_BASE = 15
	SCREEN_BASE_34K SCREEN_BASE = 16
	SCREEN_BASE_36K SCREEN_BASE = 17
	SCREEN_BASE_38K SCREEN_BASE = 18
	SCREEN_BASE_40K SCREEN_BASE = 19
	SCREEN_BASE_42K SCREEN_BASE = 20
	SCREEN_BASE_44K SCREEN_BASE = 21
	SCREEN_BASE_46K SCREEN_BASE = 22
	SCREEN_BASE_48K SCREEN_BASE = 23
	SCREEN_BASE_50K SCREEN_BASE = 24
	SCREEN_BASE_52K SCREEN_BASE = 25
	SCREEN_BASE_54K SCREEN_BASE = 26
	SCREEN_BASE_56K SCREEN_BASE = 27
	SCREEN_BASE_58K SCREEN_BASE = 28
	SCREEN_BASE_60K SCREEN_BASE = 29
	SCREEN_BASE_62K SCREEN_BASE = 30
	SCREEN_BASE_64K SCREEN_BASE = 31
)

// Specifies a screen size with support for either text or rotation/scale
type SCREEN_SIZE uint8

const (
	SCREEN_SIZE_TEXT_256x256 SCREEN_SIZE = 0
	SCREEN_SIZE_TEXT_512x256 SCREEN_SIZE = 1
	SCREEN_SIZE_TEXT_256x512 SCREEN_SIZE = 2
	SCREEN_SIZE_TEXT_512x512 SCREEN_SIZE = 3
	SCREEN_SIZE_RS_128x128   SCREEN_SIZE = 0
	SCREEN_SIZE_RS_256x256   SCREEN_SIZE = 1
	SCREEN_SIZE_RS_512x512   SCREEN_SIZE = 2
	SCREEN_SIZE_RS_1024x1024 SCREEN_SIZE = 3
)

// bg scroll
const (
	bg_scroll_max = 511
)

// bg rot/scale
const (
	bg_rs_int_offset uint8 = 8

	bg_rs_fraction uint16 = 0xFF
	bg_rs_int      uint16 = 0x7F << bg_rs_int_offset
)

// bg ref point
const (
	bg_ref_pt_int_offset uint8 = 8

	bg_ref_pt_fraction uint16 = 0xFF
	bg_ref_pt_int      uint32 = 0x7FFFF << bg_ref_pt_int_offset
)
