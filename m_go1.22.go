//go:build go1.22 && !go1.23
// +build go1.22,!go1.23

package fastrand

type m struct {
	_        [240]byte
	fastrand uint64
}
