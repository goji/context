Context
=======

[![GoDoc](https://godoc.org/github.com/goji/context?status.svg)](https://godoc.org/github.com/goji/context)

Context provides a canonical way to use `go.net`'s [`context`][context] package
with [Goji][goji]. It provides two-way bindings between `context.Context`
objects and Goji's `web.C`, giving you a convenient way to look up Goji `Env`
variables from the `context.Context` and allowing you to freely convert one
request context to the other.

**Warning!** *The two-way binding implies that the `context.Context` is bound to the lifetime of Goji's `web.C`, which is scoped to the lifetime of the web request. You will encounter data races if you refence the `context.Context` in a goroutine which outlives the web request.*

[goji]: https://github.com/zenazn/goji
[context]: http://godoc.org/code.google.com/p/go.net/context
