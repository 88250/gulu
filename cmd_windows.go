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
	"fmt"
	"os/exec"
	"syscall"
	"unicode/utf8"

	"golang.org/x/sys/windows"
	"golang.org/x/text/encoding/ianaindex"
)

func CmdAttr(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}

func DecodeCmdOutput(output []byte) string {
	if !OS.IsWindows() {
		return string(output)
	}

	if utf8.Valid(output) {
		return string(output)
	}

	acp := windows.GetACP()
	encodingName := fmt.Sprintf("CP%d", acp)
	e, err := ianaindex.MIB.Encoding(encodingName)
	if err != nil {
		return string(output)
	}

	decoded, err := e.NewDecoder().Bytes(output)
	if err != nil {
		return string(output)
	}
	return string(decoded)
}
