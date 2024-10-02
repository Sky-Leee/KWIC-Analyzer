package logic

// 事件系统结构

import (
	"sort"
	"strings"
	"sync"
)

type Event struct {
	Lines   []string
	Results []string
}

type Observer interface {
	Update(event Event)
}

type KWICGenerator struct {
	observers []Observer
}

func (k *KWICGenerator) RegisterObserver(o Observer) {
	k.observers = append(k.observers, o)
}

func (k *KWICGenerator) NotifyObservers(event Event) {
	for _, observer := range k.observers {
		observer.Update(event)
	}
}

func (k *KWICGenerator) GenerateKWIC(lines []string) {
	var results []string
	for _, line := range lines {
		words := strings.Fields(line)
		for i := 0; i < len(words); i++ {
			shifted := append(words[i:], words[:i]...)
			results = append(results, strings.Join(shifted, " "))
		}
	}
	k.NotifyObservers(Event{Lines: lines, Results: results})
}

type FileWriter struct {
	res string
}

var fw FileWriter
var mu sync.Mutex

func (fw *FileWriter) Update(event Event) {
	sort.Slice(event.Results, func(i, j int) bool {
		return strings.ToLower(event.Results[i]) < strings.ToLower(event.Results[j])
	})

	fw.res = strings.Join(event.Results, "\n")
}

func ProcessEventSystem(content string) string {
	lines := strings.Split(content, "\n")

	mu.Lock()
	defer mu.Unlock()
	
	generator := &KWICGenerator{}
	generator.RegisterObserver(&fw)
	generator.GenerateKWIC(lines)

	return fw.res
}
