package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// MoveFile moves a file from srcPath to destPath, preserving the file's metadata.
func MoveFile(srcPath, destPath string) error {
	if !Exists(srcPath) {
		return fmt.Errorf("source file does not exist")
	}

	srcFileInfo, err := os.Stat(srcPath)
	if err != nil {
		return fmt.Errorf("could not retrieve source file info: %w", err)
	}

	sourceFile, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("failed to copy file content: %w", err)
	}

	// Preserve file permissions
	err = destFile.Chmod(srcFileInfo.Mode())
	if err != nil {
		return fmt.Errorf("failed to set permissions on destination file: %w", err)
	}

	// Preserve timestamps
	err = os.Chtimes(destPath, srcFileInfo.ModTime(), srcFileInfo.ModTime())
	if err != nil {
		return fmt.Errorf("failed to set timestamps on destination file: %w", err)
	}

	// Remove the source file
	err = os.Remove(srcPath)
	if err != nil {
		os.Remove(destPath)
		return fmt.Errorf("failed to remove the source file: %w", err)
	}

	return nil
}

// Exists checks if a file or directory exists.
func Exists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// RemoveEmptyDir removes an empty directory.
func RemoveEmptyDir(dir string) error {
	return os.Remove(dir)
}

// CountFilesToMove counts the number of files in the rootDir, excluding directories.
func CountFilesToMove(rootDir string) (int, error) {
	count := 0
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Dir(path) != rootDir {
			count++
		}
		return nil
	})

	return count, err
}
