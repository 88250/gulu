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

// Package gulu implements some common utilities.
package gulu

import "os"

// Version is the version of Gulu.
const Version = "v1.0.0"

// Logger is the logger used in Gulu internally.
var logger = Log.NewLogger(os.Stdout)

type (
	GuluFile  byte // GuluFile is the receiver of file utilities
	GuluGo    byte // GuluGo is the receiver of Go utilities
	GuluNet   byte // GuluNet is the receiver of network utilities
	GuluOS    byte // GuluOS is the receiver of OS utilities
	GuluPanic byte // GuluPanic is the receiver of panic utilities
	GuluRand  byte // GuluRand is the receiver of random utilities
	GuluRet   byte // GuluRet is the receiver of result utilities
	GuluRune  byte // GuluRune is the receiver of rune utilities
	GuluStr   byte // GuluStr is the receiver of string utilities
	GuluZip   byte // GuluZip is the receiver of zip utilities
)

var (
	File  GuluFile  // File utilities
	Go    GuluGo    // Go utilities
	Net   GuluNet   // Network utilities
	OS    GuluOS    // OS utilities
	Panic GuluPanic // Panic utilities
	Rand  GuluRand  // Random utilities
	Ret   GuluRet   // Ret utilities
	Rune  GuluRune  // Rune utilities
	Str   GuluStr   // String utilities
	Zip   GuluZip   // Zip utilities
)
