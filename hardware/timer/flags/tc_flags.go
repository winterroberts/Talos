package flags

type TMCNT_flag uint32

const (
	COUNT_UP     TMCNT_flag = 1 << 18
	IRQ_ENABLE   TMCNT_flag = 1 << 22
	TIMER_ENABLE TMCNT_flag = 1 << 23
)
