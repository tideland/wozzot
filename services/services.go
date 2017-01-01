// Tideland Wozzot - Services
//
// Copyright (C) 2016 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package services

//--------------------
// IMPORTS
//--------------------

import (
	"context"
)

//--------------------
// FACTORY
//--------------------

// Factory describes a type managing the different
// Wozzot services.
type Factory interface {
	// Authentication returns the authentication service.
	Authentication() Authentication
}

//--------------------
//  CONTEXT
//--------------------

// contextKey is used to address data inside a context.
type contextKey int

// factoryKey is the context key for the service factory.
const factoryKey contextKey = 0

// newContext creates a new context containing a factory.
func newContext(ctx context.Context, factory Factory) context.Context {
	return context.WithValue(ctx, factoryKey, factory)
}

// FromContext retrieves the factory out of a context.
func FromContext(ctx context.Context) (Factory, bool) {
	return nil, false
}

// EOF
