package main

import "fmt"

// loop tradicional
func loopsNormal() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// loop estilo while
func loopsWhile() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

// loop infinito
func loopsInfinito() {
	for {
		fmt.Println("loop infinito")
		// cuidado! vai rodar para sempre
	}
}

// loop com range (for-each)
func loopsRange() {
	valores := []string{"a", "b", "c"}
	for _, v := range valores {
		fmt.Println(v)
	}
}

func main() {
	loopsNormal()
	loopsWhile()
	loopsRange()
	// loopsInfinito() // cuidado ao descomentar
}
