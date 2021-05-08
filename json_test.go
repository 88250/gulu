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

func TestMarshalJSON(t *testing.T) {
	data, err := JSON.MarshalJSON(map[string]string{"foo": "bar"})
	if nil != err {
		t.Fail()
	}

	if "{\"foo\":\"bar\"}" != (string(data)) {
		t.Fail()
	}
}

func Test(t *testing.T) {
	m := map[string]string{}
	err := JSON.UnmarshalJSON([]byte(`{"foo":"bar"}`), &m)
	if nil != err {
		t.Fail()
	}

	if "bar" != m["foo"] {
		t.Fail()
	}
}

func TestMarshalIndentJSON(t *testing.T) {
	data, err := JSON.MarshalIndentJSON(map[string]string{"foo": "bar"}, "", "  ")
	if nil != err {
		t.Fail()
	}

	if "{\n  \"foo\": \"bar\"\n}" != (string(data)) {
		t.Fail()
	}
}
