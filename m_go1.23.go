//go:build go1.23 && !go1.24
// +build go1.23,!go1.24

package fastrand

type m struct {
	_        [244]byte
	fastrand uint64
}
