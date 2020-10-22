package collectionsutils

import "math/rand"

func IsSliceEmpty(slice []interface{}) bool {
	if slice == nil || len(slice) == 0 {
		return false
	}
	return true
}

//random string
func RandomString(l uint64) []byte {
	source := []byte("0123456789abcdefghijklmnopqrstuvwxyz")
	dst := make([]byte, l)

	length := uint64(len(source))
	mark := uint64(63)
	count := uint64(0)

LOOP:
	for {
		tmp := rand.Uint64()
		for i := 0; i < 10; i++ {
			bit := tmp & mark
			tmp = tmp >> 6

			if length <= bit {
				continue
			} else {
				dst[count] = source[bit]
				count++
				if count >= l {
					break LOOP
				}
			}
		}
	}
	return dst
}
