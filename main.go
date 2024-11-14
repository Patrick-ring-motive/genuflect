package main

import (
	"fmt"
	"reflect"
)

/*func ptr[T any](v T) *T {
	return &v
}*/

func Copy[X any, Y any](dst X, src Y) int {
	return reflect.Copy(reflect.ValueOf(dst), reflect.ValueOf(src))
}

func DeepEqual[X any, Y any](x X, y Y) bool {
	return reflect.DeepEqual(reflect.ValueOf(x), reflect.ValueOf(y))
}

func Swapper[T any](slice T) func(i, j int) {
	return reflect.Swapper(reflect.ValueOf(slice))
}

func main() {
	fmt.Println(Convert[int](42.3))
	fmt.Println()
}
