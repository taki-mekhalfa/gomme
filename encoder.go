package gomme

import (
	"fmt"
	"math/bits"

	"github.com/taki-mekhalfa/bitarray"
)

const maxUint64 = 1<<63 - 1
const minUint64 = -maxUint64 - 1

func minMax(in []int64) (min int64, max int64) {
	min, max = maxUint64, minUint64
	for _, v := range in {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return
}

func Encode(in []int64) ([]byte, int64, int, int) {
	if len(in) == 0 {
		panic("input slice is empty")
	}

	min, max := minMax(in)
	bitSize := bits.Len64(uint64(max - min))

	ba := bitarray.New()
	for _, v := range in {
		ba.Append64(uint64(v-min), bitSize)
	}

	return ba.Bytes(), min, bitSize, ba.Padding()
}

type StreamEncoder struct {
	min      int64
	max      int64
	bitSize  int
	strategy string
	ba       *bitarray.BitArray
}

func NewStreamEncoder(min, max int64, strategy string) *StreamEncoder {
	if max <= min {
		panic(fmt.Sprintf("illegal arguments %d <= %d", max, min))
	}
	if strategy != "ignore" && strategy != "saturate" {
		panic("illegal argument " + strategy)
	}

	return &StreamEncoder{min: min, max: max, strategy: strategy, ba: bitarray.New(), bitSize: bits.Len64(uint64(max - min))}
}

func (se *StreamEncoder) Encode(v int64) {
	if v < se.min {
		if se.strategy == "ignore" {
			return
		}
		v = se.min
	}

	if v > se.max {
		if se.strategy == "ignore" {
			return
		}
		v = se.max
	}

	se.ba.Append64(uint64(v-se.min), se.bitSize)
}

func (se *StreamEncoder) Encoded() ([]byte, int64, int, int) {
	return se.ba.Bytes(), se.min, se.bitSize, se.ba.Padding()
}
