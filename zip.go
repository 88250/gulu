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
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// ZipFile represents a zip file.
type ZipFile struct {
	zipFile *os.File
	writer  *zip.Writer
}

// Create creates a zip file with the specified filename.
func (*GuluZip) Create(filename string) (*ZipFile, error) {
	file, err := os.Create(filename)
	if nil != err {
		return nil, err
	}
	return &ZipFile{zipFile: file, writer: zip.NewWriter(file)}, nil
}

// Close closes the zip file writer.
func (z *ZipFile) Close() error {
	if err := z.writer.Close(); nil != err {
		return err
	}
	return z.zipFile.Close() // close the underlying writer
}

// AddEntry adds an entry.
func (z *ZipFile) AddEntry(path, name string, callback ...func(filename string)) error {
	fi, err := os.Stat(name)
	if nil != err {
		return err
	}

	fh, err := zip.FileInfoHeader(fi)
	if nil != err {
		return err
	}

	fh.Name = filepath.ToSlash(filepath.Clean(path))
	fh.Method = zip.Deflate // data compression algorithm

	if fi.IsDir() {
		fh.Name = fh.Name + "/" // be care the ending separator
	}

	entry, err := z.writer.CreateHeader(fh)
	if nil != err {
		return err
	}

	if fi.IsDir() {
		return nil
	}

	file, err := os.Open(name)
	if nil != err {
		return err
	}
	defer file.Close()

	_, err = io.Copy(entry, file)

	if 0 < len(callback) {
		callback[0](name)
	}
	return err
}

// AddDirectory adds a directory.
func (z *ZipFile) AddDirectory(path, dirName string, callback ...func(string)) error {
	files, err := os.ReadDir(dirName)
	if nil != err {
		return err
	}

	if 0 == len(files) {
		err = z.AddEntry(path, dirName, callback...)
		if nil != err {
			return err
		}
		return nil
	}

	for _, file := range files {
		localPath := filepath.Join(dirName, file.Name())
		zipPath := filepath.Join(path, file.Name())

		err = nil
		if file.IsDir() {
			err = z.AddDirectory(zipPath, localPath, callback...)
		} else {
			err = z.AddEntry(zipPath, localPath, callback...)
		}

		if nil != err {
			return err
		}
	}
	return nil
}

// Unzip extracts a zip file specified by the zipFilePath to the destination.
func (*GuluZip) Unzip(zipFilePath, destination string, callback ...func(filename string)) error {
	r, err := zip.OpenReader(zipFilePath)

	if nil != err {
		return err
	}
	defer r.Close()

	var cb func(string)
	if 0 < len(callback) {
		cb = callback[0]
	}

	for _, f := range r.File {
		err = cloneZipItem(f, destination, cb)
		if nil != err {
			return err
		}
	}
	return nil
}

func cloneZipItem(f *zip.File, dest string, callback func(filename string)) error {
	// create full directory path
	fileName := f.Name

	if !utf8.ValidString(fileName) {
		data, err := io.ReadAll(transform.NewReader(bytes.NewReader([]byte(fileName)), simplifiedchinese.GB18030.NewDecoder()))
		if nil == err {
			fileName = string(data)
		} else {
			logger.Error(err)
		}
	}

	path := filepath.Join(dest, fileName)

	err := os.MkdirAll(filepath.Dir(path), 0755)
	if nil != err {
		return err
	}

	if f.FileInfo().IsDir() {
		if err = os.MkdirAll(path, 0755); nil != err {
			return err
		}
		return nil
	}

	// clone if item is a file

	rc, err := f.Open()
	if nil != err {
		return err
	}

	defer rc.Close()

	// use os.Create() since Zip don't store file permissions
	fileCopy, err := os.Create(path)
	if nil != err {
		return err
	}
	defer fileCopy.Close()

	_, err = io.Copy(fileCopy, rc)
	if nil != err {
		return err
	}

	if err = os.Chtimes(path, f.Modified, f.Modified); nil != err {
		return err
	}

	if nil != callback {
		callback(path)
	}
	return nil
}
