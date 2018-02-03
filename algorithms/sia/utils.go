package sia

import (
	"crypto/sha256"
)

// DoubleSha256 calculates sha256 twice.
func DoubleSha256(ba []byte) []byte {
	df1 := sha256.Sum256(ba)
	df2 := sha256.Sum256(df1[:32])
	return df2[:32]
}

// Reverse r.
func Reverse(r []byte) []byte {
	var v = make([]byte, len(r))
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		v[i], v[j] = r[j], r[i]
	}
	return v
}

// Swap32 swap b.
func Swap32(b []byte) []byte {
	l := len(b)
	_mx := 0
	var o = make([]byte, l)
	for i, j := 0, 0; i < l; i++ {
		if i%4 == 0 {
			j = i
			if l-j < 4 {
				_mx = l - j
			} else {
				_mx = 4
			}
		}
		_of := j + (_mx - (i - j)) - 1
		o[i] = b[_of]
	}
	return o
}
