package main

import (
	"fmt"
)

func main() {
	var x int64
	var y int64
	fmt.Print("Informe dois numeros")
	fmt.Scan(&x, &y)

	if x > y {
		fmt.Println(x, "É maior que ", y)
	} else if x < y {
		fmt.Println(x, "É menor que ", y)
	} else if x == y {
		fmt.Println("Os dois sao iguais")
	}
}
