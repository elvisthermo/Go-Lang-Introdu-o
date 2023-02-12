package main

import "fmt"

func sun(a int, b int) int {
   return a + b
}

func main() {
   result := sun(5, 3)
   fmt.Println("Result:", result)
}