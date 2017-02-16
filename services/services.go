// Tideland Wozzot - Services
//
// Copyright (C) 2016-2017 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package services

//--------------------
// IMPORTS
//--------------------

import (
	"context"

	"github.com/tideland/golib/version"
)

//--------------------
// SERVICE
//--------------------

// Service describes the interface each service has
// ro implement.
type Service interface {
	// Init tells a service to startup providing needed
	// information in a context.
	Init(ctx context.Context) error

	// Info returns the service name and it's version.
	Info() (string, version.Version)

	// Stop terminates the service.
	Stop() error
}

//--------------------
// PROVIDER
//--------------------

// Provider describes a type managing the different
// Wozzot services.
type Provider interface {
	// Loader returns the loader service.
	Loader() Loader

	// Fetcher returns the fetcher service.
	Fetcher() Fetcher

	// Renderer returns the renderer service.
	Renderer() Renderer
}

//--------------------
//  CONTEXT
//--------------------

// contextKey is used to address data inside a context.
type contextKey int

// providerKey is the context key for the service provider.
const providerKey contextKey = 0

// NewContext creates a new context containing a
// service provider.
func NewContext(ctx context.Context, provider Provider) context.Context {
	return context.WithValue(ctx, providerKey, provider)
}

// FromContext retrieves the provider out of a context.
func FromContext(ctx context.Context) (Provider, bool) {
	return nil, false
}

// EOF
