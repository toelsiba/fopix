package bitmap

import (
	"bytes"
	"math/rand"
	"testing"
	"time"
)

func TestBitmap(t *testing.T) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	n := r.Intn(1000)
	bm := New(n)

	for i := 0; i < n; i++ {
		err := bm.SetBit(i, randBit(r))
		if err != nil {
			t.Fatal(err)
		}
	}

	var bit, bitInvert uint
	var err error

	data := make([]byte, len(bm.data))

	for i := 0; i < n; i++ {

		copy(data, bm.data)

		if bit, err = bm.GetBit(i); err != nil {
			t.Fatal(err)
		}
		if err = bm.SetBit(i, not(bit)); err != nil {
			t.Fatal(err)
		}
		if bitInvert, err = bm.GetBit(i); err != nil {
			t.Fatal(err)
		}

		if bit == bitInvert {
			t.Fatalf("wrong set invert bit for offset %d", i)
		}

		if err = bm.SetBit(i, bit); err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(data, bm.data) {
			t.Fatal("data change after bit manipulate")
		}
	}
}

func not(a uint) uint {
	switch a {
	case 0:
		return 1
	case 1:
		return 0
	default:
		panic("invalid bit")
	}
}

func randBool(r *rand.Rand) bool {
	return (r.Int() & 1) == 1
}

func randBit(r *rand.Rand) uint {
	return uint(r.Int() & 1)
}
