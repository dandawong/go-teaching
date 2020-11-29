package variable

import (
	"fmt"
	"math"
)

func variableBasic() {
	var i int = 10               // int
	var f float32 = 10.23        // float
	var str string = "hello, go" // string
	var c complex64 = 1 + 2i     // complex

	fmt.Println(i, f, str, c)
}

func variableAdvance() {
	var arr1 [3]int                   // array
	arr1[0] = 1                       // set array value
	var arr2 [3]int = [3]int{1, 2, 3} // array literal

	var slice1 []int = make([]int, 3) // slice
	slice1 = append(slice1, 1)        // append a slice
	slice2 := []int{1, 2, 3, 4}       // slice literal

	var map1 map[string]int = make(map[string]int) // mapping
	map1["hello"] = 1                              // set map value
	map1["world"] = 2                              // set map value
	delete(map1, "omg")                            // remove value
	i, ok := map1["aaa"]                           // get map value

	fmt.Println(arr1, arr2, slice1, slice2, i, ok, map1)
}

func variableMaster() {
	// master variable
	type Person struct { // create struct type
		firstName string // struct data
		lastName  string
	}
	var p Person = Person{"Danda", "Wong"}

	// function type
	type Operator func(x float64) float64
	// Map applies op to each element of a.
	mapFunc := func(op Operator, a []float64) []float64 {
		res := make([]float64, len(a))
		for i, x := range a {
			res[i] = op(x)
		}
		return res
	}
	op := math.Abs
	a := []float64{1, -2}
	b := mapFunc(op, a)

	fmt.Println(p, a, b)
}

// RunVariable run the example
func RunVariable() {
	fmt.Println("--- basic ---")
	variableBasic()

	fmt.Println("--- advance ---")
	variableAdvance()

	fmt.Println("--- master ---")
	variableMaster()
}
