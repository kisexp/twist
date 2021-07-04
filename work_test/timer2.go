package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("timer 1 expired")

	timer2 := time.NewTimer(time.Second * 5)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 expired")
	}()

	stop2 := timer2.Stop()
	fmt.Println("stop2:", stop2)
	if stop2 {
		fmt.Println("timer 2 stoped")
	}
}
