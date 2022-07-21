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

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hendrik Array"
}

func TestChannel(t *testing.T) {
	channel := make(chan string, 1)
	defer close(channel)
	// go func() {
	// 	fmt.Println("channel recive")
	// 	data := <-channel
	// 	fmt.Println(data)
	// }()

	// go func() {
	// 	fmt.Println("goroutin")
	// 	// time.Sleep(2 * time.Second)
	// 	channel <- "Hendrik"
	// }()

	go func() {
		data := <-channel
		fmt.Println(data)
	}()
	channel <- "Hendrik"

	// time.Sleep(2 * time.Second)
}

func TestChannelAsParamter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel) // channel by default menganut pass by refrence, artinya jika kita memasukkan channel ke argument, tidak perlu membuat channel sebagai pointer
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	fmt.Println("Only In")
	time.Sleep(2 * time.Second)
	channel <- "Hendrik Array"
	fmt.Println("Only In pass value to channel")
}

func OnlyOut(channel <-chan string) {
	fmt.Println("Only Out")
	data := <-channel
	fmt.Println(data)
	fmt.Println("recive value from channel")
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 5) // dengan menggunakan buffer channel, channel tidak akan memblocking hingga hingga batas memasukkan data ke channel telah habis
	defer close(channel)

	channel <- "hendrik"

	fmt.Println("selesai")
}
