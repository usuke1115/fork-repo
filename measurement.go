package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	day := now.Day()
	fmt.Printf("Today's day is %d.\n", day)
}
