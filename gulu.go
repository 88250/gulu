// Gulu - Golang common utilities for everyone.
// Copyright (c) 2019-present, b3log.org
//
// LianDi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

// Package gulu implements some common utilities.
package gulu

import "os"

// Version is the version of Gulu.
const Version = "v1.0.0"

// Logger is the logger used in Gulu internally.
var logger = Log.NewLogger(os.Stdout)

type (
	// GuluFile is the receiver of file utilities
	GuluFile byte
	// GuluGo is the receiver of Go utilities
	GuluGo byte
	// GuluNet is the receiver of network utilities
	GuluNet byte
	// GuluOS is the receiver of OS utilities
	GuluOS byte
	// GuluPanic is the receiver of panic utilities
	GuluPanic byte
	// GuluRand is the receiver of random utilities
	GuluRand byte
	// GuluRet is the receiver of result utilities
	GuluRet byte
	// GuluRune is the receiver of rune utilities
	GuluRune byte
	// GuluStr is the receiver of string utilities
	GuluStr byte
	// GuluZip is the receiver of zip utilities
	GuluZip byte
)

var (
	// File utilities
	File GuluFile
	// Go utilities
	Go GuluGo
	// Net utilities
	Net GuluNet
	// OS utilities
	OS GuluOS
	// Panic utilities
	Panic GuluPanic
	// Rand utilities
	Rand GuluRand
	// Ret utilities
	Ret GuluRet
	// Rune utilities
	Rune GuluRune
	// Str utilities
	Str GuluStr
	// Zip utilities
	Zip GuluZip
)
