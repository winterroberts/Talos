package register

import mem "talos/hardware/memory"

const (
	KEYINPUT Register16 = Register16(mem.IO_START + 0x130)
	KEYCNT   Register16 = Register16(mem.IO_START + 0x132)
)
