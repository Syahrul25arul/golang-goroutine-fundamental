package goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {

	// pool := sync.Pool{}

	/*
	*  untuk mengoveride isi data pool ketika sudah habis
	 */
	pool := &sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("hendrik")
	pool.Put("Rizal")
	pool.Put("Array")

	/*
	* setelah menggunakan data pool, data nya harus dikembalikan lagi
	 */
	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			// time.Sleep(1 * time.Second)
			pool.Put(data)
			// fmt.Println("put")
		}()
	}
	// time.Sleep(5 * time.Second)
	fmt.Println("Selesai")
}
