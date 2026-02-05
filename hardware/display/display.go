package display

import (
	"talos/hardware/display/flags"
	reg "talos/hardware/register"
)

// SetMode set's the display mode the Gameboy will use for rendering.
func SetMode(m Mode) {
	reg.DISPCNT.ClearBits(mode_bits)
	reg.DISPCNT.Set(uint16(m))
}

func SetDISPCNT(flag flags.DISPCNT_flag, v bool) {
	reg.DISPCNT.SetFlag(uint16(flag), v)
}

// DISPSTAT
func SetDISPSTAT(flag flags.DISPSTAT_flag_w, v bool) {
	reg.DISPSTAT.SetFlag(uint16(flag), v)
}

func GetDISPSTAT[f flags.DISPSTAT_flag | flags.DISPSTAT_flag_w](flag f) bool {
	return reg.DISPSTAT.GetFlag(uint16(flag))
}

func SetVcountSetting(v uint8) {
	if v > vcount_max {
		v = vcount_max
	}
	reg.DISPSTAT.ModifyBits(vcount_setting_bits, uint16(vcount_setting_offset), uint16(v))
}

func GetVcountSetting() uint8 {
	return uint8(reg.DISPSTAT.GetBits(vcount_setting_bits, uint16(vcount_setting_offset)))
}

// Vcount
func GetVcount() uint8 {
	return uint8(reg.VCOUNT.Get())
}

// BG#CNT
func SetBGCNT(r reg.BGControlRegister, flag flags.BGCNT_flag, v bool) {
	reg.Register16(r).SetFlag(uint16(flag), v)
}

func GetBGCNT(r reg.BGControlRegister, flag flags.BGCNT_flag) bool {
	return reg.Register16(r).GetFlag(uint16(flag))
}

func SetBGPriority(r reg.BGControlRegister, v BG_PRIORITY) {
	reg.Register16(r).ModifyBits(bg_priority_bits, 0, uint16(v))
}

func GetBGPriority(r reg.BGControlRegister) uint16 {
	return reg.Register16(r).GetBits(bg_priority_bits, 0)
}

func SetCharacterBaseBlock(r reg.BGControlRegister, v CHAR_BASE) {
	reg.Register16(r).ModifyBits(char_base_bits, uint16(char_base_offset), uint16(v))
}

func GetCharacterBaseBlock(r reg.BGControlRegister) uint16 {
	return reg.Register16(r).GetBits(char_base_bits, uint16(char_base_offset))
}

func SetScreenBaseBlock(r reg.BGControlRegister, v SCREEN_BASE) {
	reg.Register16(r).ModifyBits(screen_base_bits, uint16(screen_base_offset), uint16(v))
}

func GetScreenBaseBlock(r reg.BGControlRegister) uint16 {
	return reg.Register16(r).GetBits(screen_base_bits, uint16(screen_base_offset))
}

func SetScreenSize(r reg.BGControlRegister, v SCREEN_SIZE) {
	reg.Register16(r).ModifyBits(screen_size_bits, uint16(screen_size_offset), uint16(v))
}

func GetScreenSize(r reg.BGControlRegister) uint16 {
	return reg.Register16(r).GetBits(screen_size_bits, uint16(screen_size_offset))
}

// BG#OFS
func SetBackgroundOffset(r reg.BGScrollRegister, v uint16) {
	if v > bg_scroll_max {
		v = bg_scroll_max
	}
	reg.Register16(r).Set(v)
}

// BG#RS
func SetBackgroundRotScaleFraction(r reg.BGRotScaleRegister, v uint16) {
	reg.Register16(r).ModifyBits(bg_rs_fraction, 0, v)
}

func SetBackgroundRotScaleInteger(r reg.BGRotScaleRegister, v uint16) {
	reg.Register16(r).ModifyBits(bg_rs_int, uint16(bg_rs_int_offset), v)
}

func SetBackgroundRotScaleFlag(r reg.BGRotScaleRegister, flag flags.BG2P_flag, v bool) {
	reg.Register16(r).SetFlag(uint16(flag), v)
}

