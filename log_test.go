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

package log_test

import (
	"github.com/goombaio/log"
)

// NilInfoLogger is a log.InfoLogger that does nothing.
type NilInfoLogger struct{}

var _ log.InfoLogger = NilInfoLogger{}

func (_ NilInfoLogger) Info(_ string) {
	// Do nothing.
}

// NilErrorLogger is a log.ErrorLogger that does nothing.
type NilErrorLogger struct{}

var _ log.ErrorLogger = NilErrorLogger{}

func (_ NilErrorLogger) Error(_ error, _ string) {
	// Do nothing.
}

// NilLogger is a log.Logger that does nothing.
type NilLogger struct{}

var _ log.Logger = NilLogger{}

func (_ NilLogger) Info(_ string) {
	// Do nothing.
}

func (_ NilLogger) Error(_ error, _ string) {
	// Do nothing.
}
