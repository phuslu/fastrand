// go test -v -cpu=8 -run=none -bench=. -benchtime=5s -benchmem

package fastrand

import (
	"testing"
	_ "unsafe"
)

func TestFastrand(t *testing.T) {
	t.Logf("runtime %v\n", fastrand())
	t.Logf("Fastrand %v\n", Uint32())
	t.Logf("runtime %v\n", fastrand())
	t.Logf("Fastrand %v\n", Uint32())
}

//go:noescape
//go:linkname fastrand runtime.fastrand
func fastrand() uint32

func BenchmarkGolangUint32(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fastrand()
		}
	})
}

func BenchmarkPhusluUint32(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Uint32()
		}
	})
}
