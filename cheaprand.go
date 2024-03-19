// Package cheaprand from golang runtime
//
//	package main
//
//	import (
//		"time"
//	)
//
//	func main() {
//		cheaprand := &CheapRand{uint64(time.Now().UnixNano())}
//		println(cheaprand.Uint32(), cheaprand.Uint64())
//	}
package cheaprand

import (
	"math/bits"
)

type CheapRand struct {
	Seed uint64
}

func (rand *CheapRand) Uint32() uint32 {
	rand.Seed += 0xa0761d6478bd642f
	hi, lo := bits.Mul64(rand.Seed, rand.Seed^0xe7037ed1a0b428db)
	return uint32(hi ^ lo)
}

func (rand *CheapRand) Uint32n(n uint32) uint32 {
	return uint32((uint64(rand.Uint32()) * uint64(n)) >> 32)
}

func (rand *CheapRand) Int64() int64 {
	return int64(rand.Uint32())<<31 ^ int64(rand.Uint32())
}

func (rand *CheapRand) Uint64() uint64 {
	return uint64(rand.Uint32())<<32 ^ uint64(rand.Uint32())
}
