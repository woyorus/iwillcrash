package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hi! I will crash in 15 seconds. Just wait for it!")
	<-time.After(15 * time.Second)
	panic(errors.New("crashed as planned! ;)"))
}
