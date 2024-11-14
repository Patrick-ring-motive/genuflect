package main

import (
	"reflect"
)

type MapIter[K comparable, V any] struct {
	value          *map[K]V
	reflectMapIter *reflect.MapIter
}

func (m *MapIter[K, V]) ReflectMapIter() *reflect.MapIter {
	return m.reflectMapIter
}

func (m *MapIter[K, V]) Key() K {
	return m.ReflectMapIter().Key().Interface().(K)
}

func (m *MapIter[K, V]) Value() V {
	return m.ReflectMapIter().Value().Interface().(V)
}

func (m *MapIter[K, V]) Next() bool {
	return m.ReflectMapIter().Next()
}

func (m *MapIter[K, V]) Reset(v any) {
	m.ReflectMapIter().Reset(reflect.ValueOf(v))
}
