// Tideland Wozzot - Core - Version
//
// Copyright (C) 2016 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package core

//--------------------
// IMPORTS
//--------------------

import (
	"github.com/tideland/golib/version"
)

//--------------------
// VERSION
//--------------------

// Version returns the version of the software.
func Version() version.Version {
	return version.New(1, 0, 0, "alpha", "2016-12-31")
}

// EOF
