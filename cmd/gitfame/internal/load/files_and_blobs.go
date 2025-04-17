package load

import (
	"fmt"
	"os/exec"
	"strings"
)

func FilesAndBlobs(repository, commit string) ([]string, []string, error) {
	cmd := exec.Command("git", "-C", repository, "ls-tree", "-r", commit)
	output, err := cmd.Output()
	if err != nil {
		return nil, nil, fmt.Errorf("ошибка при вызове git ls-tree: %v", err)
	}

	lines := strings.Split(string(output), "\n")

	files := make([]string, len(lines))
	blobs := make([]string, len(lines))

	for i, line := range lines[:len(lines)-1] {
		// 100644 blob 1b8a8bbb6e542170c363300a7cabfa0a688bda80	yamlembed/types_test.go

		fields := strings.Fields(line)
		blobs[i] = fields[2]

		index := len(fields[0]) + 1 + len(fields[1]) + 1 + len(fields[2]) + 1
		files[i] = strings.TrimSpace(line[index:])
	}

	return files, blobs, nil
}
