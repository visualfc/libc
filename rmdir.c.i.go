package libc

import unsafe "unsafe"

func rmdir(path *int8) int32 {
	return int32(__syscall1_r1(int64(137), int64(uintptr(unsafe.Pointer(path)))))
}
