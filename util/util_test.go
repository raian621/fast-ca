package util

import (
	"os"
	"path"
	"regexp"
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

func TestGenerateSalt(t *testing.T) {
  t.Parallel()
  salt, err := GenerateSalt()
  if err != nil {
    t.Fatal(err)
  }
  if len(salt) != 16 {
    t.Error("expected salt to have 16 bytes")
  }
}

func TestHashPassword(t *testing.T) {
  t.Parallel()
  salt, err := GenerateSalt()
  if err != nil {
    t.Fatal(err)
  }

  password := "password1234"
  passhash := HashPassword(password, salt)
  if len(passhash) != 97 {
    t.Error("expected passhash to be 97 characters long")
  }

  pattern := `\$argon2id\$v=19\$m=65536,t=3,p=2\$.+\$.+`
  matched, err := regexp.MatchString(
    pattern,
    passhash,
  )
  if err != nil {
    t.Error(err)
  }
  if !matched {
    t.Errorf("expected `%s` to match pattern `%s`", passhash, pattern)
  }
}

func TestValidatePassword(t *testing.T) {
  t.Parallel()
  password := "password1234"
  notThePassword := "notthepasswordlol"

  salt, err := GenerateSalt()
  if err != nil {
    t.Fatal(err)
  }
  passhash := HashPassword(password, salt)

  if valid, err := ValidatePassword(password, passhash); err != nil {
    t.Error(err)
  } else if !valid {
    t.Errorf("expected `password` to be validated")
  }

  if valid, err := ValidatePassword(notThePassword, passhash); err != nil {
    t.Error(err)
  } else if valid {
    t.Errorf("expected `notThePassword` to not be validated")
  }
}
