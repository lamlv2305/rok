package channel

type Middleware[T any] func(T) bool

func WithMiddleware[T any](middlewares ...Middleware[T]) func(*options[T]) {
	return func(o *options[T]) {
		for idx := range middlewares {
			o.middlewares = append(o.middlewares, middlewares[idx])
		}
	}
}