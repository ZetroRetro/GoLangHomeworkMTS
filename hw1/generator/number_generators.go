package generator

import "unsafe"

func SequenceGenerator(start uint64) func(...any) uint64 {
	next := start
	return func(...any) uint64 {
		current := next
		next++
		return current
	}
}

func FromAddressGenerator() func(...any) uint64 {
	return func(obj ...any) uint64 {
		if len(obj) == 0 {
			return 0
		}

		var res uint64

		for _, item := range obj {
			res *= 2
			res += uint64(uintptr(unsafe.Pointer(&item)))
		}

		return res
	}
}
