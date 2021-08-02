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
	"math/rand"
	"time"
)

// Ints returns a random integer array with the specified from, to and size.
func (*GuluRand) Ints(from, to, size int) []int {
	if to-from < size {
		size = to - from
	}

	var slice []int
	for i := from; i < to; i++ {
		slice = append(slice, i)
	}

	var ret []int
	for i := 0; i < size; i++ {
		idx := rand.Intn(len(slice))
		ret = append(ret, slice[idx])
		slice = append(slice[:idx], slice[idx+1:]...)
	}
	return ret
}

// String returns a random string ['a', 'z'] and ['0', '9'] in the specified length.
func (*GuluRand) String(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	time.Sleep(10 * time.Nanosecond)

	letter := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

// Int returns a random integer in range [min, max].
func (*GuluRand) Int(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(10 * time.Nanosecond)
	return min + rand.Intn(max-min)
}
