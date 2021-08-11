// Gulu - Golang common utilities for everyone.
// Copyright (c) 2019-present, b3log.org
//
// Gulu is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

// +build windows

package gulu

import (
	"syscall"
)

// IsHidden checks whether the file specified by the given path is hidden.
func (*GuluFile) IsHidden(path string) bool {
	pointer, err := syscall.UTF16PtrFromString(path)
	if nil != err {
		logger.Errorf("Checks file [%s] is hidden failed: [%s]", path, err)

		return false
	}
	attributes, err := syscall.GetFileAttributes(pointer)
	if nil != err {
		logger.Errorf("Checks file [%s] is hidden failed: [%s]", path, err)

		return false
	}

	return 0 != attributes&syscall.FILE_ATTRIBUTE_HIDDEN
}
