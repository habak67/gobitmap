package gobitmap

import (
	"fmt"
	"strconv"
	"strings"
)

// highBit is the highest bit in a BitMap.
const highBit = 63

// BitMap is an efficient map, with a fixed size of 8 bytes, containing 64 individual bits indexed by a numerical
// index. The indexes are usually represented by a numerical constant to be more self-explanatory. Each bit could be
// set (true) or cleared (false). A BitMap is immutable. When updating a BitMap using method Set, Clear or Toggle a
// new updated BitMap is returned.
//
// There is no new-function for creating a BitMap. Instead, you start from EmptyBitMap which is a BitMap with all
// bits cleared.
type BitMap uint64

// Set sets a specified bit (true). The bit to set is specified using its index (0 - 63). The updated bitmap, where
// the specified bit is set, is returned. If the specified bit is out-of-range a panic is raised.
func (f BitMap) Set(bit int) BitMap {
	checkIdx(bit, highBit)
	return f | (1 << bit)
}

// Clear clears a specified bit (false). The bit to clear is specified using its index (0 - 63). The updated bitmap,
// where the specified bit, is cleared is returned. If the specified bit is out-of-range a panic is raised.
func (f BitMap) Clear(bit int) BitMap {
	checkIdx(bit, highBit)
	return f &^ (1 << bit)
}

// Toggle toggles a specified bit. That is if the bit is cleared it is set and if the bit is set it is cleared.
// The bit to toggle is specified using its index (0 - 63). The updated bitmap, where the specified bit is toggled,
// is returned. If the specified bit is out-of-range a panic is raised.
func (f BitMap) Toggle(bit int) BitMap {
	checkIdx(bit, highBit)
	return f ^ (1 << bit)
}

// Has returns true if the specified bit is set. Otherwise, false is returned. The bit to check is specified using
// its index (0 - 63). If the specified bit is out-of-range a panic is raised.
func (f BitMap) Has(bit int) bool {
	checkIdx(bit, highBit)
	return f&(1<<bit) != 0
}

// Empty returns true if all bits in the bitmap are cleared. If at least bit is set false is returned.
func (f BitMap) Empty() bool {
	return f == 0
}

// String creates a string representation of a BitMap using a list where the index of each set bit is presented.
//
//	bit 3, 5 and 22 is set: [3,5,22]
func (f BitMap) String() string {
	return f.StringFunc("[", "]", ",", func(idx int) string { return strconv.Itoa(idx) })
}

func (f BitMap) StringFunc(start, end, separator string, fn func(idx int) string) string {
	var sb strings.Builder
	sb.WriteString(start)
	first := true
	for i := 0; i <= highBit; i++ {
		if f.Has(i) {
			if !first {
				sb.WriteString(separator)
			}
			sb.WriteString(fn(i))
			first = false
		}
	}
	sb.WriteString(end)
	return sb.String()
}

func checkIdx(idx, high int) {
	// We prefer fail-fast, so we panic if the bitmap index is out-of-range
	if idx < 0 || idx > high {
		panic(fmt.Sprintf("bitmap: index out-of-range (0 - %d): %d", high, idx))
	}
}

// EmptyBitMap is an empty bitmap where all bits are cleared (false). It is usually used as the starting BitMap from
// where you set the bits needed.
const EmptyBitMap = BitMap(0)

func And(left, right BitMap) BitMap {
	return left & right
}

func Or(left, right BitMap) BitMap {
	return left | right
}

func Xor(left, right BitMap) BitMap {
	return left ^ right
}
