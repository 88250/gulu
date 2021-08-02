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
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// WriteFileSaferByHandle writes the data to a temp file and writes the original file if everything else succeeds.
// Note: This function does not close the file handle after writing data.
func (GuluFile) WriteFileSaferByHandle(handle *os.File, data []byte) error {
	writePath := handle.Name()
	dir, name := filepath.Split(writePath)
	f, err := ioutil.TempFile(dir, name+"*.tmp")
	if nil != err {
		return err
	}

	if _, err = f.Write(data); nil == err {
		err = f.Sync()
	}

	if closeErr := f.Close(); nil == err {
		err = closeErr
	}

	if nil == err {
		err = handle.Truncate(0)
	}

	if nil == err {
		_, err = handle.WriteAt(data, 0)
	}

	if nil == err {
		err = handle.Sync()
	}

	if nil == err {
		os.Remove(f.Name())
	}
	return err
}

// WriteFileSafer writes the data to a temp file and atomically move if everything else succeeds.
func (GuluFile) WriteFileSafer(writePath string, data []byte, perm os.FileMode) error {
	// credits: https://github.com/vitessio/vitess/blob/master/go/ioutil2/ioutil.go

	dir, name := filepath.Split(writePath)
	f, err := ioutil.TempFile(dir, name+"*.tmp")
	if nil != err {
		return err
	}

	if _, err = f.Write(data); nil == err {
		err = f.Sync()
	}

	if closeErr := f.Close(); nil == err {
		err = closeErr
	}

	if permErr := os.Chmod(f.Name(), perm); nil == err {
		err = permErr
	}

	if nil == err {
		var renamed bool
		for i := 0; i < 3; i++ {
			err = os.Rename(f.Name(), writePath) // Windows 上重命名是非原子的
			if nil == err {
				renamed = true
				break
			}

			if errMsg := strings.ToLower(err.Error()); strings.Contains(errMsg, "access is denied") || strings.Contains(errMsg, "used by another process") { // 文件可能是被锁定
				time.Sleep(100 * time.Millisecond)
				continue
			}
			break
		}

		if !renamed {
			// 直接写入
			err = os.WriteFile(writePath, data, perm)
		}
	}

	if nil != err {
		os.Remove(f.Name())
	}
	return err
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

// Copy copies the source to the dest.
// Keep the dest access/mod time as the same as the source.
func (gl *GuluFile) Copy(source, dest string) (err error) {
	if !gl.IsExist(source) {
		return os.ErrNotExist
	}

	if gl.IsDir(source) {
		return gl.CopyDir(source, dest)
	}
	return gl.CopyFile(source, dest)
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
	return gl.copyFile(source, dest, true)
}

// CopyFileNewtimes copies the source file to the dest file.
// Do not keep the dest access/mod time as the same as the source.
func (gl *GuluFile) CopyFileNewtimes(source, dest string) (err error) {
	return gl.copyFile(source, dest, false)
}

func (*GuluFile) copyFile(source, dest string, chtimes bool) (err error) {
	if err = os.MkdirAll(filepath.Dir(dest), 0755); nil != err {
		return
	}

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destfile.Close()

	sourceinfo, err := os.Stat(source)
	if nil != err {
		return
	}

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
	return gl.copyDir(source, dest, true)
}

// CopyDirNewtimes copies the source directory to the dest directory.
// Do not keep the dest access/mod time as the same as the source.
func (gl *GuluFile) CopyDirNewtimes(source, dest string) (err error) {
	return gl.copyDir(source, dest, false)
}

func (gl *GuluFile) copyDir(source, dest string, chtimes bool) (err error) {
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
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
			err = gl.copyDir(srcFilePath, destFilePath, chtimes)
			if err != nil {
				logger.Error(err)
				return
			}
		} else {
			err = gl.copyFile(srcFilePath, destFilePath, chtimes)
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
