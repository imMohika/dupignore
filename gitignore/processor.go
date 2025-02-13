package gitignore

import (
	"strings"
)

type Processor struct {
	KeepComments bool
	KeepNewLines bool
}

func (p *Processor) ProcessContent(content string) string {
	lines := strings.Split(content, "\n")
	processed := p.ProcessLines(lines)

	result := strings.Join(processed, "\n")
	return result
}

func (p *Processor) ProcessLines(lines []string) []string {
	var processed []string
	seen := make(map[string]struct{})

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		switch {
		case trimmed == "":
			if p.KeepNewLines {
				processed = append(processed, line)
			}
		case strings.HasPrefix(trimmed, "#"):
			if p.KeepComments {
				processed = append(processed, line)
			}
		default:
			if _, exists := seen[trimmed]; !exists {
				seen[trimmed] = struct{}{}
				processed = append(processed, line)
			}
		}
	}
	return processed
}
