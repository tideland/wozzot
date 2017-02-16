// Tideland Wozzot - Services - Errors
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
	"github.com/tideland/golib/errors"
)

//--------------------
// CONSTANTS
//--------------------

// Error codes.
const (
	ErrDoesNotExist = iota + 1
	ErrCannotLoad
)

// Error messages.
var errorMessages = errors.Messages{
	ErrDoesNotExist: "document '%s' does not exist",
	ErrCannotLoad:   "cannot load document: %v",
}

// EOF
