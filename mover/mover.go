package mover

import (
	"fmt"
	"os"
	"path/filepath"

	utils "github.com/mooyg/file-mover/fileutils"
	pm "github.com/schollz/progressbar/v3"
)

type FileMover struct {
	RootDir  string
	Progress pm.ProgressBar
}

func NewFileMover(rootDir string) (*FileMover, error) {
	if _, err := os.Stat(rootDir); err != nil {
		return nil, fmt.Errorf("No path found")
	}

	total, err := utils.CountFilesToMove(rootDir)

	if err != nil {
		return nil, fmt.Errorf("error counting files: %w", err)
	}

	return &FileMover{
		RootDir:  rootDir,
		Progress: *pm.Default(int64(total)),
	}, nil
}

func (fm *FileMover) MoveFiles() error {

	err := fm.walkTree()

	if err != nil {
		return fmt.Errorf("error moving files: %w", err)
	}

	fmt.Println("Done")

	return nil
}

func (fm *FileMover) walkTree() error {
	return filepath.Walk(fm.RootDir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Dir(path) != fm.RootDir {
			destPath := filepath.Join(fm.RootDir, filepath.Base(path))
			err := utils.MoveFile(path, destPath)

			if err != nil {
				return err
			}

			fm.Progress.Add(1)
		}

		return nil
	})
}
