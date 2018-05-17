package httprouter

import (
)

type MiddlewareFunc func(next Handle) Handle

type middleware interface {
	Middleware(handler Handle) Handle
}

// Middleware allows MiddlewareFunc to implement the middleware interface.
func (mw MiddlewareFunc) Middleware(handle Handle) Handle {
	return mw(handle)
}

func (r *Router)Use(mwf ...MiddlewareFunc) {
	for _, fn := range mwf {
		r.middleware = append(r.middleware, fn)
	}
}
