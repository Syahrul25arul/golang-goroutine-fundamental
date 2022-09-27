package goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGomaxprocs(t *testing.T) {

	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	// total cpu
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	// total thread
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	// mengubah jumlah thread
	runtime.GOMAXPROCS(20)
	totalThreadNew := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread New", totalThreadNew)
	// total Goroutine
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine", totalGoroutine)

	group.Wait()
}
