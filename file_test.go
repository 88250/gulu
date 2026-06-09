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
	errors "errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestIsSubPath(t *testing.T) {
	subPath := filepath.Join(testdataDir, "subPath")
	if !File.IsSubPath(testdataDir, subPath) {
		t.Errorf("[%s] should be a sub path of [%s]", subPath, testdataDir)
		return
	}
}

func TestRemoveEmptyDirs(t *testing.T) {
	testPath := "testdata/dir"

	// case 1
	if err := os.RemoveAll(testPath); nil != err {
		t.Errorf("clear test empty dir [%s] failed: %s", testPath, err)
	}

	a := filepath.Join(testPath, "a")
	if err := os.MkdirAll(a, 0755); nil != err {
		t.Errorf("make dir [%s] failed: %s", testPath, err)
	}

	if err := File.RemoveEmptyDirs(testPath); nil != err {
		t.Errorf("remove empty dirs failed: %s", err)
	}

	if File.IsDir(a) || File.IsDir(testPath) {
		t.Errorf("empty dir [%s] exists", a)
	}

	// case 2
	if err := os.RemoveAll(testPath); nil != err {
		t.Errorf("clear test empty dir [%s] failed: %s", testPath, err)
	}

	if err := os.MkdirAll(a, 0755); nil != err {
		t.Errorf("make dir [%s] failed: %s", testPath, err)
	}
	test := filepath.Join(a, "test")
	if err := os.WriteFile(test, []byte(""), 0644); nil != err {
		t.Errorf("write file [%s] failed: %s", test, err)
	}

	if err := File.RemoveEmptyDirs(testPath); nil != err {
		t.Errorf("remove empty dirs failed: %s", err)
	}

	if !File.IsDir(a) || !File.IsDir(testPath) {
		t.Errorf("empty dir [%s] exists", a)
	}

	// case 3
	if err := os.RemoveAll(testPath); nil != err {
		t.Errorf("clear test empty dir [%s] failed: %s", testPath, err)
	}

	if err := os.MkdirAll(a, 0755); nil != err {
		t.Errorf("make dir [%s] failed: %s", testPath, err)
	}

	if err := File.RemoveEmptyDirs(testPath, "a"); nil != err {
		t.Errorf("remove empty dirs failed: %s", err)
	}

	if !File.IsDir(a) || !File.IsDir(testPath) {
		t.Errorf("empty dir [%s] exists", a)
	}

	if err := os.RemoveAll(testPath); nil != err {
		t.Errorf("clear test empty dir [%s] failed: %s", testPath, err)
	}
}

func TestIsValidFilename(t *testing.T) {
	if !File.IsValidFilename("hello.go") {
		t.Errorf("[hello.go] should be a valid filename")
	}
	if File.IsValidFilename("hello?.go") {
		t.Errorf("[hello?.go] should not be a valid filename")
	}
}

func TestWriteFileSaferByReader(t *testing.T) {
	writePath := "testdata/filewrite.go"
	defer os.RemoveAll(writePath)
	if err := File.WriteFileSaferByReader(writePath, strings.NewReader("test"), 0644); nil != err {
		t.Errorf("write file [%s] failed: %s", writePath, err)
	}
}

func TestWriteFileSafer(t *testing.T) {
	writePath := "testdata/filewrite.go"
	defer os.RemoveAll(writePath)

	if err := os.WriteFile(writePath, []byte("0"), 0644); nil != err {
		t.Fatalf("write file [%s] failed: %s", writePath, err)
	}

	info, err := os.Stat(writePath)
	if nil != err {
		t.Fatalf("stat file [%s] failed: %s", writePath, err)
	}

	if err = File.WriteFileSafer(writePath, []byte("test"), 0644); nil != err {
		t.Errorf("write file [%s] failed: %s", writePath, err)
	}

	info, err = os.Stat(writePath)
	if nil != err {
		t.Fatalf("stat file [%s] failed: %s", writePath, err)
	}
	modTime2 := info.ModTime()
	t.Logf("file mod time [%v]", modTime2)
}

func TestIsHidden(t *testing.T) {
	filename := "./file.go"
	isHidden := File.IsHidden(filename)
	if isHidden {
		t.Error("file [" + filename + "] is not hidden")
	}
}

func TestGetFileSize(t *testing.T) {
	filename := "./file.go"
	size := File.GetFileSize(filename)
	if 0 > size {
		t.Error("file [" + filename + "] size is [" + strconv.FormatInt(size, 10) + "]")
	}
}

func TestIsExist(t *testing.T) {
	if !File.IsExist(".") {
		t.Error(". must exist")
		return
	}
}

func TestIdBinary(t *testing.T) {
	if File.IsBinary("not binary content") {
		t.Error("The content should not be binary")
		return
	}
}

