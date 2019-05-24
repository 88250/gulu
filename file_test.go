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

import (
	"path/filepath"
	"strconv"
	"testing"
)

func TestGetFileSize(t *testing.T) {
	size := File.GetFileSize(".")

	t.Log("size of file [.] is [" + strconv.FormatInt(size, 10) + "]")
}

func TestIsExist(t *testing.T) {
	if !File.IsExist(".") {
		t.Error(". must exist")

		return
	}
}

func TestIdBinary(t *testing.T) {
	if File.IsBinary("not binary content") {
		t.Error("The content should not be binary")

		return
	}
}

func TestIsImg(t *testing.T) {
	if !File.IsImg(".jpg") {
		t.Error(".jpg should be a valid extension of a image file")

		return
	}
}

func TestIsDir(t *testing.T) {
	if !File.IsDir(".") {
		t.Error(". should be a directory")

		return
	}
}

func TestCopyDir(t *testing.T) {
	dest := filepath.Join(testDir, "util")

	err := File.CopyDir(".", dest)
	if nil != err {
		t.Error("Copy dir error: ", err)

		return
	}
}

func TestCopyFile(t *testing.T) {
	dest := filepath.Join(testDir, "file.go")

	err := File.CopyFile("./file.go", dest)
	if nil != err {
		t.Error("Copy file error: ", err)

		return
	}
}
