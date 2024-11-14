package main

//package genuflect
import (
	"reflect"
	"unsafe"
)

func ptr[T any](v T) *T {
	return &v
}

func valueAny(r []reflect.Value) []any {
	ret := make([]any, len(r))
	for i, v := range r {
		ret[i] = v.Interface()
	}
	return ret
}

type Value[T any] struct {
	value        *T
	reflectValue *reflect.Value
}

type ValuePointer[T any, U *T] struct {
	value        *T
	reflectValue *reflect.Value
}

func ValueOf[T any](value *T) *Value[T] {
	return &Value[T]{value: value, reflectValue: ptr(reflect.ValueOf(value))}
}

func (v *Value[T]) ReflectValue() *reflect.Value {
	return v.reflectValue
}

func (v *Value[T]) Value() T {
	return reflect.Indirect(*v.ReflectValue()).Interface().(T)
}

func Append[T any, U any](s T, x ...U) T {
	arr := make([]reflect.Value, len(x))
	for i, v := range x {
		arr[i] = reflect.ValueOf(v)
	}
	return reflect.Append(reflect.ValueOf(s), arr...).Interface().(T)
}

func (s *Value[T]) Append(x ...any) T {
	arr := make([]reflect.Value, len(x))
	for i, v := range x {
		arr[i] = reflect.ValueOf(v)
	}
	return reflect.Append(*s.ReflectValue(), arr...).Interface().(T)
}

func AppendSlice[T any, U any](t T, u U) T {
	return reflect.AppendSlice(reflect.ValueOf(t), reflect.ValueOf(u)).Interface().(T)
}
func (t *Value[T]) AppendSlice(u any) T {
	return reflect.AppendSlice(*t.ReflectValue(), reflect.ValueOf(u)).Interface().(T)
}

func Indirect[T any](t *T) T {
	return reflect.Indirect(reflect.ValueOf(t)).Interface().(T)
}

/* this is not exactly the same since you can't use a &T as a return type */
func (t *Value[T]) Indirect() T {
	return reflect.Indirect(*t.ReflectValue()).Interface().(T)
}

func Elem[T any](t *T) T {
	return reflect.ValueOf(t).Elem().Interface().(T)
}

/* this is not exactly the same since you can't use a &T as a return type */
func (t *Value[T]) Elem() T {
	return (*t.ReflectValue()).Elem().Interface().(T)
}

func MakeChan[T any](buffer int, typ Type[T]) chan T {
	var t T
	return reflect.MakeChan(reflect.ValueOf(t).Type(), buffer).Interface().(chan T)
}

func (typ *Type[T]) MakeChan(buffer int) chan T {
	var t T
	return reflect.MakeChan(reflect.ValueOf(t).Type(), buffer).Interface().(chan T)
}

/*The MakeFunc function doesn't really make sense when using generics*/
func MakeFunc[T any](fn T) T {
	return fn
}

func MakeMap[T any](m ...Type[T]) T {
	var t T
	return reflect.MakeMap(reflect.ValueOf(t).Type()).Interface().(T)
}

func (typ *Type[T]) MakeMap() T {
	var t T
	return reflect.MakeMap(reflect.ValueOf(t).Type()).Interface().(T)
}

func MakeMapWithSize[T any](n int, m ...Type[T]) T {
	var t T
	return reflect.MakeMapWithSize(reflect.ValueOf(t).Type(), n).Interface().(T)
}

func (typ *Type[T]) MakeMapWithSize(n ...int) T {
	var sz int
	if len(n) > 0 {
		sz = n[0]
	}
	var t T
	return reflect.MakeMapWithSize(reflect.ValueOf(t).Type(), sz).Interface().(T)
}

func MakeSlice[T any](leng int, cap int, typ ...Type[T]) T {
	var t T
	return reflect.MakeSlice(reflect.ValueOf(t).Type(), leng, cap).Interface().(T)
}

func (typ *Type[T]) MakeSlice(a ...int) T {
	var leng int
	var cap int
	if len(a) > 0 {
		leng = a[0]
	}
	if len(a) > 1 {
		cap = a[1]
	}
	var t T
	return reflect.MakeSlice(reflect.ValueOf(t).Type(), leng, cap).Interface().(T)
}

func New[T any](typ ...Type[T]) *T {
	var t T
	return reflect.New(reflect.ValueOf(t).Type()).Interface().(*T)
}

func (typ *Type[T]) New() *T {
	var t T
	return reflect.New(reflect.ValueOf(t).Type()).Interface().(*T)
}
func NewAt[T any](p unsafe.Pointer, typ ...Type[T]) T {
	var t T
	return reflect.NewAt(reflect.ValueOf(t).Type(), p).Interface().(T)
}

