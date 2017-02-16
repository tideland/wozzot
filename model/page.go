// Tideland Wozzot - Model - Page
//
// Copyright (C) 2016-2017 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package model

//--------------------
// PAGE
//--------------------

// Page describes a rendered page.
type Page struct {
	Header  map[string]string
	Content []byte
}

// EOF
