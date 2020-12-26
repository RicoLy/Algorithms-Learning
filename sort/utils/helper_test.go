package utils

import (
	"fmt"
	"testing"
)

func TestGetArrayOfSize(t *testing.T) {
	arr := GetArrayOfSize(10000)
	fmt.Println(arr)
	fmt.Println(len(arr))
}
