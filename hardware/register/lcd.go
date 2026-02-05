package register

import mem "talos/hardware/memory"

type BGControlRegister Register16

type BGScrollRegister Register16

type BGRotScaleRegister Register16

type BGRefPointRegister Register32

type WinDimHRegister Register16

type WinDimVRegister Register16

const (
	DISPCNT  Register16 = Register16(mem.IO_START)       // (R/W) lcd control
	DISPSTAT Register16 = Register16(mem.IO_START + 0x4) // (R/W) lcd stats
	VCOUNT   Register16 = Register16(mem.IO_START + 0x6) // (R) vertical counter

	// (R/W) background controls
	BG0CNT BGControlRegister = BGControlRegister(mem.IO_START + 0x8)
	BG1CNT BGControlRegister = BGControlRegister(mem.IO_START + 0xA)
	BG2CNT BGControlRegister = BGControlRegister(mem.IO_START + 0xC)
	BG3CNT BGControlRegister = BGControlRegister(mem.IO_START + 0xE)

	// (W) background offsets h/x, v/y
	BG0HOFS BGScrollRegister = BGScrollRegister(mem.IO_START + 0x10)
	BG0VOFS BGScrollRegister = BGScrollRegister(mem.IO_START + 0x12)
	BG1HOFS BGScrollRegister = BGScrollRegister(mem.IO_START + 0x14)
	BG1VOFS BGScrollRegister = BGScrollRegister(mem.IO_START + 0x16)
	BG2HOFS BGScrollRegister = BGScrollRegister(mem.IO_START + 0x18)
	BG2VOFS BGScrollRegister = BGScrollRegister(mem.IO_START + 0x1A)
	BG3HOFS BGScrollRegister = BGScrollRegister(mem.IO_START + 0x1C)
	BG3VOFS BGScrollRegister = BGScrollRegister(mem.IO_START + 0x1E)

	// (W) bg2 rotation/scale
	BG2PA BGRotScaleRegister = BGRotScaleRegister(mem.IO_START + 0x20)
	BG2PB BGRotScaleRegister = BGRotScaleRegister(mem.IO_START + 0x22)
	BG2PC BGRotScaleRegister = BGRotScaleRegister(mem.IO_START + 0x24)
	BG2PD BGRotScaleRegister = BGRotScaleRegister(mem.IO_START + 0x26)
	BG2X  BGRefPointRegister = BGRefPointRegister(mem.IO_START + 0x28)
	BG2Y  BGRefPointRegister = BGRefPointRegister(mem.IO_START + 0x2C)

	// (W) bg3 rotation scale
	BG3PA BGRotScaleRegister = BGRotScaleRegister(mem.IO_START + 0x30)
	BG3PB BGRotScaleRegister = BGRotScaleRegister(mem.IO_START + 0x32)
	BG3PC BGRotScaleRegister = BGRotScaleRegister(mem.IO_START + 0x34)
	BG3PD BGRotScaleRegister = BGRotScaleRegister(mem.IO_START + 0x36)
	BG3X  BGRefPointRegister = BGRefPointRegister(mem.IO_START + 0x38)
	BG3Y  BGRefPointRegister = BGRefPointRegister(mem.IO_START + 0x3C)

	// (W) window h/v dimensions
	WIN0H WinDimHRegister = WinDimHRegister(mem.IO_START + 0x40)
	WIN1H WinDimHRegister = WinDimHRegister(mem.IO_START + 0x42)
	WIN0V WinDimVRegister = WinDimVRegister(mem.IO_START + 0x44)
	WIN1V WinDimVRegister = WinDimVRegister(mem.IO_START + 0x46)

	// (R/W) in/out of windows 0/1
	WININ  Register16 = Register16(mem.IO_START + 0x48)
	WINOUT Register16 = Register16(mem.IO_START + 0x4A)

	MOSAIC   Register16 = Register16(mem.IO_START + 0x4C) // (W) mosaic size
	BLDCNT   Register16 = Register16(mem.IO_START + 0x50) // (R/W) color special effects
	BLDALPHA Register16 = Register16(mem.IO_START + 0x52) // (R/W) alpha blend
	BLDY     Register16 = Register16(mem.IO_START + 0x54) // (W) brightness
)
