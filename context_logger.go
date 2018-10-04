// Copyright 2018 Goomba Authors. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package log

import (
	"fmt"
	"io"
)

// ContextLogger represents a log.Logger that encodes keyvals to a io.Writer in
// logfmt format in tree parts: prefixes, arguments and suffixes.
type ContextLogger struct {
	output io.Writer

	prefixes []interface{}
	suffixes []interface{}
}

// NewContextLogger returns a new ContextLogger type instance
func NewContextLogger(w io.Writer) Logger {
	contextLogger := &ContextLogger{
		output: w,

		prefixes: make([]interface{}, 0),
		suffixes: make([]interface{}, 0),
	}

	return contextLogger
}

// AddPrefix adds a prefix to the logger prefixes set.
func (l ContextLogger) AddPrefix(prefix interface{}) error {

	l.prefixes = append(l.prefixes, prefix)

	return nil
}

// AddSuffix adds a prefix to the logger prefixes set.
func (l ContextLogger) AddSuffix(suffix interface{}) error {

	l.suffixes = append(l.suffixes, suffix)

	return nil
}

// Log encodes encodes keyvals to a io.Writer in logfmt format.
func (l ContextLogger) Log(keyvals ...interface{}) error {
	// Prefixes
	for _, prefix := range l.prefixes {
		_, err := fmt.Fprintf(l.output, "%v ", prefix)
		if err != nil {
			return err
		}
	}

	// Arguments
	for _, keyval := range keyvals {
		_, err := fmt.Fprintf(l.output, "%v ", keyval)
		if err != nil {
			return err
		}
	}

	// Suffixes
	for _, suffix := range l.suffixes {
		_, err := fmt.Fprintf(l.output, "%v ", suffix)
		if err != nil {
			return err
		}
	}

	_, err := fmt.Fprintf(l.output, "\n")

	return err
}
