package repository

import "context"

// Key is the key name of the store in the Gin context.
const Key = "store"

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext returns the Store associated with this context.
func FromContext(c context.Context) Repository {
	return c.Value(Key).(Repository)
}

// ToContext adds the Store to this context if it supports
// the Setter interface.
func ToContext(c Setter, repo Repository) {
	c.Set(Key, repo)
}
