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
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// IsSubPath checks whether the toCheckPath is a sub path of absPath. Both paths should be absolute paths.
func (*GuluFile) IsSubPath(absPath, toCheckPath string) bool {
	if 1 > len(absPath) || 1 > len(toCheckPath) {
		return false
	}
	if absPath == toCheckPath { // 相同路径时不认为是子路径
		return false
	}

	if OS.IsWindows() {
		if filepath.IsAbs(absPath) && filepath.IsAbs(toCheckPath) {
			if strings.ToLower(absPath)[0] != strings.ToLower(toCheckPath)[0] {
				// 不在一个盘
				return false
			}
		}
	}

	up := ".." + string(os.PathSeparator)
	rel, err := filepath.Rel(absPath, toCheckPath)
	if err != nil {
		return false
	}
	if !strings.HasPrefix(rel, up) && rel != ".." {
		return true
	}
	return false
}

// RemoveEmptyDirs removes all empty dirs under the specified dir path.
func (*GuluFile) RemoveEmptyDirs(dir string, excludes ...string) (err error) {
	_, err = removeEmptyDirs(dir, excludes...)
	return
}

func removeEmptyDirs(dir string, excludes ...string) (removed bool, err error) {
	// Credit to: https://github.com/InfuseAI/ArtiVC/blob/main/internal/core/utils.go
	// LICENSE Apache License 2.0 https://github.com/InfuseAI/ArtiVC/blob/main/LICENSE

	dirName := filepath.Base(dir)
	if Str.Contains(dirName, excludes) {
		return
	}

	var hasEntries bool
	entires, err := os.ReadDir(dir)
	if err != nil {
		return false, err
	}
	for _, entry := range entires {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			removed, err = removeEmptyDirs(subdir, excludes...)
			if err != nil {
				return false, err
			}
			if !removed {
				hasEntries = true
			}
		} else {
			hasEntries = true
		}
	}

	if !hasEntries && !Str.Contains(dirName, excludes) {
		err = os.Remove(dir)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func (*GuluFile) IsValidFilename(name string) bool {
	reserved := []string{"\\", "/", ":", "*", "?", "\"", "'", "<", ">", "|"}
	for _, r := range reserved {
		if strings.Contains(name, r) {
			return false
		}
	}
	return true
}

// WriteFileSaferByReader writes the data to a temp file and atomically move if everything else succeeds.
func (*GuluFile) WriteFileSaferByReader(writePath string, reader io.Reader, perm os.FileMode) (err error) {
	dir, name := filepath.Split(writePath)
	tmp := filepath.Join(dir, name+Rand.String(7)+".tmp")
	f, err := os.OpenFile(tmp, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if nil != err {
		return
	}

	if _, err = io.Copy(f, reader); nil != err {
		return
	}

	if err = f.Sync(); nil != err {
		return
	}

	if err = f.Close(); nil != err {
		return
	}

	if err = os.Chmod(f.Name(), perm); nil != err {
		return
	}

	for i := 0; i < 3; i++ {
		err = os.Rename(f.Name(), writePath) // Windows 上重命名是非原子的
		if nil == err {
			os.Remove(f.Name())
			return
		}

		if errMsg := strings.ToLower(err.Error()); strings.Contains(errMsg, "access is denied") || strings.Contains(errMsg, "used by another process") { // 文件可能是被锁定
			time.Sleep(200 * time.Millisecond)
			continue
		}
		break
	}
	return
}

// WriteFileSafer writes the data to a temp file and atomically move if everything else succeeds.
func (*GuluFile) WriteFileSafer(writePath string, data []byte, perm os.FileMode) (err error) {
	// credits: https://github.com/vitessio/vitess/blob/master/go/ioutil2/ioutil.go

	dir, name := filepath.Split(writePath)
	tmp := filepath.Join(dir, name+Rand.String(7)+".tmp")
	f, err := os.OpenFile(tmp, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if nil != err {
		return
	}

	if _, err = f.Write(data); nil != err {
		return
	}

	if err = f.Sync(); nil != err {
		return
	}

	if err = f.Close(); nil != err {
		return
	}

	if err = os.Chmod(f.Name(), perm); nil != err {
		return
	}

	for i := 0; i < 3; i++ {
		err = os.Rename(f.Name(), writePath) // Windows 上重命名是非原子的
		if nil == err {
			os.Remove(f.Name())
			return
		}

		if errMsg := strings.ToLower(err.Error()); strings.Contains(errMsg, "access is denied") || strings.Contains(errMsg, "used by another process") { // 文件可能是被锁定
			time.Sleep(200 * time.Millisecond)
			continue
		}
		break
	}
	return
}

// GetFileSize get the length in bytes of file of the specified path.
func (*GuluFile) GetFileSize(path string) int64 {
	fi, err := os.Stat(path)
	if nil != err {
		logger.Error(err)

		return -1
	}

	return fi.Size()
}

// IsExist determines whether the file spcified by the given path is exists.
func (*GuluFile) IsExist(path string) bool {
	_, err := os.Stat(path)

	return err == nil || os.IsExist(err)
}

// IsBinary determines whether the specified content is a binary file content.
func (*GuluFile) IsBinary(content string) bool {
	for _, b := range content {
		if 0 == b {
			return true
		}
	}

	return false
}

// IsImg determines whether the specified extension is a image.
func (*GuluFile) IsImg(extension string) bool {
	ext := strings.ToLower(extension)

	switch ext {
	case ".jpg", ".jpeg", ".bmp", ".gif", ".png", ".svg", ".ico":
		return true
	default:
		return false
	}
}

// IsDir determines whether the specified path is a directory.
func (*GuluFile) IsDir(path string) bool {
	fio, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return false
	}

	if nil != err {
		logger.Warnf("determines whether [%s] is a directory failed: [%v]", path, err)
		return false
	}
	return fio.IsDir()
}

var (
	ErrCopyToSub = errors.New("cannot copy to a sub path of the source")
)

// Copy copies the source to the dest.
// Keep the dest access/mod time as the same as the source.
func (gl *GuluFile) Copy(source, dest string) (err error) {
	if !gl.IsExist(source) {
		return os.ErrNotExist
	}

	if gl.IsDir(source) {
		if gl.IsSubPath(source, dest) {
			return ErrCopyToSub
		}
		return gl.copyDir(source, dest, false, true)
	}
	return gl.copyFile(source, dest, false, true)
}

// CopyWithoutHidden copies the source to the dest without hidden files.
func (gl *GuluFile) CopyWithoutHidden(source, dest string) (err error) {
	if !gl.IsExist(source) {
		return os.ErrNotExist
	}

	if gl.IsDir(source) {
		if gl.IsSubPath(source, dest) {
			return ErrCopyToSub
		}
		return gl.copyDir(source, dest, true, true)
	}
	return gl.copyFile(source, dest, true, true)
}

// CopyNewtimes copies the source to the dest.
// Do not keep the dest access/mod time as the same as the source.
func (gl *GuluFile) CopyNewtimes(source, dest string) (err error) {
	if !gl.IsExist(source) {
		return os.ErrNotExist
	}

	if gl.IsDir(source) {
		return gl.CopyDirNewtimes(source, dest)
	}
	return gl.CopyFileNewtimes(source, dest)
}

// CopyFile copies the source file to the dest file.
// Keep the dest access/mod time as the same as the source.
func (gl *GuluFile) CopyFile(source, dest string) (err error) {
	return gl.copyFile(source, dest, false, true)
}

// CopyFileNewtimes copies the source file to the dest file.
// Do not keep the dest access/mod time as the same as the source.
func (gl *GuluFile) CopyFileNewtimes(source, dest string) (err error) {
	return gl.copyFile(source, dest, false, false)
}

func (gl *GuluFile) copyFile(source, dest string, ignoreHidden, chtimes bool) (err error) {
	sourceinfo, err := os.Lstat(source)
	if nil != err {
		return
	}

	if 0 != sourceinfo.Mode()&os.ModeSymlink {
		// 忽略符号链接
		return
	}

	if ignoreHidden && gl.IsHidden(source) {
		return
	}

	if err = os.MkdirAll(filepath.Dir(dest), 0755); nil != err {
		return
	}

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destfile.Close()

	if err = os.Chmod(dest, sourceinfo.Mode()); nil != err {
		return
	}

	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourcefile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if nil != err {
		return
	}

	if chtimes {
		if err = os.Chtimes(dest, sourceinfo.ModTime(), sourceinfo.ModTime()); nil != err {
			return
		}
	}
	return
}

// CopyDir copies the source directory to the dest directory.
// Keep the dest access/mod time as the same as the source.
func (gl *GuluFile) CopyDir(source, dest string) (err error) {
	if gl.IsSubPath(source, dest) {
		return ErrCopyToSub
	}
	return gl.copyDir(source, dest, false, true)
}

// CopyDirNewtimes copies the source directory to the dest directory.
// Do not keep the dest access/mod time as the same as the source.
func (gl *GuluFile) CopyDirNewtimes(source, dest string) (err error) {
	if gl.IsSubPath(source, dest) {
		return ErrCopyToSub
	}
	return gl.copyDir(source, dest, false, false)
}

func (gl *GuluFile) copyDir(source, dest string, ignoreHidden, chtimes bool) (err error) {
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	if ignoreHidden && gl.IsHidden(source) {
		return
	}

	if err = os.MkdirAll(dest, 0755); err != nil {
		return err
	}

	dirs, err := os.ReadDir(source)
	if err != nil {
		return err
	}

	for _, f := range dirs {
		srcFilePath := filepath.Join(source, f.Name())
		destFilePath := filepath.Join(dest, f.Name())

		if f.IsDir() {
			err = gl.copyDir(srcFilePath, destFilePath, ignoreHidden, chtimes)
			if err != nil {
				logger.Error(err)
				return
			}
		} else {
			err = gl.copyFile(srcFilePath, destFilePath, ignoreHidden, chtimes)
			if err != nil {
				logger.Error(err)
				return
			}
		}
	}

	if chtimes {
		if err = os.Chtimes(dest, sourceinfo.ModTime(), sourceinfo.ModTime()); nil != err {
			return
		}
	}
	return nil
}

// GrepResult represents a single match of a Grep search.
type GrepResult struct {
	File    string // file path where the match was found
	Line    int    // 1-based line number
	Text    string // the full line content
	Context bool   // true if this is a context line (neither match nor actual result)
}

// Grep searches files matching regexPattern, optionally filtered by includeGlob.
// root can be a file or directory path. includeGlob is a file glob pattern like
// "*.go" or "*.{ts,tsx}". regexPattern is the regular expression to match.
// context specifies how many lines of context before and after each match to include
// (0 means only matching lines). maxResults limits the number of returned results
// (0 or negative defaults to 64). Hidden directories and binary files are skipped.
func (*GuluFile) Grep(root string, includeGlob string, regexPattern string, context int, maxResults int) ([]*GrepResult, error) {
	re, err := regexp.Compile(regexPattern)
	if err != nil {
		return nil, err
	}

	if maxResults <= 0 {
		maxResults = 64
	}

	var results []*GrepResult
	info, err := os.Stat(root)
	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		err = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return nil
			}

			if d.IsDir() {
				if skipDir(d.Name()) {
					return filepath.SkipDir
				}
				return nil
			}

			if !d.Type().IsRegular() {
				return nil
			}

			if includeGlob != "" && !matchInclude(d.Name(), includeGlob) {
				return nil
			}

			grepFile(path, re, context, &results, maxResults)
			if len(results) > maxResults {
				results = results[:maxResults]
			}

			return nil
		})
	} else {
		grepFile(root, re, context, &results, maxResults)
	}

	if err != nil {
		return results, err
	}

	if len(results) > maxResults {
		results = results[:maxResults]
	}

	return results, nil
}

