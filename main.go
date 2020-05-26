package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"

	"sigs.k8s.io/yaml"
)

// Scores is a map of names to scores
type Scores map[string][]int

// Results is the output format
type Results struct {
	Winners           []Winner `json:"winners"`
	Disqualifications []string `json:"disqualifications"`
}

// Winner is a name and their average score
type Winner struct {
	Name string  `json:"name"`
	Avg  float32 `json:"avg"`
}

func main() {
	input, err := readScores("scores.yaml")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	out := Results{
		Winners:           []Winner{},
		Disqualifications: []string{},
	}

	for name, scores := range *input {
		// Disqualify contestents without AT LEAST 7 scores.
		if len(scores) < 7 {
			out.Disqualifications = append(out.Disqualifications, name)
			continue
		}
		// Exclude the best and worst scores.
		// sort.Ints uses a variant of quicksort
		// perf is avg O(n log n) and worst O(n^2)
		// Could be replaced with a O(n) solution, if needed.
		sort.Ints(scores)
		scores = scores[1 : len(scores)-1]

		avg := float32(sum(scores)) / float32(len(scores))

		out.Winners = append(out.Winners, Winner{
			Name: name,
			Avg:  avg,
		})
	}

	// Select the top 3 avg scores.
	// sort.Slice uses a variant of quicksort
	// perf is avg O(n log n) and worst O(n^2)
	// Could be replaced with a O(n) solution, if needed.
	sort.Slice(out.Winners, func(i, j int) bool {
		return out.Winners[i].Avg > out.Winners[j].Avg
	})
	if len(out.Winners) > 3 {
		out.Winners = out.Winners[:3]
	}

	// sort disqualifications, so the result is stable
	sort.Slice(out.Disqualifications, func(i, j int) bool {
		return strings.Compare(out.Disqualifications[i], out.Disqualifications[j]) < 0
	})

	err = writeResults(out, "results.yaml")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func readScores(path string) (*Scores, error) {
	byteValue, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read file: %v", err)
	}
	var scores Scores
	err = yaml.UnmarshalStrict(byteValue, &scores)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse scores: %v", err)
	}
	return &scores, nil
}

func writeResults(r Results, path string) error {
	pretty, err := yaml.Marshal(r)
	if err != nil {
		return fmt.Errorf("Failed to marshal results to yaml: %v", err)
	}
	err = ioutil.WriteFile(path, pretty, 0644)
	if err != nil {
		return fmt.Errorf("Failed to write results to file: %v", err)
	}
	return nil
}
