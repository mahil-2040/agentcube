/*
Copyright The Volcano Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mtls

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestWaitForCertificateFiles(t *testing.T) {
	dir := t.TempDir()
	cfg := Config{
		CertFile: filepath.Join(dir, "cert.pem"),
		KeyFile:  filepath.Join(dir, "key.pem"),
		CAFile:   filepath.Join(dir, "ca.pem"),
	}
	for _, path := range []string{cfg.CertFile, cfg.KeyFile, cfg.CAFile} {
		if err := os.WriteFile(path, []byte("test"), 0600); err != nil {
			t.Fatalf("write test file: %v", err)
		}
	}

	if err := WaitForCertificateFiles(cfg, time.Second); err != nil {
		t.Fatalf("WaitForCertificateFiles() error = %v", err)
	}
}

func TestCertificateFilesExist_MissingFile(t *testing.T) {
	dir := t.TempDir()
	cfg := Config{
		CertFile: filepath.Join(dir, "cert.pem"),
		KeyFile:  filepath.Join(dir, "key.pem"),
		CAFile:   filepath.Join(dir, "ca.pem"),
	}
	for _, path := range []string{cfg.CertFile, cfg.KeyFile} {
		if err := os.WriteFile(path, []byte("test"), 0600); err != nil {
			t.Fatalf("write test file: %v", err)
		}
	}

	if certificateFilesExist(cfg) {
		t.Fatal("certificateFilesExist() = true, want false")
	}
}