// skipDir returns true if the directory should be skipped during traversal.
func skipDir(name string) bool {
	return name == ".git" || name == ".svn" || name == ".hg" || strings.HasPrefix(name, ".")
}

// matchInclude checks whether filename matches the include glob pattern.
// Supports brace expansion like "*.{go,ts}".
func matchInclude(filename, includeGlob string) bool {
	patterns := expandBrace(includeGlob)
	for _, p := range patterns {
		if matched, _ := filepath.Match(p, filename); matched {
			return true
		}
	}
	return false
}

// expandBrace expands brace patterns like "*.{go,ts}" into ["*.go", "*.ts"].
func expandBrace(pattern string) []string {
	i := strings.Index(pattern, "{")
	if i < 0 {
		return []string{pattern}
	}

	j := strings.Index(pattern[i:], "}")
	if j < 0 {
		return []string{pattern}
	}
	j += i

	prefix := pattern[:i]
	body := pattern[i+1 : j]
	suffix := pattern[j+1:]

	var result []string
	for _, opt := range strings.Split(body, ",") {
		result = append(result, expandBrace(prefix+strings.TrimSpace(opt)+suffix)...)
	}
	return result
}

// grepFile reads a single file line by line and appends matching lines to results.
// context specifies how many lines before and after each match to include as context.
func grepFile(path string, re *regexp.Regexp, context int, results *[]*GrepResult, maxResults int) {
	if len(*results) >= maxResults {
		return
	}

	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	type bufEntry struct {
		lineNum int
		text    string
	}

	beforeBuf := make([]bufEntry, 0, context+1)
	afterRemaining := 0

	flushBeforeBuf := func() {
		for _, e := range beforeBuf {
			if len(*results) >= maxResults {
				return
			}
			*results = append(*results, &GrepResult{
				File:    path,
				Line:    e.lineNum,
				Text:    e.text,
				Context: true,
			})
		}
		beforeBuf = beforeBuf[:0]
	}

	emit := func(lineNum int, text string, isContext bool) {
		if len(*results) >= maxResults {
			return
		}
		*results = append(*results, &GrepResult{
			File:    path,
			Line:    lineNum,
			Text:    text,
			Context: isContext,
		})
	}

	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	lineNum := 0
	for scanner.Scan() {
		if len(*results) >= maxResults {
			return
		}
		lineNum++
		line := scanner.Text()

		if re.MatchString(line) {
			flushBeforeBuf()
			emit(lineNum, line, false)
			afterRemaining = context
		} else if afterRemaining > 0 {
			emit(lineNum, line, true)
			afterRemaining--
		} else {
			beforeBuf = append(beforeBuf, bufEntry{lineNum, line})
			if len(beforeBuf) > context {
				copy(beforeBuf, beforeBuf[1:])
				beforeBuf = beforeBuf[:len(beforeBuf)-1]
			}
		}
	}
}
