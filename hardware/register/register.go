package register

import mem "talos/hardware/memory"

type Register16 mem.Addr

func (r Register16) Set(val uint16) {
	mem.Write16(uint32(r), val)
}

func (r Register16) Get() uint16 {
	return mem.Read16(uint32(r))
}

func (r Register16) ClearBits(mask uint16) {
	r.Set(r.Get() &^ mask)
}

func (r Register16) SetBits(value uint16) {
	r.Set(r.Get() | value)
}

func (r Register16) SetFlag(flag uint16, v bool) {
	if v {
		r.SetBits(flag)
	} else {
		r.ClearBits(flag)
	}
}

func (r Register16) GetFlag(flag uint16) bool {
	return r.Get()&flag != 0
}

func (r Register16) ModifyBits(bits, offset, v uint16) {
	r.ClearBits(bits)
	r.SetBits(v << offset)
}

func (r Register16) GetBits(bits, offset uint16) uint16 {
	return r.Get() & bits >> offset
}

type Register32 mem.Addr

func (r Register32) Set(val uint32) {
	mem.Write32(uint32(r), val)
}

func (r Register32) Get() uint32 {
	return mem.Read32(uint32(r))
}

func (r Register32) ClearBits(mask uint32) {
	r.Set(r.Get() &^ mask)
}

func (r Register32) SetBits(value uint32) {
	r.Set(r.Get() | value)
}

func (r Register32) SetFlag(flag uint32, v bool) {
	if v {
		r.SetBits(flag)
	} else {
		r.ClearBits(flag)
	}
}

func (r Register32) GetFlag(flag uint32) bool {
	return r.Get()&flag != 0
}

func (r Register32) ModifyBits(bits, offset, v uint32) {
	r.ClearBits(bits)
	r.SetBits(v << offset)
}

func (r Register32) GetBits(bits, offset uint32) uint32 {
	return r.Get() & bits >> offset
}
