package main

import (
   "fmt"
   "math"
)

type Matematica struct{}

func (m Matematica) RaizQuadrada(x float64) float64 {
   return math.Sqrt(x)
}

func (m Matematica) Potencia(x float64, y float64) float64 {
   return math.Pow(x, y)
}

func (m Matematica) Fatorial(x int) int {
   if x == 0 {
      return 1
   }
   return x * m.Fatorial(x-1)
}

func main() {
   m := Matematica{}
   fmt.Println("Raiz quadrada de 9:", m.RaizQuadrada(9))
   fmt.Println("2 elevado ao cubo:", m.Potencia(2, 3))
   fmt.Println("Fatorial de 5:", m.Fatorial(5))
}
