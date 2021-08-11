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
	"runtime"
	"testing"
)

func TestIsWindows(t *testing.T) {
	goos := runtime.GOOS

	if "windows" == goos && !OS.IsWindows() {
		t.Error("runtime.GOOS returns [windows]")

		return
	}
}

func TestIsLinux(t *testing.T) {
	goos := runtime.GOOS

	if "linux" == goos && !OS.IsLinux() {
		t.Error("runtime.GOOS returns [linux]")

		return
	}
}

func TestIsDarwin(t *testing.T) {
	goos := runtime.GOOS

	if "darwin" == goos && !OS.IsDarwin() {
		t.Error("runtime.GOOS returns [darwin]")

		return
	}
}

func TestPwd(t *testing.T) {
	if "" == OS.Pwd() {
		t.Error("Working directory should not be empty")

		return
	}
}

func TestHome(t *testing.T) {
	home, err := OS.Home()
	if nil != err {
		t.Error("Can not get user home")

		return
	}

	t.Log(home)
}
