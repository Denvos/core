package policies

import (
	"sync"
)

type Policy struct {
	ID          string
	Description string
	Effect      string // allow or deny
	Resources   []string
	Actions     []string
	Conditions  map[string]interface{}
}

type Engine struct {
	mu       sync.RWMutex
	policies []Policy
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) Add(policy Policy) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.policies = append(e.policies, policy)
}

func (e *Engine) Evaluate(resource, action string, context map[string]interface{}) bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	for _, p := range e.policies {
		if !matchResource(p, resource) {
			continue
		}
		if !matchAction(p, action) {
			continue
		}
		if !matchConditions(p, context) {
			continue
		}
		return p.Effect == "allow"
	}
	return false
}

func matchResource(p Policy, resource string) bool {
	if len(p.Resources) == 0 {
		return true
	}
	for _, r := range p.Resources {
		if r == resource || r == "*" {
			return true
		}
	}
	return false
}

func matchAction(p Policy, action string) bool {
	if len(p.Actions) == 0 {
		return true
	}
	for _, a := range p.Actions {
		if a == action || a == "*" {
			return true
		}
	}
	return false
}

func matchConditions(p Policy, context map[string]interface{}) bool {
	if len(p.Conditions) == 0 {
		return true
	}
	for k, v := range p.Conditions {
		if val, ok := context[k]; !ok || val != v {
			return false
		}
	}
	return true
}
