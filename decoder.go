package gomme

import (
	"github.com/taki-mekhalfa/bitarray"
)

func Decode(encoded []byte, shift int64, bitSize, padding int) []int64 {
	ba := bitarray.New()
	ba.AppendBytes(encoded, padding)

	result := make([]int64, 0, ba.Len()/bitSize)

	for i := 0; i+bitSize <= ba.Len(); i += bitSize {
		result = append(result, shift+int64(ba.Extract(i, i+bitSize)))
	}

	return result
}
