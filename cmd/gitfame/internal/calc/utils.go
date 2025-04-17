package calc

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func isEmptyFile(repoPath, blobHash string) (bool, error) {
	cmd := exec.Command("git", "-C", repoPath, "cat-file", "-s", blobHash)
	sizeOutput, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("ошибка при вызове git cat-file: %v", err)
	}

	sizeStr := strings.TrimSpace(string(sizeOutput))
	size, _ := strconv.Atoi(sizeStr)

	return size == 0, nil
}

func lastModifyingCommit(repoPath, file, commitHash string, useCommitter bool) (string, string, error) {
	var cmd *exec.Cmd

	if useCommitter {
		cmd = exec.Command("git", "-C", repoPath, "log", commitHash, "-n", "1", "--pretty=format:%H %cn", "--", file)
	} else {
		cmd = exec.Command("git", "-C", repoPath, "log", commitHash, "-n", "1", "--pretty=format:%H %an", "--", file)
	}

	output, err := cmd.Output()
	if err != nil {
		return "", "", fmt.Errorf("ошибка при вызове git log: %v", err)
	}

	// 9e2fbf54f1be4a013ff20366ec9226812db200d1 Fedor Korotkiy

	line := string(output)
	borderPos := len("9e2fbf54f1be4a013ff20366ec9226812db200d1")
	commit := line[:borderPos]
	author := strings.TrimSpace(line[borderPos+1:])

	return commit, author, nil
}
