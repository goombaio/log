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
	"io/ioutil"
	"testing"

	"github.com/goombaio/log"
)

func TestContextLogger(t *testing.T) {
	output := ioutil.Discard
	_ = log.NewContextLogger(output)
}

func TestContextLogger_Log(t *testing.T) {
	output := ioutil.Discard
	logger := log.NewContextLogger(output)

	err := logger.Log("foo")
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	err = logger.Log("foo", 1)
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}

func TestContextLogger_AddPrefix(t *testing.T) {
	output := ioutil.Discard
	logger := log.NewContextLogger(output)

	logger.AddPrefix("prefix1")

	err := logger.Log("foo")
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	err = logger.Log("foo", 1)
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}

func TestContextLogger_AddSuffix(t *testing.T) {
	output := ioutil.Discard
	logger := log.NewContextLogger(output)

	logger.AddSuffix("suffix1")

	err := logger.Log("foo")
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	err = logger.Log("foo", 1)
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}
