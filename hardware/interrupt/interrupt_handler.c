#include <stdint.h>

volatile uint32_t interrupt_flags = 0;

#define IRQ_VBLANK   (1 << 0)
#define IRQ_HBLANK   (1 << 1)
#define IRQ_VCOUNT   (1 << 2)
#define IRQ_TIMER0   (1 << 3)
#define IRQ_TIMER1   (1 << 4)
#define IRQ_TIMER2   (1 << 5)
#define IRQ_TIMER3   (1 << 6)
#define IRQ_SERIAL   (1 << 7)
#define IRQ_DMA0     (1 << 8)
#define IRQ_DMA1     (1 << 9)
#define IRQ_DMA2     (1 << 10)
#define IRQ_DMA3     (1 << 11)
#define IRQ_KEYPAD   (1 << 12)
#define IRQ_GAMEPAK  (1 << 13)

// ARM interrupt handler
__attribute__((target("arm")))
void interruptHandler(void) {
    volatile uint16_t *IF = (uint16_t*)0x04000202;
    volatile uint16_t *BIOS_IF = (uint16_t*)0x03007FF8;
    
    uint16_t flags = *IF;
    
    // ack
    *IF = flags;
    *BIOS_IF |= flags;
    
    // go flags
    interrupt_flags |= flags;
}

uint32_t getHandlerAddr(void) {
    return (uint32_t)&interruptHandler;
}

// Check flag
uint32_t peekIRQ(uint16_t flag) {
    return (interrupt_flags & flag) ? 1 : 0;
}

// Ack a specific flag
void ackIRQ(uint16_t flag) {
    interrupt_flags &= ~flag;
}