// Package context provides Goji integration with go.net/context.
package context

import (
	"code.google.com/p/go.net/context"
	"github.com/zenazn/goji/web"
)

type ctx struct {
	c *web.C
	context.Context
}

func (c ctx) Value(key interface{}) interface{} {
	if key == &ckey {
		return c.c
	}
	if s, ok := key.(string); c.c.Env != nil && ok {
		if v, ok := c.c.Env[s]; ok {
			return v
		}
	}
	return c.Context.Value(key)
}
