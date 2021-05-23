package gulu

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func (*GuluTar) Tar(source, target string) error {
	filename := filepath.Base(source)
	target = filepath.Join(target, fmt.Sprintf("%s.tar", filename))
	if err := os.MkdirAll(filepath.Dir(target), 0755); nil != err {
		return err
	}

	tarfile, err := os.Create(target)
	if nil != err {
		return err
	}
	defer tarfile.Close()

	tarball := tar.NewWriter(tarfile)
	defer tarball.Close()

	info, err := os.Stat(source)
	if nil != err {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	return filepath.Walk(source,
		func(path string, info os.FileInfo, err error) error {
			if nil != err {
				return err
			}
			header, err := tar.FileInfoHeader(info, info.Name())
			if nil != err {
				return err
			}

			if baseDir != "" {
				header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
			}

			if err := tarball.WriteHeader(header); nil != err {
				return err
			}

			if info.IsDir() {
				return nil
			}

			file, err := os.Open(path)
			if nil != err {
				return err
			}
			defer file.Close()
			_, err = io.Copy(tarball, file)
			return err
		})
}

func (*GuluTar) Untar(tarball, target string) error {
	reader, err := os.Open(tarball)
	if nil != err {
		return err
	}
	defer reader.Close()
	tarReader := tar.NewReader(reader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if nil != err {
			return err
		}

		path := filepath.Join(target, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); nil != err {
				return err
			}
			continue
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if nil != err {
			return err
		}
		defer file.Close()
		_, err = io.Copy(file, tarReader)
		if nil != err {
			return err
		}
	}
	return nil
}
