package cmd

import (
	"fmt"
	"github.com/immohika/dupignore/internal/config"
	"strings"
)

type ConfigCmd struct {
	Key   string `arg:"" help:"Config key. Leave empty to show value of all keys" optional:""`
	Value string `arg:"" help:"Update value of specified key (true/false)" optional:""`
}

func (c *ConfigCmd) Run(cfg *config.Config, _ *GlobalVars) error {
	fmt.Printf("Config location: %s\n", cfg.Path)

	if c.Key == "" {
		fmt.Printf("Keep Comments (keepcomments/kc): %t\n", cfg.KeepComments)
		fmt.Printf("Keep New Lines (keepnewlines/kl): %t\n", cfg.KeepNewLines)
		return nil
	}

	key, err := config.ExpandKey(c.Key)
	if err != nil {
		return err
	}

	if c.Value == "" {
		switch key {
		case config.KeepComments:
			fmt.Printf("Keep Comments: %t\n", cfg.KeepComments)
		case config.KeepNewLines:
			fmt.Printf("Keep New Lines: %t\n", cfg.KeepNewLines)
		}
		return nil
	}

	var value bool
	switch strings.ToLower(c.Value) {
	case "true", "1", "yes":
		value = true
	case "false", "0", "no":
		value = false
	default:
		return fmt.Errorf("invalid boolean value: %s (must be true/false)", c.Value)
	}

	err = cfg.Update(key, value)
	if err != nil {
		return err
	}

	return nil
}
