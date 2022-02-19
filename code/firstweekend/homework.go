package main

import (
	"fmt"
	"time"
)

func product(th chan<- int) {
	for i := 0; i <= 9; i++ {
		th <- i
		time.Sleep(1 * time.Second)
		fmt.Printf("生产编号：%d\n", i)
	}
	close(th)
}

func consumer(th <-chan int) {
	//for i := 0; i <= 9; i++ {
	for num := range th {
		//num := <-th
		time.Sleep(1 * time.Second)
		fmt.Printf("消费编号：%d\n", num)
	}
}

func main() {
	th := make(chan int, 5)
	go product(th)
	consumer(th)

	array := []string{"I", "am", "stupid", "and", "weak"}

	for char := range array {
		array[2] = "smart"
		array[4] = "strong"
		fmt.Println(array[char])
	}
	fmt.Println(array)
}
