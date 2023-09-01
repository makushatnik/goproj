package utils

import (
	"fmt"

	"github.com/thanhpk/randstr"
)

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func ContainsInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func getRandomStr() string {
	MyString := randstr.String(20)
	fmt.Println(MyString)
	return MyString
}
