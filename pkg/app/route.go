package app

import (
	"regexp"
	"sync"
)

type router struct {
	mu               sync.RWMutex
	routes           map[string]func() Composer
	routesWithRegexp []regexpRoute
}

func makeRouter() router { _ = "STUB: not implemented"; return *new(router) }

func (r *router) route(path string, newComponent func() Composer) {
	_ = "STUB: not implemented"
	return
}

func (r *router) routeWithRegexp(pattern string, newComponent func() Composer) {
	_ = "STUB: not implemented"
	return
}

func (r *router) routed(path string) bool { _ = "STUB: not implemented"; return false }

func (r *router) createComponent(path string) (Composer, bool) {
	_ = "STUB: not implemented"
	return *new(Composer), false
}

type regexpRoute struct {
	regexp       *regexp.Regexp
	newComponent func() Composer
}
