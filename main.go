package main

import (
	// "bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/fatih/color"
	"github.com/huandu/xstrings"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	os.Getenv("S3_BUCKET")

	some_print()

	go spinner()
	res := fib(44)
	fmt.Printf("\rs44 num of fib = %d\n", res)
}

func spinner() {
	for {
		for _, c := range `-\|/` {
			fmt.Printf("\r%c", c)
			time.Sleep(100 * time.Microsecond)
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func some_print() {
	m := map[string]int{}
	m["third"] = 3

	fmt.Println(m)
	// os.Open()
	// bufio.NewScanner()
	color.Red("This is a Red String")
	fmt.Println(xstrings.Shuffle("Shuffled String! big and beauty"))
}
