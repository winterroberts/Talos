package flags

type WININ_flag uint16

const (
	WIN0_BG0_E    WININ_flag = 1
	WIN0_BG1_E    WININ_flag = 1 << 1
	WIN0_BG2_E    WININ_flag = 1 << 2
	WIN0_BG3_E    WININ_flag = 1 << 3
	WIN0_OBJ_E    WININ_flag = 1 << 4
	WIN0_EFFECT_E WININ_flag = 1 << 5

	WIN1_BG0_E    WININ_flag = 1 << 8
	WIN1_BG1_E    WININ_flag = 1 << 9
	WIN1_BG2_E    WININ_flag = 1 << 10
	WIN1_BG3_E    WININ_flag = 1 << 11
	WIN1_OBJ_E    WININ_flag = 1 << 12
	WIN1_EFFECT_E WININ_flag = 1 << 13
)

type WINOUT_flag uint16

const (
	OUTW_BG0_E    WINOUT_flag = 1
	OUTW_BG1_E    WINOUT_flag = 1 << 1
	OUTW_BG2_E    WINOUT_flag = 1 << 2
	OUTW_BG3_E    WINOUT_flag = 1 << 3
	OUTW_OBJ_E    WINOUT_flag = 1 << 4
	OUTW_EFFECT_E WINOUT_flag = 1 << 5

	OBJW_BG0_E    WINOUT_flag = 1 << 8
	OBJW_BG1_E    WINOUT_flag = 1 << 9
	OBJW_BG2_E    WINOUT_flag = 1 << 10
	OBJW_BG3_E    WINOUT_flag = 1 << 11
	OBJW_OBJ_E    WINOUT_flag = 1 << 12
	OBJW_EFFECT_E WINOUT_flag = 1 << 13
)
