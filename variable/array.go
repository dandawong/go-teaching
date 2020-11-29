package variable

import "fmt"

func arrayBasic() {
	// Create array
	var array1 [3]int = [3]int{1, 2, 3}
	// Copy array
	array2 := array1
	// Set array data
	array3 := array2
	array3[1] = 17

	fmt.Println(array1, array2, array3)
}

func arrayAdvance() {
	var array1 [5]int = [5]int{1, 3, 5, 7, 9}
	array2 := array1[1:3]

	fmt.Println(len(array1), array1, len(array2), array2)
}

func arrayMaster() {
	var array1 [5]int = [5]int{1, 3, 5, 7, 9}
	var array2 [5]int = [5]int{2, 4, 6, 8, 10}
	var array3 [10]int

	copy(array3[:], append(array1[:], array2[:]...))

	fmt.Println(array1, array2, array3)
}

// RunArray run the example
func RunArray() {
	fmt.Println("--- basic ---")
	arrayBasic()
	fmt.Println("--- advance ---")
	arrayAdvance()
	fmt.Println("--- master ---")
	arrayMaster()
}
