package main

import (
	"fmt"
	"math/big"
)

func factorial(x *big.Int) *big.Int {
	if x.Cmp(new(big.Int).SetInt64(1)) == 1 {
		return new(big.Int).Mul(x, factorial(new(big.Int).Sub(x, new(big.Int).SetInt64(1))))
	} else {
		return new(big.Int).SetInt64(1)
	}
	
}

func main() {
	var temp big.Int
	fmt.Print("ENTER THE NUMBER TO GET FACTORIAL:")
	fmt.Scanln(&temp)
	fmt.Println(factorial(&temp))
}

