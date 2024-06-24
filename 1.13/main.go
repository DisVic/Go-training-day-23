package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n, a int
	_, _ = fmt.Scan(&n, &a)
	number := strings.ReplaceAll(strconv.Itoa(n), strconv.Itoa(a), "")
	fmt.Print(number)
}
