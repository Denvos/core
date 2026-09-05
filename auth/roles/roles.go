package roles

import (
	"sync"
)

type Role struct {
	Name        string
	Permissions []string
	Description string
}

type Registry struct {
	mu    sync.RWMutex
	roles map[string]Role
}

func NewRegistry() *Registry {
	return &Registry{roles: make(map[string]Role)}
}

func (r *Registry) Register(role Role) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.roles[role.Name] = role
}

func (r *Registry) Get(name string) (Role, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	role, ok := r.roles[name]
	return role, ok
}

func (r *Registry) HasPermission(roleName, permission string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	role, ok := r.roles[roleName]
	if !ok {
		return false
	}
	for _, p := range role.Permissions {
		if p == permission {
			return true
		}
	}
	return false
}
