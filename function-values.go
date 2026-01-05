package main

import (
	"fmt"
	"math"
)

// compute は「2つの float64 を受け取って float64 を返す関数」を受け取り、
// その関数に (3, 4) を渡して結果を返す
func compute(fn func(float64, float64) float64) float64 {
	// fn(3, 4)
	return fn(3, 4)
}

func main() {
	// hypot は直角三角形の斜辺の長さを求める関数
	// 数学的には：√(x² + y²)
	hypot := func(x, y float64) float64 {
		// x*x + y*y  = x² + y²
		// math.Sqrt = √
		return math.Sqrt(x*x + y*y)
	}

	// √(5² + 12²) = √(25 + 144) = √169 = 13
	fmt.Println(hypot(5, 12))

	// compute(hypot)
	// = hypot(3, 4)
	// = √(3² + 4²) = √25 = 5
	fmt.Println(compute(hypot))

	// compute(math.Pow)
	// = math.Pow(3, 4)
	// = 3⁴ = 81
	fmt.Println(compute(math.Pow))
}