func (typ *Type[T]) NewAt(p unsafe.Pointer) T {
	var t T
	return reflect.NewAt(reflect.ValueOf(t).Type(), p).Interface().(T)
}

func CanInt[V any](v V) bool {
	return reflect.ValueOf(v).CanInt()
}

func (v *Value[T]) CanInt() bool {
	return reflect.Indirect(*v.ReflectValue()).CanInt()
}

func Int[V any](v V) int64 {
	return reflect.ValueOf(v).Int()
}

func (v *Value[T]) Int() int64 {
	return v.ReflectValue().Int()
}

func CanConvert[T any, V any](v V, r ...Type[T]) bool {
	if len(r) > 0 {
		return reflect.ValueOf(v).CanConvert(*r[0].ReflectType())
	}
	var t T
	return reflect.ValueOf(v).CanConvert(reflect.ValueOf(t).Type())
}

func Convert[T any, V any](v V, r ...reflect.Type) T {
	if len(r) > 0 {
		return reflect.ValueOf(v).Convert(r[0]).Interface().(T)
	}
	var t T
	return reflect.ValueOf(v).Convert(reflect.ValueOf(t).Type()).Interface().(T)
}

func Zero[T any](typ ...Type[T]) T {
	var t T
	return reflect.Zero(reflect.ValueOf(t).Type()).Interface().(T)
}

func (typ *Type[T]) Zero() T {
	var t T
	return reflect.Zero(reflect.ValueOf(t).Type()).Interface().(T)
}

func Addr[T any](t T) *T {
	return reflect.ValueOf(t).Addr().Interface().(*T)
}

func (t *Value[T]) Attr() *T {
	return (*t.ReflectValue()).Addr().Interface().(*T)
}

func Bool[T any](t T) bool {
	return reflect.ValueOf(t).Bool()
}

func (t *Value[T]) Bool() bool {
	return (*t.ReflectValue()).Bool()
}

func Bytes[T any](t T) []byte {
	return reflect.ValueOf(t).Bytes()
}

func (t *Value[T]) Bytes() []byte {
	return (*t.ReflectValue()).Bytes()
}

func Call[T any, X any](fn T, args ...any) []any {
	var argv []reflect.Value = make([]reflect.Value, len(args))
	for i, a := range args {
		switch v := a.(type) {
		case reflect.Value:
			argv[i] = v
		case Value[X]:
			argv[i] = *v.ReflectValue()
		default:
			argv[i] = reflect.ValueOf(v)
		}
	}
	return valueAny(reflect.ValueOf(fn).Call(argv))
}

/*
	Using generic type T because of generics contraints on methods.

This doesn't give the best outcome but I will investigate if fixable.
For now recommend using the function version over the method.
*/
func (t *Value[T]) Call(args ...any) []any {
	var argv []reflect.Value = make([]reflect.Value, len(args))
	for i, a := range args {
		switch v := a.(type) {
		case reflect.Value:
			argv[i] = v
		case Value[T]:
			argv[i] = *v.ReflectValue()
		default:
			argv[i] = reflect.ValueOf(v)
		}
	}
	return valueAny((*t.ReflectValue()).Call(argv))
}

func CallSlice[T any, X any](fn T, args ...any) []any {
	var argv []reflect.Value = make([]reflect.Value, len(args))
	for i, a := range args {
		switch v := a.(type) {
		case reflect.Value:
			argv[i] = v
		case Value[X]:
			argv[i] = *v.ReflectValue()
		default:
			argv[i] = reflect.ValueOf(v)
		}
	}
	return valueAny(reflect.ValueOf(fn).CallSlice(argv))
}

func (t *Value[T]) CallSlice(args ...any) []any {
	var argv []reflect.Value = make([]reflect.Value, len(args))
	for i, a := range args {
		switch v := a.(type) {
		case reflect.Value:
			argv[i] = v
		case Value[T]:
			argv[i] = *v.ReflectValue()
		default:
			argv[i] = reflect.ValueOf(v)
		}
	}
	res := (*t.ReflectValue()).CallSlice(argv)
	ret := make([]any, len(res))
	for i, v := range res {
		ret[i] = v.Interface()
	}
	return ret
}

func CanAddr[V any](v V) bool {
	return reflect.ValueOf(v).CanAddr()
}

func (v *Value[T]) CanAddr() bool {
	return reflect.Indirect(*v.ReflectValue()).CanAddr()
}

func CanComplex[V any](v V) bool {
	return reflect.ValueOf(v).CanComplex()
}

