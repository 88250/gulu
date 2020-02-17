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
	"testing"
)

var testdataDir = "testdata"
var packageName = filepath.Join(testdataDir, "test_zip")

func TestCreate(t *testing.T) {
	zipFile, err := Zip.Create(packageName + ".zip")
	if nil != err {
		t.Error(err)

		return
	}

	zipFile.AddDirectoryN(".", "testdata")
	if nil != err {
		t.Error(err)

		return
	}

	err = zipFile.Close()
	if nil != err {
		t.Error(err)

		return
	}
}

func TestUnzip(t *testing.T) {
	err := Zip.Unzip(packageName+".zip", packageName)
	if nil != err {
		t.Error(err)

		return
	}
}

func _TestEmptyDir(t *testing.T) {
	dir1 := "/dir/subDir1"
	dir2 := "/dir/subDir2"

	err := os.MkdirAll(packageName+dir1, os.ModeDir)
	if nil != err {
		t.Error(err)

		return
	}

	err = os.MkdirAll(packageName+dir2, os.ModeDir)
	if nil != err {
		t.Error(err)

		return
	}

	f, err := os.Create(packageName + dir2 + "/file")
	if nil != err {
		t.Error(err)

		return
	}
	f.Close()

	zipFile, err := Zip.Create(packageName + "/dir.zip")
	if nil != err {
		t.Error(err)

		return
	}

	zipFile.AddDirectoryN("dir", packageName+"/dir")
	if nil != err {
		t.Error(err)

		return
	}

	err = zipFile.Close()
	if nil != err {
		t.Error(err)

		return
	}

	err = Zip.Unzip(packageName+"/dir.zip", packageName+"/unzipDir")
	if nil != err {
		t.Error(err)

		return
	}

	if !File.IsExist(packageName+"/unzipDir") || !File.IsDir(packageName+"/unzipDir") {
		t.Error("Unzip failed")

		return
	}

	if !File.IsExist(packageName+"/unzipDir"+dir1) || !File.IsDir(packageName+"/unzipDir"+dir1) {
		t.Error("Unzip failed")

		return
	}

	if !File.IsExist(packageName+"/unzipDir"+dir2) || !File.IsDir(packageName+"/unzipDir"+dir2) {
		t.Error("Unzip failed")

		return
	}

	if !File.IsExist(packageName+"/unzipDir"+dir2+"/file") || File.IsDir(packageName+"/unzipDir"+dir2+"/file") {
		t.Error("Unzip failed")

		return
	}
}

func TestMain(m *testing.M) {
	retCode := m.Run()

	// clean test data
	os.RemoveAll(packageName + ".zip")
	os.RemoveAll(packageName)

	os.Exit(retCode)
}
