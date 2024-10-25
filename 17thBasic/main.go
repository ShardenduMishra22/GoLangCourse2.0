package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("This is a deferred statement executed at last")

	CurrentTime := time.Now()
	formattedTime := CurrentTime.Format("02-01-2006 3:04:05 PM Monday")
	println(formattedTime)
}
