package main

import (
	"fmt"
	metrics "github.com/Akilan1999/Go-Metrics-Simple"
	"time"
)

func f(from string) {
	for i := 0; i < 1000000; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Add this to the starting point of your go program
	err := metrics.RunCollector(metrics.DefaultConfig)

	if err != nil {
		fmt.Println(err)
	}

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

	// Add this to the end point of your go program
	metrics.ComputeDefaultFile()
}
