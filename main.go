package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Введите строку")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	a := scanner.Text()
	fmt.Printf("%d\n", strings.Count(a, " "))
	for j := 0; j < strings.Count(a, " ")+1; j++ {
		for i := 0; i < len(a); i++ {
			if string(a[i]) != " " {
				fmt.Print(string(a[i]))
			} else {
				fmt.Println()
				fmt.Print("\033[H\033[2J")
			}
			time.Sleep(2000 * time.Millisecond)
		}
	}
}
