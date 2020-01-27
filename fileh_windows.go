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
