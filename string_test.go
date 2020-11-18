// Gulu - Golang common utilities for everyone.
// Copyright (c) 2019-present, b3log.org
//
// LianDi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package gulu

import "testing"

func TestToBytes(t *testing.T) {
	str := "Gulu 你好！"
	bytes := Str.ToBytes(str)
	if str2 := Str.FromBytes(bytes); str != str2 {
		t.Errorf("Str Bytes convert failed [str=%s, str2=%s]", str, str2)
	}
}

func TestContains(t *testing.T) {
	if !Str.Contains("123", []string{"123", "345"}) {
		t.Error("[\"123\", \"345\"] should contain \"123\"")
		return
	}
}

func TestReplaceIgnoreCase(t *testing.T) {
	expected := "Foabcdr"
	got := Str.ReplaceIgnoreCase("Foobar", "oBa", "abcd")
	if expected != got {
		t.Errorf("expected [%s], got [%s]", expected, got)
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

func TestSubStr(t *testing.T) {
	expected := "foo测"
	got := Str.SubStr("foo测试bar", 4)
	if expected != got {
		t.Errorf("expected [%s], got [%s]", expected, got)
		return
	}
}
