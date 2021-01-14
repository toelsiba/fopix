package bitmap

import (
	"fmt"
)

const bitsPerByte = 8

func byteSetBit(x byte, i int, b uint) byte {
	y := byte(1) << i
	switch b {
	case 0:
		x &^= y
	case 1:
		x |= y
	default:
		panic(fmt.Errorf("bit has invalid value %d", b))
	}
	return x
}

func byteGetBit(x byte, i int) uint {
	return uint((x >> i) & 1)
}

func bytesSetBit(xs []byte, i int, b uint) {
	quo, rem := quoRem(i, bitsPerByte)
	xs[quo] = byteSetBit(xs[quo], rem, b)
}

func bytesGetBit(xs []byte, i int) uint {
	quo, rem := quoRem(i, bitsPerByte)
	return byteGetBit(xs[quo], rem)
}

func quoRem(x, y int) (quo, rem int) {
	quo = x / y
	rem = x % y
	return
}
