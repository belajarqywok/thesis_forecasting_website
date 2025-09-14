package helpers

import (
	"fmt"
	"runtime"
)

func MemoryUsage(tag string) {
	var memory runtime.MemStats
	runtime.ReadMemStats(&memory)

	fmt.Printf(
		"[Tag: %s] Alloc = %v MB | TotalAlloc = %v MB | Sys = %v MB | NumGC = %v\n",
		tag, memory.Alloc/1024/1024, 
		memory.TotalAlloc/1024/1024, 
		memory.Sys/1024/1024, memory.NumGC,
	)
}
