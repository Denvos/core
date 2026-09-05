package scopes

import (
	"fmt"
	"strings"
)

type Scope struct {
	Name        string
	Description string
	Permissions []string
}

type Registry struct {
	scopes map[string]Scope
}

func NewRegistry() *Registry {
	return &Registry{scopes: make(map[string]Scope)}
}

func (r *Registry) Register(scope Scope) {
	r.scopes[scope.Name] = scope
}

func (r *Registry) Get(name string) (Scope, bool) {
	s, ok := r.scopes[name]
	return s, ok
}

func (r *Registry) Validate(scopes []string) bool {
	for _, s := range scopes {
		if _, ok := r.scopes[s]; !ok {
			return false
		}
	}
	return true
}

func Parse(scopeStr string) []string {
	if scopeStr == "" {
		return []string{}
	}
	return strings.Split(scopeStr, " ")
}

func Join(scopes []string) string {
	return strings.Join(scopes, " ")
}
