package utils

import (
	"fmt"
	"path/filepath"
	"time"
)

func GenerateUniqueFileName(original string) string {
	ext := filepath.Ext(original)
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%d%s", timestamp, ext)

}

func Join(strs []string, sep string) string {
	result := ""
	for i, s := range strs {
		if i > 0 {
			result += sep
		}
		result += s
	}
	return result
}
