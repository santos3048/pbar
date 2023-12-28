package main

import (
	"fmt"
	"pbar/pbar"
	"time"
)

func main() {
	b := pbar.Create("Doing some task", 400)
	go b.Print()
	i := 0
	for i <= 300 {
		time.Sleep(5 * time.Millisecond)
		b.Msg(fmt.Sprintf("Performing task #%d", i))
		b.Up()
		i++
	}
	b.Stop()
	time.Sleep(1 * time.Second) 
	go b.Print()
	for i <= 400 {
		time.Sleep(50 * time.Millisecond)
		b.Msg(fmt.Sprintf("Performing task #%d", i))
		b.Up()
		i++
	}
	b.Finish("Performed all tasks succesfully!")
}
