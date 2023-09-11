package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	i := 123
	s := strconv.Itoa(i)
	
	strings.HasPrefix(s, "1")

	fmt.Printf("A string é %s\n", s)
}
