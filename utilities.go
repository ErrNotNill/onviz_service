package main

import (
	"fmt"
	"time"
)

func Timer(count int) {
	var timer int
	seconds := time.Second * 1
	for {
		timer++
		time.Sleep(seconds)
		if timer > 7 && timer < 18 {
			count++
			fmt.Println("count: ", count)
			fmt.Println("is working time: ", timer)
			if count > 12 {
				fmt.Println("count was stopped in: ", count)
				return
			}
		}
		if timer > 24 {
			timer = 0
		}
		fmt.Println(timer)
	}
}
