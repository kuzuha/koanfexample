package result

import "github.com/samber/mo"

func FlatMap[T any, U any](m mo.Result[T], f func(T) mo.Result[U]) mo.Result[U] {
	if m.IsOk() {
		return f(m.MustGet())
	}
	return mo.Err[U](m.Error())
}
