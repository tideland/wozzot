// Tideland Wozzot - Main - Daemon - Unit Tests
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
	"testing"

	"github.com/tideland/golib/audit"
	"github.com/tideland/golib/etc"
)

//--------------------
// TESTS
//--------------------

// TestNewDaeomon tests the publinda daemon initialization.
func TestNewDaemon(t *testing.T) {
	assert := audit.NewTestingAssertion(t, true)
	cfgStr := "{etc}"
	ctx := prepareTestContext(assert, cfgStr)
	server := func(sctx context.Context, h http.Handler) error {
		return nil
	}

	d, err := NewDaemon(ctx, server)
	assert.NotNil(d)
	assert.Nil(err)
}

//--------------------
// HELPERS
//--------------------

// prepareTestContext prepares a context for the tests.
func prepareTestContext(assert audit.Assertion, cfgStr string) context.Context {
	cfg, err := etc.ReadString(cfgStr)
	assert.Nil(err)
	ctx := context.Background()
	return etc.NewContext(ctx, cfg)
}

// EOF
