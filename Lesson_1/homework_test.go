package main

import (
	"unsafe"

	"github.com/stretchr/testify/assert"

	"testing"
)

// go test -v homework_test.go

func ToLittleEndian[T uint16 | uint32 | uint64](number T) T {
	size := int(unsafe.Sizeof(number))

	var result T
	ptr := unsafe.Pointer(&number)
	resultPtr := unsafe.Pointer(&result)

	for i := 0; i < size; i++ {
		srcByte := *(*byte)(unsafe.Pointer(uintptr(ptr) + uintptr(size-i-1)))
		dstByte := (*byte)(unsafe.Pointer(uintptr(resultPtr) + uintptr(i)))
		*dstByte = srcByte
	}

	return result
}

func TestConversionUint32(t *testing.T) {
	tests := map[string]struct {
		number uint32
		result uint32
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x04030201,
		},
		"test case #6": {
			number: 0xAABBCCDD,
			result: 0xDDCCBBAA,
		},
		"test case #7": {
			number: 0x12345678,
			result: 0x78563412,
		},
		"test case #8": {
			number: 0x00000001,
			result: 0x01000000,
		},
		"test case #9": {
			number: 0x01000000,
			result: 0x00000001,
		},
		"test case #10": {
			number: 0x00010000,
			result: 0x00000100,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestConversionUint64(t *testing.T) {
	tests := map[string]struct {
		number uint64
		result uint64
	}{
		"test case #1": {
			number: 0x0000000000000000,
			result: 0x0000000000000000,
		},
		"test case #2": {
			number: 0xFFFFFFFFFFFFFFFF,
			result: 0xFFFFFFFFFFFFFFFF,
		},
		"test case #3": {
			number: 0x01_23_45_67_89_AB_CD_EF,
			result: 0xEF_CD_AB_89_67_45_23_01,
		},
		"test case #4": {
			number: 0x0000FFFF0000FFFF,
			result: 0xFFFF0000FFFF0000,
		},
		"test case #5": {
			number: 0xFFFF0000FFFF0000,
			result: 0x0000FFFF0000FFFF,
		},
		"test case #6": {
			number: 0xAAAACCCCEEEEFFFF,
			result: 0xFFFFEEEECCCCAAAA,
		},
		"test case #7": {
			number: 0x000000000000000F,
			result: 0x0F00000000000000,
		},
		"test case #8": {
			number: 0x0F00000000000000,
			result: 0x000000000000000F,
		},
		"test case #9": {
			number: 0x00_00_00_0F_00_00_00_00,
			result: 0x00_00_00_00_0F_00_00_00,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestConversionUint16(t *testing.T) {
	tests := map[string]struct {
		number uint16
		result uint16
	}{
		"test case #1": {
			number: 0x0000,
			result: 0x0000,
		},
		"test case #2": {
			number: 0xFFFF,
			result: 0xFFFF,
		},
		"test case #3": {
			number: 0x0FF0,
			result: 0xF00F,
		},
		"test case #5": {
			number: 0x00FF,
			result: 0xFF00,
		},
		"test case #6": {
			number: 0x000F,
			result: 0x0F00,
		},
		"test case #7": {
			number: 0x1234,
			result: 0x3412,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
