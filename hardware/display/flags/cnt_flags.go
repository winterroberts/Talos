package flags

type DISPCNT_flag uint16

const (
	FRAME_SELECT     DISPCNT_flag = 1 << 4
	HBLANK_FREE      DISPCNT_flag = 1 << 5
	OBJ_VRAM_MAPPING DISPCNT_flag = 1 << 6
	FORCED_BLANK     DISPCNT_flag = 1 << 7
	DISP_BG0         DISPCNT_flag = 1 << 8
	DISP_BG1         DISPCNT_flag = 1 << 9
	DISP_BG2         DISPCNT_flag = 1 << 10
	DISP_BG3         DISPCNT_flag = 1 << 11
	DISP_OBJ         DISPCNT_flag = 1 << 12
	WIN0             DISPCNT_flag = 1 << 13
	WIN1             DISPCNT_flag = 1 << 14
	WIN_OBJ          DISPCNT_flag = 1 << 15
)
