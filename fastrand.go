package fastrand

import (
	"math/bits"
	"unsafe"
)

//go:nosplit
func getm() uintptr

//go:nosplit
func Uint32() uint32 {
	mp := (*m)(unsafe.Pointer(getm()))
	mp.fastrand += 0xa0761d6478bd642f
	hi, lo := bits.Mul64(mp.fastrand, mp.fastrand^0xe7037ed1a0b428db)
	return uint32(hi ^ lo)
}
