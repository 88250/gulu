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
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

const (
	pathSeparator     = string(os.PathSeparator)     // OS-specific path separator
	pathListSeparator = string(os.PathListSeparator) // OS-specific path list separator
)

// GetAPIPath gets the Go source code path $GOROOT/src.
func (*GuluGo) GetAPIPath() string {
	return filepath.FromSlash(path.Clean(runtime.GOROOT() + "/src"))
}

// IsAPI determines whether the specified path belongs to Go API.
func (*GuluGo) IsAPI(path string) bool {
	apiPath := Go.GetAPIPath()

	return strings.HasPrefix(filepath.FromSlash(path), apiPath)
}

// GetGoFormats gets Go format tools. It may return ["gofmt", "goimports"].
func (*GuluGo) GetGoFormats() []string {
	ret := []string{"gofmt"}

	p := Go.GetExecutableInGOBIN("goimports")
	if File.IsExist(p) {
		ret = append(ret, "goimports")
	}

	sort.Strings(ret)

	return ret
}

// GetExecutableInGOBIN gets executable file under GOBIN path.
//
// The specified executable should not with extension, this function will append .exe if on Windows.
func (*GuluGo) GetExecutableInGOBIN(executable string) string {
	if OS.IsWindows() {
		executable += ".exe"
	}

	gopaths := filepath.SplitList(os.Getenv("GOPATH"))

	for _, gopath := range gopaths {
		// $GOPATH/bin/$GOOS_$GOARCH/executable
		ret := gopath + pathSeparator + "bin" + pathSeparator +
			os.Getenv("GOOS") + "_" + os.Getenv("GOARCH") + pathSeparator + executable
		if File.IsExist(ret) {
			return ret
		}

		// $GOPATH/bin/{runtime.GOOS}_{runtime.GOARCH}/executable
		ret = gopath + pathSeparator + "bin" + pathSeparator +
			runtime.GOOS + "_" + runtime.GOARCH + pathSeparator + executable
		if File.IsExist(ret) {
			return ret
		}

		// $GOPATH/bin/executable
		ret = gopath + pathSeparator + "bin" + pathSeparator + executable
		if File.IsExist(ret) {
			return ret
		}
	}

	// $GOBIN/executable
	gobin := os.Getenv("GOBIN")
	if "" != gobin {
		ret := gobin + pathSeparator + executable
		if File.IsExist(ret) {
			return ret
		}
	}

	return "./" + executable
}
