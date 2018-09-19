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

/*
Package log implements interfaces and objects for logging and it provides a
structured logger.

Package log provides a structured logger.

Structured logging produces logs easily consumed by humans or machines. Humans
might be interested in debugging errors, or tracing specific requests. Machines
might be interested in counting interesting events, or aggregating information
for off-line processing. Package log is designed to encourage both of these
best practices.

Basic Usage

The fundamental interface is Logger. Loggers create log events from key/value
data. The Logger interface has a single method, Log, which accepts a sequence
of alternating key/value pairs.

	type Logger interface {
		Log(keyvals ...interface{}) error
	}

Example

This is an example using log.FmtLogger which prints all given keyvalues:

	package main

	import (
		"github.com/goombaio/log"
	)

	func main() {
		output := os.Stdout
		logger := log.NewFmtLogger(output)

		err := logger.Log("foo", "bar", "1", true, nil)
		if err != nil {
			panic(err)
		}
		// Output:
		// foo bar 1 true <nil>
	}
*/
package log
