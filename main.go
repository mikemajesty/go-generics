package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	~int | ~int64 | ~float64
}

type MyNumber int

func SumGeneric[T Number](m map[string]T) T {
	var sum T
	for _, value := range m {
		sum += value
	}
	return sum
}

func Sum[T Number](numer1 T, number2 T) T {
	if numer1 == number2 {
		return numer1

	}
	return numer1 + number2
}

func Max[T constraints.Ordered](numer1 T, number2 T) T {
	if numer1 > number2 {
		return numer1
	}
	return number2
}

type Stringer interface {
	String() string
}

type MyString string

func (m MyString) String() string {
	return string(m)
}

func Concat[T Stringer](vals []T) string {
	var result string = ""
	for _, val := range vals {
		result += val.String()
	}
	return result
}

func main() {
	fmt.Println(Concat([]MyString{"1", "2", "3"}))
	fmt.Println(SumGeneric(map[string]MyNumber{"a": 1, "b": 2, "c": 3}))
	fmt.Println(SumGeneric(map[string]int64{"a": 1, "b": 2, "c": 3}))
	fmt.Println(SumGeneric(map[string]float64{"a": 1., "b": 2.5, "c": 3}))
}
