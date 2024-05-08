package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

/*
NOTE:
range interates over an array and returned 2 values
	-index and value
since we didn't want to focus on index you can leave it blank
	with "_" (blank identifier)
the size of array is encoded with it type
	try to pass a [4]int in a function expecting a [5]int and compile will fail, this also applies with types
*/

/*
NOTE: Sum with old for loop
func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += number[i]
	}
	return sum
}
*/
