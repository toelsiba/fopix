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
		if err := bm.Set(i, randBool(r)); err != nil {
			t.Fatal(err)
		}
	}

	var bit, bitInvert bool
	var err error

	data := make([]byte, len(bm.data))

	for i := 0; i < n; i++ {

		copy(data, bm.data)

		if bit, err = bm.Get(i); err != nil {
			t.Fatal(err)
		}
		if err = bm.Set(i, !bit); err != nil {
			t.Fatal(err)
		}
		if bitInvert, err = bm.Get(i); err != nil {
			t.Fatal(err)
		}

		if bit == bitInvert {
			t.Fatalf("wrong set invert bit for offset %d", i)
		}

		if err = bm.Set(i, bit); err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(data, bm.data) {
			t.Fatal("data change after bit manipulate")
		}
	}
}

func randBool(r *rand.Rand) bool {
	return (r.Int() & 1) == 1
}
