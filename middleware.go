package main

import (
	"fmt"
	"net/http"
)

type AuthMiddleware struct {
	next *MiddlewareHandler
}

func (m *AuthMiddleware) GetAuthUser() *User {
	return &User{
		ID:   1,
		Name: "Giorno Giovanna",
	}
}

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := m.GetAuthUser()
	if user == nil {
		fmt.Println("Unauthorized User")
	}
	m.next.Context.SetUser(user)
	m.next.Handler.ServeHTTP(w, r)
}
