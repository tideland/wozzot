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
	"fmt"

	"github.com/tideland/golib/errors"
	"github.com/tideland/golib/etc"
	"github.com/tideland/golib/version"
	"github.com/tideland/wozzot/model"
)

//--------------------
// LOADER
//--------------------

// Loader provides access to the document base.
type Loader interface {
	Service

	// ReadDocument returns the document with the given ID.
	ReadDocument(id string) (model.Document, error)
}

//--------------------
// STUB LOADER
//--------------------

// stubLoader simulates a loadeer for tests.
type stubLoader struct {
	cfg      etc.Etc
	provider Provider
}

// newStubLoader creates a new stub loader.
func newStubLoader() Loader {
	return &stubLoader{}
}

// Init implements the Service interface.
func (l *stubLoader) Init(cfg etc.Etc, p Provider) error {
	l.cfg = cfg
	l.provider = p
	return nil
}

// Info implements the Service interface.
func (l *stubLoader) Info() (string, version.Version) {
	return "Stub Loader", version.New(0, 4, 0)
}

// Stop implements the Service interface.
func (l *stubLoader) Stop() error {
	return nil
}

// ReadDocument implements the Loader interface.
func (l *stubLoader) ReadDocument(id string) (model.Document, error) {
	switch id {
	case "does-not-exist":
		return model.Document{}, errors.New(ErrDoesNotExist, errorMessages, id)
	case "fail-during-load":
		err := fmt.Errorf("ouch")
		return model.Document{}, errors.New(ErrCannotLoad, errorMessages, err)
	case "markdown-document":
		return model.Document{
			ID:      id,
			Format:  model.FormatMarkdown,
			Tags:    []string{"test", "foo", "bar"},
			Title:   "MArkdown document by stub loader",
			Content: "The quick brown **fox** jumps over the lazy dog.",
		}, nil
	case "sml-document":
		return model.Document{
			ID:      id,
			Format:  model.FormatSML,
			Tags:    []string{"test", "foo", "bar"},
			Title:   "SML document by stub loader",
			Content: "{p The quick brown {b fox} jumps over the lazy dog.}",
		}, nil
	default:
		return model.Document{
			ID:      id,
			Format:  model.FormatText,
			Tags:    []string{"test", "foo", "bar"},
			Title:   "Text document by stub loader",
			Content: "The quick brown fox jumps over the lazy dog.",
		}, nil
	}
}

// EOF
