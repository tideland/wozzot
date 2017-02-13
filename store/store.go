// Tideland Wozzot - Store
//
// Copyright (C) 2016-2017 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package store

//--------------------
// IMPORTS
//--------------------

import (
	"github.com/tideland/wozzot/model"
)

//--------------------
// STORE
//--------------------

// Store describes the access to the document base.
type Store interface {
	// ReadDocument returns the document with the given ID.
	ReadDocument(id string) (model.Document, error)
}

// EOF
