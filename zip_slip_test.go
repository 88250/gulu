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
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestUnzipRejectsPathTraversal(t *testing.T) {
	tmp := t.TempDir()
	zipPath := filepath.Join(tmp, "mal_path_traversal.zip")
	outDir := filepath.Join(tmp, "out")

	entries := map[string]string{
		"../outside.txt": "pwn",
		"safe.txt":       "ok",
	}
	createZip(t, zipPath, entries)

	err := Zip.Unzip(zipPath, outDir)
	if err == nil {
		t.Fatalf("expected error when unzipping path-traversal zip")
	}

	// ensure file outside was not created (relative to outDir parent)
	outsidePath := filepath.Join(tmp, "outside.txt")
	if _, statErr := os.Stat(outsidePath); statErr == nil {
		t.Fatalf("unexpectedly found outside file %s extracted", outsidePath)
	}
}

// Helper: create a zip file at zipPath with given entries (name -> content).
func createZip(t *testing.T, zipPath string, entries map[string]string) {
	f, err := os.Create(zipPath)
	if err != nil {
		t.Fatalf("create zip file: %v", err)
	}
	defer f.Close()

	w := zip.NewWriter(f)
	defer func() {
		if err := w.Close(); err != nil {
			t.Fatalf("close zip writer: %v", err)
		}
	}()

	for name, body := range entries {
		h := &zip.FileHeader{
			Name:   name,
			Method: zip.Deflate,
		}
		// write entry
		wr, err := w.CreateHeader(h)
		if err != nil {
			t.Fatalf("create header %s: %v", name, err)
		}
		if _, err := io.Copy(wr, strings.NewReader(body)); err != nil {
			t.Fatalf("write entry %s: %v", name, err)
		}
	}
}
