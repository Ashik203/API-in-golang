package middlerware

import (
	"net/http"
)

type MiddleWare func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []MiddleWare
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]MiddleWare, 0),
	}
}

func (m Manager) User(middlerware ...MiddleWare) Manager {
	m.globalMiddlewares = append(m.globalMiddlewares, middlerware...)
	return m

}

func (m *Manager) With(handler http.Handler, middlerware ...MiddleWare) http.Handler {

	var h http.Handler
	h = handler

	for _, m := range middlerware {
		h = m(h)
	}

	for _, m := range m.globalMiddlewares {
		h = m(h)
	}
	return h
}
