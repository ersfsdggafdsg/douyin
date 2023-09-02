package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	t, _ := strconv.ParseInt(os.Args[1], 10, 64)
	fmt.Println(time.Unix(t / 1000, t % 1000 * 1000000))
}

