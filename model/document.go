// Tideland Wozzot - Model - Document
//
// Copyright (C) 2016-2017 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package model

//--------------------
// FORMAT
//--------------------

// Format describes the content format of a document.
type Format int

// Supported document formats.
const (
	FormatText Format = iota + 1
	FormatMarkdown
	FormatSML
)

//--------------------
// DOCUMENT
//--------------------

// Document describes a single document.
type Document struct {
	ID      string
	Format  Format
	Tags    []string
	Title   string
	Content string
}

// EOF
