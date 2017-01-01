// Tideland Wozzot - Services - Authentication
//
// Copyright (C) 2016 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package services

//--------------------
// AUTHENTiCATION
//--------------------

// Authentication provides functionality for login and
// user management.
type Authentication interface {
	// Authenticate authenticates a user by his password.
	Authenticate(userID, password string) error
}

// EOF
