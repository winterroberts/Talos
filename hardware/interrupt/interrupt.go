package interrupt

/*
#cgo CFLAGS: -I${SRCDIR}/hardware/interrupt
#include <stdint.h>

extern uint32_t getHandlerAddr(void);
extern uint32_t peekIRQ(uint16_t flag);
extern void ackIRQ(uint16_t flag);
*/
import "C"

import (
	"talos/hardware/interrupt/flags"
	"talos/hardware/memory"
	reg "talos/hardware/register"
)

const num_handlers = 14

var handlers [num_handlers]func()

// SetHandler specifies a function to enter when an interruptis fired
func SetHandler(flag flags.INTERRUPT_flag, handler func()) {
	handlers[uint16(flag)] = handler
}

// ProcessInterrupts calls handlers for any enabled interrupt requests
func ProcessInterrupts() {
	for i := range num_handlers {
		flag := flags.INTERRUPT_flag(i)
		if Peek(flag) {
			Ack(flag)
			if h := handlers[i]; h != nil {
				h()
			}
		}
	}
}

// EnableMaster enables interrupts
func EnableMaster() {
	reg.INTERRUPT_MASTER_ENABLE.Set(1)
}

// DisableMaster disables interrupts
func DisableMaster() {
	reg.INTERRUPT_MASTER_ENABLE.Set(0)
}

// Enable enables a single interrupt to be fired
func Enable(flag flags.INTERRUPT_flag) {
	reg.INTERRUPT_ENABLE.SetFlag(uint16(1<<flag), true)
}

// Disable disables a single interrupt being fired
func Disable(flag flags.INTERRUPT_flag) {
	reg.INTERRUPT_ENABLE.SetFlag(uint16(1<<flag), false)
}

// c stuff
func Setup() {
	handlerAddr := uint32(C.getHandlerAddr())
	memory.Write32(0x03007FFC, handlerAddr)
}

// Check determines if a given enabled interrupt has fired and acknowledges it
func Check(flag flags.INTERRUPT_flag) bool {
	peek := Peek(flag)
	Ack(flag)
	return peek
}

// Peek determines if a given enabled interrupt has fired without acknowledging
func Peek(flag flags.INTERRUPT_flag) bool {
	return C.peekIRQ(C.uint16_t(uint16(1<<flag))) != 0
}

// Ack acknowledges a given enabled interrupt was fired
func Ack(flag flags.INTERRUPT_flag) {
	C.ackIRQ(C.uint16_t(uint16(1 << flag)))
}