func TestIsImg(t *testing.T) {
	if !File.IsImg(".jpg") {
		t.Error(".jpg should be a valid extension of a image file")
		return
	}
}

func TestIsDir(t *testing.T) {
	if !File.IsDir(".") {
		t.Error(". should be a directory")
		return
	}
}

func TestCopyDir(t *testing.T) {
	source := "testcopydir"
	os.Mkdir(source, 0644)
	dest := filepath.Join(testdataDir, source)
	defer os.Remove(dest)

	err := File.CopyDir(source, dest)
	if nil != err {
		t.Error("Copy dir failed: ", err)
		return
	}

	sourceStat, _ := os.Stat(source)
	destStat, _ := os.Stat(dest)
	if sourceStat.ModTime() != destStat.ModTime() {
		t.Error("mod time is not equal")
		return
	}
}

func TestCopyFile(t *testing.T) {
	source := "./file.go"
	dest := filepath.Join(testdataDir, "file.go")
	defer os.Remove(dest)
	err := File.CopyFile(source, dest)
	if nil != err {
		t.Error("Copy file failed: ", err)
		return
	}

	sourceStat, _ := os.Stat(source)
	destStat, _ := os.Stat(dest)
	if sourceStat.ModTime() != destStat.ModTime() {
		t.Error("mod time is not equal")
		return
	}
}

func TestCopy(t *testing.T) {
	dest := filepath.Join(testdataDir, "file.go")
	defer os.Remove(dest)
	err := File.Copy("./file.go", dest)
	if nil != err {
		t.Error("Copy failed: ", err)
		return
	}

	sourceStat, _ := os.Stat("./file.go")
	destStat, _ := os.Stat(dest)
	if sourceStat.ModTime() != destStat.ModTime() {
		t.Error("mod time is not equal")
		return
	}

	dest = filepath.Join(testdataDir, "subPath")
	err = File.Copy(testdataDir, dest)
	if !errors.Is(err, ErrCopyToSub) {
		t.Error("Copy should fail")
		return
	}
}

func TestCopyWithoutHidden(t *testing.T) {
	dest := filepath.Join(testdataDir, ".gitignore")
	err := File.CopyWithoutHidden("./.gitignore", dest)
	if nil != err {
		t.Error("Copy failed: ", err)
		return
	}

	if File.IsExist(dest) {
		t.Error(".gitignore should not exist")
		return
	}
}

func TestCopyDirNewtimes(t *testing.T) {
	source := "testcopydir"
	os.Mkdir(source, 0644)
	dest := filepath.Join(testdataDir, source)
	defer os.Remove(dest)

	time.Sleep(100 * time.Millisecond) // CI

	err := File.CopyDirNewtimes(source, dest)
	if nil != err {
		t.Error("Copy dir failed: ", err)
		return
	}

	sourceStat, _ := os.Stat(source)
	destStat, _ := os.Stat(dest)
	if sourceStat.ModTime() == destStat.ModTime() {
		t.Error("mod time is equal")
		return
	}
}

func TestCopyFileNewtimes(t *testing.T) {
	source := "./file.go"
	dest := filepath.Join(testdataDir, "file.go")
	defer os.Remove(dest)
	err := File.CopyFileNewtimes(source, dest)
	if nil != err {
		t.Error("Copy file failed: ", err)
		return
	}

	sourceStat, _ := os.Stat(source)
	destStat, _ := os.Stat(dest)
	if sourceStat.ModTime() == destStat.ModTime() {
		t.Error("mod time is equal")
		return
	}
}

