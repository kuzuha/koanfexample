package result

import "github.com/samber/mo"

func Map[T any, U any](m mo.Result[T], f func(T) (U, error)) mo.Result[U] {
	if m.IsOk() {
		var err error
		res, err := f(m.MustGet())
		if err != nil {
			return mo.Err[U](err)
		}
		return mo.Ok(res)
	}
	return mo.Err[U](m.Error())
}
