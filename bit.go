package consentstring

import (
	"fmt"
	"time"
)

// custom error types
var (
	ErrISO639      = fmt.Errorf("language was not entered in two-letter ISO639-1")
	ErrPurposeList = fmt.Errorf("incorrect purposes array size")
)

// bit is a convience type to abstract []byte
type bit []byte

// setBit sets a single byte at a specified position
func (b bit) setBit(index int) {
	byteIndex := index / 8
	shift := uint((byteIndex+1)*8 - index - 1)
	b[byteIndex] |= byte(1 << shift)
}

// unsetBit unsets a single byte at a specified position
func (b bit) unsetBit(index int) {
	byteIndex := index / 8
	shift := uint((byteIndex+1)*8 - index - 1)
	b[byteIndex] &= ^(1 << shift)
}

func (b bit) setPurposes(purposes []uint) error {
	if len(purposes) != 5 {
		return ErrPurposeList
	}

	for i, p := range purposes {
		if p == 1 {
			b.setBit(purposesOffset + i)
		} else {
			b.unsetBit(purposesOffset + i)
		}
	}

	return nil
}

// setInt sets an interger with specified padding
func (b bit) setInt(startInclusive int, size int, to int) {
	b.setNumber(startInclusive, size, to)
}

// setDateToDeciseconds converts time.Time to deciseconds and sets it's bit values
func (b bit) setDateToDeciseconds(startInclusive int, size int, t time.Time) {
	deciseconds := int64(t.UnixNano() / 1e8)
	b.setNumberInt64(startInclusive, size, deciseconds)
}

// setSixBitString converts the a two-letter ISO639-1 language code to it's bit values
func (b bit) setSixBitString(startInclusive int, size int, lang string) error {
	if len(lang) != 2 {
		return ErrISO639
	}
	for i, r := range lang {
		charCode := int(r) - 65
		b.setInt(startInclusive+(i*6), 6, charCode)
	}
	return nil
}

func (b bit) setNumberInt64(startInclusive int, size int, to int64) {
	for i := size - 1; i >= 0; i-- {
		index := startInclusive + i
		byteIndex := index / 8
		shift := uint((byteIndex+1)*8 - index - 1)
		b[byteIndex] |= byte((to % 2) << shift)
		to /= 2
	}
}

func (b bit) setNumber(startInclusive int, size int, to int) {
	for i := size - 1; i >= 0; i-- {
		index := startInclusive + i
		byteIndex := index / 8
		shift := uint((byteIndex+1)*8 - index - 1)
		b[byteIndex] |= byte((to % 2) << shift)
		to /= 2
	}
}
