package flags

type KEYINPUT_flag uint16

const (
	KEY_A      KEYINPUT_flag = 1
	KEY_B      KEYINPUT_flag = 1 << 1
	KEY_SELECT KEYINPUT_flag = 1 << 2
	KEY_START  KEYINPUT_flag = 1 << 3
	KEY_RIGHT  KEYINPUT_flag = 1 << 4
	KEY_LEFT   KEYINPUT_flag = 1 << 5
	KEY_UP     KEYINPUT_flag = 1 << 6
	KEY_DOWN   KEYINPUT_flag = 1 << 7
	KEY_RBUMP  KEYINPUT_flag = 1 << 8
	KEY_LBUMP  KEYINPUT_flag = 1 << 9
)
