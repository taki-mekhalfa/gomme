# gomme
Min-Max Time Series Encoding

Gomme allows to encode a sequence of integer values using a fixed (for all values) number of bits but minimal with regards to the data range. For example: for a series of boolean values only one bit is needed, for a series of integer percentages 7 bits are needed, etc.

Gomme is useful when:

Time series is integer-valued. (It doesn't work with floats :))
The range of the data is known in advance (if not streaming, this is not necessary).
The data range is relatively small.
The data does not have properties that would make other compression algorithms useful, or these other algorithms have an unacceptable cost for the use case.
Gomme can also be used as a baseline to calculate the true compression ratio of a compression algorithm on data of a certain nature.

## Usage
The following shows some examples:
```go
// Import bitarray into your code and refer to it as `bitarray`
import "github.com/taki-mekhalfa/gomme"
```

#### Encode & Decode 
```go
  	in := []int64{0, 1, 0, 1, 0, 0, 0, 0}
	encoded, shift, bitSize, padding := gomme.Encode(in)
	out := gomme.Decode(encoded, shift, bitSize, padding)

	fmt.Printf("len=%d, content:%08b\n", len(encoded), encoded)
	fmt.Printf("shift: %d, bitSize: %d, padding: %d\n", shift, bitSize, padding)
	fmt.Println(out)

    /* Output
        len=1, content:[01010000]
        shift: 0, bitSize: 1, padding: 0
        [0 1 0 1 0 0 0 0]
    */

```

#### Encode a stream & Decode 
```go
  	in := []int64{0, 1, 0, 1, 0, 0, 0, 0}

	se := gomme.NewStreamEncoder(0, 1, "ignore")
	for _, v := range in {
		se.Encode(v) // Encode a stream of values
	}
	encoded, shift, bitSize, padding := se.Encoded()
	out := gomme.Decode(encoded, shift, bitSize, padding)

	fmt.Printf("len=%d, content:%08b\n", len(encoded), encoded)
	fmt.Printf("shift: %d, bitSize: %d, padding: %d\n", shift, bitSize, padding)
	fmt.Println(out)

    /* Output
        len=1, content:[01010000]
        shift: 0, bitSize: 1, padding: 0
        [0 1 0 1 0 0 0 0]
    */

```