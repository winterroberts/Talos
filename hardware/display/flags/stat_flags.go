package flags

// Specifies read-only display status flags
type DISPSTAT_flag uint16

// Specifies writeable displat status flags
type DISPSTAT_flag_w DISPSTAT_flag

const (
	// 1 if not drawing
	VBLANK DISPSTAT_flag = 1
	// 1 at the end of a scan line
	HBLANK DISPSTAT_flag = 1 << 1
	// 1 if drawing line as matched by vcount setting
	VCOUNTER DISPSTAT_flag = 1 << 2
	// Enables vblank interrupt
	VBLANK_I DISPSTAT_flag_w = 1 << 3
	// Enables hblank interrupt
	HBLANK_I DISPSTAT_flag_w = 1 << 4
	// Enables vcount interrupt
	VCOUNTER_I DISPSTAT_flag_w = 1 << 5
)
