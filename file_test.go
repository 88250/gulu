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
	"path/filepath"
	"strconv"
	"testing"
)

func TestIsHidden(t *testing.T) {
	filename := "./file.go"
	isHidden := File.IsHidden(filename)
	if isHidden {
		t.Error("file [" + filename + "] is not hidden")
	}
}

func TestGetFileSize(t *testing.T) {
	filename := "./file.go"
	size := File.GetFileSize(filename)
	if 0 > size {
		t.Error("file [" + filename + "] size is [" + strconv.FormatInt(size, 10) + "]")
	}
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
	testcopydir := "testcopydir"
	os.Mkdir(testcopydir, 0644)
	dest := filepath.Join(testdataDir, testcopydir)
	defer os.Remove(dest)

	err := File.CopyDir(testcopydir, dest)
	if nil != err {
		t.Error("Copy dir failed: ", err)

		return
	}
}

func TestCopyFile(t *testing.T) {
	dest := filepath.Join(testdataDir, "file.go")
	defer os.Remove(dest)
	err := File.CopyFile("./file.go", dest)
	if nil != err {
		t.Error("Copy file failed: ", err)

		return
	}
}
