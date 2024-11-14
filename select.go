package main

import (
	"reflect"
)

type SelectDir = reflect.SelectDir

type SelectCase[T any] struct {
	selectCase        *chan T
	reflectSelectCase *reflect.SelectCase
}

func (s *SelectCase[T]) ReflectSelectCase() *reflect.SelectCase {
	return s.reflectSelectCase
}

func (s *SelectCase[T]) Dir(i ...SelectDir) SelectDir {
	if len(i) > 0 {
		s.ReflectSelectCase().Dir = reflect.SelectDir(i[0])
	}
	return s.ReflectSelectCase().Dir
}

func (s *SelectCase[T]) Chan(v ...chan T) chan T{
  if len(v) > 0 {
    s.ReflectSelectCase().Chan = reflect.ValueOf(v[0])
  }
  return s.ReflectSelectCase().Chan.Interface().(chan T)
}

func (s *SelectCase[T]) Send(v ...T) T{
  if len(v) > 0 {
    s.ReflectSelectCase().Send = reflect.ValueOf(v[0])
  }
  return s.ReflectSelectCase().Send.Interface().(T)
}
