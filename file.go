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

// WriteFileSafer writes the data to a temp file and atomically move if everything else succeeds.
func (GuluFile) WriteFileSafer(writePath string, data []byte, perm os.FileMode) error {
	// credits: https://github.com/vitessio/vitess/blob/master/go/ioutil2/ioutil.go

	dir, name := filepath.Split(writePath)
	f, err := ioutil.TempFile(dir, name)
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
		logger.Warnf("Determines whether [%s] is a directory failed: [%v]", path, err)

		return false
	}

	return fio.IsDir()
}

// Copy copies the source to the dest.
func (gl *GuluFile) Copy(source, dest string) (err error) {
	if !gl.IsExist(source) {
		return os.ErrNotExist
	}

	if gl.IsDir(source) {
		return gl.CopyDir(source, dest)
	}
	return gl.CopyFile(source, dest)
}

// CopyFile copies the source file to the dest file.
func (*GuluFile) CopyFile(source, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourcefile.Close()

	if err = os.MkdirAll(filepath.Dir(dest), 0755); nil != err {
		return
	}

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if nil == err {
		sourceinfo, err := os.Stat(source)
		if nil == err {
			os.Chmod(dest, sourceinfo.Mode())
			os.Chtimes(dest, sourceinfo.ModTime(), sourceinfo.ModTime())
		}
	}
	return nil
}

// CopyDir copies the source directory to the dest directory.
func (*GuluFile) CopyDir(source, dest string) (err error) {
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir
	if err = os.MkdirAll(dest, sourceinfo.Mode()); err != nil {
		return err
	}

	directory, err := os.Open(source)
	if err != nil {
		return err
	}
	defer directory.Close()

	objects, err := directory.Readdir(-1)
	if err != nil {
		return err
	}

	for _, obj := range objects {
		srcFilePath := filepath.Join(source, obj.Name())
		destFilePath := filepath.Join(dest, obj.Name())

		if obj.IsDir() {
			// create sub-directories - recursively
			err = File.CopyDir(srcFilePath, destFilePath)
			if err != nil {
				logger.Error(err)
			}
		} else {
			err = File.CopyFile(srcFilePath, destFilePath)
			if err != nil {
				logger.Error(err)
			}
		}
	}

	os.Chtimes(dest, sourceinfo.ModTime(), sourceinfo.ModTime())
	return nil
}
