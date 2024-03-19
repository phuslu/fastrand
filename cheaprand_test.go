// go test -v -cpu=8 -run=none -bench=. -benchtime=5s -benchmem

package cheaprand

import (
	"testing"
	"time"
	_ "unsafe"
)

//go:noescape
//go:linkname cheaprand runtime.cheaprand
func cheaprand() uint32

//go:noescape
//go:linkname cheaprandn runtime.cheaprandn
func cheaprandn(x uint32) uint32

func BenchmarkGolangUint32(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cheaprand()
		}
	})
}

func BenchmarkGolangUint32n(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cheaprandn(1 << 31)
		}
	})
}

func BenchmarkPhusluUint32(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		phuslu := &CheapRand{uint64(time.Now().UnixNano())}
		for pb.Next() {
			phuslu.Uint32()
		}
	})
}

func BenchmarkPhusluUint32n(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		phuslu := &CheapRand{uint64(time.Now().UnixNano())}
		for pb.Next() {
			phuslu.Uint32n(1 << 31)
		}
	})
}