func (v *Value[T]) CanComplex() bool {
	return reflect.Indirect(*v.ReflectValue()).CanComplex()
}

func CanFloat[V any](v V) bool {
	return reflect.ValueOf(v).CanFloat()
}

func (v *Value[T]) CanFloat() bool {
	return reflect.Indirect(*v.ReflectValue()).CanFloat()
}

func CanInterface[V any](v V) bool {
	return reflect.ValueOf(v).CanInterface()
}

func (v *Value[T]) CanInterface() bool {
	return reflect.Indirect(*v.ReflectValue()).CanInterface()
}

func CanSet[V any](v V) bool {
	return reflect.ValueOf(v).CanSet()
}

func (v *Value[T]) CanSet() bool {
	return reflect.Indirect(*v.ReflectValue()).CanSet()
}

func CanUint[V any](v V) bool {
	return reflect.ValueOf(v).CanUint()
}

func (v *Value[T]) CanUint() bool {
	return reflect.Indirect(*v.ReflectValue()).CanUint()
}

func Cap[V any](v V) int {
	return reflect.ValueOf(v).Cap()
}

func (v *Value[T]) Cap() int {
	return reflect.Indirect(*v.ReflectValue()).Cap()
}

func Clear[V any](v V) {
	reflect.ValueOf(v).Clear()
}

func (v *Value[T]) Clear() {
	reflect.Indirect(*v.ReflectValue()).Clear()
}

func Close[V any](v V) {
	reflect.ValueOf(v).Close()
}

func (v *Value[T]) Close() {
	reflect.Indirect(*v.ReflectValue()).Close()
}

func Comparable[V any](v V) bool {
	return reflect.ValueOf(v).Comparable()
}

func (v *Value[T]) Comparable() bool {
	return reflect.Indirect(*v.ReflectValue()).Comparable()
}

func Complex[V any](v V) complex128 {
	return reflect.ValueOf(v).Complex()
}

func (v *Value[T]) Complex() complex128 {
	return reflect.Indirect(*v.ReflectValue()).Complex()
}

func Equal[V any, T any](v V, t T) bool {
	return reflect.ValueOf(v).Equal(reflect.ValueOf(t))
}

func (v *Value[T]) Equal(t any) bool {
	var a reflect.Value
	switch x := t.(type) {
	case reflect.Value:
		a = x
	case Value[T]:
		a = *x.ReflectValue()
	default:
		a = reflect.ValueOf(x)
	}
	return reflect.Indirect(*v.ReflectValue()).Equal(a)
}

func (v *Value[T]) Field(i int) any {
	return (*v.ReflectValue()).Field(i).Interface()
}

func Field[V any](v V, i int) any {
	return reflect.ValueOf(v).Field(i).Interface()
}

func (v *Value[T]) FieldByIndex(i []int) any {
	return (*v.ReflectValue()).FieldByIndex(i).Interface()
}

func FieldByIndex[V any](v V, i []int) any {
	return reflect.ValueOf(v).FieldByIndex(i).Interface()
}

func (v *Value[T]) FieldByIndexErr(i []int) (any, error) {
	a, err := (*v.ReflectValue()).FieldByIndexErr(i)
	return a.Interface(), err
}

func FieldByIndexErr[V any](v V, i []int) (any, error) {
	a, err := reflect.ValueOf(v).FieldByIndexErr(i)
	return a.Interface(), err
}

func (v *Value[T]) FieldByName(name string) any {
	return (*v.ReflectValue()).FieldByName(name).Interface()
}

func FieldByName[V any](v V, name string) any {
	return reflect.ValueOf(v).FieldByName(name).Interface()
}

func (v *Value[T]) FieldByNameFunc(match func(string) bool) any {
	return (*v.ReflectValue()).FieldByNameFunc(match).Interface()
}

func FieldByNameFunc[V any](v V, match func(string) bool) any {
	return reflect.ValueOf(v).FieldByNameFunc(match).Interface()
}

func Float[V any](v V) float64 {
	return reflect.ValueOf(v).Float()
}

func (v *Value[T]) Float() float64 {
	return reflect.Indirect(*v.ReflectValue()).Float()
}

func Grow[V any](v V, n int) {
	reflect.ValueOf(v).Grow(n)
}

func (v *Value[T]) Grow(n int) {
	reflect.Indirect(*v.ReflectValue()).Grow(n)
}

func Index[V any](v V, i int) any {
	return reflect.ValueOf(v).Index(i).Interface()
}

func (v *Value[T]) Index(i int) any {
	return reflect.Indirect(*v.ReflectValue()).Index(i).Interface()
}

