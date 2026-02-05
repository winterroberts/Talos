package register

import mem "talos/hardware/memory"

type TimerCounterRegister Register16

type TimerControlRegister Register16

const (
	TM0CNT_L TimerCounterRegister = TimerCounterRegister(mem.IO_START + 0x100)
	TM0CNT_H TimerControlRegister = TimerControlRegister(mem.IO_START + 0x102)
	TM1CNT_L TimerCounterRegister = TimerCounterRegister(mem.IO_START + 0x104)
	TM1CNT_H TimerControlRegister = TimerControlRegister(mem.IO_START + 0x106)
	TM2CNT_L TimerCounterRegister = TimerCounterRegister(mem.IO_START + 0x108)
	TM2CNT_H TimerControlRegister = TimerControlRegister(mem.IO_START + 0x10A)
	TM3CNT_L TimerCounterRegister = TimerCounterRegister(mem.IO_START + 0x10C)
	TM3CNT_H TimerControlRegister = TimerControlRegister(mem.IO_START + 0x10E)
)
