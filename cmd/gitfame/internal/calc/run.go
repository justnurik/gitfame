package calc

import "math"

func Run(repoPath, commitHash string, files, blobs []string, workersCount uint, useCommitter bool) (*Result, error) {
	chunkSize := max((len(files)+int(workersCount)-1)/int(workersCount), 1)
	gorutineCount := int(math.Ceil(float64(len(files)) / float64(chunkSize)))

	results := make([]*Result, gorutineCount)
	for i := range results {
		results[i] = newResult()
	}

	chanErr := make(chan error, gorutineCount)
	defer close(chanErr)

	for i, j := 0, 0; i < len(files); i, j = i+chunkSize, j+1 {
		end := min(i+chunkSize, len(files))

		go worker(repoPath, commitHash, files[i:end], blobs[i:end], useCommitter, results[j], chanErr)
	}

	for range gorutineCount {
		err := <-chanErr
		if err != nil {
			return nil, err
		}
	}

	//! assert: all gorutine stopped

	//* reduce
	result := newResult()
	for _, res := range results {
		result.Merge(res)
	}

	return result, nil
}
