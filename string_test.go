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

func TestContains(t *testing.T) {
	if !Str.Contains("123", []string{"123", "345"}) {
		t.Error("[\"123\", \"345\"] should contain \"123\"")

		return
	}
}

func TestLCS(t *testing.T) {
	str := Str.LCS("123456", "abc34def")

	if "34" != str {
		t.Error("[\"123456\"] and [\"abc34def\"] should have the longest common substring [\"34\"]")

		return
	}
}
