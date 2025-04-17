package programm

type Config struct {
	Repository   string
	Revision     string
	OrderBy      OrderType
	UseCommitter bool
	Format       FormatType
	Extensions   []string
	Languages    []string
	Exclude      []string
	RestrictTo   []string
	ProgressBar  bool
	WorkersCount PositiveInt
}
