package main

import (
	"reflect"
)

type StructField[T any] struct {
	value              *T
	reflectStructField *reflect.StructField
}

type StructTag = reflect.StructTag

func (m *StructField[T]) ReflectStructField() *reflect.StructField {
	return m.reflectStructField
}

func (m *StructField[T]) IsExported() bool {
	return m.ReflectStructField().IsExported()
}

func (m *StructField[T]) Name(name ...string) string {
	if len(name) > 0 {
		m.ReflectStructField().Name = name[0]
	}
	return m.ReflectStructField().Name
}

func (m *StructField[T]) Type() Type[T] {
	return TypeOf[T]()
}

func SetStructFieldType[To any, From any](m *StructField[From], typ ...Type[To]) *StructField[To] {
	var v To
	m.ReflectStructField().Type = reflect.ValueOf(v).Type()
	return &StructField[To]{value: &v, reflectStructField: m.ReflectStructField()}
}

func (m *StructField[T]) PkgPath(path ...string) string {
	if len(path) > 0 {
		m.ReflectStructField().PkgPath = path[0]
	}
	return m.ReflectStructField().PkgPath
}

func (m *StructField[T]) Index(i ...[]int) []int {
	if len(i) > 0 {
		m.ReflectStructField().Index = i[0]
	}
	return m.ReflectStructField().Index
}

func (m *StructField[T]) Tag(tag ...StructTag) StructTag {
	if len(tag) > 0 {
		m.ReflectStructField().Tag = tag[0]
	}
	return m.ReflectStructField().Tag
}

func (m *StructField[T]) Offset(i ...uintptr) uintptr {
	if len(i) > 0 {
		m.ReflectStructField().Offset = i[0]
	}
	return m.ReflectStructField().Offset
}

func VisibleFields[T any](typ ...Type[T]) []StructField[any] {
	var t T
	fields := reflect.VisibleFields(reflect.ValueOf(t).Type())
	sfields := make([]StructField[any], len(fields))
	for i, f := range fields {
		sfields[i] = StructField[any]{value: ptr(any(t)), reflectStructField: &f}
	}
	return sfields
}

