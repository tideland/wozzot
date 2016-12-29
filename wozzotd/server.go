// Tideland Wozzot - Daemon - Server
//
// Copyright (C) 2016 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package main

//--------------------
// IMPORTS
//--------------------

import (
	"context"
	"net/http"

	"github.com/tideland/golib/etc"
	"github.com/tideland/golib/logger"
)

//--------------------
// SERVER
//--------------------

// ServerFunc is a function starting one or more HTTP and/or
// WS server with the passed handler.
type ServerFunc func(ctx context.Context, h http.Handler) error

// StandardServer is the standard implementation for running
// the server. Currently it only provides standard HTTP, later
// also TLS, HTTP/2, and potentially WS.
func StandardServer(ctx context.Context, h http.Handler) error {
	// Retrieve configuration.
	httpAddr := ":12345"
	cfg, ok := etc.FromContext(ctx)
	if ok {
		httpAddr = cfg.ValueAsString("general/server/http-address", httpAddr)
	}

	// Start servers.
	errc := make(chan error)

	go func() {
		// Standard HTTP daemon.
		logger.Infof("Starting the HTTP daemon on %q ...", httpAddr)
		err := http.ListenAndServe(httpAddr, h)
		errc <- err
	}()

	return <-errc
}

// EOF
