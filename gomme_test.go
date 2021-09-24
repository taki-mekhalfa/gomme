package gomme

import (
	"math/rand"
	"testing"
	"time"
)

func equal(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestEncodeDecode(t *testing.T) {

	sizes := []int{10, 100, 200, 500, 1000, 10000, 100000}
	mins := []int64{0, 0, -50}
	maxes := []int64{1, 100, 50}

	for _, size := range sizes {
		for i, min := range mins {
			max := maxes[i]

			seed := time.Now().UnixNano()
			rand.Seed(seed)

			in := make([]int64, size)
			for i := 0; i < size; i++ {
				in[i] = min + rand.Int63n(max-min+1)
			}

			encoded, shift, bitSize, padding := Encode(in)
			out := Decode(encoded, shift, bitSize, padding)

			if !equal(in, out) {
				t.Errorf("seed: %d, size: %d, min:%d, max:%d", seed, size, min, max)
			}
		}
	}

}

func TestEncodeStream(t *testing.T) {

	sizes := []int{10, 100, 200, 500, 1000, 10000, 100000}
	mins := []int64{0, 0, -50}
	maxes := []int64{1, 100, 50}

	for _, size := range sizes {
		for i, min := range mins {
			max := maxes[i]

			seed := time.Now().UnixNano()
			rand.Seed(seed)

			in := make([]int64, size)
			for i := 0; i < size; i++ {
				in[i] = min + rand.Int63n(max-min+1)
			}

			se := NewStreamEncoder(min, max, "ignore")
			for _, v := range in {
				se.Encode(v)
			}

			encoded, shift, bitSize, padding := se.Encoded()
			out := Decode(encoded, shift, bitSize, padding)

			if !equal(in, out) {
				t.Errorf("seed: %d, size: %d, min:%d, max:%d", seed, size, min, max)
			}
		}
	}

}
