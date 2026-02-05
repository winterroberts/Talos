package input

import (
	"talos/hardware/input/flags"
	reg "talos/hardware/register"
)

// GetInputStatus returns whether a key input is currently active
func GetInputStatus(flag flags.KEYINPUT_flag) bool {
	return !reg.KEYINPUT.GetFlag(uint16(flag))
}
