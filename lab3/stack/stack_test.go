package stack_test

import (
	"runtime"
	"testing"
)

func TestStack(t *testing.T) {
	var m runtime.MemStats
	var testValue int
	runtime.ReadMemStats(&m)
	alloc1 := m.Alloc
	_ = testValue
	testValue = 10
	runtime.ReadMemStats(&m)
	alloc2 := m.Alloc
	var flag = false
	if alloc2/1024/1024 != 0 {
		t.Errorf("BEFORE ADD TO STACK %d AFTER ADD TO STACK (EXPECTED 0) \t RESULT \t %s%d%s \t❌", alloc1/1024/1024, "\033[31m", alloc2/1024/1024, "\033[0m")
	} else {
		t.Logf("BEFORE ADD TO STACK %d AFTER ADD TO STACK (EXPECTED 0) \t RESULT \t %s%d%s \t✅", alloc1/1024/1024, "\u001B[32m", alloc2/1024/1024, "\033[0m")
		flag = true
	}
	if flag {
		runtime.ReadMemStats(&m)
		alloc1 := m.Alloc
		arr := make([]int, 1000000)
		_ = arr
		runtime.ReadMemStats(&m)
		alloc2 := m.Alloc
		if alloc1 >= alloc2 {
			t.Errorf("BEFORE ADD TO HEAP %d AFTER ADD TO HEAP (EXPECTED >) \t RESULT \t %s%d%s \t❌", alloc1/1024/1024, "\033[31m", alloc2/1024/1024, "\033[0m")
		} else {
			t.Logf("BEFORE ADD TO HEAP %d AFTER ADD TO HEAP (EXPECTED >) \t RESULT \t %s%d%s \t✅", alloc1/1024/1024, "\u001B[32m", alloc2/1024/1024, "\033[0m")
		}
	}
}
