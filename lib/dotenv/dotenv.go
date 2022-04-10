package dotenv

import (
	"fmt"
	"os"
	"strings"
)

func Load(path string) error {
	f, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Load: read file error: %v", err)
	}
	data := strings.ReplaceAll(strings.ReplaceAll(string(f), "\r\n", "\n"), "\r", "\n")
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}
		kv := strings.SplitN(line, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("Load: invalid line: %s", line)
		}
		os.Setenv(kv[0], kv[1])
	}
	return nil
}
