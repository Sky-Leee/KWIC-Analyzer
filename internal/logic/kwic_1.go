package logic

// 面向对象结构

import (
	"strings"
)

type KWIC struct {
	lines   []string
	results []string
}

func (k *KWIC) read(content string) error {
	k.lines = strings.Split(content, "\n")
	return nil
}

func (k *KWIC) generate() {
	for _, line := range k.lines {
		words := strings.Fields(line)
		for i := 0; i < len(words); i++ {
			shifted := append(words[i:], words[:i]...)
			k.results = append(k.results, strings.Join(shifted, " "))
		}
	}
	sortResults(k.results)
}

func (k *KWIC) write() string {
	return strings.Join(k.results, "\n")
}

func ProcessObjectOriented(content string) string {
	k := &KWIC{}
	k.read(content)
	k.generate()
	return k.write()
}
