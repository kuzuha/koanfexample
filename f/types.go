package f

import "github.com/samber/mo"

type (
	OptionalOrResult[T any] interface {
		mo.Option[T] | mo.Result[T]
	}
)
