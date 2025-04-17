//go:build !solution

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"gitlab.com/justnurik/gitfame/cmd/gitfame/internal/args"
	"gitlab.com/justnurik/gitfame/cmd/gitfame/internal/calc"
	"gitlab.com/justnurik/gitfame/cmd/gitfame/internal/filter"
	"gitlab.com/justnurik/gitfame/cmd/gitfame/internal/load"
	"gitlab.com/justnurik/gitfame/cmd/gitfame/internal/print"
	"gitlab.com/justnurik/gitfame/cmd/gitfame/internal/programm"
)

var programmConfig programm.Config

func main() {
	cmd := &cobra.Command{
		Use:   "gitfame",
		Short: "Calculate Git repository statistics per author",
		RunE:  run,
	}

	args.AddFlags(cmd, &programmConfig)
	err := cmd.Execute()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, _ []string) error {
	err := args.Validate(&programmConfig)
	if err != nil {
		return err
	}

	// if programmConfig.ProgressBar {
	// 	//! TODO
	// }

	if programmConfig.Languages != nil {
		load.StartLoadLanguageExtensions()
	}

	files, blobs, err := load.FilesAndBlobs(programmConfig.Repository, programmConfig.Revision)
	if err != nil {
		return err
	}

	filteredFiles, err := filter.All(files, programmConfig.Extensions, programmConfig.Languages, programmConfig.Exclude, programmConfig.RestrictTo)
	if err != nil {
		return err
	}

	result, err := calc.Run(programmConfig.Repository, programmConfig.Revision, filteredFiles, blobs, uint(programmConfig.WorkersCount), programmConfig.UseCommitter)
	if err != nil {
		return err
	}

	authToLineCount, authToCommitCount, authToFileCount := result.Get()
	return print.Result(programmConfig.Format, print.NewConfig(programmConfig.OrderBy, authToLineCount, authToCommitCount, authToFileCount))
}
