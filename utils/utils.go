package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/otiai10/copy"
)

// Exists returns true if the path exists.
func Exists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

// ReadFile rerads a file.
func ReadFile(path string) []byte {
	content, _ := os.ReadFile(path)

	return content
}

// IsFile returns true if the specific path is a file.
func IsFile(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

// IsDirectory returns true if the specified path is a directory.
func IsDirectory(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

// RemoveFile deletes a file and all its children.
//nolint: wrapcheck
func RemoveFile(path string) error {
	return os.RemoveAll(path)
}

// WriteFile writes the content to the destination file.
//nolint: wrapcheck
func WriteFile(content []byte, dest string, permissions int) error {
	perm := fs.FileMode(permissions)

	if err := ioutil.WriteFile(dest, content, perm); err != nil {
		return err
	}

	return nil
}

// MD5Sum computes the md5sum of a byte array.
func MD5Sum(content []byte) string {
	hash := md5.Sum(content)

	return hex.EncodeToString(hash[:])
}

// CopyFile copies the content of the source file to the specified destination.
//nolint: wrapcheck,gosec
func CopyFile(src, dest string, permissions int) error {
	content, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	if err := WriteFile(content, dest, permissions); err != nil {
		return err
	}

	return nil
}

// CopyDir recursively copies the content of a source dir to the specified destination.
//nolint: wrapcheck
func CopyDir(src, dest string) error {
	return copy.Copy(src, dest)
}

// LinkFile symlinks a source file to its destination.
//nolint: wrapcheck
func LinkFile(src, dest string) error {
	return os.Symlink(src, dest)
}
