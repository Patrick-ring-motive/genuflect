package main

import (
	"reflect"
)

type Method[T any] struct {
	value         *T
	reflectMethod *reflect.Method
}

func (m *Method[T]) ReflectMethod() *reflect.Method {
	return m.reflectMethod
}

func (m *Method[T]) IsExported() bool {
	return m.ReflectMethod().IsExported()
}

func (m *Method[T]) Name(name ...string) string {
	if len(name) > 0 {
		m.ReflectMethod().Name = name[0]
	}
	return m.ReflectMethod().Name
}

func (m *Method[T]) Type() Type[T] {
	return TypeOf[T]()
}

func SetMethodType[To any, From any](m *Method[From], typ ...Type[To]) *Method[To] {
	var v To
	m.ReflectMethod().Type = reflect.ValueOf(v).Type()
	return &Method[To]{value: &v, reflectMethod: m.ReflectMethod()}
}

func (m *Method[T]) PkgPath(path ...string) string {
	if len(path) > 0 {
		m.ReflectMethod().PkgPath = path[0]
	}
	return m.ReflectMethod().PkgPath
}

func (m *Method[T]) Index(i ...int) int {
	if len(i) > 0 {
		m.ReflectMethod().Index = i[0]
	}
	return m.ReflectMethod().Index
}

func (m *Method[T]) Func(fn ...T) T {
	if len(fn) > 0 {
		m.ReflectMethod().Func = reflect.ValueOf(fn[0])
	}
	return m.ReflectMethod().Func.Interface().(T)
}
