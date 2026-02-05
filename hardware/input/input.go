package input

import (
	"talos/hardware/input/flags"
	reg "talos/hardware/register"
)

func GetInputStatus(flag flags.KEYINPUT_flag) bool {
	return !reg.KEYINPUT.GetFlag(uint16(flag))
}
