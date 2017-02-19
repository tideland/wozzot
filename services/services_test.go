// Tideland Wozzot - Services - Unit Tests
//
// Copyright (C) 2016-2017 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package services_test

//--------------------
// IMPORTS
//--------------------

import (
	"context"
	"testing"

	"github.com/tideland/golib/audit"
	"github.com/tideland/golib/etc"

	"github.com/tideland/wozzot/services"
)

//--------------------
// TESTS
//--------------------

// TestProvider tests starting the provider.
func TestNewDaemon(t *testing.T) {
	assert := audit.NewTestingAssertion(t, true)
	cfgStr := "{etc}"
	ctx := prepareTestContext(assert, cfgStr)

	p, err := services.NewProvider(ctx)
	assert.NotNil(p)
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
