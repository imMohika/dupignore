package cmd

import (
	"fmt"
	"github.com/immohika/dupignore/gitignore"
	"github.com/immohika/dupignore/internal/config"
	"strings"
)

type DedupCmd struct {
	FilePath     string `arg:"" help:"Path to gitignore file" type:"path" default:"./.gitignore"`
	KeepComments *bool  `help:"Keep comments" short:"c" type:"bool"`
	KeepNewLines *bool  `help:"Keep newlines" short:"l" type:"bool"`
}

func (cmd *DedupCmd) Run(cfg *config.Config, _ *GlobalVars) error {
	keepComments := cfg.KeepComments
	if cmd.KeepComments != nil {
		keepComments = *cmd.KeepComments
	}

	keepNewLines := cfg.KeepNewLines
	if cmd.KeepNewLines != nil {
		keepNewLines = *cmd.KeepNewLines
	}

	processor := gitignore.Processor{
		KeepComments: keepComments,
		KeepNewLines: keepNewLines,
	}

	content, err := gitignore.ReadFile(cmd.FilePath)
	if err != nil {
		return err
	}

	processed := processor.ProcessContent(content)
	if err := gitignore.WriteFile(cmd.FilePath, processed); err != nil {
		return fmt.Errorf("save error: %w", err)
	}

	originalCount := len(strings.Split(content, "\n"))
	newCount := len(strings.Split(processed, "\n"))
	fmt.Printf("Processed: %s\n", cmd.FilePath)
	fmt.Printf("Done (%d → %d lines)\n", originalCount, newCount)

	return nil
}
