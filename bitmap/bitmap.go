package bitmap

import "errors"

var ErrOutOfRange = errors.New("bitmap: index out of range.")

type Bitmap struct {
	size int
	data []byte
}

func New(size int) *Bitmap {
	if size <= 0 {
		return &Bitmap{}
	}
	n, rem := quoRem(size, 8)
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

func (bm *Bitmap) Set(offset int, val bool) error {
	if err := bm.checkOutOfRange(offset); err != nil {
		return err
	}
	index, shift := quoRem(offset, 8)
	if val {
		bm.data[index] |= 1 << uint(shift)
	} else {
		bm.data[index] &= ^(1 << uint(shift))
	}
	return nil
}

func (bm *Bitmap) Get(offset int) (val bool, err error) {
	if err = bm.checkOutOfRange(offset); err != nil {
		return
	}
	index, shift := quoRem(offset, 8)
	val = ((bm.data[index] >> uint(shift)) & 1) == 1
	return
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

	quo, rem := quoRem(bm.size, 8)

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

func quoRem(x, y int) (quo, rem int) {
	quo = x / y
	rem = x - quo*y
	return
}
