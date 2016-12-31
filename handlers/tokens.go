// Tideland Wozzot - Handlers - Tokins
//
// Copyright (C) 2016 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of t

package handlers

//--------------------
// IMPORTS
//--------------------

import (
	"time"

	"github.com/tideland/golib/logger"
	"github.com/tideland/gorest/jwt"
	"github.com/tideland/gorest/rest"
)

//--------------------
// TOKENS HANDLER
//--------------------

// tokensHandler handles the user authentication and creation
// of JWT.
type tokensHandler struct{}

// NewTokensHandler creates a new handler responsible
// for the verification of logins and creation of JSON
// Web Tokens used during a session.
func NewTokensHandler() rest.ResourceHandler {
	return &tokensHandler{}
}

// ID is specified on the ResourceHandler interface.
func (h *tokensHandler) ID() string {
	return "tokens"
}

// Init is specified on the ResourceHandler interface.
func (h *tokensHandler) Init(env rest.Environment, domain, resource string) error {
	return nil
}

// Get is specified on the GetResourceHandler interface. It creates a token for
// the sent user ID and password.
func (h *tokensHandler) Get(job rest.Job) (bool, error) {
	// Retrieve data for token creation.
	userID := job.Query().ValueAsString("userid", "")
	password := job.Query().ValueAsString("password", "")
	// Authenticate the user and create the token.
	token, err := h.createToken(userID, password)
	if err != nil {
		msg := "authentication and token creation failed for user '%s': %v"
		logger.Errorf(msg, userID, err)
		return rest.NegativeFeedback(job.JSON(true), rest.StatusUnauthorized, msg, userID, err)
	}
	// Create response.
	response := token.String()
	err = job.JSON(true).Write(rest.StatusCreated, response)
	if err != nil {
		msg := "failed during response: %v"
		logger.Errorf(msg, err)
		return rest.NegativeFeedback(job.JSON(true), rest.StatusBadRequest, msg, err)
	}
	return true, nil
}

// createToken authenticates the user, retrieves the authorization, checks
// the licensing, and creates the according JWT.
func (h *tokensHandler) createToken(userID, password string) (jwt.JWT, error) {
	// Authenticate user.
	authentication, err := h.authenticate(userID, password)
	if err != nil {
		return nil, err
	}
	// Authorize user.
	authorization, err := h.authorize(userID)
	if err != nil {
		return nil, err
	}
	// Create token.
	claims := jwt.NewClaims()
	claims.SetIssuer("Tideland Wozzot")
	claims.SetIssuedAt(time.Now().UTC())
	claims.SetExpiration(time.Now().Add(60 * time.Minute).UTC())
	claims.SetSubject(userID)
	token, err := jwt.Encode(claims, []byte("secret"), jwt.HS512)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// EOF
