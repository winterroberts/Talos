package flags

type DISPSTAT_flag uint16

type DISPSTAT_flag_w DISPSTAT_flag

const (
	VBLANK     DISPSTAT_flag   = 1
	HBLANK     DISPSTAT_flag   = 1 << 1
	VCOUNTER   DISPSTAT_flag   = 1 << 2
	VBLANK_I   DISPSTAT_flag_w = 1 << 3
	HBLANK_I   DISPSTAT_flag_w = 1 << 4
	VCOUNTER_I DISPSTAT_flag_w = 1 << 5
)
