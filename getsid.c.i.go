package libc

func getsid(pid int32) int32 {
	return int32(__syscall1_r1(int64(310), int64(pid)))
}
