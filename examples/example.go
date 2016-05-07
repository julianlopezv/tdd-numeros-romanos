package main

import (
    "bufio"
    "fmt"
    "os"
	
	"github.com/julianlopezv/tdd-roman-numbers"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Ingrese numero romano: ")
    text, _ := reader.ReadString('\n')
    fmt.Println(roman.RomanToDecimal(text))
}