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

import (
	"strings"
	"testing"
)

func TestGetAPIPath(t *testing.T) {
	apiPath := Go.GetAPIPath()

	if !strings.HasSuffix(apiPath, "src") {
		t.Error("api path should end with \"src\"")

		return
	}
}

func TestIsAPI(t *testing.T) {
	apiPath := Go.GetAPIPath()

	if !Go.IsAPI(apiPath) {
		t.Error("api path root should belong to api path")

		return
	}

	root := "/root"

	if Go.IsAPI(root) {
		t.Error("root should not belong to api path")

		return
	}
}

func TestGetGoFormats(t *testing.T) {
	formats := Go.GetGoFormats()

	if len(formats) < 1 {
		t.Error("should have one go format tool [gofmt] at least")
	}
}

func TestGetExecutableInGOBIN(t *testing.T) {
	bin := Go.GetExecutableInGOBIN("test")

	if OS.IsWindows() {
		if !strings.HasSuffix(bin, ".exe") {
			t.Error("Executable binary should end with .exe")

			return
		}
	}
}
