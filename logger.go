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

// logger represents an implementation of log.Logger.
type logger struct {
	output io.Writer
}

// NewLogger creates a local instance of log.Logger.
func NewLogger(out io.Writer) Logger {
	return &logger{
		output: out,
	}
}

// SetOutput sets the logger output destination.
func (l *logger) SetOutput(out io.Writer) {
	l.output = out
}

// Info logs an information message.
func (l *logger) Info(msg string) {
	_, _ = fmt.Fprintf(l.output, "%s %s\n", "[INFO]", msg)
}

// Error logs an error and a message.
func (l *logger) Error(err error, msg string) {
	_, _ = fmt.Fprintf(l.output, "%s %s, %s\n", "[ERROR]", msg, err)
}
