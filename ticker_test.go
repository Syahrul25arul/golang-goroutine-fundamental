package goroutines

import (
	"fmt"
	"testing"
	"time"
)

// ticker adalah reperesentasi kejadian yang berulang

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second) // akan di jalankan berulang setiap satu detik. mengembalikan object ticker yang mempunyai channel
	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()
	for time := range ticker.C { // terjadi deadlock karna ketika ticker di stop, time masih menunggu data dari channel
		fmt.Println(time)
	}

}

func TestTickObject(t *testing.T) {
	ticker := time.Tick(1 * time.Second) // hanya mengembalikan channel

	for time := range ticker {
		fmt.Println(time)
	}

}
