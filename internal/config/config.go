package config

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"strings"
)

const (
	KeepComments = "keepcomments"
	KeepNewLines = "keepnewlines"
)

type Config struct {
	KeepNewLines bool `toml:"keep_new_lines"`
	KeepComments bool `toml:"keep_comments"`

	Path string
}

func From(path string) (*Config, error) {
	c := Config{
		KeepNewLines: false,
		KeepComments: true,
		Path:         path,
	}

	_, err := toml.DecodeFile(path, &c)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &c, nil
		}

		return nil, fmt.Errorf("toml decode: %w", err)
	}

	err = c.validate()
	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	return &c, nil
}

func (c *Config) validate() error {
	return nil
}

func (c *Config) Update(key string, value bool) error {
	k, err := ExpandKey(key)
	if err != nil {
		return err
	}

	switch k {
	case KeepComments:
		c.KeepComments = value
	case KeepNewLines:
		c.KeepNewLines = value
	}

	file, err := os.Create(c.Path)
	if err != nil {
		return fmt.Errorf("create config file: %w", err)
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	if err := encoder.Encode(c); err != nil {
		return fmt.Errorf("encode config: %w", err)
	}

	fmt.Printf("Updated value of %s to %v\n", key, value)
	return nil
}

func ExpandKey(key string) (string, error) {
	k := strings.ToLower(key)
	switch k {
	case KeepComments, "kc":
		return KeepComments, nil
	case KeepNewLines, "kl":
		return KeepNewLines, nil
	}

	return "", fmt.Errorf("unknown config key: %s", key)
}
