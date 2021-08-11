// Gulu - Golang common utilities for everyone.
// Copyright (c) 2019-present, b3log.org
//
// Gulu is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

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
