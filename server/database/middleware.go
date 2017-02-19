package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Context Middleware
type contextTimeout struct {
	duration time.Duration
	parent   context.Context
}

func (ct *contextTimeout) ContextTimeout(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(ct.parent, ct.duration)
		defer cancel()
		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

// -----------------------

// Logging Middleware
type routeLogger struct {
	logger io.Writer
}

func (rl *routeLogger) RouteLog(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		h.ServeHTTP(w, r)
		if rl.logger != nil {
			rl.logger.Write([]byte(fmt.Sprintf("[%s] %q %v\n", r.Method, r.URL.String(), time.Since(t))))
		}

	}
	return http.HandlerFunc(fn)
}

func admincheck(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// logic for checking if admin from the request.
		// fetch token

		// find token in database
		// if found approve action
		// if not a admin return.
		// if is admin continue request
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
