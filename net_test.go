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

import "testing"

func TestLocalIP(t *testing.T) {
	ip, err := Net.LocalIP()

	if nil != err {
		t.Error(err)
	}

	t.Log(ip)
}

func TestLocalMac(t *testing.T) {
	mac, err := Net.LocalMac()

	if nil != err {
		t.Error(err)
	}

	t.Log(mac)
}
