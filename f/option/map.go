package result

import "github.com/samber/mo"

func Map[T any, U any](m mo.Option[T], f func(T) (U, bool)) mo.Option[U] {
	if m.IsPresent() {
		return mo.TupleToOption(f(m.MustGet()))
	}
	return mo.None[U]()
}
