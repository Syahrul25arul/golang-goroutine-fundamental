package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // channel harus di close jika selesai digunakan, jika tidak dia akan tetap ada menggantung di memori dan akan menyebabkan memory leak
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hendrik array"
		fmt.Println("selesai mengirim data ke channel")
	}()

	/*
	* error : all goroutines are asleep - deadlock
	* deadlock biasanya adalah proses goroutine yang menunggu menerima data tapi data tidak pernah dikirim dari goroutine lain
	 */

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}
