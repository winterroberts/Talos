package flags

// Specifies color effect blending for background, obj, and backdrop
type BLDCNT_flag uint16

const (
	PX1_BG0 BLDCNT_flag = 1
	PX1_BG1 BLDCNT_flag = 1 << 1
	PX1_BG2 BLDCNT_flag = 1 << 2
	PX1_BG3 BLDCNT_flag = 1 << 3
	PX1_OBJ BLDCNT_flag = 1 << 4
	PX1_BD  BLDCNT_flag = 1 << 5

	PX2_BG0 BLDCNT_flag = 1 << 8
	PX2_BG1 BLDCNT_flag = 1 << 9
	PX2_BG2 BLDCNT_flag = 1 << 10
	PX2_BG3 BLDCNT_flag = 1 << 11
	PX2_OBJ BLDCNT_flag = 1 << 12
	PX2_BD  BLDCNT_flag = 1 << 13
)
