package flags

// Specifies background contorl flags
type BGCNT_flag uint16

const (
	// Enables/disables mosaic effects
	MOSAIC BGCNT_flag = 1 << 6
	// Sets color palettes to either 0=16x16 or 1=256x1
	PALETTES BGCNT_flag = 1 << 7
	// Contorls display area overflow 0=transparent, 1=wrap
	OVERFLOW BGCNT_flag = 1 << 13
)

// Specifies roation and scale flags
type BG2P_flag uint16

const (
	// Sets the rotation/scale sign to 0=positive, 1=negative
	RS_SIGN BG2P_flag = 1 << 15
)

// Specifies background reference pointer flags
type BGRP_flag uint32

const (
	// Set the background reference sign 0=positive, 1=negative
	RP_SIGN BGRP_flag = 1 << 27
)
