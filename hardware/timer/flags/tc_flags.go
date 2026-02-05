package flags

// Specifies timer control flags
type TMCNT_flag uint32

const (
	// Enables/disables timer overflow causing the next timer to increment
	COUNT_UP TMCNT_flag = 1 << 18
	// Enables/disables an interrupt when a timer overflows
	IRQ_ENABLE TMCNT_flag = 1 << 22
	// Enables/disables a timer
	TIMER_ENABLE TMCNT_flag = 1 << 23
)
