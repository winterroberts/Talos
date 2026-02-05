package register

import mem "talos/hardware/memory"

const (
	INTERRUPT_MASTER_ENABLE Register32 = Register32(mem.IO_START + 0x208)
	INTERRUPT_ENABLE        Register16 = Register16(mem.IO_START + 0x200)
	INTERRUPT_REQUEST       Register16 = Register16(mem.IO_START + 0x202)
)
