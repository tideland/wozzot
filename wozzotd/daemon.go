// Tideland Wozzot - Daemon
//
// Copyright (C) 2016-2017 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package main

//--------------------
// IMPORTS
//--------------------

import (
	"context"
	"errors"

	"github.com/tideland/golib/etc"
	"github.com/tideland/golib/logger"
	"github.com/tideland/gorest/rest"
	"github.com/tideland/wozzot/core"

	"github.com/tideland/wozzot/handlers"
)

//--------------------
// DAEMON
//--------------------

// Daemon is a server running all configured services.
// Configured means they only have to have their section
// inside the system configuration, even without any
// settings in it.
type Daemon interface {
	// Run lets the daemon start the servers and all
	// the services.
	Run() error

	// Finalize cleans up when the daemon stops working.
	Finalize() error
}

// daemon implements the Daemon interface.
type daemon struct {
	ctx    context.Context
	cfg    etc.Etc
	mux    rest.Multiplexer
	server ServerFunc
}

// NewDaemon creates a new daemon and prepares the services.
func NewDaemon(ctx context.Context, server ServerFunc) (Daemon, error) {
	logger.Infof("starting Tideland Wozzot daemon %v...", core.Version())
	d := &daemon{
		ctx:    ctx,
		server: server,
	}
	// Init parts of the daemon, order based on dependencies.
	if err := d.initConfiguration(); err != nil {
		return nil, err
	}
	if err := d.initMultiplexer(); err != nil {
		return nil, err
	}
	if err := d.initHandlers(); err != nil {
		return nil, err
	}
	return d, nil
}

// Run implements the Daemon interface.
func (d *daemon) Run() error {
	// Start the server.
	return d.server(d.ctx, d.mux)
}

// Finalize implements the Daemon interface.
func (d *daemon) Finalize() error {
	return nil
}

// initConfiguration retrieves the configuration for the daemon.
func (d *daemon) initConfiguration() error {
	cfg, ok := etc.FromContext(d.ctx)
	if !ok {
		return errors.New("daemon cannot run without configuration")
	}
	d.cfg = cfg
	return nil
}

// initMultiplexer starts a configured multiplexer in
// the given context.
func (d *daemon) initMultiplexer() error {
	cfg, err := d.cfg.Split("general/multiplexer")
	if err != nil {
		return err
	}
	d.mux = rest.NewMultiplexer(d.ctx, cfg)
	return nil
}

// initHandlers starts and registers the different
// different web handlers.
func (d *daemon) initHandlers() error {
	// Helper to stop registering handlers after the
	// first error.
	register := func(err error, domain, resource string, handler rest.ResourceHandler) error {
		if err != nil {
			return err
		}
		return d.mux.Register(domain, resource, handler)
	}
	// Now register the handlers.
	err := register(nil, "system", "informations", handlers.NewInformationsHandler(d.ctx))
	return err
}

// EOF