// BG#RP
func SetBackgroundRefPointFraction(r reg.BGRefPointRegister, v uint16) {
	reg.Register32(r).ModifyBits(uint32(bg_ref_pt_fraction), 0, uint32(v))
}

func SetBackgroundRefPointInteger(r reg.BGRefPointRegister, v uint32) {
	reg.Register32(r).ModifyBits(bg_ref_pt_int, uint32(bg_ref_pt_int_offset), v)
}

func SetBackgrounfRefPointFlag(r reg.BGRefPointRegister, flag flags.BGRP_flag, v bool) {
	reg.Register32(r).SetFlag(uint32(flag), v)
}

// WINH/V
func SetWindowDimensionHNear(r reg.WinDimHRegister, v uint8) {
	reg.Register16(r).ModifyBits(win_near, 0, uint16(v))
}

func SetWindowDimensionHFar(r reg.WinDimHRegister, v uint8) {
	reg.Register16(r).ModifyBits(win_far, win_far_offset, uint16(v))
}

func SetWindowDimensionVNear(r reg.WinDimVRegister, v uint8) {
	reg.Register16(r).ModifyBits(win_near, 0, uint16(v))
}

func SetWindowDimensionVFar(r reg.WinDimVRegister, v uint8) {
	reg.Register16(r).ModifyBits(win_far, win_far_offset, uint16(v))
}

// WININ/OUT
func SetWindowInFlag(flag flags.WININ_flag, v bool) {
	reg.WININ.SetFlag(uint16(flag), v)
}

func GetWindowInFlag(flag flags.WININ_flag) bool {
	return reg.WININ.GetFlag(uint16(flag))
}

func SetWindowOutFlag(flag flags.WINOUT_flag, v bool) {
	reg.WINOUT.SetFlag(uint16(flag), v)
}

func GetWindowOutFlag(flag flags.WINOUT_flag) bool {
	return reg.WINOUT.GetFlag(uint16(flag))
}

// mosaic
func SetMosaicBackgroundHSize(v uint8) {
	reg.MOSAIC.ModifyBits(mosaic_bg_h, 0, uint16(v))
}

func SetMosaicBackgroundVSize(v uint8) {
	reg.MOSAIC.ModifyBits(mosaic_bg_v, mosaic_bg_v_offset, uint16(v))
}

func SetMosaicObjHSize(v uint8) {
	reg.MOSAIC.ModifyBits(mosaic_obj_h, mosaic_obj_h_offset, uint16(v))
}

func SetMosaicObjVSize(v uint8) {
	reg.MOSAIC.ModifyBits(mosaic_obj_v, mosaic_obj_v_offset, uint16(v))
}

// effect
func SetColorEffectFlag(flag flags.BLDCNT_flag, v bool) {
	reg.BLDCNT.SetFlag(uint16(flag), v)
}

func GetColorEffectFlag(flag flags.BLDCNT_flag) bool {
	return reg.BLDCNT.GetFlag(uint16(flag))
}

func SetColorEffect(effect BLD_EFFECT) {
	reg.BLDCNT.ModifyBits(bld_effect_bits, uint16(bld_effect_offset), uint16(effect))
}

// alpha
func SetBlendAlphaEVACoeff(v uint8) {
	reg.BLDALPHA.ModifyBits(bld_alpha_eva_bits, 0, uint16(v))
}

func GetBlendAlphaEVACoeff() uint8 {
	return uint8(reg.BLDALPHA.GetBits(bld_alpha_eva_bits, 0))
}

func SetBlendAlphaEVBCoeff(v uint8) {
	reg.BLDALPHA.ModifyBits(bld_alpha_evb_bits, bld_alpha_evb_offset, uint16(v))
}

func GetBlendAlphaEVBCoeff() uint8 {
	return uint8(reg.BLDALPHA.GetBits(bld_alpha_evb_bits, bld_alpha_evb_offset))
}

// brightness
func SetBlendBrightnessEVYCoeff(v uint8) {
	reg.BLDY.SetBits(uint16(v))
}
