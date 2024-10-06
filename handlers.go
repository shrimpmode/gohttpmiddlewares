package main

import "net/http"

type RouteHandler func(http.ResponseWriter, *http.Request, *Context)

type Context struct {
	User *User
}

func (ctx *Context) SetUser(user *User) {
	ctx.User = user
}

func NewContext() *Context {
	return &Context{
		User: &User{},
	}
}

type MiddlewareHandler struct {
	Context *Context
	Handler http.Handler
}

func (f *MiddlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := &AuthMiddleware{next: f}
	// Apply more middlewares here
	m.ServeHTTP(w, r)
}

func NewRouteHandler(handler RouteHandler) http.Handler {
	ctx := NewContext()
	h := func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, ctx)
	}
	return &MiddlewareHandler{
		Context: ctx,
		Handler: http.HandlerFunc(h),
	}
}
