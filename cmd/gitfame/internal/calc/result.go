package calc

type Result struct {
	authorToLinesCount map[string]int
	authorToCommits    map[string]map[string]struct{}
	authorToFilesCount map[string]int
}

func newResult() *Result {
	return &Result{
		authorToLinesCount: make(map[string]int),
		authorToCommits:    make(map[string]map[string]struct{}),
		authorToFilesCount: make(map[string]int),
	}
}

func (r *Result) Get() (map[string]int, map[string]int, map[string]int) {
	authorToCommitsCount := make(map[string]int)

	for author, commits := range r.authorToCommits {
		authorToCommitsCount[author] += len(commits)
	}

	return r.authorToLinesCount, authorToCommitsCount, r.authorToFilesCount
}

func (r *Result) Add(author string, lines, files int) {
	r.authorToLinesCount[author] += lines
	r.authorToFilesCount[author] += files
}

func (r *Result) AddCommit(author string, commit string) {
	if _, ok := r.authorToCommits[author]; !ok {
		r.authorToCommits[author] = make(map[string]struct{})
	}

	r.authorToCommits[author][commit] = struct{}{}
}

func (r *Result) Merge(other *Result) {
	for author := range other.authorToLinesCount {
		r.Add(author, other.authorToLinesCount[author], other.authorToFilesCount[author])
	}

	for author, commits := range other.authorToCommits {
		for commit := range commits {
			r.AddCommit(author, commit)
		}
	}
}
