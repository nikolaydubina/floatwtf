package floatwtf_test

import (
	"fmt"
	"math"
	"sort"
)

func Example_compare_144_96() {
	var x float32 = 144.96
	var y float64 = 144.96
	fmt.Print(float64(x) > y, x > float32(y))
	// Output: true false
}

func Example_unequalOneTenth() {
	var x float32 = 0.1 // 0.100000001490116119384765625
	var y float64 = 0.1 // 0.1000000000000000055511151231257827021181583404541015625
	fmt.Print(float64(x) == y)
	// Output: false
}

func Example_threeTimesOneThird() {
	var x float64 = 0.1
	fmt.Println(x + x + x)
	// Output: 0.30000000000000004
}

func Example_incrementOnAssignment() {
	var fx float32 = 99999999
	var ix int32 = int32(fx)
	fmt.Println(ix)
	// Output: 100000000
}

func Example_nineTimesOneNinth() {
	var x float64 = 1 / float64(9)
	fmt.Println(x + x + x + x + x + x + x + x + x)
	// Output: 1.0000000000000002
}

func Example_nan() {
	var a float32 = 1
	var x float32 = .0 / (a - 1) // NaN
	fmt.Print(x == x)
	// Output: false
}

func Example_nanMap() {
	m := make(map[float64]int)
	m[math.NaN()] = 1
	m[math.NaN()] = 2
	m[math.NaN()] = 3

	_, ok := m[math.NaN()]
	fmt.Println("does map contain NaN?", ok)

	var ks []float64
	var vs []int
	for k, v := range m {
		ks = append(ks, k)
		vs = append(vs, v)
	}

	sort.Float64s(ks)
	sort.Ints(vs)

	fmt.Println(ks)
	fmt.Println(vs)
	// Output:
	// does map contain NaN? false
	// [NaN NaN NaN]
	// [1 2 3]
}
