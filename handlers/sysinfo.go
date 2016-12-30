// Tideland Wozzot - Handlers - System Information
//
// Copyright (C) 2016 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package handlers

//--------------------
// IMPORTS
//--------------------

import (
	"context"

	"github.com/tideland/gorest/rest"

	"github.com/tideland/wozzot/core"
)

//--------------------
// SYSTEM INFORMATION HANDLER
//--------------------

// sysinfo contains the returned system information.
type sysinfo struct {
	Manufacturer string
	Software     string
	Version      string
}

// sysinfoHandler returns information about the system.
type sysinfoHandler struct {
}

// NewSysInfoHandler creates a new system information handler.
func NewSysInfoHandler(ctx context.Context) rest.ResourceHandler {
	return &sysinfoHandler{}
}

// ID is specified on the ResourceHandler interface.
func (h *sysinfoHandler) ID() string {
	return "system information"
}

// Init is specified on the ResourceHandler interface.
func (h *sysinfoHandler) Init(env rest.Environment, domain, resource string) error {
	return nil
}

// Get is specified on the GetResourceHandler interface.
func (h *sysinfoHandler) Get(job rest.Job) (bool, error) {
	info := &sysinfo{
		Manufacturer: "Tideland",
		Software:     "Wozzot",
		Version:      core.Version().String(),
	}
	job.JSON(true).Write(rest.StatusOK, info)
	return true, nil
}

// EOF
