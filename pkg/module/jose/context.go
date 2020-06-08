package jose

import (
	"context"
)

type key string

const authKey key = "auth"

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext returns the Store associated with this context.
func FromContext(c context.Context) *UserToken {
	user := c.Value(string(authKey))
	if user == nil {
		user = c.Value(authKey)
		if user == nil {
			return nil
		}
	}
	return user.(*UserToken)
}

// ToContext adds the Store to this context if it supports the Setter interface.
func ToContext(c Setter, user *UserToken) {
	c.Set(string(authKey), user)
}

// SetContext add user to context
func SetContext(c context.Context, user *UserToken) context.Context {
	return context.WithValue(c, authKey, user)
}
