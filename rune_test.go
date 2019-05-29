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

func TestIsNumOrLetter(t*testing.T){
	if !Rune.IsNumOrLetter(rune('0')) {
		t.Fail()
	}
	if Rune.IsNumOrLetter(rune('@')) {
		t.Fail()
	}
}

func TestIsLetter(t*testing.T){
	if !Rune.IsLetter(rune('a')) {
		t.Fail()
	}
	if Rune.IsLetter(rune('0')) {
		t.Fail()
	}
}
