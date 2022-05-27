package main

import (
	"fmt"
	"math"
)

func gcd(a, b float64) float64 {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

func pollard(n float64) float64{
	var a, i, d float64
	a = 2
	i = 2
	for {
		a = math.Mod(math.Pow(a, i), n)
		d = gcd((a-1), n)
		if (1 < d) && (d < n) {
			return d
			break 
		}
		i += 1 
	}
	return 0
}

func main(){
	var a, b int
	var n float64

	fmt.Print("Enter the number to factorize : ")

	fmt.Scanln(&n)
	
	a = int(pollard(n))
	b = int(int(n) / a) 

	fmt.Printf("The factors of %d are : %d and %d \n",int(n), a, b)
}