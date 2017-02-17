// Tideland Wozzot - Core - Version
//
// Copyright (C) 2016-2017 Frank Mueller / Tideland / Oldenburg / Germany
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
	return version.New(0, 5, 0)
}

// EOF
