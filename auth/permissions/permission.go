package permissions

import (
	"sync"
)

type Permission struct {
	Resource string
	Action   string
}

type Registry struct {
	mu          sync.RWMutex
	permissions map[string]Permission
}

func NewRegistry() *Registry {
	return &Registry{
		permissions: make(map[string]Permission),
	}
}

func (r *Registry) Register(resource, action string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	key := resource + ":" + action
	r.permissions[key] = Permission{Resource: resource, Action: action}
}

func (r *Registry) Has(resource, action string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	key := resource + ":" + action
	_, ok := r.permissions[key]
	return ok
}

func (r *Registry) List() []Permission {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var perms []Permission
	for _, p := range r.permissions {
		perms = append(perms, p)
	}
	return perms
}
