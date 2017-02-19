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

	"github.com/tideland/golib/errors"
	"github.com/tideland/golib/logger"
	"github.com/tideland/golib/version"

	"github.com/tideland/wozzot/model"
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
	// Fetcher returns the fetcher service.
	Fetcher() (Fetcher, error)

	// Loader returns the loader service.
	Loader() (Loader, error)

	// Renderer returns the renderer service.
	Renderer(format model.Format) (Renderer, error)
}

// provider implements Provider.
type provider struct {
	ctx    context.Context
	loader Loader
}

// NewProvider returns a new provider.
func NewProvider(ctx context.Context) (Provider, error) {
	logger.Infof("starting the service provider ...")
	p := &provider{}
	p.ctx = NewContext(ctx, p)
	return p, nil
}

// Fetcher implements the Provider interface.
func (p *provider) Fetcher() (Fetcher, error) {
	return nil, nil
}

// Loader implements the Provider interface.
func (p *provider) Loader() (Loader, error) {
	if p.loader == nil {
		p.loader = newStubLoader()
		err := p.loader.Init(p.ctx)
		if err != nil {
			p.loader = nil
			return nil, errors.Annotate(err, ErrStartingService, errorMessages)

		}
		info, ver := p.loader.Info()
		logger.Infof("started service %s version %v", info, ver)
	}
	return p.loader, nil
}

// Renderer implements the Provider interface.
func (p *provider) Renderer(format model.Format) (Renderer, error) {
	return nil, nil
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
