package main

import (
	"fmt"
	"time"
)

func main() {
	go start()
	follow()

	time.Sleep(time.Second * 5)
}

func start() {
	fmt.Println("starter func")
}

func follow() {
	fmt.Println("follow func")

}
