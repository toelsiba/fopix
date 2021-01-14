package bitmap

import (
	"errors"
)

var ErrOutOfRange = errors.New("bitmap: index out of range.")

type Bitmap struct {
	size int
	data []byte
}

func New(size int) *Bitmap {
	if size <= 0 {
		return &Bitmap{}
	}
	n, rem := quoRem(size, bitsPerByte)
	if rem > 0 {
		n++
	}
	return &Bitmap{size, make([]byte, n)}
}

func (bm *Bitmap) Size() int {
	return bm.size
}

func (bm *Bitmap) checkOutOfRange(offset int) error {
	if offset < 0 {
		return ErrOutOfRange
	}
	if offset >= bm.size {
		return ErrOutOfRange
	}
	return nil
}

func (bm *Bitmap) SetBit(i int, b uint) error {
	err := bm.checkOutOfRange(i)
	if err != nil {
		return err
	}
	bytesSetBit(bm.data, i, b)
	return nil
}

func (bm *Bitmap) GetBit(i int) (uint, error) {
	err := bm.checkOutOfRange(i)
	if err != nil {
		return 0, err
	}
	b := bytesGetBit(bm.data, i)
	return b, nil
}

func (bm *Bitmap) SetAll(val bool) {
	var b byte
	if val {
		b = 0xFF
	}
	data := bm.data
	for i := range data {
		data[i] = b
	}
}

func (bm *Bitmap) String() string {

	var bs []byte

	quo, rem := quoRem(bm.size, bitsPerByte)

	for i := 0; i < quo; i++ {
		b := bm.data[i]
		for j := 0; j < 8; j++ {
			if (b & 1) == 1 {
				bs = append(bs, '1')
			} else {
				bs = append(bs, '0')
			}
			b >>= 1
		}
	}

	if rem > 0 {
		b := bm.data[quo]
		for j := 0; j < rem; j++ {
			if (b & 1) == 1 {
				bs = append(bs, '1')
			} else {
				bs = append(bs, '0')
			}
			b >>= 1
		}
	}

	return string(bs)
}
