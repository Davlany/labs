package garbageCollector_test

import (
	"runtime"
	"testing"
)

func TestGarbageColl(t *testing.T) {
	var m runtime.MemStats
	s := make([]int, 10000000)
	_ = s
	s = nil
	runtime.GC()
	runtime.ReadMemStats(&m)

	if m.Alloc/1024/1024 != 0 {
		t.Errorf("EXPECTED 0 \t RESULT \t %s%d%s \t❌", "\033[31m", m.Alloc/1024/1024, "\033[0m")
	} else {
		t.Logf("EXPECTED 0 \t RESULT \t %s%d%s \t✅", "\u001B[32m", m.Alloc/1024/1024, "\033[0m")
	}
}
