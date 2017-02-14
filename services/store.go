// Tideland Wozzot - Services - Store
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
	"github.com/tideland/wozzot/model"
)

//--------------------
// LOADER
//--------------------

// Loader provides access to the document base.
type Loader interface {
	// ReadDocument returns the document with the given ID.
	ReadDocument(id string) (model.Document, error)
}

// EOF
