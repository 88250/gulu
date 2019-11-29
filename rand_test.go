// Gulu - Golang common utilities for everyone.
// Copyright (c) 2019-present, b3log.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gulu

import "testing"

func TestInts(t *testing.T) {
	ints := Rand.Ints(10, 19, 20)
	if 9 != len(ints) {
		t.Fail()
	}
	ints = Rand.Ints(10, 19, 5)
	if 5 != len(ints) {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	r1 := Rand.String(16)
	r2 := Rand.String(16)

	if r1 == r2 {
		t.Fail()
	}
}

func TestInt(t *testing.T) {
	r1 := Rand.Int(0, 65535)
	r2 := Rand.Int(0, 65535)

	if r1 == r2 {
		t.Fail()
	}
}
