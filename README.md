# float.. wait, what?

`144.96 > 144.96`... but also `144.96 < 144.96`
```go
package main

import "fmt"

func main() {
	var x float32 = 144.96
	var y float64 = 144.96
	fmt.Print(float64(x) > y, x > float32(y))
	// Output: true false
}
```

`0.1 != 0.1`
```go
package main

import "fmt"

func main() {
	var x float32 = 0.1 // 0.100000001490116119384765625
	var y float64 = 0.1 // 0.1000000000000000055511151231257827021181583404541015625
	fmt.Print(float64(x) == y)
	// Output: false
}
```

`0.1 + 0.1 + 0.1 != 0.3`
```go
package main

import "fmt"

func main() {
	var x float64 = 0.1
	fmt.Println(x + x + x)
	// Output: 0.30000000000000004
}
```

Increment on assignment
```go
package main

import "fmt"

func main() {
	var fx float32 = 99999999
	var ix int32 = int32(fx)
	fmt.Println(ix)
	// Output: 100000000
}
```

`9 * 1/9 != 1`
```go
package main

import (
	"fmt"
)

func main() {
	var x float64 = 1 / float64(9)
	fmt.Println(x + x + x + x + x + x + x + x + x)
}
```

`NaN != NaN`
```go
package main

import "fmt"

func main() {
	var a float32 = 1
	var x float32 = .0 / (a - 1) // NaN
	fmt.Print(x == x)
	// Output: false
}
```

`NaN` is not in the map... and yet map is filled with `NaN`s
```go
package main

import (
	"fmt"
	"math"
)

func main() {
	m := make(map[float64]int)
	m[math.NaN()] = 1
	m[math.NaN()] = 2
	m[math.NaN()] = 3

	_, ok := m[math.NaN()]
	fmt.Println("does map contain NaN?", ok) // false, NaN != NaN

	for k, v := range m {
		fmt.Println(k, v)
	}
	// Output:
	// NaN 1
	// NaN 2
	// NaN 3
}
```