func Interface[V any](v V) any {
	return reflect.ValueOf(v).Interface()
}

func (v *Value[T]) Interface() any {
	return reflect.Indirect(*v.ReflectValue()).Interface()
}

func IsNil[V any](v V) bool {
	return reflect.ValueOf(v).IsNil()
}

func (v *Value[T]) IsNil() bool {
	return reflect.Indirect(*v.ReflectValue()).IsNil()
}

func IsValid[V any](v V) bool {
	return reflect.ValueOf(v).IsValid()
}

func (v *Value[T]) IsValid() bool {
	return reflect.Indirect(*v.ReflectValue()).IsValid()
}

func IsZero[V any](v V) bool {
	return reflect.ValueOf(v).IsZero()
}

func (v *Value[T]) IsZero() bool {
	return reflect.Indirect(*v.ReflectValue()).IsZero()
}

func Kind[V any](v V) reflect.Kind {
	return reflect.ValueOf(v).Kind()
}

func (v *Value[T]) Kind() reflect.Kind {
	return reflect.Indirect(*v.ReflectValue()).Kind()
}

func Len[V any](v V) int {
	return reflect.ValueOf(v).Len()
}

func (v *Value[T]) Len() int {
	return reflect.Indirect(*v.ReflectValue()).Len()
}

func MapIndex[V any](v V, key any) any {
	return reflect.ValueOf(v).MapIndex(reflect.ValueOf(key)).Interface()
}

func (v *Value[T]) MapIndex(key any) any {
	return reflect.Indirect(*v.ReflectValue()).MapIndex(reflect.ValueOf(key)).Interface()
}

func MapKeys[V any](v V) []any {
	return valueAny(reflect.ValueOf(v).MapKeys())
}

func (v *Value[T]) MapKeys() []any {
	return valueAny(reflect.Indirect(*v.ReflectValue()).MapKeys())
}

func MapRange[V any](v V) *reflect.MapIter {
	return reflect.ValueOf(v).MapRange()
}

func (v *Value[T]) MapRange() *reflect.MapIter {
	return reflect.Indirect(*v.ReflectValue()).MapRange()
}


func GetMethod[V any](v V, i int) any {
	return reflect.ValueOf(v).Method(i).Interface()
}

func (v *Value[T]) Method(i int) any {
	return reflect.Indirect(*v.ReflectValue()).Method(i).Interface()
}

func MethodByName[V any](v V, name string) any {
	return reflect.ValueOf(v).MethodByName(name).Interface()
}

func (v *Value[T]) MethodByName(name string) any {
	return reflect.Indirect(*v.ReflectValue()).MethodByName(name).Interface()
}

func NumField[V any](v V) int {
	return reflect.ValueOf(v).NumField()
}

func (v *Value[T]) NumField() int {
	return reflect.Indirect(*v.ReflectValue()).NumField()
}


func NumMethod[V any](v V) int {
	return reflect.ValueOf(v).NumMethod()
}

func (v *Value[T]) NumMethod() int {
	return reflect.Indirect(*v.ReflectValue()).NumMethod()
}

/*
:
func (v Value) OverflowComplex(x complex128) bool
func (v Value) OverflowFloat(x float64) bool
func (v Value) OverflowInt(x int64) bool
func (v Value) OverflowUint(x uint64) bool
func (v Value) Pointer() uintptr
func (v Value) Recv() (x Value, ok bool)
func (v Value) Send(x Value)
func (v Value) Seq() iter.Seq[Value]
func (v Value) Seq2() iter.Seq2[Value, Value]
func (v Value) Set(x Value)
func (v Value) SetBool(x bool)
func (v Value) SetBytes(x []byte)
func (v Value) SetCap(n int)
func (v Value) SetComplex(x complex128)
func (v Value) SetFloat(x float64)
func (v Value) SetInt(x int64)
func (v Value) SetIterKey(iter *MapIter)
func (v Value) SetIterValue(iter *MapIter)
func (v Value) SetLen(n int)
func (v Value) SetMapIndex(key, elem Value)
func (v Value) SetPointer(x unsafe.Pointer)
func (v Value) SetString(x string)
func (v Value) SetUint(x uint64)
func (v Value) SetZero()
func (v Value) Slice(i, j int) Value
func (v Value) Slice3(i, j, k int) Value
func (v Value) String() string
func (v Value) TryRecv() (x Value, ok bool)
func (v Value) TrySend(x Value) bool
func (v Value) Type() Type
func (v Value) Uint() uint64
func (v Value) UnsafeAddr() uintptr
func (v Value) UnsafePointer() unsafe.Pointer
*/
