package result

import "github.com/samber/mo"

func FlatMap[T any, U any](m mo.Option[T], f func(T) mo.Option[U]) mo.Option[U] {
	if m.IsPresent() {
		return f(m.MustGet())
	}
	return mo.None[U]()
}
