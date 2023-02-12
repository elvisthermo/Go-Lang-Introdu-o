package main

import "fmt"

func main() {
   // Operadores de igualdade e não-igualdade
   a := 5
   b := 3
   fmt.Println("a == b:", a == b)
   fmt.Println("a != b:", a != b)

   // Operador E lógico
   fmt.Println("true && true:", true && true)
   fmt.Println("true && false:", true && false)
   fmt.Println("false && false:", false && false)

   // Operador OU lógico
   fmt.Println("true || true:", true || true)
   fmt.Println("true || false:", true || false)
   fmt.Println("false || false:", false || false)

   // Operador de negação
   fmt.Println("!true:", !true)
   fmt.Println("!false:", !false)
}