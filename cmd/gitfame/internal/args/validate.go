package args

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"gitlab.com/justnurik/gitfame/cmd/gitfame/internal/programm"
)

func isValidPath(path string) (bool, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return false, err
	}

	info, err := os.Stat(absPath)
	if err != nil {
		return false, err
	}

	return info.IsDir(), nil
}

func isValidGitCommit(repoPath, commit string) (bool, error) {
	cmd := exec.Command("git", "-C", repoPath, "rev-parse", "--verify", commit)
	_, err := cmd.CombinedOutput()

	if err != nil {
		return false, err
	}

	return true, nil
}

func Validate(config *programm.Config) error {
	{ // config.Repository
		ok, err := isValidPath(config.Repository)
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("некорректный путь: %s", config.Repository)
		}
	}

	{ // config.Revision
		valid, err := isValidGitCommit(config.Repository, config.Revision)
		if err != nil {
			return fmt.Errorf("это не гит репозитории: %s", config.Repository)
		}
		if !valid {
			return fmt.Errorf("некорректный коммит: %s", config.Revision)
		}
	}

	// config.Exclude, config.RestrictTo
	for _, pattern := range append(config.Exclude, config.RestrictTo...) {
		if _, err := filepath.Match(pattern, ""); err != nil {
			return fmt.Errorf("некорректный Glob-паттерн: %s", pattern)
		}
	}

	return nil
}
