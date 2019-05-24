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
	logger.Info(testdataDir)
	retCode := m.Run()

	// clean test data
	os.RemoveAll(packageName + ".zip")
	os.RemoveAll(packageName)

	os.Exit(retCode)
}
