package main

import (
	"fmt"
	"time"
)

type Process struct {
	Id    uint64
	Value uint64
}

func Run(p Process, kill chan uint64, show chan bool) {
	for {
		select {
		case k := <-kill:
			if k == p.Id {
				return
			} else {
				kill <- k
			}
		case s := <-show:
			if s {
				fmt.Printf("\n%d: %d", p.Id, p.Value)
			}
			p.Value += 1
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func Show(in, out chan bool) {
	s := false
	for {
		select {
		case <-in:
			s = !s
		default:
			out <- s
		}
	}
}

func main() {
	op := -1
	var id uint64 = 0
	var killId uint64
	out := make(chan bool)
	in := make(chan bool)
	kill := make(chan uint64)
	go Show(in, out)
	for op != 4 {
		fmt.Print("1. Add\n2. Show\n3. Kill\n4. Exit\n > ")
		fmt.Scan(&op)
		switch op {
		case 1:
			p := Process{Id: id}
			id += 1
			go Run(p, kill, out)
		case 2:
			in <- true
		case 3:
			fmt.Scan(&killId)
			kill <- killId
		case 4:
			fmt.Println("Goodbye")
		}
	}
}
