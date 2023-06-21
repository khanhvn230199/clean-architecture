package ctx

import (
	"context"
	"net/http"
)

type key string

const (
	// CustomerKey ...
	CustomerKey key = "customer"
	// EmployeeKey ...
	EmployeeKey key = "employee"
	// AdminKey ...
	AdminKey key = "admin"
)

// Set ...
func Set(r *http.Request, k key, v interface{}) *http.Request {
	if v == nil {
		return r
	}

	return r.WithContext(context.WithValue(r.Context(), k, v))
}

// Get ...
func Get(ctx context.Context, key interface{}) interface{} {
	return ctx.Value(key)
}
