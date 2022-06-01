package dm

import "time"

type FuncBool func() bool
type Func func()

func LoopUntil(f FuncBool) bool {
	for {
		if f() {
			return true
		}
	}
	return false
}

func LoopUntilWithDuration(f FuncBool, d time.Duration) {
	for {
		if f() {
			break
		}
		time.Sleep(d)
	}
}

func LoopTimes(times uint, f Func) {
	var i uint
	for i = 0; i < times; i++ {
		f()
	}
}

func LoopTimesUntil(times uint, f FuncBool) bool {
	var i uint
	for i = 0; i < times; i++ {
		if f() {
			return true
			break
		}
	}
	return false
}

func LoopTimesUntilWithDuration(times uint, f FuncBool, d time.Duration) bool {
	var i uint
	for i = 0; i < times; i++ {
		if f() {
			return true
			break
		}
		time.Sleep(d)
	}
	return false
}
