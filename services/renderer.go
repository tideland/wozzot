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
// RENDERER
//--------------------

// Renderer provides rendering of a document with a template.
type Renderer interface {
	Service

	// Render renders a document with an internally defined
	// template, otherwise with a default template.
	Render(doc model.Document) (model.Page, error)

	// RenderTemplate renders a document with a given template,
	// regardles of an internally defined template.
	RenderTemplate(doc model.Document, tmpl model.Template) (model.Page, error)
}

// EOF
