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

// FmtLogger ...
type FmtLogger struct {
	output io.Writer
}

// NewFmtLogger ...
func NewFmtLogger(w io.Writer) Logger {
	logger := &FmtLogger{
		output: w,
	}

	return logger
}

// Log ...
func (l FmtLogger) Log(keyvals ...interface{}) error {
	for _, keyval := range keyvals {
		_, err := fmt.Fprintf(l.output, "%v ", keyval)
		if err != nil {
			return err
		}
	}
	_, err := fmt.Fprintf(l.output, "\n")

	return err
}
