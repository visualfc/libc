package libc

import unsafe "unsafe"

func Copysignf(x float32, y float32) float32 {
	type _cgoa_200 struct {
		f float32
	}
	var ux _cgoa_200
	ux.f = x
	var uy _cgoa_200
	uy.f = y
	*(*uint32)(unsafe.Pointer(&ux)) &= uint32(2147483647)
	*(*uint32)(unsafe.Pointer(&ux)) |= *(*uint32)(unsafe.Pointer(&uy)) & uint32(2147483648)
	return ux.f
}
