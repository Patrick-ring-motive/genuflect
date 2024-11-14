package main

import (
	"reflect"
)

type ChanDir = reflect.ChanDir

type Type[T any] struct {
	typ         *T
	reflectType *reflect.Type
}

func (t *Type[T]) ReflectType() *reflect.Type {
	return t.reflectType
}

func TypeOf[T any](value ...T) Type[T] {
	var t T
	return Type[T]{typ: ptr(t), reflectType: ptr(reflect.ValueOf(t).Type())}
}

func ArrayOf[T any](length int, elem ...Type[T]) Type[[]T] {
	t := make([]T, length)
	return Type[[]T]{typ: ptr(t), reflectType: ptr(reflect.ValueOf(t).Type())}
}

func ChanOf[T any](dir ChanDir, typ ...Type[T]) Type[chan T] {
	var t T
	var ct chan T
	return Type[chan T]{typ: ptr(ct), reflectType: ptr(reflect.ChanOf(dir, reflect.ValueOf(t).Type()))}
}

func FuncOf(in, out []Type[any], variadic bool) Type[any] {
	var fn Type[any]
	ins := make([]reflect.Type, len(in))
	for i, v := range in {
		ins[i] = *v.reflectType
	}
	outs := make([]reflect.Type, len(out))
	for i, v := range in {
		outs[i] = *v.reflectType
	}
	return Type[any]{typ: ptr(any(fn)), reflectType: ptr(reflect.FuncOf(ins, outs, variadic))}
}
func MapOf[K comparable, V any](key K, val V) Type[map[K]V] {
	return TypeOf[map[K]V]()
}
func PointerTo[T any](t ...Type[T]) Type[*T] {
	return TypeOf[*T]()
}
func PtrTo[T any](t ...Type[T]) Type[*T] {
	return TypeOf[*T]()
}
func SliceOf[T any](elem ...Type[T]) Type[[]T] {
	var t []T
	return Type[[]T]{typ: ptr(t), reflectType: ptr(reflect.ValueOf(t).Type())}
}
func StructOf[T any](fields []StructField[any]) Type[any] {
	f := make([]reflect.StructField, len(fields))
	for i, v := range fields {
		f[i] = *v.reflectStructField
	}
	return Type[any]{typ: ptr(any(f)), reflectType: ptr(reflect.StructOf(f))}
}
func TypeFor[T any]() Type[T] {
	return TypeOf[T]()
}
//func SliceAt(typ Type, p unsafe.Pointer, n int) Value