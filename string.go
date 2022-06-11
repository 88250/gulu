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
	"bytes"
	"strings"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

// RemoveInvisible removes invisible characters from string str.
func RemoveInvisible(str string) string {
	str = strings.ReplaceAll(str, "\u00A0", " ") // NBSP 转换为普通空格
	str = RemoveZeroWidthCharacters(str)
	str = RemoveCtl(str)
	return str
}

// RemoveCtl removes all control characters from string str.
func RemoveCtl(str string) string {
	return strings.Map(func(r rune) rune {
		if r >= 32 && r != 127 {
			return r
		}
		return -1
	}, str)
}

const (
	// ZWSP represents zero-width space.
	ZWSP = '\u200B'

	// ZWNBSP represents zero-width no-break space.
	ZWNBSP = '\uFEFF'

	// ZWJ represents zero-width joiner.
	ZWJ = '\u200D'

	// ZWNJ represents zero-width non-joiner.
	ZWNJ = '\u200C'

	empty = ""
)

var replacer = strings.NewReplacer(string(ZWSP), empty,
	string(ZWNBSP), empty,
	string(ZWJ), empty,
	string(ZWNJ), empty)

// HasZeroWidthCharacters reports whether string s contains zero-width characters.
func HasZeroWidthCharacters(s string) bool {
	return strings.ContainsRune(s, ZWSP) ||
		strings.ContainsRune(s, ZWNBSP) ||
		strings.ContainsRune(s, ZWJ) ||
		strings.ContainsRune(s, ZWNJ)
}

// RemoveZeroWidthCharacters removes all zero-width characters from string s.
func RemoveZeroWidthCharacters(s string) string {
	return replacer.Replace(s)
}

// RemoveZeroWidthSpace removes zero-width space characters from string s.
func RemoveZeroWidthSpace(s string) string {
	return strings.Replace(s, string(ZWSP), empty, -1)
}

// RemoveZeroWidthNoBreakSpace removes zero-width no-break space characters from string s.
func RemoveZeroWidthNoBreakSpace(s string) string {
	return strings.Replace(s, string(ZWNBSP), empty, -1)
}

// RemoveZeroWidthJoiner removes zero-width joiner characters from string s.
func RemoveZeroWidthJoiner(s string) string {
	return strings.Replace(s, string(ZWJ), empty, -1)
}

// RemoveZeroWidthNonJoiner removes zero-width non-joiner characters from string s.
func RemoveZeroWidthNonJoiner(s string) string {
	return strings.Replace(s, string(ZWNJ), empty, -1)
}

func IsASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

// SubstringsBetween returns a slice of sub strings between the start and end.
func SubstringsBetween(str, start, end string) (ret []string) {
	parts := strings.Split(str, start)
	for _, p := range parts {
		if !strings.Contains(p, end) {
			continue
		}
		parts2 := strings.Split(p, end)
		ret = append(ret, parts2[0])
	}
	return
}

// FromBytes converts the specified byte array to a string.
func (*GuluStr) FromBytes(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

// ToBytes converts the specified str to a byte array.
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
			idx := strings.Index(sub, strings.ToLower(searchStrRepl[j]))
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

// Enclose encloses search strings with open and close, case-insensitively.
func (*GuluStr) EncloseIgnoreCase(text, open, close string, searchStrs ...string) string {
	buf := &bytes.Buffer{}
	textLower := strings.ToLower(text)
	for i := 0; i < len(textLower); i++ {
		sub := textLower[i:]
		var found bool
		for j := 0; j < len(searchStrs); j++ {
			idx := strings.Index(sub, strings.ToLower(searchStrs[j]))
			if 0 != idx {
				continue
			}
			buf.WriteString(open)
			buf.WriteString(text[i : i+len(searchStrs[j])])
			buf.WriteString(close)
			i += len(searchStrs[j]) - 1
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
