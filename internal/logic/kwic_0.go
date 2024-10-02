package logic

// 主程序-子程序结构
import (
	"strings"
)

func ProcessMainSub(content string) string {
	lines := strings.Split(content, "\n")
	res := generate(lines)
	sortResults(res)
	return strings.Join(res, "\n")
}

func generate(lines []string) []string {
	var results []string
	for _, line := range lines {
		words := strings.Fields(line)
		for i := 0; i < len(words); i++ {
			shifted := append(words[i:], words[:i]...)
			results = append(results, strings.Join(shifted, " "))
		}
	}
	return results
}
