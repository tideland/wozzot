// Tideland Wozzot - Daemon - Main
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
	"os"

	"github.com/tideland/golib/etc"
	"github.com/tideland/golib/logger"
)

//--------------------
// CONST
//--------------------

const cfgStr = `{etc
	{general
		{server
			{http-address localhost:12345}
		}
		{multiplexer
			{basepath /}
		}
	}
}`

//--------------------
// DAEMON CONTEXT
//--------------------

// prepareContext initializes the context the daemon
// is running in.
func prepareContext() context.Context {
	// TODO Mue 2016-12-28 Reading from file later.
	cfg, err := etc.ReadString(cfgStr)
	if err != nil {
		logger.Errorf("cannot read configuration: %v", err)
		os.Exit(-1)
	}
	ctx := etc.NewContext(context.Background(), cfg)
	return ctx
}

//--------------------
// MAIN
//--------------------

// main is the main programm (as usual)
func main() {
	ctx := prepareContext()
	d, err := NewDaemon(ctx, StandardServer)
	if err != nil {
		logger.Errorf("cannot create daemon: %v", err)
		os.Exit(-1)
	}
	defer d.Finalize()
	if err = d.Run(); err != nil {
		logger.Fatalf("fatal error during runtime: %v", err)
	}
}

// EOF
