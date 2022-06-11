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

import (
	"testing"
)

func TestSubstringsBetween(t *testing.T) {
	got := Str.SubstringsBetween("foo<bar>baz<bar2>", "<", ">")
	if 2 != len(got) {
		t.Errorf("substrings between [%s] should have 2 elements", got)
		return
	}
}

func TestIsASCII(t *testing.T) {
	if !Str.IsASCII("foo") {
		t.Error("[foo] should be ASCII")
		return
	}

	if Str.IsASCII("foo测试") {
		t.Error("[foo测试] should not be ASCII")
		return
	}
}

func TestRemoveInvisible(t *testing.T) {
	expected := "foo测试barbaz"
	got := Str.RemoveInvisible("foo\u200b测试\nbar\tbaz")
	if expected != got {
		t.Errorf("expected [%s], got [%s]", expected, got)
		return
	}
}

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

func TestReplacesIgnoreCase(t *testing.T) {
	expected := "abcdbarefgh"
	got := Str.ReplacesIgnoreCase("Foobarbaz", "foo", "abcd", "baz", "efgh")
	if expected != got {
		t.Errorf("expected [%s], got [%s]", expected, got)
		return
	}

	expected = "bar baz baz"
	got = Str.ReplacesIgnoreCase("foo bar baz", "foo", "bar", "bar", "baz")
	if expected != got {
		t.Errorf("expected [%s], got [%s]", expected, got)
		return
	}

	expected = "bar baz baz"
	got = Str.ReplacesIgnoreCase("foo bar baz", "Bar", "baz", "foo", "bar")
	if expected != got {
		t.Errorf("expected [%s], got [%s]", expected, got)
		return
	}

	expected = "fazz bar barr"
	got = Str.ReplacesIgnoreCase("foo bar baz", "oo", "azz", "az", "arr")
	if expected != got {
		t.Errorf("expected [%s], got [%s]", expected, got)
		return
	}
}

func TestEncloseIgnoreCase(t *testing.T) {
	var expected, got string
	expected = "<mark>Foo</mark>bar<mark>baz</mark>"
	got = Str.EncloseIgnoreCase("Foobarbaz", "<mark>", "</mark>", "foo", "baz")
	if expected != got {
		t.Errorf("expected [%s], got [%s]", expected, got)
		return
	}

	expected = "F<mark>oo</mark><mark>ba</mark>r<mark>ba</mark>z"
	got = Str.EncloseIgnoreCase("Foobarbaz", "<mark>", "</mark>", "Oo", "Ba")
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
