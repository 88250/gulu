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
	"bytes"
	"strings"
	"unicode/utf8"
	"unsafe"
)

// FromBytes converts the specified byte array to a string.
func (*GuluStr) FromBytes(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

// Bytes converts the specified str to a byte array.
func (*GuluStr) ToBytes(str string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&str))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// Contains determines whether the str is in the strs.
func (*GuluStr) Contains(str string, strs []string) bool {
	for _, v := range strs {
		if v == str {
			return true
		}
	}
	return false
}

// ReplaceIgnoreCase replace searchStr with repl in the text, case-insensitively.
func (*GuluStr) ReplaceIgnoreCase(text, searchStr, repl string) string {
	buf := &bytes.Buffer{}
	textLower := strings.ToLower(text)
	searchStrLower := strings.ToLower(searchStr)
	searchStrLen := len(searchStr)
	var end int
	for {
		idx := strings.Index(textLower, searchStrLower)
		if 0 > idx {
			break
		}

		buf.WriteString(text[:idx])
		buf.WriteString(repl)
		end = idx + searchStrLen
		textLower = textLower[end:]
	}
	buf.WriteString(text[end:])
	return buf.String()
}

// ReplacesIgnoreCase replace searchStr-repl pairs in the text, case-insensitively.
func (*GuluStr) ReplacesIgnoreCase(text string, searchStrRepl ...string) string {
	if 1 == len(searchStrRepl)%2 {
		return text
	}

	buf := &bytes.Buffer{}
	textLower := strings.ToLower(text)
	for i := 0; i < len(textLower); i++ {
		sub := textLower[i:]
		var found bool
		for j := 0; j < len(searchStrRepl); j += 2 {
			idx := strings.Index(sub, searchStrRepl[j])
			if 0 != idx {
				continue
			}
			buf.WriteString(searchStrRepl[j+1])
			i += len(searchStrRepl[j]) - 1
			found = true
			break
		}
		if !found {
			buf.WriteByte(text[i])
		}
	}
	return buf.String()
}

// LCS gets the longest common substring of s1 and s2.
//
// Refers to http://en.wikibooks.org/wiki/Algorithm_Implementation/Strings/Longest_common_substring.
func (*GuluStr) LCS(s1 string, s2 string) string {
	var m = make([][]int, 1+len(s1))

	for i := 0; i < len(m); i++ {
		m[i] = make([]int, 1+len(s2))
	}

	longest := 0
	xLongest := 0

	for x := 1; x < 1+len(s1); x++ {
		for y := 1; y < 1+len(s2); y++ {
			if s1[x-1] == s2[y-1] {
				m[x][y] = m[x-1][y-1] + 1
				if m[x][y] > longest {
					longest = m[x][y]
					xLongest = x
				}
			} else {
				m[x][y] = 0
			}
		}
	}
	return s1[xLongest-longest : xLongest]
}

// SubStr decode str into runes and get substring with the specified length.
func (*GuluStr) SubStr(str string, length int) (ret string) {
	var count int
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		i += size
		ret += string(r)
		count++
		if length <= count {
			break
		}
	}
	return
}
