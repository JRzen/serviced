// Copyright 2014 The Serviced Authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build integration,!quick

package zzk_test

import (
	"testing"

	. "github.com/control-center/serviced/zzk"
	. "gopkg.in/check.v1"
)

var _ = Suite(&ZZKTest{})

type ZZKTest struct {
	ZZKTestSuite
}

func Test(t *testing.T) {
	TestingT(t)
}
