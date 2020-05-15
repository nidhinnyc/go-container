package main

import (
	"fmt"
	"math"

	"github.com/nidhinnyc/go-learning/03_packages/sampkg"
	"github.com/nidhinnyc/go-learning/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(100.6))
	fmt.Println(math.Ceil(3.3))
	fmt.Println(math.Sqrt(36))
	fmt.Println(strutil.Reverse("hello"))
	fmt.Println(sampkg.Greeting("Namaskaaram !!!"))
}
