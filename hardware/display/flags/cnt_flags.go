package flags

// Specifies flags for display control
type DISPCNT_flag uint16

const (
	// Select frame 0 or 1 for double-buffered rendering
	FRAME_SELECT DISPCNT_flag = 1 << 4
	// Enables or disables OAM access during HBLANK
	HBLANK_FREE DISPCNT_flag = 1 << 5
	// 0=1D, 1=2D object VRAM mapping
	OBJ_VRAM_MAPPING DISPCNT_flag = 1 << 6
	// Enable fast VRAM, palette, OAM access
	FORCED_BLANK DISPCNT_flag = 1 << 7
	// Enable/disable background 0
	DISP_BG0 DISPCNT_flag = 1 << 8
	// Enable/disable background 1
	DISP_BG1 DISPCNT_flag = 1 << 9
	// Enable/disable background 2
	DISP_BG2 DISPCNT_flag = 1 << 10
	// Enable/disable background 3
	DISP_BG3 DISPCNT_flag = 1 << 11
	// Enable/disable object display
	DISP_OBJ DISPCNT_flag = 1 << 12
	// Enable/disable window 0 display
	WIN0 DISPCNT_flag = 1 << 13
	// Enable/disable window 1 display
	WIN1 DISPCNT_flag = 1 << 14
	// Enable/disable object window displat
	WIN_OBJ DISPCNT_flag = 1 << 15
)
