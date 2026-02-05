package flags

// Specifies control inside of windows
type WININ_flag uint16

const (
	// Enable/disable background 0 in window 0
	WIN0_BG0_E WININ_flag = 1
	// Enable/disable background 1 in window 0
	WIN0_BG1_E WININ_flag = 1 << 1
	// Enable/disable background 2 in window 0
	WIN0_BG2_E WININ_flag = 1 << 2
	// Enable/disable background 3 in window 0
	WIN0_BG3_E WININ_flag = 1 << 3
	// Enable/disable objs in window 0
	WIN0_OBJ_E WININ_flag = 1 << 4
	// Enable/disable effects in window 0
	WIN0_EFFECT_E WININ_flag = 1 << 5

	// Enable/disable background 0 in window 1
	WIN1_BG0_E WININ_flag = 1 << 8
	// Enable/disable background 1 in window 1
	WIN1_BG1_E WININ_flag = 1 << 9
	// Enable/disable background 2 in window 1
	WIN1_BG2_E WININ_flag = 1 << 10
	// Enable/disable background 3 in window 1
	WIN1_BG3_E WININ_flag = 1 << 11
	// Enable/disable objs in window 1
	WIN1_OBJ_E WININ_flag = 1 << 12
	// Enable/disable effects in window 1
	WIN1_EFFECT_E WININ_flag = 1 << 13
)

// Specifies control flags outside of windows
type WINOUT_flag uint16

const (
	// Enable/disable background 0 out of window 0
	OUTW_BG0_E WINOUT_flag = 1
	// Enable/disable background 1 out of window 0
	OUTW_BG1_E WINOUT_flag = 1 << 1
	// Enable/disable background 2 out of window 0
	OUTW_BG2_E WINOUT_flag = 1 << 2
	// Enable/disable background 3 out of window 0
	OUTW_BG3_E WINOUT_flag = 1 << 3
	// Enable/disable objs out of window 0
	OUTW_OBJ_E WINOUT_flag = 1 << 4
	// Enable/disable effects out of window 0
	OUTW_EFFECT_E WINOUT_flag = 1 << 5

	// Enable/disable background 0 out of window 1
	OBJW_BG0_E WINOUT_flag = 1 << 8
	// Enable/disable background 1 out of window 1
	OBJW_BG1_E WINOUT_flag = 1 << 9
	// Enable/disable background 2 out of window 1
	OBJW_BG2_E WINOUT_flag = 1 << 10
	// Enable/disable background 3 out of window 1
	OBJW_BG3_E WINOUT_flag = 1 << 11
	// Enable/disable objs out of window 1
	OBJW_OBJ_E WINOUT_flag = 1 << 12
	// Enable/disable effects out of window 1
	OBJW_EFFECT_E WINOUT_flag = 1 << 13
)
