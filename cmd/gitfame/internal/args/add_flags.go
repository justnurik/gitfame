package args

import (
	"github.com/spf13/cobra"
	"gitlab.com/justnurik/gitfame/cmd/gitfame/internal/programm"
)

// Добавляем обрабатываемые флаги
func AddFlags(cmd *cobra.Command, config *programm.Config) {
	config.WorkersCount = 4 // default: 4

	cmd.Flags().StringVar(&config.Repository, "repository", ".", flagDocRepository)
	cmd.Flags().StringVar(&config.Revision, "revision", "HEAD", flagDocRevision)
	cmd.Flags().Var(&config.OrderBy, "order-by", flagDocOrderBy) // default: "line"
	cmd.Flags().BoolVar(&config.UseCommitter, "use-committer", false, flagDocUseCommitter)
	cmd.Flags().Var(&config.Format, "format", flagDocFormat) // default: "tabular"
	cmd.Flags().StringSliceVar(&config.Extensions, "extensions", nil, flagDocExtensions)
	cmd.Flags().StringSliceVar(&config.Languages, "languages", nil, flagDocLanguages)
	cmd.Flags().StringSliceVar(&config.Exclude, "exclude", nil, flagDocExclude)
	cmd.Flags().StringSliceVar(&config.RestrictTo, "restrict-to", nil, flagDocRestrictTo)
	cmd.Flags().BoolVar(&config.ProgressBar, "progress-bar", true, flagDocProgressBar)
	cmd.Flags().Var(&config.WorkersCount, "workers-count", flagDocWorkersCount) // default: 4
}
