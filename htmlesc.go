package main

import (
	"bufio"
	"fmt"
	"html"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	t, _ := r.ReadString('\n')
	fmt.Printf("%s", html.UnescapeString(t))
}
