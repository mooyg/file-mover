package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func MoveFile(srcPath string, destPath string) error {
	fileExists := fileExists(destPath)

	if !fileExists {
		return fmt.Errorf("file already exists")
	}

	sourceFile, err := os.Open(srcPath)
	defer sourceFile.Close()

	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}

	destFile, err := os.Create(destPath)
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)

	if err != nil {
		return fmt.Errorf("failed to copy file content: %w", err)
	}

	err = os.Remove(srcPath)

	if err != nil {
		os.Remove(destPath)
		return fmt.Errorf("failed to remove the source file reverting: %w", err)
	}
	return nil

}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func RemoveEmptyDir(dir string) error {
	return os.Remove(dir)
}

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
