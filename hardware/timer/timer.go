package timer

import (
	reg "talos/hardware/register"
	"talos/hardware/timer/flags"
)

type TIMER_PRESCALER uint8

const (
	prescaler_bits = 0x3

	PRESCALER_1    TIMER_PRESCALER = 0
	PRESCALER_64   TIMER_PRESCALER = 1
	PRESCALER_256  TIMER_PRESCALER = 2
	PRESCALER_1024 TIMER_PRESCALER = 3
)

func SetDuration(r reg.TimerCounterRegister, ticks uint16) {
	reg.Register16(r).Set((0xFFFF - ticks) & 0xFFFE)
}

func SetTimerPrescaler(r reg.TimerControlRegister, p TIMER_PRESCALER) {
	reg.Register16(r).ModifyBits(prescaler_bits, 0, uint16(p))
}

func SetTimerControlFlag(r reg.TimerControlRegister, flag flags.TMCNT_flag, v bool) {
	reg.Register16(r).SetFlag(uint16(flag), v)
}
