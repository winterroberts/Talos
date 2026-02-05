package display

const (
	mosaic_bg_v_offset  uint16 = 4
	mosaic_obj_h_offset uint16 = 8
	mosaic_obj_v_offset uint16 = 12

	mosaic_bg_h  uint16 = 0xF
	mosaic_bg_v  uint16 = 0xF << mosaic_bg_v_offset
	mosaic_obj_h uint16 = 0xF << mosaic_obj_h_offset
	mosaic_obj_v uint16 = 0xF << mosaic_obj_v_offset
)
