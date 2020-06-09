package fastrand

import "testing"

type Rng uint64

// Next returns the next random number
func (r *Rng) Next() uint64 {
	*r ^= *r >> 12 // a
	*r ^= *r << 25 // b
	*r ^= *r >> 27 // c
	return uint64(*r * 2685821657736338717)
}

// Intn returns a uniform random integer [0,bound)
func (r *Rng) Intn(bound uint64) uint64 {
	threshold := -bound % bound
	n := r.Next()
	for n < threshold {
		n = r.Next()
	}
	return n % bound
}

func (r *Rng) Int32n(n uint32) uint32 {

	bits := r.Next()

	random32bit := uint32(bits)
	bits >>= 32
	full := false

	multiresult := uint64(random32bit) * uint64(n)
	leftover := uint32(multiresult)
	if leftover < n {
		threshold := uint32(-n) % n
		for leftover < threshold {
			random32bit = uint32(bits)
			bits >>= 32
			if !full {
				bits = r.Next()
			}
			full = !full
			multiresult = uint64(random32bit) * uint64(n)
			leftover = uint32(multiresult)
		}
	}
	return uint32(multiresult >> 32) // [0, n)
}

func TestInt32n(t *testing.T) {

	var counts [1000]int

	r := Rng(0xdeadbeef)
	for i := 0; i < 100e6; i++ {
		counts[r.Int32n(1000)]++
	}

	t.Logf("counts = %+v\n", counts)
}

func main() {
	TestInt32n(32)
}
