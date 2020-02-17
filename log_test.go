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
	"testing"
)

func TestSetLevel(t *testing.T) {
	Log.SetLevel("trace")
}

func TestTrace(t *testing.T) {
	Log.SetLevel("trace")
	logger.Trace("trace")
	logger.SetLevel("off")
	logger.Trace("trace")
}

func TestTracef(t *testing.T) {
	logger.SetLevel("trace")
	logger.Tracef("tracef")
	logger.SetLevel("off")
	logger.Tracef("tracef")
}

func TestDebug(t *testing.T) {
	logger.SetLevel("debug")
	logger.Debug("debug")
	logger.SetLevel("off")
	logger.Debug("debug")
}

func TestDebugf(t *testing.T) {
	logger.SetLevel("debug")
	logger.Debugf("debugf")
	logger.SetLevel("off")
	logger.Debug("debug")
}

func TestInfo(t *testing.T) {
	logger.SetLevel("info")
	logger.Info("info")
	logger.SetLevel("off")
	logger.Info("info")
}

func TestInfof(t *testing.T) {
	logger.SetLevel("info")
	logger.Infof("infof")
	logger.SetLevel("off")
	logger.Infof("infof")
}

func TestWarn(t *testing.T) {
	logger.SetLevel("warn")
	logger.Warn("warn")
	logger.SetLevel("off")
	logger.Warn("warn")
}

func TestWarnf(t *testing.T) {
	logger.SetLevel("warn")
	logger.Warnf("warnf")
	logger.SetLevel("off")
	logger.Warnf("warnf")
}

func TestError(t *testing.T) {
	logger.SetLevel("error")
	logger.Error("error")
	logger.SetLevel("off")
	logger.Error("error")
}

func TestErrorf(t *testing.T) {
	logger.SetLevel("error")
	logger.Errorf("errorf")
	logger.SetLevel("off")
	logger.Errorf("errorf")
}

func TestGetLevel(t *testing.T) {
	if getLevel("trace") != Trace {
		t.FailNow()

		return
	}

	if getLevel("debug") != Debug {
		t.FailNow()

		return
	}

	if getLevel("info") != Info {
		t.FailNow()

		return
	}

	if getLevel("warn") != Warn {
		t.FailNow()

		return
	}

	if getLevel("error") != Error {
		t.FailNow()

		return
	}
}

func TestLoggerSetLevel(t *testing.T) {
	logger.SetLevel("trace")

	if logger.level != Trace {
		t.FailNow()

		return
	}
}

func TestIsTraceEnabled(t *testing.T) {
	logger.SetLevel("trace")

	if !logger.IsTraceEnabled() {
		t.FailNow()

		return
	}
}

func TestIsDebugEnabled(t *testing.T) {
	logger.SetLevel("debug")

	if !logger.IsDebugEnabled() {
		t.FailNow()

		return
	}
}

func TestIsWarnEnabled(t *testing.T) {
	logger.SetLevel("warn")

	if !logger.IsWarnEnabled() {
		t.FailNow()

		return
	}
}
