package utils

import (
	"bufio"
	"os"
	"strconv"
)

func GetArrayOfSize(n int) []int {
	//p, err := build.Default.Import("Algorithms-Learning-With-Go/sort/utils", "I:/info", build.FindOnly)
	//if err != nil {
	//	panic(err)
	//}
	//fname := filepath.Join( "IntegerArray.txt")
	//f, _ := os.Open(fname)
	f,_ := os.Open("IntegerArray.txt")
	defer f.Close()
	numbers := make([]int, n)
	scanner := bufio.NewScanner(f)
	//for scanner.Scan() {
	//	s, _ := strconv.Atoi(scanner.Text())
	//	numbers = append(numbers, s)
	//}
	for i := 0; i < n; i++ {
		scanner.Scan()
			//s, _ := strconv.Atoi(scanner.Text())
			numbers[i],_ = strconv.Atoi(scanner.Text())
	}

	return numbers[0:n]
}