func TestGrep(t *testing.T) {
	grepDir := filepath.Join(testdataDir, "grepdir")
	defer os.RemoveAll(grepDir)

	if err := os.MkdirAll(grepDir, 0755); nil != err {
		t.Fatalf("mkdir failed: %s", err)
	}

	os.WriteFile(filepath.Join(grepDir, "a.go"), []byte("package main\nfunc foo() {}\nfunc bar() {}\n"), 0644)
	os.WriteFile(filepath.Join(grepDir, "b.ts"), []byte("const x = 1;\nconst y = () => {};\n"), 0644)
	os.WriteFile(filepath.Join(grepDir, "c.txt"), []byte("hello world\nfoo bar\nHELLO AGAIN\n"), 0644)

	// case 1: grep all files in dir
	results, err := File.Grep(grepDir, "", "func", 0, 100)
	if nil != err {
		t.Fatalf("grep failed: %s", err)
	}
	if 2 != len(results) {
		t.Errorf("expected 2 func matches, got %d", len(results))
	}

	// case 2: grep with .go include
	results, err = File.Grep(grepDir, "*.go", "func", 0, 100)
	if nil != err {
		t.Fatalf("grep failed: %s", err)
	}
	if 2 != len(results) {
		t.Errorf("expected 2 func matches in .go, got %d", len(results))
	}

	// case 3: grep with brace expansion include
	results, err = File.Grep(grepDir, "*.{go,ts}", "const", 0, 100)
	if nil != err {
		t.Fatalf("grep failed: %s", err)
	}
	if 2 != len(results) { // b.ts has 2 "const" lines
		t.Errorf("expected 2 const matches, got %d", len(results))
	}

	// case 4: grep single file
	results, err = File.Grep(filepath.Join(grepDir, "c.txt"), "", "hello", 0, 100)
	if nil != err {
		t.Fatalf("grep single file failed: %s", err)
	}
	if 1 != len(results) {
		t.Errorf("expected 1 hello match, got %d", len(results))
	}

	// case 4b: case-insensitive grep
	results, err = File.Grep(filepath.Join(grepDir, "c.txt"), "", "(?i)hello", 0, 100)
	if nil != err {
		t.Fatalf("grep ci single file failed: %s", err)
	}
	if 2 != len(results) {
		t.Errorf("expected 2 (?i)hello matches, got %d: %+v", len(results), results)
	}

	// case 5: max results limit
	results, err = File.Grep(grepDir, "", "func|const|hello|foo|bar|HELLO", 0, 2)
	if nil != err {
		t.Fatalf("grep with limit failed: %s", err)
	}
	if 2 != len(results) {
		t.Errorf("expected 2 max results, got %d", len(results))
	}

	// case 6: no match
	results, err = File.Grep(grepDir, "", "nonexistent12345", 0, 100)
	if nil != err {
		t.Fatalf("grep no match failed: %s", err)
	}
	if 0 != len(results) {
		t.Errorf("expected 0 matches, got %d", len(results))
	}

	// case 7: regex pattern
	results, err = File.Grep(grepDir, "*.go", `func \w+`, 0, 100)
	if nil != err {
		t.Fatalf("grep regex failed: %s", err)
	}
	if 2 != len(results) {
		t.Errorf("expected 2 regex matches, got %d", len(results))
	}
}

func TestGrepContext(t *testing.T) {
	ctxDir := filepath.Join(testdataDir, "grepctxdir")
	defer os.RemoveAll(ctxDir)

	if err := os.MkdirAll(ctxDir, 0755); nil != err {
		t.Fatalf("mkdir failed: %s", err)
	}

	os.WriteFile(filepath.Join(ctxDir, "d.txt"), []byte("line 1 a\nline 2 b\nline 3 MATCH\nline 4 d\nline 5 e\nline 6 f\nline 7 MATCH\nline 8 h\nline 9 i\nline 10 j\n"), 0644)

	// case 1: context=2, verify before/after context lines present
	results, err := File.Grep(ctxDir, "", "MATCH", 2, 100)
	if nil != err {
		t.Fatalf("grep context failed: %s", err)
	}
	// lines 1,2 (before match 3), 3:: (match), 4,5 (after match 3),
	// 6 (before match 7), 7:: (match), 8,9 (after match 7)
	// line 5 is after-context of match 3, consumed before match 7's before-buf
	if 9 != len(results) {
		t.Errorf("expected 7 results with context, got %d: %+v", len(results), results)
	}
	if results[2].Context {
		t.Errorf("line 3 (MATCH) should not be marked as context")
	}
	if !results[0].Context {
		t.Errorf("line 1 should be marked as context")
	}

	// case 2: context=0 returns only matches
	results, err = File.Grep(ctxDir, "", "MATCH", 0, 100)
	if nil != err {
		t.Fatalf("grep no-context failed: %s", err)
	}
	if 2 != len(results) {
		t.Errorf("expected 2 matches without context, got %d", len(results))
	}
	for _, r := range results {
		if r.Context {
			t.Errorf("no result should be marked as context when context=0")
		}
	}
}

func TestExpandBrace(t *testing.T) {
	patterns := expandBrace("*.{go,ts}")
	if 2 != len(patterns) {
		t.Errorf("expected 2 patterns, got %d: %v", len(patterns), patterns)
	}
	if "*.go" != patterns[0] || "*.ts" != patterns[1] {
		t.Errorf("unexpected patterns: %v", patterns)
	}

	patterns = expandBrace("*.{go,ts,tsx}")
	if 3 != len(patterns) {
		t.Errorf("expected 3 patterns, got %d: %v", len(patterns), patterns)
	}
}

func TestCopyNewtimes(t *testing.T) {
	source := "./file.go"
	dest := filepath.Join(testdataDir, "file.go")
	defer os.Remove(dest)
	err := File.CopyNewtimes(source, dest)
	if nil != err {
		t.Error("Copy failed: ", err)
		return
	}

	sourceStat, _ := os.Stat(source)
	destStat, _ := os.Stat(dest)
	if sourceStat.ModTime() == destStat.ModTime() {
		t.Error("mod time is equal")
		return
	}
}
