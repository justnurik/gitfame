package calc

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func nextLine(scanner *bufio.Scanner) ([]string, bool) {
	ok := scanner.Scan()
	if !ok {
		return nil, false
	}

	line := scanner.Text()
	return strings.Fields(line), true
}

func skipLine(scanner *bufio.Scanner, count int) {
	for range count {
		scanner.Scan()
	}
}

func worker(repoPath, commitHash string, files, blobs []string, useCommitter bool, result *Result, chanErr chan<- error) {

	for i, file := range files {

		uniqAuthorInFile := make(map[string]struct{})
		commitToAuthor := make(map[string]string)

		if ok, err := isEmptyFile(repoPath, blobs[i]); ok {
			if err != nil {
				chanErr <- err
				return
			}

			lastCommit, author, err := lastModifyingCommit(repoPath, file, commitHash, useCommitter)
			if err != nil {
				chanErr <- fmt.Errorf("ошибка при получении последнего коммита для %s: %v", file, err)
				return
			}

			result.Add(author, 0, 1)
			result.AddCommit(author, lastCommit)
			continue
		}

		cmd := exec.Command("git", "-C", repoPath, "blame", "--porcelain", commitHash, "--", file)
		output, err := cmd.StdoutPipe()
		if err != nil {
			chanErr <- fmt.Errorf("ошибка при создании stdout для git blame (%s): %v", file, err)
			return
		}

		if err := cmd.Start(); err != nil {
			chanErr <- fmt.Errorf("ошибка при запуске git blame (%s): %v", file, err)
			return
		}

		var ok bool
		var commitHash string
		var authorOrCommitter string
		var lineCount int

		scanner := bufio.NewScanner(output)

		for {

			commitHash, lineCount, ok = headHandle(scanner)
			if !ok {
				break
			}

			name, ok := authorBlockHandle(scanner, useCommitter, lineCount)
			if ok {
				authorOrCommitter = name
				skipBlockHandle(scanner, lineCount)
			} else {
				authorOrCommitter = commitToAuthor[commitHash]
			}

			// Подсчет
			result.Add(authorOrCommitter, lineCount, 0)
			if _, exists := uniqAuthorInFile[authorOrCommitter]; !exists {
				uniqAuthorInFile[authorOrCommitter] = struct{}{}
				result.Add(authorOrCommitter, 0, 1)
			}
			commitToAuthor[commitHash] = authorOrCommitter
			result.AddCommit(authorOrCommitter, commitHash)

		}

	}

	chanErr <- nil
}

func headHandle(scanner *bufio.Scanner) (string, int, bool) {
	// b668f2179e493d71b7ad47f36d1e29438dcd68a8 1 1 4

	var commitHash string
	var lineCount int

	fields, ok := nextLine(scanner)
	if !ok {
		return "", 0, false
	}

	commitHash = fields[0]
	lineCount, _ = strconv.Atoi(fields[3])

	return commitHash, lineCount, true
}

func authorBlockHandle(scanner *bufio.Scanner, useCommitter bool, lineCount int) (string, bool) {
	//! case 1
	// author Arseny Balobanov
	// author-mail <verytable@yandex.ru>
	// author-time 1584005571
	// author-tz +0300
	// committer Arseny Balobanov
	// committer-mail <verytable@yandex.ru>
	// committer-time 1584005571
	// committer-tz +0300

	//! case 2
	// "sort"
	// b668f2179e493d71b7ad47f36d1e29438dcd68a8 6 7
	// 				"testing"
	// b668f2179e493d71b7ad47f36d1e29438dcd68a8 7 8
	//
	// b668f2179e493d71b7ad47f36d1e29438dcd68a8 8 9
	//

	scanner.Scan()
	line := scanner.Text()
	if strings.HasPrefix(line, "\t") { //! case 2
		skipLine(scanner, 2*(lineCount-1))
		return "", false
	}

	//! case 1
	skipLine(scanner, 3) // mail, time, tz

	if useCommitter {

		scanner.Scan()
		line := scanner.Text()
		skipLine(scanner, 3) // mail, time, tz

		return line[len("committer "):], true
	}

	// !useCommitter
	authorOrCommitter := line[len("author "):]
	skipLine(scanner, 4) // committer block

	return authorOrCommitter, true
}

func skipBlockHandle(scanner *bufio.Scanner, lineCount int) {
	//! case 1
	// summary Adding lrucache task readme+solution+tests.
	// filename lrucache/cache_test.go
	// 		package lrucache
	// b668f2179e493d71b7ad47f36d1e29438dcd68a8 2 2
	//
	// b668f2179e493d71b7ad47f36d1e29438dcd68a8 3 3
	// 		import (
	// b668f2179e493d71b7ad47f36d1e29438dcd68a8 4 4
	// 				"math/rand"

	//! case 2
	// summary Better error message in cache_test.go
	// previous 2d99b2b3361d4aedad561824b1ce7296aad57e2b lrucache/cache_test.go
	// filename lrucache/cache_test.go
	// 		package lrucache
	// b668f2179e493d71b7ad47f36d1e29438dcd68a8 2 2
	//
	// b668f2179e493d71b7ad47f36d1e29438dcd68a8 3 3
	// 		import (
	// b668f2179e493d71b7ad47f36d1e29438dcd68a8 4 4
	// 				"math/rand"

	skipLine(scanner, 1) // summary
	for {
		fields, _ := nextLine(scanner)
		if fields[0] == "previous" {
			skipLine(scanner, 1) // filename
			break
		}

		if fields[0] == "filename" {
			break
		}
	}

	skipLine(scanner, 1)               // первая строка кода
	skipLine(scanner, 2*(lineCount-1)) // код
}
