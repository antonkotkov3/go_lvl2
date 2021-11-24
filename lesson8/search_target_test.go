package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewSearchTarget(t *testing.T) {
	tempFile, err := os.CreateTemp(".", "temp")
	if err != nil {
		t.Fatal("cant create temp file for test")
	}
	if err = tempFile.Close(); err != nil {
		t.Fatal("cant closing temp file")
	}
	defer func() {
		if err = os.Remove(tempFile.Name()); err != nil {
			t.Fatal("cant remove temp file")
		}
	}()

	var (
		testTable = []struct {
			name        string
			fileName    string
			expectError bool
		}{
			{name: "empty file name", fileName: "", expectError: true},
			{name: "current folder", fileName: ".", expectError: false},
			{name: "prev-folder", fileName: "../", expectError: false},
			{name: "for temp file", fileName: tempFile.Name(), expectError: false},
		}
	)

	for _, tc := range testTable {
		targetStruct, err := NewSearchTarget(tc.fileName)

		if (err != nil && !tc.expectError) || (err == nil && tc.expectError) {
			t.Fatalf(
				"Name: %q, expect err: %t, but got: %v\n",
				tc.name,
				tc.expectError,
				err,
			)
		}

		if targetStruct != nil {
			tcAbs, err := filepath.Abs(tc.fileName)
			if err != nil {
				t.Fatalf("Name: %q. Cant read abs for %q", tc.name, tc.fileName)
			}
			targetAbs, err := filepath.Abs(targetStruct.Name)
			if err != nil {
				t.Fatalf("Name: %q. Cant read abs-target for %q", tc.name, targetStruct.Name)
			}

			if tcAbs != targetAbs {
				t.Fatalf(
					"Name: %q, expected struct-file name: %s, but got: %s",
					tc.name,
					tc.fileName,
					targetStruct.Name,
				)
			}
		}
	}

}
