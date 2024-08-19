package util

import (
	"os"
	"path"
	"testing"
)

func TestRelativeToAbsolutePath(t *testing.T) {
	testDir, err := os.Executable()
	if err != nil {
		t.Fatal(testDir)
	}
	testDir = path.Dir(testDir)

	testCases := []struct {
		name         string
		relativePath string
		baseDir      string
		useBaseDir   bool
		wantErr      error
		wantPath     string
	}{
		{
			name:         "with default baseDir",
			useBaseDir:   false,
			relativePath: "cool/directory/here",
			wantPath:     path.Join(testDir, "cool/directory/here"),
		},
		{
			name:         "with default baseDir and relative path",
			useBaseDir:   false,
			relativePath: "./cool/directory/here",
			wantPath:     path.Join(testDir, "cool/directory/here"),
		},
		{
			name:         "with BASE_DIR environment variable",
			useBaseDir:   true,
			baseDir:      "/custom/directory",
			relativePath: "cool/directory/here",
			wantPath:     "/custom/directory/cool/directory/here",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			if tc.useBaseDir {
				oldBaseDir := os.Getenv("BASE_DIR")
				if err := os.Setenv("BASE_DIR", tc.baseDir); err != nil {
					t.Fatal(err)
				}
				defer func() {
					t.Log(os.Setenv("BASE_DIR", oldBaseDir))
				}()
			}

			absPath, err := RelativeToAbsolutePath(tc.relativePath)
			if err != tc.wantErr {
				t.Errorf("expected error `%v`, got `%v`", tc.wantErr, err)
			}
			if absPath != tc.wantPath {
				t.Errorf("expected absPath `%s`, got `%s`", tc.wantPath, absPath)
			}
		})
	}
}
