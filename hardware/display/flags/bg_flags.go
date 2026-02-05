package flags

type BGCNT_flag uint16

const (
	MOSAIC   BGCNT_flag = 1 << 6
	PALETTES BGCNT_flag = 1 << 7
	OVERFLOW BGCNT_flag = 1 << 13
)

type BG2P_flag uint16

const (
	RS_SIGN BG2P_flag = 1 << 15
)

type BGRP_flag uint32

const (
	RP_SIGN BGRP_flag = 1 << 27
)
