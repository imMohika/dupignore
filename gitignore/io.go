package gitignore

import (
	"fmt"
	"os"
)

func ReadFile(path string) (string, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", fmt.Errorf("file does not exist: %s", path)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read error: %w", err)
	}
	return string(content), nil
}

func WriteFile(path string, content string) error {
	err := os.WriteFile(path, []byte(content), 0600)
	if err != nil {
		return fmt.Errorf("write error: %w", err)
	}
	return nil
}
