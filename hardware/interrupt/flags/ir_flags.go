package flags

// Specifies BIOS-requested interrupts
type INTERRUPT_flag uint16

const (
	// Not drawing
	IR_VBLANK INTERRUPT_flag = 0
	// End of scan line
	IR_HBLANK INTERRUPT_flag = 1
	// Hit vcount setting
	IR_VCOUNTER INTERRUPT_flag = 2
	// Timer 0 overflow
	IR_TIMER0 INTERRUPT_flag = 3
	// Timer 1 overflow
	IR_TIMER1 INTERRUPT_flag = 4
	// Timer 2 overflow
	IR_TIMER2 INTERRUPT_flag = 5
	// Timer 3 overflow
	IR_TIMER3 INTERRUPT_flag = 6
	// Serial data available
	IR_SERIAL INTERRUPT_flag = 7
	// DMA0 data available
	IR_DMA0 INTERRUPT_flag = 8
	// DMA1 data available
	IR_DMA1 INTERRUPT_flag = 9
	// DMA2 data available
	IR_DMA2 INTERRUPT_flag = 10
	// DMA3 data available
	IR_DMA3 INTERRUPT_flag = 11
	// Keypad key pressed
	IR_KEYPAD INTERRUPT_flag = 12
	// Gamepak-generated interrupt
	IR_GAMEPAK INTERRUPT_flag = 13
)
