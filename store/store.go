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
	"context"
)

//--------------------
// STORE
//--------------------

// Store describes the persistency methods of Wozzot
// used by the different handlers. 
type Store interface {
	// ReadDocument returns the document with the given ID.
	ReadDocument(id string) (Document, error)
}

// EOF