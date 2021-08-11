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
	"path/filepath"
	"testing"
)

var tarDirPath = testdataDir + "_tar"
var untarDirPath = testdataDir + "_untar"

func TestTar(t *testing.T) {
	if err := Tar.Tar(testdataDir, tarDirPath); nil != err {
		t.Error(err)
		return
	}
}

func TestUntar(t *testing.T) {
	err := Tar.Untar(filepath.Join(tarDirPath, testdataDir+".tar"), untarDirPath)
	if nil != err {
		t.Error(err)
		return
	}
}
