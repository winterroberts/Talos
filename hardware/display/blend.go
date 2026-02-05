package display

const (
	bld_effect_offset uint8  = 6
	bld_effect_bits   uint16 = 0b11 << bld_effect_offset
)

// Specifies blending effect flags
type BLD_EFFECT uint8

const (
	BLD_EFFECT_NONE           BLD_EFFECT = 0
	BLD_EFFECT_ALPHA_BLEND    BLD_EFFECT = 1
	BLD_EFFECT_BRIGHTNESS_INC BLD_EFFECT = 2
	BLD_EFFECT_BRIGHTNESS_DEC BLD_EFFECT = 3
)

const (
	bld_alpha_evb_offset = 8

	bld_alpha_eva_bits = 0xF
	bld_alpha_evb_bits = 0xF << bld_alpha_evb_offset
)
