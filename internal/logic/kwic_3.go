package logic

// 管道-过滤结构

import (
	"sort"
	"strings"
)

type Filter func([]string) []string

func sortResults(results []string) {
	sort.Slice(results, func(i, j int) bool {
		return strings.ToLower(results[i]) < strings.ToLower(results[j])
	})
}

func ProcessPipeFilter(content string) string {
	lines := strings.Split(content, "\n")

	results := generate(lines)
	sortResults(results)

	return strings.Join(results, "\n")
}
